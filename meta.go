// ===============================================
// Description  : nicego/meta.go
// Author       : StevE.Z
// Email        : stevzhang01@gmail.com
// Date         : 2019-04-22 10:33:10
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
