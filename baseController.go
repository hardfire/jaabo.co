package main

import "net/http"

// Credits : http://codegangsta.gitbooks.io/building-web-apps-with-go/content/controllers/README.html

// Action defines a standard function signature for us to use when creating
// controller actions. A controller action is basically just a method attached to
// a controller.
type Action func(rw http.ResponseWriter, r *http.Request) error

// This is our Base Controller
type BaseController struct{}

// The action function helps with error handling in a controller
func (c *BaseController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}
