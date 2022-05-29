package fakejson

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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

func getFakeJSONObjectById(id int) (*FakeJson, error) {
	response, err := http.Get(base_url + "/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject *FakeJson
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}

func deleteTodoItem(id int) error {
	req, err := http.NewRequest(http.MethodDelete, base_url+"/"+strconv.Itoa(id), nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
