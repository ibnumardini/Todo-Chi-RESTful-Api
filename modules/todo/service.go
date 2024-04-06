package todo

type Service interface {
	FindAll() ([]TodoResponse, error)
}

type service struct {
	repo *repo
}

func newService(repo *repo) service {
	return service{repo}
}

func (s *service) FindAll() ([]TodoResponse, error) {
	todos, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := []TodoResponse{}

	for _, todo := range todos {
		newResponse := TodoResponse{
			Id:        todo.Id,
			Task:      todo.Task,
			Note:      todo.Note.String,
			DoneAt:    todo.DoneAt.Time,
			CreatedAt: todo.CreatedAt,
		}

		response = append(response, newResponse)
	}

	return response, nil
}
