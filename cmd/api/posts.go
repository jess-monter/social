package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jess-monter/social/internal/store"
)

type CreatePostPayload struct {
	Content string   `json:"content"`
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()

	post := &store.Post{
		Content: payload.Content,
		Title:   payload.Title,
		Tags:    payload.Tags,
		UserID:  1, // TODO: Replace with authenticated user ID
	}

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func parseID(param string) (int64, error) {
	intID, err := strconv.ParseInt(param, 10, 64)
	if err != nil || intID < 1 {
		return 0, err
	}
	return intID, nil
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	postIDParam := chi.URLParam(r, "postID")
	postID, err := parseID(postIDParam)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid post ID")
		return
	}

	ctx := r.Context()
	post, err := app.store.Posts.GetPostByID(ctx, postID)
	if err != nil {
		if err == store.ErrRecordNotFound {
			writeJSONError(w, http.StatusNotFound, "post not found")
		} else {
			writeJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
