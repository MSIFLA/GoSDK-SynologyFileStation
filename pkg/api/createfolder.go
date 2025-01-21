package api

import "github.com/MSIFLA/GoSDK-SynologyFileStation/internal/data"

// CreateFolderCreate :  SYNO.FileStation.CreateFolder.create
func (api *FileStation) CreateFolderCreate(request *CreateFolderCreateRequest) (*CreateFolderCreateResponse, error) {
	const (
		API     = "CreateFolder"
		Version = 2
		Method  = "create"
	)
	var out CreateFolderCreateResponse

	req, resp, err := api.doRequest(API, Version, Method, &request, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, api.error(req, out.Error)
	}

	output := resp.(*CreateFolderCreateResponse)

	return output, nil
}

type CreateFolderCreateRequest struct {
	FolderPath  string `json:"folder_path"`
	Name        string `json:"name"`
	ForceParent bool   `json:"force_parent"`
	Additional  string `json:"additional"`
}

type CreateFolderCreateResponse struct {
	Data struct {
		Folders []struct {
			IsDir bool   `json:"isdir"`
			Name  string `json:"name"`
			Path  string `json:"path"`
		} `json:"folders"`
	} `json:"data"`
	Error   data.Error `json:"error"`
	Success bool       `json:"success"`
}
