// ===============================================
// Description  : nicego/route.go
// Author       : StevE.Z
// Email        : stevzhang01@gmail
// Date         : 2019-04-20 20:51:30
// ================================================
package nicego

import (
	"context"
	"net/http"
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

// From
func (rt *route) From(pattern string) *router {
	return &router{route: rt, pattern: pattern}
}

// ServeHTTP
func (rt *route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.mux.ServeHTTP(w, r)
}
