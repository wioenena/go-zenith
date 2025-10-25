package zenith

import (
	"net/http"

	"github.com/wioenena/go-zenith/controller"
	zenithHTTP "github.com/wioenena/go-zenith/http"
)

type defaultApp struct {
	controllerManager controller.Manager
}

func (d *defaultApp) Run(addr string) error {
	return http.ListenAndServe(addr, d)
}

func (d *defaultApp) RunTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, d)
}

func (d *defaultApp) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	firstPathSegment := zenithHTTP.GetFirstPathSegment(request.URL.Path)
	cont, err := d.controllerManager.GetController(firstPathSegment)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	trimmedPath := request.URL.Path[len(firstPathSegment):]
	handlers, err := cont.GetHandlers(zenithHTTP.HttpMethod(request.Method), trimmedPath)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	for _, handler := range handlers {
		if err := handler(nil); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (d *defaultApp) GetControllerManager() controller.Manager {
	return d.controllerManager
}

func NewDefaultApp() App {
	app := new(defaultApp)
	app.controllerManager = controller.NewDefaultManager()
	return app
}
