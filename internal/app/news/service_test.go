package news

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/mock"
	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {

	// setup database
	db, mock := mock.DBConnection()

	news := &model.News{
			Title: "Title",
			Description: "Description",
			Image: "Image",
	}

	rows := sqlmock.NewRows([]string{"title", "description", "image"}).
		AddRow(news.Title, news.Description, news.Image)

	query := "SELECT * FROM `news` WHERE id = ? AND `news`.`deleted_at` IS NULL ORDER BY `news`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)
	
	f := factory.NewFactory(db)
	service := NewService(f)

	payload := &dto.ByIDRequest{
		ID : 1,
	}

	user, err := service.FindByID(context.Background(), payload)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, user)
	assert.NoError(t, err)

}