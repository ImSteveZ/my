// ===============================================
// Description  : nicego/meta.go
// Author       : StevE.Z
// Email        : stevzhang01@gmail
// Date         : 2019-04-20 20:51:30
// ================================================
package nicego

import (
	"context"
	"net/http"
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
	if mt := ctx.Value(metaKey{}); mt != nil {
		mtv := mt.(metaVal)
		return mtv.w, mtv.r
	} else {
		return nil, nil
	}
}
