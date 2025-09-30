package service

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ProcessData(data string) (string, error) {
	return data, nil
}
