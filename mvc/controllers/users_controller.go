package controllers

import (
	"log"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		return
	}
	userId, err :=
}
