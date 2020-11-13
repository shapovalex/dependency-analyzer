package license

type pypiProjectUrls struct {
	Source   string
	Homepage string
}

type pypiInfo struct {
	License      string
	Project_urls map[string]string
	Description  string
}

type pypiResponse struct {
	Info pypiInfo
}
