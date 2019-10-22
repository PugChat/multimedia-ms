package repository

import (
	"github.com/pugchat_multimedia/domain/model"
)

type EntryRepository interface {
	FindByID(id int) (*model.Archivo, error)
	Store(entry *model.Archivo) error
	Update(entry *model.Archivo) error
	Delete(entry *model.Archivo) error
	FindAll() ([]*model.Archivo, error)
}
