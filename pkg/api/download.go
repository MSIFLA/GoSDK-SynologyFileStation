package api

import (
	"io"
)

func (api *FileStation) Download(request *DownloadRequest) ([]byte, error) {

	const (
		API     = "Download"
		Version = 2
		Method  = "download"
	)

	params, err := api.buildParamsFromStruct(request)
	if err != nil {
		return nil, err
	}

	req, err := api.buildStandardRequest(
		API,
		Version,
		Method,
		params,
	)
	if err != nil {
		return nil, err
	}
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	output, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return output, nil
}

type DownloadRequest struct {
	Path string `json:"path"`
	Mode string `json:"mode"` // open or download
}
