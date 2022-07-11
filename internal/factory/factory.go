package factory

import (
	"gorm.io/gorm"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
)

type Factory struct {
	NewsRepository repository.News
}

func NewFactory(db *gorm.DB) *Factory {
	return &Factory{
		NewsRepository: repository.NewNews(db),
	}
}
