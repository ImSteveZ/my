// ===============================================
// Description  : route.go
// Author       : StevE.Z
// Email        : stevzhang01@gmail
// Date         : 2019-04-20 20:51:30
// LastEditTime : 2019-04-20 22:34:48
// ================================================
package nicego

import (
	"net/http"
	"context"
)

// route
type route struct {
	ctx context.Context
	mux *http.ServeMux
}

// NewRoute
func NewRoute(ctx context.Context) *route {
	return &route{ctx: ctx, mux: http.NewServeMux()}
}

// Static
func (rt *route) Static(pattern, dir string) {
	fileServer := http.FileServer(http.Dir(dir))
	rt.mux.Handle(pattern, http.StripPrefix(pattern, fileServer))
}

// From
func (rt *route) From(pattern string) *router {
	return &router{route: rt, pattern: pattern}
}

// ServeHTTP
func (rt *route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.mux.ServeHTTP(w, r)
}