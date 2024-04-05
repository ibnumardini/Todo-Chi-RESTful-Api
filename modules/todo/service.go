package todo

type Service interface {
	FindAll() ([]Todo, error)
}

type service struct {
	repo *repo
}

func newService(repo *repo) service {
	return service{repo}
}

func (s *service) FindAll() ([]Todo, error) {
	return s.repo.FindAll()
}
