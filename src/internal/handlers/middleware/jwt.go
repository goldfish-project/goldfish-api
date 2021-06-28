package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"goldfish-api/internal/handlers"
	"net/http"
	"time"
)

const (
	AUTHORIZED = "authorized"
	USER_ID    = "user_id"
	EXPIRATION = "exp"
)

// AuthenticatedEndpoint defines the function header for an endpoint using the JWT Middleware
type AuthenticatedEndpoint = func(writer http.ResponseWriter, request *http.Request, userId uuid.UUID)

type JWTMiddleware struct {
	secretKey           string
	headerField string
	validityPeriodInMin time.Duration
}

// NewJWTMiddleware creates a new JWT handler instance with a given secret key
func NewJWTMiddleware(secretKey, headerField string, validityPeriodInMin int) JWTMiddleware {
	return JWTMiddleware{secretKey: secretKey, validityPeriodInMin: time.Duration(validityPeriodInMin), headerField: headerField}
}

// WithAuthentication represents the actual middleware function which needs to be applied on an endpoint
func (mw *JWTMiddleware) WithAuthentication(next AuthenticatedEndpoint) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// check if header is even present
		if request.Header[mw.headerField] == nil {
			http.Error(writer, handlers.UNAUTHORIZED, http.StatusUnauthorized)
			return
		}

		// validate token
		token, err := jwt.Parse(request.Header[mw.headerField][0],
			func(token *jwt.Token) (interface{}, error) {
				//Make sure that the token method conform to "SigningMethodHMAC"
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(mw.secretKey), nil
		})

		if err != nil {
			http.Error(writer, handlers.UNAUTHORIZED, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		// check if claims are valid
		if !ok || !token.Valid {
			http.Error(writer, handlers.UNAUTHORIZED, http.StatusUnauthorized)
			return
		}

		userId, err := extractUserId(claims)

		if err != nil {
			http.Error(writer, handlers.UNAUTHORIZED, http.StatusUnauthorized)
			return
		}

		// check if token is already expired
		if isTokenExpired(claims) {
			http.Error(writer, handlers.UNAUTHORIZED, http.StatusUnauthorized)
			return
		}

		// call next
		next(writer, request, userId)
	})
}

// GetToken returns an encoded and signed JWT
func (mw *JWTMiddleware) GetToken(userId string) (string, time.Time, error) {
	// create the jwt claims
	expiration := time.Now().Add(time.Minute * mw.validityPeriodInMin)

	claims := jwt.MapClaims{}
	claims[AUTHORIZED] = true
	claims[USER_ID] = userId
	claims[EXPIRATION] = expiration.Unix()

	// sign the claims & generate the JWT
	signed := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := signed.SignedString([]byte(mw.secretKey))

	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiration, nil
}

// extractUserId extracts the user id from the JWT claims and parses it into a UUID
func extractUserId(claims jwt.MapClaims) (uuid.UUID, error) {
	// extract user id
	userIdRaw, ok := claims[USER_ID].(string)

	if !ok {
		return uuid.UUID{}, errors.New("Invalid user id claims")
	}

	// parse user id tu uuid
	userId, err := uuid.Parse(userIdRaw)

	if !ok {
		return uuid.UUID{}, err
	}

	return userId, nil
}

// extractUserId extracts the user id from the JWT claims and parses it into a UUID
func isTokenExpired(claims jwt.MapClaims) bool {
	// extract user id
	expirationRaw, ok := claims[EXPIRATION].(int64)

	if !ok {
		return false
	}

	// parse to time
	expiration := time.Unix(expirationRaw, 0)

	// check is token is epired
	return expiration.Before(time.Now())
}