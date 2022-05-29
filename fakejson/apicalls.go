package fakejson

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type FakeJson struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getFakeJSONObjects() ([]*FakeJson, error) {
	response, err := http.Get(base_url)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject []*FakeJson
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}
