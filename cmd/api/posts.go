package main

import (
	"gopher/internal/store"
	"net/http"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHanlder(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJson(w, r, &payload); err != nil {
		writeJsonError(w, http.StatusBadRequest, err.Error())
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
		writeJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := WriteJson(w, http.StatusCreated, post); err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

}
func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx = r.Context()
	if post, err := app.store.Posts.GetByID(ctx, r.PathValue("postID")); err!=nil{
		
	}
}
