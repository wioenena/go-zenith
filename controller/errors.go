package controller

import (
	"errors"
	"fmt"
)

var ErrControllerRouteEmpty = errors.New("controller route cannot be empty")
var ErrControllerHandlersEmpty = errors.New("controller must have at least one handler")

type ErrControllerAlreadyRegistered struct{ Route string }

func (e ErrControllerAlreadyRegistered) Error() string {
	return fmt.Sprintf("controller route %s already registered", e.Route)
}

type ErrControllerNotFound struct{ Route string }

func (e ErrControllerNotFound) Error() string {
	return fmt.Sprintf("controller route %s not found", e.Route)
}
