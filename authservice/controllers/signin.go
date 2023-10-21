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
	"errors"
	"fmt"
	"net/http"
	"time"

	"authservice/jwt"
	"authservice/models"

	"go.uber.org/zap"
)

// SigninRequest represents the request body for the signin API.
type SigninRequest struct {
	Email        string `json:"email"`
	PasswordHash string `json:"passwordhash"`
}

type SigninController struct {
	logger *zap.Logger
}
type SigninResponse struct {
	Token string `json:"token"`
}

// NewSigninController returns a frsh Signin controller

func NewSigninController(logger *zap.Logger) *SigninController {
	return &SigninController{
		logger: logger,
	}
}

// we need this function to be private
func getSignedToken() (string, error) {
	// we make a JWT Token here with signing method of ES256 and claims.
	// claims are attributes.
	// Aud - audience
	// Iss - issuer
	// Exp - expiration of the Token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"Aud": "frontend.knowsearch.ml",
	// 	"Iss": "knowsearch.ml",
	// 	"Exp": string(time.Now().Add(time.Minute * 1).Unix()),
	// })
	claimsMap := jwt.ClaimsMap{
		Aud: "frontend.knowsearch.ml",
		Iss: "knowsearch.ml",
		Exp: fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
	}

	secret := jwt.GetSecret()
	if secret == "" {
		return "", errors.New("empty JWT secret")
	}

	header := "HS256"
	tokenString, err := jwt.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

// searches the user in the database.
func validateUser(email string, passwordHash string) (bool, error) {
	usr, exists := models.GetUserObject(email)
	if !exists {
		return false, errors.New("user does not exist")
	}
	passwordCheck := usr.ValidatePasswordHash(passwordHash)

	if !passwordCheck {
		return false, nil
	}
	return true, nil
}

// This will be supplied to the MUX router. It will be called when signin request is sent
// if user not found or not validates, returns the Unauthorized error
// if found, returns the JWT back. [How to return this?]
// Signin - sign in a user
// @Summary Sign in a user and obtain a JWT token
// @Description Sign in with email and password
// @Accept json
// @Produce json
// @Param input body SigninRequest true "User signin info"
// @Success 200 {object} SigninResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /signin [post]
func (ctrl *SigninController) SigninHandler(rw http.ResponseWriter, r *http.Request) {
	// Create a struct to hold the request data
	var signinRequest SignupRequest

	// Decode the JSON request body into the SignupRequest struct
	err := json.NewDecoder(r.Body).Decode(&signinRequest)
	if err != nil {
		ctrl.logger.Error("Error decoding request body", zap.Error(err))
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid request body"))
		return
	}

	// Now you can access the data from signinRequest
	email := signinRequest.Email
	passwordHash := signinRequest.PasswordHash

	// Validate the request and check if the user exists
	valid, err := validateUser(email, passwordHash)
	if err != nil {
		// This means either the user does not exist
		ctrl.logger.Warn("User does not exist", zap.String("email", email))
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("User Does not Exist"))
		return
	}

	if !valid {
		// This means the password is wrong
		ctrl.logger.Warn("Password is wrong", zap.String("email", email))
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Incorrect Password"))
		return
	}

	tokenString, err := getSignedToken()
	if err != nil {
		ctrl.logger.Error("Unable to sign the token", zap.Error(err))
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Internal Server Error"))
		return
	}
	ctrl.logger.Info("Token sign", zap.String("token", tokenString), zap.String("email", email))

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(tokenString))
}
