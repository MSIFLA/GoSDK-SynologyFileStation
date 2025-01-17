package api

import "github.com/MSIFLA/GoSDK-SynologyFileStation/internal/data"

// InfoGet : SYNO.FileStation.Info.get
func (api *FileStation) InfoGet() (*InfoGetResponse, error) {
	const (
		API     = "Info"
		Version = 2
		Method  = "get"
	)
	var out InfoGetResponse

	req, resp, err := api.doRequest(API, Version, Method, nil, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, api.error(req, out.Error)
	}

	output := resp.(*InfoGetResponse)

	return output, nil
}

type InfoGetResponse struct {
	Data struct {
		EnableListUserGrp bool   `json:"enable_list_usergrp"`
		Hostname          string `json:"hostname"`
		IsManager         bool   `json:"is_manager"`
		Items             []struct {
			GID int `json:"gid"`
		} `json:"items"`
		SupportFileRequest bool `json:"support_file_request"`
		SupportSharing     bool `json:"support_sharing"`
		SupportVFS         bool `json:"support_vfs"`
		SupportVirtual     struct {
			EnableISOMount    bool `json:"enable_iso_mount"`
			EnableRemoteMount bool `json:"enable_remote_mount"`
		} `json:"support_virtual"`
		SupportVirtualProtocol []string `json:"support_virtual_protocol"`
		SystemCodepage         string   `json:"system_codepage"`
		UID                    int      `json:"uid"`
	} `json:"data"`
	Error   data.Error `json:"error"`
	Success bool       `json:"success"`
}
