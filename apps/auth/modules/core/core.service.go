package core

type CoreService interface {
	AuthorizeEndpoint() string
}

type coreService struct{}

func NewCoreService() CoreService {
	return &coreService{}
}

func (cs coreService) AuthorizeEndpoint() string {
	return "Hello, World!"
}
