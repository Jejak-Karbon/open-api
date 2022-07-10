package news

import (
	"os"
	"bytes"
	"io/ioutil"
	"log"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/constant"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
	aws_util "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/aws"
	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/str"
)

type service struct {
	NewsRepository repository.News
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.News], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.News, error)
	Create(ctx context.Context, payload *dto.CreateNewsRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateNewsRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.News, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		NewsRepository: f.NewsRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.News], error) {

	News, info, err := s.NewsRepository.Find(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.News])
	result.Datas = News
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.News, error) {

	data, err := s.NewsRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateNewsRequest) (string, error) {

	// upload image if exist
	var uploader *s3manager.Uploader
	uploader = aws_util.NewUploader()
	log.Printf("up %+v\n", uploader)

	log.Println("uploading...")
	file, err := ioutil.ReadFile("Soal.pdf")
	if err != nil {
		log.Fatal(err)
	}

	file_destination := "avatar/"+str.GenerateRandString(10)+".pdf"

	upInput := &s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")), // bucket's name
		Key:         aws.String(file_destination),        // files destination location
		Body:        bytes.NewReader(file),               // content of the file
		ContentType: aws.String("file/pdf"),              // content type
	}
	resp, err := uploader.UploadWithContext(context.Background(), upInput)
	log.Printf("res %+v\n", resp)
	log.Printf("err %+v\n", err)

	var News = model.News{
		Title:       payload.Title,
		Image:       payload.Image,
		Description: payload.Description,
		IsActive:    payload.IsActive,
	}

	err2 := s.NewsRepository.Create(ctx, News)
	if err2 != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err2)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateNewsRequest) (string, error) {

	var data = make(map[string]interface{})

	if payload.Title != nil {
		data["title"] = payload.Title
	}
	if payload.Image != nil {
		data["image"] = payload.Image
	}
	if payload.Description != nil {
		data["description"] = payload.Description
	}
	if payload.IsActive != 0 {
		data["is_active"] = payload.IsActive
	}

	err := s.NewsRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.News, error) {

	data, err := s.NewsRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.NewsRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil

}
