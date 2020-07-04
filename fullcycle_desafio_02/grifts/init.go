package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/viniciusbsneto/fullcycle_desafio_02/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
