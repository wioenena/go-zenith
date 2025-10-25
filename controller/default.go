package controller

type controllerCache map[string]Controller

type defaultManager struct {
	controllers controllerCache
}

func (d *defaultManager) RegisterController(controller Controller) error {
	route := controller.GetRoute()
	if route == "" {
		return ErrControllerRouteEmpty
	}
	if _, exists := d.controllers[route]; exists {
		return ErrControllerAlreadyRegistered{route}
	}

	controller.Setup()
	d.controllers[route] = controller
	return nil
}

func (d *defaultManager) RegisterControllers(controllers ...Controller) error {
	for _, controller := range controllers {
		if err := d.RegisterController(controller); err != nil {
			return err
		}
	}

	return nil
}

func (d *defaultManager) GetController(route string) (Controller, error) {
	if controller, exists := d.controllers[route]; exists {
		return controller, nil
	}

	return nil, ErrControllerNotFound{route}
}

func NewDefaultManager() Manager {
	m := new(defaultManager)
	m.controllers = make(controllerCache)
	return m
}
