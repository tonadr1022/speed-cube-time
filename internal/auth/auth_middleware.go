package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

func WithJWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authString := r.Header.Get("Authorization")

		// attempt to parse the token
		if authString == "" {
			util.WriteApiError(w, http.StatusUnauthorized, "unauthenticated")
			return
		}
		strings := strings.Split(authString, "Bearer ")
		if len(strings) != 2 {
			util.WriteApiError(w, http.StatusUnauthorized, "malformed token")
		}
		tokenString := strings[1]
		token, err := validateJWT(tokenString)

		// handle the token result
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			util.WriteApiError(w, http.StatusUnauthorized, "malformed token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			util.WriteApiError(w, http.StatusUnauthorized, "invalid signature")
		case errors.Is(err, jwt.ErrTokenExpired):
			util.WriteApiError(w, http.StatusUnauthorized, "token expired")
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			util.WriteApiError(w, http.StatusUnauthorized, "token not valid yet")
		case token.Valid:
			// get the id and username from the token and attach to the context
			claims := token.Claims.(jwt.MapClaims)
			userID := claims["id"].(string)
			username := claims["username"].(string)
			if userID == "" || username == "" {
				util.WriteApiError(w, http.StatusUnauthorized, "claims not found in token")
				return
			}
			// attach context and call the next middleware or handler
			ctx := withUser(r.Context(), userID, username)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		default:
			// should not occur generally
			util.WriteApiError(w, http.StatusUnauthorized, "could not handle this token")
		}
	}
}

type contextKey int

const userKey contextKey = iota

// returns a context that contains the user identity
func withUser(ctx context.Context, id string, username string) context.Context {
	return context.WithValue(ctx, userKey, entity.User{ID: id, Username: username})
}

// returns the user identity from the given context, otherwise nil if not found
func CurrentUser(ctx context.Context) Identity {
	if user, ok := ctx.Value(userKey).(entity.User); ok {
		return user
	}
	return nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// typically shouldn't happen
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}
