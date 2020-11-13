package service

import (
	"encoding/json"
	"fmt"
	"github.com/shapovalex/depAnalyzer/model"
	"io/ioutil"
	"net/http"
)

func ResolvePypiInfoForPackage(name string) model.PypiInfo {
	baseAddress := "https://pypi.python.org/pypi/%s/json"
	address := fmt.Sprintf(baseAddress, name)
	return resolvePypiInfo(address)
}

func resolvePypiInfo(address string) model.PypiInfo {
	response, err := http.Get(address)
	if err != nil {
		fmt.Println("Unable to get license, error during http request")
		return getErrorPyPiInfo()
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Unable to get license, error during reading data")
		return getErrorPyPiInfo()
	}

	var pypiResponse model.PypiResponse
	err = json.Unmarshal(body, &pypiResponse)
	if err != nil {
		fmt.Println("Unable to get license for package", address, ". Error during parsing json", err.Error())
		return getErrorPyPiInfo()
	}

	return pypiResponse.Info
}

func getErrorPyPiInfo() model.PypiInfo {
	info := new(model.PypiInfo)
	info.License = "Unable to resolve version"
	return *info
}
