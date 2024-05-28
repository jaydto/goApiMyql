package users

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jaydto/goApiMyql/config"
	"github.com/jaydto/goApiMyql/service/auth"
	"github.com/jaydto/goApiMyql/types"
	"github.com/jaydto/goApiMyql/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}
// handleLogin handles user login
// @Summary User login
// @Description Logs in a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body types.LoginUserPayload true "Login payload"
// @Success 200 {object} types.Message "message"
// @Failure 400 {object} types.Error "error"
// @Failure 404 {object} map[string]string "error"
// @Router /login [post]
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// user
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user with email %v not found or invalid email or password", payload.Email))
		return
	}

	// compare passwords

	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid password or email"))
		return

	}
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJwt(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteMessage(w, http.StatusOK, map[string]string{"token": token})

}

//handleRegister handles user Registration
// @Summary User registration
// @Description Register User with Email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body types.RegisterUserPayload true "Register payload"
// @Success 201 {object} types.Message 
// @Failure 400 {object} types.Error 
// @Failure 404 {object} types.Error "error"
// @Failure 500 {object} map[string]string "error"
// @Router /register [post]
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get json payload

	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Return error if user does exist
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	// Create if does not exist
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteMessage(w, http.StatusCreated, "User created successfully")

}
