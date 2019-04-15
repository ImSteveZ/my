package main

import (
	"fmt"
	"log"
	"mygo"
	"net/http"
)

func main() {
	myx := mygo.NewMyx()
	myx.ServeFile("/", "./public")
	myx.HandleFunc("/news", News)
	myx.HandleFunc("/usr/:username", Auth, Usr)
	log.Fatal(http.ListenAndServe(":3000", myx))
}

var users = map[string]struct{}{
	"alice": struct{}{},
	"bob":   struct{}{},
}

func Auth(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	if user, ok := ctx.Get("username"); ok {
		if _, ok := users[user.(string)]; ok {
			ctx.Set("login", true)
			return true
		}
	}
	http.NotFound(w, r)
	return false
}

func News(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	w.Write([]byte("welcome to news category"))
	return true
}

func Usr(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	user, _ := ctx.Get("username")
	fmt.Fprintf(w, "Hello, %s\n", user.(string))
	return true
}
