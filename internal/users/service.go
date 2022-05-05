package users

type Service interface {
	GetAll() ([]User, error)
	Store(nombre, apellido, email, fechaDeCreacion string, edad int, altura float64, activo *bool) (User, error)
	Update(id int, nombre, apellido, email, fechaDeCreacion string, edad int, altura float64, activo *bool) (User, error)
	UpdateNameAndSurname(id int, nombre string, apellido string) (User, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func (s *service) Update(id int, nombre string, apellido string, email string, fechaDeCreacion string, edad int, altura float64, activo *bool) (User, error) {
	return s.repository.Update(id, nombre, apellido, email, fechaDeCreacion, edad, altura, activo)
}

func (s *service) UpdateNameAndSurname(id int, nombre string, apellido string) (User, error) {
	return s.repository.UpdateNameAndSurname(id, nombre, apellido)
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Store(nombre, apellido, email, fechaDeCreacion string, edad int, altura float64, activo *bool) (User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return User{}, err
	}

	lastID++

	user, err := s.repository.Store(lastID, nombre, apellido, email, fechaDeCreacion, edad, altura, activo)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
