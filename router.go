// ===============================================
// Description  : routeR.go
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

// router
type router struct {
	route *route
	pattern string
	middlewares []func(context.Context, func(context.Context))
}

// Use
func (rtr *router) Use(middlewares ...func(context.Context, func(context.Context))) {
	rtr.middlewares = append(rtr.middlewares, middlewares...)
	return rtr
}

// Do
func (rtr *router) Do(controller func(conterxt.Context)) {
	rtr.route.mux.HandleFunc(rtr.pattern, func(w http.ResponseWriter, r *http.Request) {
		var (
			next func(context.Context)
			i int
		)
		next = func(ctx context.Context) {
			if i < len(rtr.middlewares) {
				i++
				middlewares[i-1](ctx, next)
			} else {
				if controller != nil {
					controller(ctx)
				}
			}
		}
		ctx := context.WithValue(rtr.route.ctx, metaKey{}, metaVal{w: w, r: r})
		next(ctx)
	})
}

