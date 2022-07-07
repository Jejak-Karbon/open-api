package repository

import (
	"context"
	"strings"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type News interface {
	Create(ctx context.Context, data model.News) error
	Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.News, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.News, error)
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type news struct {
	Db *gorm.DB
}

func NewNews(db *gorm.DB) *news {
	return &news{
		db,
	}
}

func (n *news) Create(ctx context.Context, data model.News) error {
	return n.Db.WithContext(ctx).Model(&model.News{}).Create(&data).Error
}

func (n *news) Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.News, *dto.PaginationInfo, error) {
	var Newss []model.News
	var count int64

	query := n.Db.WithContext(ctx).Model(&model.News{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(title) LIKE ?  ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&Newss).Error

	return Newss, dto.CheckInfoPagination(paginate, count), err
}

func (n *news) FindByID(ctx context.Context, ID uint) (model.News, error) {

	var data model.News
	err := n.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	return data, err
}

func (n *news) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := n.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.News{}).Updates(data).Error
	return err
}

func (n *news) Delete(ctx context.Context, ID uint) error {

	err := n.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.News{}).Error
	return err
}
