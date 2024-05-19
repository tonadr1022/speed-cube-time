package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling JWT auth middleware")
		tokenString := r.Header.Get("X-Session-Token")
		token, err := validateJWT(tokenString)
		switch {
		case token.Valid:
			// get the id and username from the token and attach to the context
			claims := token.Claims.(jwt.MapClaims)
			userID := claims["id"].(string)
			username := claims["username"].(string)
			ctx := withUser(r.Context(), userID, username)
			r = r.WithContext(ctx)
			// allow next
			next.ServeHTTP(w, r)
		case errors.Is(err, jwt.ErrTokenMalformed):
			util.WriteJson(w, http.StatusForbidden, "malformed token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			util.WriteJson(w, http.StatusForbidden, "invalid signature")
		case errors.Is(err, jwt.ErrTokenExpired):
			util.WriteJson(w, http.StatusForbidden, "token expired")
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			util.WriteJson(w, http.StatusForbidden, "token not valid yet")
		default:
			util.WriteJson(w, http.StatusForbidden, "could not handle this token")
		}
	})
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
