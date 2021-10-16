package group

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Group {
	return allEntities
}

func (s *Service) Get(idx int) (*Group, error) {
	return &allEntities[idx], nil
}
