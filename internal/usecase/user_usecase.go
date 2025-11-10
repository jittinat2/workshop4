package usecase

import (
	"strings"
	"workshop4/internal/entity"
	"workshop4/internal/repository"
)

type UserUsecase interface {
	Create(u *entity.User) error
	List() ([]entity.User, error)
	Get(id string) (*entity.User, error)
	Update(u *entity.User) error
	Delete(id string) error
}

type userUsecase struct {
	repo repository.UserRepo
}

func NewUserUsecase(r repository.UserRepo) UserUsecase {
	return &userUsecase{repo: r}
}

func (s *userUsecase) Create(u *entity.User) error {
	if strings.TrimSpace(u.ID) == "" {
		id, err := s.repo.NextID()
		if err != nil {
			return err
		}
		u.ID = id
	}
	return s.repo.Create(u)
}

func (s *userUsecase) List() ([]entity.User, error)        { return s.repo.GetAll() }
func (s *userUsecase) Get(id string) (*entity.User, error) { return s.repo.GetByID(id) }
func (s *userUsecase) Update(u *entity.User) error         { return s.repo.Update(u) }
func (s *userUsecase) Delete(id string) error              { return s.repo.Delete(id) }
