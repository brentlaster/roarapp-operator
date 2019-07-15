package controller

import (
	"github.com/redhat/roarapp-operator/pkg/controller/roarapp"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, roarapp.Add)
}
