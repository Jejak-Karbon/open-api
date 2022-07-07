package factory

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/database"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
)

type Factory struct {
	NewsRepository repository.News
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		NewsRepository: repository.NewNews(db),
	}
}
