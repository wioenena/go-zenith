package zenith

import "github.com/wioenena/go-zenith/controller"

type App interface {
	// Run starts the web server on the specified address (e.g., ":8080")
	// using standard HTTP. This is a blocking operation.
	// It returns an error if the server fails to start (e.g., "port already in use").
	Run(addr string) error
	// RunTLS starts the web server on the specified address using
	// HTTPS (TLS). It requires paths to the SSL certificate (certFile)
	// and key (keyFile) files. This is also a blocking operation.
	RunTLS(addr, certFile, keyFile string) error
	GetControllerManager() controller.Manager
}
