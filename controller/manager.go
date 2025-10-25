package controller

type Manager interface {
	RegisterController(controller Controller) error
	RegisterControllers(controllers ...Controller) error
	GetController(route string) (Controller, error)
}
