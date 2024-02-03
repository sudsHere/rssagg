package main

import (
	"github/sudsHere/rssagg/internal/database"
	"net/http"
	"github/sudsHere/rssagg/internal/auth"
	"fmt"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithErr(w, 403, fmt.Sprintf("Auth err: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserFromApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithErr(w, 403, fmt.Sprintf("Couldn't find user: %v", err))
			return
		}

		handler(w, r, user)
	}
}