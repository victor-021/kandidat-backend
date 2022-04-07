package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

// Use a single instance of Validate, it caches struct info.
var validate *validator.Validate

func TestMain(m *testing.M) {
	// Run test suite
	os.Exit(RunTests(m))
}

func RunTests(m *testing.M) int {
	setupConfig()
	gin.SetMode(gin.TestMode)

	dbPool = setupDBPool()
	router = setupRouter()

	defer dbPool.Close()

	validate = validator.New()

	return m.Run()
}

func TestPingRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetCommunities(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/communities", nil)
	router.ServeHTTP(w, req)

	var communityarray []Community

	err := json.Unmarshal(w.Body.Bytes(), &communityarray)
	if err != nil {
		t.Errorf("Error unmarshalling json: %v", err)
	}

	// Validate all Community structs in the array communityarray
	for _, community := range communityarray {
		err = validate.Struct(community)
		if err != nil {
			t.Errorf("Error validating struct: %v", err)
		}
	}

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUser(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(w, req)

	var user User

	err := json.Unmarshal(w.Body.Bytes(), &user)
	if err != nil {
		t.Errorf("Error unmarshalling json: %v", err)
	}

	// Validate User struct
	err = validate.Struct(user)
	if err != nil {
		t.Errorf("Error validating struct: %v", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserCommunities(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1/communities", nil)
	router.ServeHTTP(w, req)

	var communityarray []Community

	err := json.Unmarshal(w.Body.Bytes(), &communityarray)
	if err != nil {
		t.Errorf("Error unmarshalling json: %v", err)
	}

	// Validate all Community structs in the array communityarray
	for _, community := range communityarray {
		err = validate.Struct(community)
		if err != nil {
			t.Errorf("Error validating struct: %v", err)
		}
	}

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserFollowers(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1/followers", nil)
	router.ServeHTTP(w, req)

	var userarray []User

	err := json.Unmarshal(w.Body.Bytes(), &userarray)
	if err != nil {
		t.Errorf("Error unmarshalling json: %v", err)
	}

	// Validate all User structs in the array userarray
	for _, user := range userarray {
		err = validate.Struct(user)
		if err != nil {
			t.Errorf("Error validating struct: %v", err)
		}
	}

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProduct(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/1", nil)
	router.ServeHTTP(w, req)

	var product Product

	err := json.Unmarshal(w.Body.Bytes(), &product)
	if err != nil {
		t.Errorf("Error unmarshalling json: %v", err)
	}

	// Validate Product struct
	err = validate.Struct(product)
	if err != nil {
		t.Errorf("Error validating struct: %v", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserGetReviews(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1/reviews", nil)
	router.ServeHTTP(w, req)

	var reviewarray []Review

	err := json.Unmarshal(w.Body.Bytes(), &reviewarray)
	if err != nil {
		t.Errorf("Error unmarshalling json: %v", err)
	}

	// Validate all Review structs in the array reviewarray
	for _, review := range reviewarray {
		err = validate.Struct(review)
		if err != nil {
			t.Errorf("Error validating struct: %v", err)
		}
	}

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserProducts(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1/products", nil)
	router.ServeHTTP(w, req)

	var productarray []Product

	err := json.Unmarshal(w.Body.Bytes(), &productarray)
	if err != nil {
		t.Errorf("Error unmarshalling json: %v", err)
	}

	// Validate all Product structs in the array productarray
	for _, product := range productarray {
		err = validate.Struct(product)
		if err != nil {
			t.Errorf("Error validating struct: %v", err)
		}
	}

	assert.Equal(t, http.StatusOK, w.Code)
}
