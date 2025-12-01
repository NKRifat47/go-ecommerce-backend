package user

import "ecommerce/domain"

type service struct {
	usRepo UserRepo
}

func NewService(usRepo UserRepo) Service {
	return &service{
		usRepo: usRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.usRepo.Create(user)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}
func (svc *service) Find(email string, pass string) (*domain.User, error) {
	usr, err := svc.usRepo.Find(email, pass)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil

}
