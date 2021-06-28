package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goldfish-api/internal/core/ports"
	"net/http"
	"strings"
)

type HTTPUserHandler struct {
	userService ports.UserService
	router *mux.Router
}

// NewHTTPUserHandler creates a new instance of the user handler
func NewHTTPUserHandler(router *mux.Router, userService ports.UserService) *HTTPUserHandler {
	userHandler := HTTPUserHandler{
		userService: userService,
		router:      router,
	}

	// set routes of user handler
	userHandler.init()

	return &userHandler
}

// init initializes the handler and sets up the routes
func (handler *HTTPUserHandler) init() {
	handler.router.HandleFunc("/register", handler.Register)
	handler.router.HandleFunc("/login", handler.Login)
}

// Register endpoint -> create user account action
func (handler *HTTPUserHandler) Register(writer http.ResponseWriter, request *http.Request) {
	var requestBody UserRegisterRequest

	// parse and validate request body
	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil || !requestBody.IsValid() {
		http.Error(writer, INVALID_BODY, http.StatusBadRequest)
		return
	}

	// save user
	user := requestBody.ToUser()
	jwt, expiration, err := handler.userService.Create(&user)

	if err != nil {
		// database unique key violation => user with given email already registered
		if strings.Contains(err.Error(), "23505") {
			http.Error(writer, DUPLICATE_EMAIL, http.StatusBadRequest)
			return
		}

		// internal server error = not good xD
		fmt.Println(err)
		http.Error(writer, INTERNAL_SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	// send back JSON response
	sendJSON(writer, UserAuthenticatedResponse{
		Token:      jwt,
		Expiration: expiration,
	})
}

// Login endpoint -> authenticate user to get JWT
func (handler *HTTPUserHandler) Login(writer http.ResponseWriter, request *http.Request) {
	var requestBody UserAuthenticateRequest

	// parse and validate request body
	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil || !requestBody.IsValid() {
		http.Error(writer, INVALID_BODY, http.StatusBadRequest)
		return
	}

	jwt, expiration, err := handler.userService.Authenticate(requestBody.Email, requestBody.Password)

	if err != nil {
		http.Error(writer, UNAUTHORIZED, http.StatusUnauthorized)
		return
	}

	// send back JSON response
	sendJSON(writer, UserAuthenticatedResponse{
		Token:      jwt,
		Expiration: expiration,
	})
}