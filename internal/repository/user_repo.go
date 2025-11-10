package repository

import (
	"fmt"
	"strconv"
	"strings"

	"workshop4/internal/db"
	"workshop4/internal/entity"

	"gorm.io/gorm"
)

const idPrefix = "LBK"

type UserRepo interface {
	Create(u *entity.User) error
	GetAll() ([]entity.User, error)
	GetByID(id string) (*entity.User, error)
	Update(u *entity.User) error
	Delete(id string) error
	NextID() (string, error)
}

type sqliteUserRepo struct {
	db *gorm.DB
}

func NewSQLiteUserRepo() UserRepo {
	return &sqliteUserRepo{db: db.DB}
}

func (r *sqliteUserRepo) Create(u *entity.User) error {
	return r.db.Create(u).Error
}

func (r *sqliteUserRepo) GetAll() ([]entity.User, error) {
	var users []entity.User
	return users, r.db.Find(&users).Error
}

func (r *sqliteUserRepo) GetByID(id string) (*entity.User, error) {
	var u entity.User
	if err := r.db.First(&u, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *sqliteUserRepo) Update(u *entity.User) error {
	return r.db.Save(u).Error
}

func (r *sqliteUserRepo) Delete(id string) error {
	return r.db.Delete(&entity.User{}, "id = ?", id).Error
}

func (r *sqliteUserRepo) NextID() (string, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return "", err
	}
	max := 0
	for _, u := range users {
		if strings.HasPrefix(u.ID, idPrefix) && len(u.ID) > len(idPrefix) {
			nStr := u.ID[len(idPrefix):]
			if n, err := strconv.Atoi(nStr); err == nil {
				if n > max {
					max = n
				}
			}
		}
	}
	return fmt.Sprintf("%s%06d", idPrefix, max+1), nil
}
