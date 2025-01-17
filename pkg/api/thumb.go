package api

import (
	"io"
)

// ThumbGet : SYNO.FileStation.Thumb.get
func (api *FileStation) ThumbGet(request *ThumbGetRequest) ([]byte, error) {
	const (
		API     = "Thumb"
		Version = 2
		Method  = "get"
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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type ThumbGetRequest struct {
	Path   string `json:"path"`
	Size   string `json:"size"`
	Rotate int    `json:"rotate"`
}
