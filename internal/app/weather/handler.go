package weather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type Wilayah struct {
	Id        string
	Propinsi  string
	Kota      string
	Kecamatan string
	Lat       string
	Lon       string
}

type Cuaca struct {
	JamCuaca  string
	KodeCuaca string
	Cuaca     string
	Humidity  string
	TempC     string
	TempF     string
}

func Get(c echo.Context) error {

	url := "https://ibnux.github.io/BMKG-importer/cuaca/wilayah.json"
	resp, err := http.Get(url)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	var dataWilayah []Wilayah
	datas := []byte(body)

	_err := json.Unmarshal(datas, &dataWilayah)
	if _err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(dataWilayah).Send(c)

}

func GetByID(c echo.Context) error {

	rootUrl := "https://ibnux.github.io/BMKG-importer/cuaca/"
	ext := ".json"
	url := rootUrl + c.Param("id") + ext
	resp, err := http.Get(url)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	var dataCuaca []Cuaca
	datas := []byte(body)

	_err := json.Unmarshal(datas, &dataCuaca)
	if _err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(dataCuaca).Send(c)
}
