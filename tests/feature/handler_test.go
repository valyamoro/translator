package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/valyamoro/internal/domain"
	"github.com/valyamoro/internal/handler"
)

type MockUsersService struct {
	CreateFunc func(user domain.User) (domain.User, error)
}

func (m *MockUsersService) Create(user domain.User) (domain.User, error) {
	return m.CreateFunc(user)
}

func TestCreateUser_Success(t *testing.T) {
    gin.SetMode(gin.TestMode)

    mockUsersService := &MockUsersService{
        CreateFunc: func(user domain.User) (domain.User, error) {
            return user, nil
        },
    }
    handler := handler.NewUserHandler(mockUsersService)
    router := gin.Default()
    handler.InitRoutes(router)

    user := domain.User{
        ID:          1,
        Username:    "Test user",
        Dictionaries: []domain.Dictionary{},
    }

    body, _ := json.Marshal(user)
    req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    router.ServeHTTP(w, req)

    if w.Code != http.StatusCreated {
        t.Errorf("Expected status %d, but not %d", http.StatusCreated, w.Code)
    }

    var createdUser domain.User
    if err := json.Unmarshal(w.Body.Bytes(), &createdUser); err != nil {
        t.Fatalf("Failed to unmarshal response: %v", err)
    }
    if !user.Equals(createdUser) {
        t.Errorf("Expected user %+v, but got %+v", user, createdUser)
    }
}


func TestCreateUser_BindError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUsersService := &MockUsersService{}
	handler := handler.NewUserHandler(mockUsersService)
	router := gin.Default()
	handler.InitRoutes(router)

	invalidBody := `{"invalud_json"}`
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(invalidBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, but got %d", http.StatusBadRequest, w.Code)
	}
	if !bytes.Contains(w.Body.Bytes(), []byte("invalid character")) {
        t.Errorf("Expected error message to contain 'invalid character', got %s", w.Body.String())
	}
}

func TestCreateUser_CreateError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUsersService := &MockUsersService{
		CreateFunc: func(user domain.User) (domain.User, error) {
			return domain.User{}, errors.New("service error")
		},
	}
	handler := handler.NewUserHandler(mockUsersService)
	router := gin.Default()
	handler.InitRoutes(router)

	user := domain.User{
		ID: 1,
		Username: "Test user",
	}

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
        t.Errorf("Expected status %d, but got %d", http.StatusInternalServerError, w.Code)
	}
	if !bytes.Contains(w.Body.Bytes(), []byte("service error")) {
		t.Errorf("Expected error message to contain service error, got %s", w.Body)
	}
}
