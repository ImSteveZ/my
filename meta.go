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

// metaKey
type metaKey struct{}

// metaVal
type metaVal struct {
	w http.ResponseWriter
	r *http.Request
}

// GetMeta
func GetMeta(ctx context.Context) (http.ResponseWriter, *http.Request) {
	if meta := ctx.Value(metaKey{}).(metaVal); meta != nil {
		return meta.w, meta.r
	} else {
		return nil, nil
	}
}