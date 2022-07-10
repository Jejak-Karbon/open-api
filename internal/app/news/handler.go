package news

import (
	"context"
	"os"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	_ "net/http"
	"fmt"
	"strconv"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
	aws_util "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/aws"
	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/str"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {

	payload := new(dto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Find(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get news success", &result.PaginationInfo).Send(c)
}

func (h *handler) GetByID(c echo.Context) error {

	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.FindByID(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Create(c echo.Context) error {

	payload := new(dto.CreateNewsRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	upload,_ := c.FormFile("image")

	src, err := upload.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(upload.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// upload image if exist
	var uploader *s3manager.Uploader
	uploader = aws_util.NewUploader()

	log.Println("uploading...")
	file, err := ioutil.ReadFile(upload.Filename)
	if err != nil {
		log.Fatal(err)
	}

	file_destination := str.GenerateRandString(10)+upload.Filename

	upInput := &s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")), // bucket's name
		Key:         aws.String(file_destination),        // files destination location
		Body:        bytes.NewReader(file),               // content of the file
		ContentType: aws.String(upload.Header["Content-Type"][0]),              // content type
	}

	resp, err := uploader.UploadWithContext(context.Background(), upInput)
	log.Printf("res %+v\n", resp)
	log.Printf("err %+v\n", err)

	img := resp.Location

	result, err := h.service.Create(c.Request().Context(),img, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)

}

func (h *handler) Update(c echo.Context) error {
	payload := new(dto.UpdateNewsRequest)
	if err := c.Bind(payload); err != nil {
		fmt.Println("bind", err.Error())
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		fmt.Println("validate", err.Error())
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	strID := c.Param("id")
	ID, err := strconv.Atoi(strID)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	result, err := h.service.Update(c.Request().Context(), uint(ID), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Delete(c echo.Context) error {

	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.Delete(c.Request().Context(), payload.ID)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
