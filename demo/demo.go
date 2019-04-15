package main

import (
	"log"
	"mygo"
	"net/http"
)

func main() {
	myx := mygo.NewMyx()
	myx.ServeFile("/", "./public")
	myx.HandleFunc("/index/", Auth, Index)
	myx.HandleFunc("/usr/", Auth, Usr)
	log.Fatal(http.ListenAndServe(":3000", myx))
}

func Auth(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	ctx.Set("auth", "login")
	return true
}

func Index(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	authval, _ := ctx.Get("auth")
	w.Write([]byte("index_" + authval.(string)))
	return true
}

func Usr(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	authval, _ := ctx.Get("auth")
	w.Write([]byte("usr_" + authval.(string)))
	return true
}
