package main

import (
	"encoding/json"
	"fmt"
	"github/sudsHere/rssagg/internal/database"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed_follows, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't create feed follows: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowstoFeedFollows(feed_follows))
}

func (apiCfg apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	user_feed, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't find feed for given user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsFollowstoFeedsFollows(user_feed))
}

func (apiCgf apiConfig) handleDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't parse feedFollow ID: %v", err))
		return
	}

	err = apiCgf.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID: feedFollowId,
		UserID: user.ID,
	})

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't delete feed follows: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}