package servers

type serviceConfigInterface interface {
	GetServerPort() string
	GetServerAddress() string
}
