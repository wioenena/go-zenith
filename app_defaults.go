package zenith

import "net/http"

type defaultApp struct{}

func (d *defaultApp) Run(addr string) error {
	return http.ListenAndServe(addr, d)
}

func (d *defaultApp) RunTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, d)
}

func (d *defaultApp) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte("Hello, Zenith!")); err != nil {
		panic(err)
	}
}

func DefaultApp() App {
	return &defaultApp{}
}
