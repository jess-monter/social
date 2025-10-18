package main

import (
	"net/http"

	"github.com/jess-monter/social/internal/store"
)

// GetUserFeed godoc
//
//	@Summary		Get User Feed
//	@Description	Retrieve the feed for the authenticated user.
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int		false	"Number of posts to return"		default(20)		minimum(1)	maximum(100)
//	@Param			offset	query		int		false	"Number of posts to skip"		default(0)		minimum(0)
//	@Param			sort	query		string	false	"Sort order: 'asc' or 'desc'"	default(desc)	enum(asc, desc)
//	@Param			search	query		string	false	"Search term for post content or title"
//	@Param			tags	query		array	false	"Filter by tags"
//	@Success		200		{array}		store.FeedPost
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error	"Resource Not Found"
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/{id}/feed [get]
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	// user := getUserFromCtx(r)

	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(4), fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
