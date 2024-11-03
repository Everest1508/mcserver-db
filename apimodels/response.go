package apimodels

type TypeResponse struct {
	Status   string              `json:"status"`
	Response map[string][]string `json:"response"`
}

type FileResponse struct {
	Version   string   `json:"version"`
	File      string   `json:"file"`
	Size      FileSize `json:"size"`
	MD5       string   `json:"md5"`
	Built     int64    `json:"built"`
	Stability string   `json:"stability"`
}

type FileSize struct {
	Display string `json:"display"`
	Bytes   int    `json:"bytes"`
}

type SubTypeResponse struct {
	Status   string `json:"status"`
	Response map[string][]FileResponse
}
