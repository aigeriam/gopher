package main

import (
	"log"
	"net/http"
)

func (app *application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error %s path, %s error %s", r.Method, r.URL.Path, err)
	writeJsonError(w, http.StatusInternalServerError, "the server encountered a problem ")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad request error: %s path %s error %s", r.Method, r.URL.Path, err)
	writeJsonError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path %s error %s", r.Method, r.URL.Path, err)
	writeJsonError(w, http.StatusNotFound, err.Error())
}
