package license

import (
	"encoding/json"
	"fmt"
	"github.com/shapovalex/depAnalyzer/helper"
	"github.com/shapovalex/depAnalyzer/processor"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type PyPiLicense struct {
}

func (p PyPiLicense) Process(params processor.Params) {
	lines, errR := helper.ReadLines(params.InputFiles)
	if errR != nil {
		fmt.Println("Unable to read input files")
		os.Exit(1)
	}
	var result []string
	for _, line := range lines {
		dependencyParts := strings.Split(line, "==")
		license := resolvePackageLicense(dependencyParts[0], dependencyParts[1])
		result = append(result, line+" => "+license)
	}
	errW := helper.WriteLines(result, params.OutputFiles)
	if errW != nil {
		fmt.Println("Unable to write to output files")
		os.Exit(1)
	}
}

func (p PyPiLicense) GetSupportedDependencyManager() string {
	return "pip"
}

func (p PyPiLicense) GetSupportedOperation() string {
	return "license"
}

type pypiInfo struct {
	License string
}

type pypiResponse struct {
	Info pypiInfo
}

func resolvePackageLicense(name string, version string) string {
	baseAddress := "https://pypi.python.org/pypi/%s/%s/json"
	address := fmt.Sprintf(baseAddress, name, version)
	response, err := http.Get(address)
	if err != nil {
		fmt.Println("Unable to get license, error during http request")
		return "Unable to resolve version"
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Unable to get license, error during reading data")
		return "Unable to resolve version"
	}

	var pypiResponse pypiResponse
	err = json.Unmarshal(body, &pypiResponse)
	if err != nil {
		fmt.Println("Unable to get license for package", name, version, ". Error during parsing json")
		return "Unable to resolve version"
	}

	return pypiResponse.Info.License
}
