package controller

import (
	"github.com/apaarshrm39/k8s-Operator-sdk-prac/pkg/controller/presentation"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, presentation.Add)
}
