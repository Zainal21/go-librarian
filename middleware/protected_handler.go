package middleware

import (
	"go-book-management/utils"
	"log"
	"net/http"
)

func ProtectedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			utils.JsonResponse(w, nil, "Unauthorized", http.StatusUnauthorized)
			return
		}

		username, err := utils.VerifyToken(token)
		if err != nil {
			utils.JsonResponse(w, nil, "Unauthorized", http.StatusUnauthorized)
			return
		}

		log.Println(username)

		next.ServeHTTP(w, r)
	})
}
