package controller

import (
	"github.com/example-inc/minizenko-operator/pkg/controller/minizenko"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, minizenko.Add)
}
