package api

import (
	"encoding/json"
	"fmt"
	entity "go-mobile/internal/entitiy"
	sl "go-mobile/package/logger/slog"
	"io"
	"log/slog"
	"net/http"
)

type userApi struct {
	apiUrl string
	log    *slog.Logger
	client *http.Client
}

func NewUserApi(apiUrl string, log *slog.Logger) *userApi {
	return &userApi{
		apiUrl: apiUrl,
		log:    log,
		client: &http.Client{},
	}
}

func (ua *userApi) GetPeopleInfo(passportSeries, passportNumber string) (*entity.PeopleApiResponse, error) {
	url := fmt.Sprintf("%s/info?passportSerie=%s&passportNumber=%s", ua.apiUrl, passportSeries, passportNumber)
	res, err := ua.client.Get(url)
	if err != nil {
		ua.log.Error("PEOPLE API CLIENT - GetPeopleInfo", sl.Err(err))
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			ua.log.Error("PEOPLE API CLIENT - GetPeopleInfo", sl.Err(err))
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		ua.log.Error("PEOPLE API CLIENT - GetPeopleInfo", sl.Err(err))
		return nil, fmt.Errorf("error with status: %s", res.Status)
	}

	var people entity.PeopleApiResponse
	if err := json.NewDecoder(res.Body).Decode(&people); err != nil {
		ua.log.Error("PEOPLE API CLIENT - GetPeopleInfo", sl.Err(err))
		return nil, err
	}

	return nil, nil
}
