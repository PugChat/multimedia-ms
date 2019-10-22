package persistance

import (
	"github.com/jinzhu/gorm"
	"github.com/pugchat_multimedia/domain/model"
	"github.com/pugchat_multimedia/usecase/repository"
)

type entryRepository struct {
	db *gorm.DB
}

// NewEntryRepository creates a repository using MySQL for datasotre.
func NewEntryRepository(db *gorm.DB) repository.EntryRepository {
	return &entryRepository{db}
}

func (eR *entryRepository) FindByID(id int) (*model.Archivo, error) {
	archivo := model.Archivo{ID: id}
	err := eR.db.First(&archivo).Error
	if err != nil {
		return nil, err
	}

	return &archivo, nil
}

func (eR *entryRepository) Store(archivo *model.Archivo) error {
	return eR.db.Save(archivo).Error
}

func (eR *entryRepository) Update(archivo *model.Archivo) error {
	// Save will include all fields when perform the Updating SQL, even it is not changed
	return eR.db.Model(&model.Archivo{ID: archivo.ID}).Updates(archivo).Error
}

func (eR *entryRepository) Delete(archivo *model.Archivo) error {
	return eR.db.Delete(archivo).Error
}

func (eR *entryRepository) FindAll() ([]*model.Archivo, error) {
	archivos := []*model.Archivo{}

	err := eR.db.Find(&archivos).Error
	if err != nil {
		return nil, err
	}

	return archivos, nil
}
