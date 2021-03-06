package routes

import (
	"net/http"

	controller "api/controller"

	"github.com/gorilla/mux"
)

// Route route
type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("POST", "/api/addBug", controller.AddBug, nil)
	register("GET", "/api/getBugList", controller.GetBugList, nil)
	register("POST", "/api/updateBug", controller.UpdateBug, nil)
	register("DELETE", "/api/deleteBug", controller.DeleteBug, nil)
	register("GET", "/api/readImage", controller.ReadImage, nil)
}

// NewRouter NewRouter
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
