package api

import "github.com/MSIFLA/GoSDK-SynologyFileStation/internal/data"

// VirtualFolderList : SYNO.FileStation.VirtualFolder.list
func (api *FileStation) VirtualFolderList(request *VirtualFolderListRequest) (*VirtualFolderListResponse, error) {
	const (
		API     = "VirtualFolder"
		Version = 2
		Method  = "list"
	)
	var out VirtualFolderListResponse

	req, resp, err := api.doRequest(API, Version, Method, &request, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, api.error(req, out.Error)
	}

	output := resp.(*VirtualFolderListResponse)

	return output, nil
}

type VirtualFolderListRequest struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	SortBy        string `json:"sort_by"`
	SortDirection string `json:"sort_direction"`
	Additional    string `json:"additional"`
}

type VirtualFolderListResponse struct {
	Data struct {
		Total   int `json:"total"`
		Offset  int `json:"offset"`
		Folders []struct {
			Path       string `json:"path"`
			Name       string `json:"name"`
			Additional struct {
				RealPath string `json:"real_path"`
				Owner    struct {
					UserName  string `json:"user_name"`
					GroupName string `json:"group_name"`
					UID       int    `json:"uid"`
					GID       int    `json:"gid"`
				} `json:"owner"`
				Time struct {
					LastAccessTime   string `json:"last_access_time"`
					LastModifiedTime string `json:"last_modified_time"`
					LastChangeTime   string `json:"last_change_time"`
					CreateTime       string `json:"create_time"`
				} `json:"time"`
				Perm struct {
					Read    bool `json:"read"`
					Write   bool `json:"write"`
					Execute bool `json:"execute"`
				} `json:"perm"`
				MountPointType string `json:"mount_point_type"`
				VolumeStatus   struct {
					FreeSpace  int64 `json:"free_space"`
					TotalSpace int64 `json:"total_space"`
					ReadOnly   bool  `json:"read_only"`
				} `json:"volume_status"`
			} `json:"additional"`
		} `json:"folders"`
	}
	Error   data.Error `json:"error"`
	Success bool       `json:"success"`
}
