package place

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	Status string `json:"status"`
	Data   []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

func getAPIResponse(pref int) (response, error) {
	queryArea := fmt.Sprintf("%2d", pref)
	url := "https://www.land.mlit.go.jp/webland/api/CitySearch?area=" + queryArea
	resp, err := http.Get(url)
	if err != nil {
		return response{}, err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response{}, err
	}

	jsonBytes := ([]byte)(byteArray)
	data := new(response)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return response{}, err
	}

	return *data, nil
}
