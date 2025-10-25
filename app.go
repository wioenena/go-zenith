package zenith

type App interface {
	Run(addr string) error
	RunTLS(addr, certFile, keyFile string) error
}
