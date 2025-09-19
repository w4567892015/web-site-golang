package index

type IndexService interface {
	GetHello() string
}

type indexService struct{}

func NewIndexService() IndexService {
	return &indexService{}
}

func (is indexService) GetHello() string {
	return "Hello, World!"
}
