package api

import (
	"github.com/MSIFLA/GoSDK-SynologyFileStation/internal/data"
	"io"
)

func (api *FileStation) CopyMoveStart(request *CopyMoveStartRequest) (*CopyMoveStartResponse, error) {
	const (
		API     = "CopyMove"
		Version = 3
		Method  = "start"
	)
	var out CopyMoveStartResponse

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

	if err = api.decodeResponse(resp, &out); err != nil {
		return nil, err
	}
	if !out.Success {
		return nil, api.error(req, out.Error)
	}

	return &out, nil
}

type CopyMoveStartRequest struct {
	Path             string  `json:"path"`
	DestFolderPath   string  `json:"dest_folder_path"`
	Overwrite        *bool   `json:"overwrite"`
	RemoveSrc        *bool   `json:"remove_src"`
	AccurateProgress *bool   `json:"accurate_progress"`
	SearchTaskId     *string `json:"search_taskid"`
}

type CopyMoveStartResponse struct {
	Data struct {
		TaskID string `json:"taskid"`
	} `json:"data"`
	Error   data.Error `json:"error"`
	Success bool       `json:"success"`
}

func (api *FileStation) CopyMoveStatus(request *CopyMoveStatusRequest) (*CopyMoveStatusResponse, error) {
	const (
		API     = "CopyMove"
		Version = 3
		Method  = "status"
	)
	var out CopyMoveStatusResponse

	req, output, err := api.doRequest(API, Version, Method, &request, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, api.error(req, out.Error)
	}

	resp := output.(*CopyMoveStatusResponse)

	return resp, nil
}

type CopyMoveStatusRequest struct {
	TaskId string `json:"taskid"`
}

type CopyMoveStatusResponse struct {
	Data struct {
		DestFolderPath string  `json:"dest_folder_path"`
		Finished       bool    `json:"finished"`
		FoundDirNum    int     `json:"found_dir_num"`
		FoundFileNum   int     `json:"found_file_num"`
		FoundFileSize  int64   `json:"found_file_size"`
		Path           string  `json:"path"`
		ProcessedSize  int64   `json:"processed_size"`
		ProcessingPath string  `json:"processing_path"`
		Progress       float64 `json:"progress"`
		SkipStatus     struct {
			Status string `json:"status"`
		}
		Status       string  `json:"status"`
		Total        int64   `json:"total"`
		TransferRate float64 `json:"transfer_rate"`
	} `json:"data"`
	Error   data.Error `json:"error"`
	Success bool       `json:"success"`
}

func (api *FileStation) CopyMoveStop(request *CopyMoveStopRequest) (*data.GenericResponse, error) {
	const (
		API     = "CopyMove"
		Version = 3
		Method  = "stop"
	)
	var out data.GenericResponse

	req, output, err := api.doRequest(API, Version, Method, &request, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, api.error(req, out.Error)
	}

	resp := output.(*data.GenericResponse)

	return resp, nil
}

type CopyMoveStopRequest struct {
	TaskId string `json:"taskid"`
}
