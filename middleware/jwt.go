package middleware

import (
	"github.com/alvinamartya/go-bukuibu-be/utils"
	"net/http"
	"strconv"
	"strings"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("Authorization")
		if jwtToken == "" {
			utils.HttpResponse(w, map[string]interface{}{
				"error": "Missing Authorization Token",
			}, http.StatusUnauthorized)
		}

		jwtToken = strings.Replace(jwtToken, "Bearer ", "", 1)
		claims, err := utils.VerifyJWTToken(jwtToken)
		if err != nil {
			utils.HttpResponse(w, map[string]interface{}{
				"error": "Invalid Authorization Token",
			}, http.StatusUnauthorized)
			return
		}

		userId := strconv.Itoa(int(claims.Id))
		r.Header.Set("UserId", userId)
		r.Header.Set("Username", claims.Name)
		next.ServeHTTP(w, r)
	})
}
