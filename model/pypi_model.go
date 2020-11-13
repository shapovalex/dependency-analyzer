package model

type PypiProjectUrls struct {
	Source   string
	Homepage string
}

type PypiInfo struct {
	License      string
	Project_urls map[string]string
	Description  string
}

type PypiResponse struct {
	Info PypiInfo
}
