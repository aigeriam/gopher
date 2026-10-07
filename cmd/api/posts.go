package main

import (
	"errors"
	"gopher/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHanlder(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJson(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//todo change after auth
	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserId:  1,
		Tags:    payload.Tags,
	}
	if post.Tags == nil {
		post.Tags = []string{}
	}
	ctx := r.Context()
	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
	if err := WriteJson(w, http.StatusCreated, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

}
func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idparam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idparam, 10, 64)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}
	post, err := app.store.Posts.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundError(w, r, err)
		default:
			app.InternalServerError(w, r, err)
		}
		return
	}
	if err := WriteJson(w, http.StatusAccepted, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
