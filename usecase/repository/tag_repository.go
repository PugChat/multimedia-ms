package repository

import (
	"github.com/pugchat_multimedia/domain/model"
)

type TagRepository interface {
	FindByID(id int) (*model.Tag, error)
	Store(tag *model.Tag) error
	Update(tag *model.Tag) error
	Delete(tag *model.Tag) error
	FindAll() ([]*model.Tag, error)
	FindOrCreateAll(tags []*model.Tag) error
}
