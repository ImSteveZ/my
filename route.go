// ===============================================
// Description  : nicego/route.go
// Author       : StevE.Z
// Email        : stevzhang01@gmail.com
// Date         : 2019-04-22 10:30:30
// ================================================
package nicego

import (
	"context"
	"net/http"
)

// route
type Route struct {
	ctx context.Context
	mux *http.ServeMux
}

// NewRoute
func NewRoute(ctx context.Context) *Route {
	return &Route{ctx: ctx, mux: http.NewServeMux()}
}

// From
func (rt *Route) From(pattern string) *Router {
	return &Router{route: rt, pattern: pattern}
}

// ServeHTTP
func (rt *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.mux.ServeHTTP(w, r)
}
