package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Make_Get_Request(url string, headers map[string]string) ([]byte, error) {

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	//set the header(s) from header map to request variable
	for headerName, headerValue := range headers {
		fmt.Println("headerName", headerName, "headerValue", headerValue)
		request.Header.Set(headerName, headerValue)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		return bodyBytes, fmt.Errorf("Invalid response code: %d", response.StatusCode)
	}

	if err != nil {
		return nil, err
	}

	return bodyBytes, nil

}
