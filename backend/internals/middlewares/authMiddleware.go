package middlewares

import (
	"backend/internals/utils"
	"context"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Cookies : ", r.Cookies())
		token, err := r.Cookie("token")
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := utils.VerifyToken(token.Value)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "user", *user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
