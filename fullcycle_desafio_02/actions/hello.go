package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HelloHandler is a handler to serve up
// the /hello endpoint a hello page
func HelloHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("hello.html"))
}
