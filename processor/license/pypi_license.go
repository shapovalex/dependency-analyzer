package license

import (
	"fmt"
	"github.com/shapovalex/depAnalyzer/helper"
	"github.com/shapovalex/depAnalyzer/model"
	"github.com/shapovalex/depAnalyzer/processor"
	"github.com/shapovalex/depAnalyzer/service"
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
		//releasePypiInfo := resolvePypiInfoForRelease(dependencyParts[0], dependencyParts[1])
		//packagePypiInfo := resolvePypiInfoForPackage(dependencyParts[0])
		//result = append(result, line+","+extractLicenseType(releasePypiInfo)+","+extractLicenseLineAddress(dependencyParts[0], dependencyParts[1], releasePypiInfo, packagePypiInfo))

		pypiInfoForPackage := service.ResolvePypiInfoForPackage(dependencyParts[0])
		result = append(result, line+","+extractLicenseType(pypiInfoForPackage)+","+service.FindLicenseFile(extractGithubAddress(pypiInfoForPackage)))
	}
	errW := helper.WriteLines(result, params.OutputFiles)
	if errW != nil {
		fmt.Println("Unable to write to output files")
		os.Exit(1)
	}
}

func extractLicenseType(pypiInfo model.PypiInfo) string {
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

func extractGithubAddress(packageInfo model.PypiInfo) string {
	var baseUrl string

	baseUrl = checkAndExtractBaseUrl(packageInfo, "Source")
	if baseUrl != "" {
		return baseUrl
	}

	baseUrl = checkAndExtractBaseUrl(packageInfo, "Source Code")
	if baseUrl != "" {
		return baseUrl
	}

	baseUrl = checkAndExtractBaseUrl(packageInfo, "Homepage")
	if baseUrl != "" {
		return baseUrl
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

func checkAndExtractBaseUrl(info model.PypiInfo, urlString string) string {
	if strings.Contains(info.Project_urls[urlString], "github.com") {
		return info.Project_urls[urlString]
	}
	return ""
}
