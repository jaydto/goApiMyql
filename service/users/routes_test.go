package users

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jaydto/goApiMyql/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	
	t.Run("Should Fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Alex",
			LastName:  "Mike",
			Password:  "1234",
			Email:     "invalid",
		}
		marshaled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshaled))
		if err != nil {
			t.Fatal(err)

		}
		
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		// router.HandleFunc("/login", handler.handleLogin)
		router.ServeHTTP(rr, req)

		log.Default().Printf("\n The code is %v", rr.Code)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d got %d", http.StatusBadRequest, rr.Code)
		}

	})

	t.Run("Should correctly register the user", func(t *testing.T){
		payload:=types.RegisterUserPayload{
			FirstName: "user",
			LastName: "kim2w",
			Email:"alex@gmail.com",
			Password: "1234",
		}
		marshaled,_:=json.Marshal(payload)
		req, err:=http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshaled))
		if err!=nil{
			t.Fatal(err)

		}
		rr:=httptest.NewRecorder()
		router:=mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		log.Default().Printf("\n The code status is %v", rr.Code)

		if rr.Code!=http.StatusCreated{
			t.Errorf("Expected status code %d but got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct {
}

// CreateUser implements types.UserStore.
func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}

// GetUserByEmail implements types.UserStore.
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

// GetUserById implements types.UserStore.
func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}
