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
		releasePypiInfo := resolvePypiInfoForRelease(dependencyParts[0], dependencyParts[1])
		packagePypiInfo := resolvePypiInfoForPackage(dependencyParts[0])
		result = append(result, line+","+extractLicenseType(releasePypiInfo)+","+extractLicenseLineAddress(dependencyParts[0], dependencyParts[1], releasePypiInfo, packagePypiInfo))
	}
	errW := helper.WriteLines(result, params.OutputFiles)
	if errW != nil {
		fmt.Println("Unable to write to output files")
		os.Exit(1)
	}
}

func extractLicenseType(pypiInfo pypiInfo) string {
	if pypiInfo.License != "" {
		return pypiInfo.License
	}
	return "Undefined license type"
}

func (p PyPiLicense) GetSupportedDependencyManager() string {
	return "pip"
}

func (p PyPiLicense) GetSupportedOperation() string {
	return "license"
}

func checkAndExtractBaseUrl(info pypiInfo, urlString string) string {
	if info.Project_urls[urlString] != "" {
		return info.Project_urls[urlString]
	}
	return ""
}

func extractLicenseLineAddress(name string, version string, releaseInfo pypiInfo, packageInfo pypiInfo) string {
	var baseUrl string

	baseUrl = checkAndExtractBaseUrl(releaseInfo, "Source")
	if baseUrl != "" {
		return baseUrl + "/blob/" + version + "/LICENSE"
	}
	baseUrl = checkAndExtractBaseUrl(packageInfo, "Source")
	if baseUrl != "" {
		return baseUrl + "/blob/master/LICENSE"
	}

	baseUrl = checkAndExtractBaseUrl(releaseInfo, "Source Code")
	if baseUrl != "" {
		return baseUrl + "/blob/" + version + "/LICENSE"
	}
	baseUrl = checkAndExtractBaseUrl(packageInfo, "Source Code")
	if baseUrl != "" {
		return baseUrl + "/blob/master/LICENSE"
	}

	if strings.Contains(releaseInfo.Project_urls["Homepage"], "github") {
		baseUrl = checkAndExtractBaseUrl(releaseInfo, "Homepage")
		if baseUrl != "" {
			return baseUrl + "/blob/" + version + "/LICENSE"
		}
		baseUrl = checkAndExtractBaseUrl(packageInfo, "Homepage")
		if baseUrl != "" {
			return baseUrl + "/blob/master/LICENSE"
		}
	}

	githubIndex := strings.Index(packageInfo.Description, "github.com")
	if githubIndex != -1 {
		githubString := packageInfo.Description[githubIndex:]
		githubIndexSpace := strings.Index(githubString, " ")
		githubIndexNewLine := strings.Index(githubString, "\n")
		if githubIndexSpace != -1 || githubIndexNewLine != -1 {
			if githubIndexSpace == -1 {
				githubIndex = githubIndexNewLine
			} else if githubIndexNewLine == -1 {
				githubIndex = githubIndexSpace
			} else if githubIndexNewLine > githubIndexSpace {
				githubIndex = githubIndexSpace
			} else {
				githubIndex = githubIndexNewLine
			}
			githubString = githubString[:githubIndex]
		}
		return githubString
	}

	return "Undefined License URL"
}

func resolvePypiInfoForPackage(name string) pypiInfo {
	baseAddress := "https://pypi.python.org/pypi/%s/json"
	address := fmt.Sprintf(baseAddress, name)
	return resolvePypiInfo(address)
}

func resolvePypiInfoForRelease(name string, version string) pypiInfo {
	baseAddress := "https://pypi.python.org/pypi/%s/%s/json"
	address := fmt.Sprintf(baseAddress, name, version)
	return resolvePypiInfo(address)
}

func resolvePypiInfo(address string) pypiInfo {
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

	var pypiResponse pypiResponse
	err = json.Unmarshal(body, &pypiResponse)
	if err != nil {
		fmt.Println("Unable to get license for package", address, ". Error during parsing json")
		return getErrorPyPiInfo()
	}

	return pypiResponse.Info
}

func getErrorPyPiInfo() pypiInfo {
	info := new(pypiInfo)
	info.License = "Unable to resolve version"
	return *info
}
