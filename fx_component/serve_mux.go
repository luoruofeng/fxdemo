package fx_component

import "net/http"

// NewServeMux builds a ServeMux that will route requests
// to the given Route.
func NewServeMux(route Route) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(route.Pattern(), route)
	return mux
}
