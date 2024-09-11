package e2e_test

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"upsider-coding-test/infrastructure/handler"
	"upsider-coding-test/infrastructure/server"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAndLogin(t *testing.T) {
	// Setup
	router := gin.Default()
	server.Route(router)

	// Register#Setup
	randomEmail := strconv.Itoa(rand.Intn(1000000)) + "@sample.com"

	w := httptest.NewRecorder()
	registerUser := handler.UserCreateParams{
		Username:    "John Doe",
		Email:       randomEmail,
		RawPassword: "password",
		CompanyID:   "b8e7fce5-77a5-4e64-9e3c-90e0c5b4c17d",
	}
	payload, err := json.Marshal(registerUser)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/users", strings.NewReader(string(payload)))
	if err != nil {
		t.Fatal(err)
	}
	// Register#Execise
	router.ServeHTTP(w, req)
	// Register#Verify
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "id")

	// Login#Setup
	w = httptest.NewRecorder()
	loginUser := handler.UserLoginParams{
		Email:       randomEmail,
		RawPassword: "password",
	}
	payload, err = json.Marshal(loginUser)
	if err != nil {
		t.Fatal(err)
	}
	req, err = http.NewRequest("POST", "/api/login", strings.NewReader(string(payload)))
	if err != nil {
		t.Fatal(err)
	}

	// Login#Execise
	router.ServeHTTP(w, req)

	// Login#Verify
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "token")
	responseBody := make(map[string]interface{})
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, responseBody["token"])
}
