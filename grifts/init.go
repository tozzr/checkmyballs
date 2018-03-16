package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/tozzr/checkmyballs/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
