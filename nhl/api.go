package nhl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseStatsUrl = "https://statsapi.web.nhl.com/api/v1/"

func GetResponse(path string, model interface{}) (interface{}, error) {
	response, err := http.Get(baseStatsUrl + path)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, model); err != nil {
		return nil, err
	}

	return model, nil
}

func ParseResponse(path string, model interface{}) error {
	response, err := http.Get(baseStatsUrl + path)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &model); err != nil {
		return err
	}

	return nil
}
