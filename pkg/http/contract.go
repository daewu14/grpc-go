package httppkg

import "net/http"

type Http interface {
	Method() string
	Handle(w http.ResponseWriter, r *http.Request)
}
