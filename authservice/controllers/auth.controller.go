// Package Authentication of Product API
//
// # Documentation for Authentication of Product API
//
// Schemes : http
// BasePath : /auth
// Version : 1.0.0
//
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// swagger:meta
package controllers

import (
	"encoding/json"
	"net/http"

	"authservice/models"
	"authservice/services"

	"go.uber.org/zap"
)

// LoginRequest represents the request body for the login API.

type AuthController struct {
	authservice *services.AuthService
	logger      *zap.Logger
}

// NewLoginController returns a fresh Login controller
func NewAuthController(logger *zap.Logger, authservice *services.AuthService) *AuthController {
	return &AuthController{
		authservice: authservice,
		logger:      logger,
	}
}

// if user not found or not validates, returns the Unauthorized error
// if found, returns the JWT back. [How to return this?]
// Login - log in a user
// @Summary Log in a user and obtain a JWT token
// @Description Log in with email and password
// @Accept json
// @Produce json
// @Param input body models.LoginRequest true "User login info"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (ctrl *AuthController) LoginHandler(rw http.ResponseWriter, r *http.Request) {
	// Create a struct to hold the request data
	var loginRequest models.LoginRequest

	// Decode the JSON request body into the LoginRequest struct
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		ctrl.logger.Error("Error decoding request body", zap.Error(err))
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid request body"))
		return
	}

	// Now you can access the data from loginRequest
	email := loginRequest.Email
	password := loginRequest.Password
	role := loginRequest.Role
	// Validate the request and check if the user exists
	user, err := ctrl.authservice.Login(email, password, role)
	if err != nil {
		// This means either the user does not exist
		ctrl.logger.Warn("User does not exist", zap.String("email", email))

		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("User Does not Exist"))
		return
	}

	// If the user exists, we sign the token
	tokenString, err := ctrl.authservice.GetSignedToken()
	if err != nil {
		ctrl.logger.Error("Unable to sign the token", zap.Error(err))
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Internal Server Error"))
		return
	}
	ctrl.logger.Info("Token sign", zap.String("token", tokenString), zap.String("email", email))

	// Return the token in the response
	rw.WriteHeader(http.StatusOK)
	// Return the user object in the response
	response := models.LoginResponse{User: user, Token: tokenString}
	json.NewEncoder(rw).Encode(response)

}

// adds the user to the database of users
// Register - register a new user
// @Summary Register a new user
// @Description Register a new user with email and password
// @Accept json
// @Produce json
// @Param input body models.RegisterRequest true "User register info"
// @Success 200 {object} models.RegisterResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (ctrl *AuthController) RegisterHandler(rw http.ResponseWriter, r *http.Request) {
	// Create a variable to hold the request data
	var registerRequest models.RegisterRequest

	// Decode the JSON request body into the RegisterRequest struct
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		ctrl.logger.Error("Error decoding request body", zap.Error(err))
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid request body"))
		return
	}

	// Now you can access the data from registerRequest
	email := registerRequest.Email
	password := registerRequest.Password
	role := registerRequest.Role
	// You can also add validation for other fields if needed

	// Validate and then add the user
	user, err := ctrl.authservice.Register(email, password, role)
	if err != nil {
		ctrl.logger.Error("Error registering user", zap.Error(err))
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error registering user"))
		return
	}

	ctrl.logger.Info("User created", zap.String("email", email))
	rw.WriteHeader(http.StatusOK)
	// Return the user object in the response
	json.NewEncoder(rw).Encode(user)
}
