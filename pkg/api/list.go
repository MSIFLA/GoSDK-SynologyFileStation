package api

import (
	"github.com/MSIFLA/GoSDK-SynologyFileStation/internal/data"
)

func (api *FileStation) List(request *ListDirRequest) (*ListDirResponse, error) {
	const (
		API     = "List"
		Version = 2
		Method  = "list"
	)
	var out ListDirResponse

	req, resp, err := api.doRequest(API, Version, Method, &request, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, api.error(req, out.Error)
	}

	output := resp.(*ListDirResponse)

	return output, nil
}

type ListDirResponse struct {
	Data struct {
		Files  []Files `json:"files"`
		Offset int     `json:"offset"`
		Total  int     `json:"total"`
	} `json:"data"`
	Error   data.Error `json:"error"`
	Success bool       `json:"success"`
}

type Files struct {
	IsDir      bool   `json:"isdir"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Additional struct {
		Owner struct {
			Gid   int    `json:"gid"`
			Group string `json:"group"`
			Uid   int    `json:"uid"`
			User  string `json:"user"`
		} `json:"owner"`
		Perm struct {
			Acl struct {
				Append bool `json:"append"`
				Del    bool `json:"del"`
				Exec   bool `json:"exec"`
				Read   bool `json:"read"`
				Write  bool `json:"write"`
			} `json:"acl"`
			IsAclMode bool `json:"is_acl_mode"`
			Posix     int  `json:"posix"`
		} `json:"perm"`
		RealPath string `json:"real_path"`
		Size     int    `json:"size"`
		Time     struct {
			Atime  int `json:"atime"`
			Crtime int `json:"crtime"`
			Ctime  int `json:"ctime"`
			Mtime  int `json:"mtime"`
		} `json:"time"`
		Type string `json:"type"`
	} `json:"additional"`
	Children []struct {
		Total  int     `json:"total"`
		Offset int     `json:"offset"`
		Files  []Files `json:"files"`
	} `json:"children"`
}

type ListDirRequest struct {
	Path          string `json:"folder_path"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	SortBy        string `json:"sort_by"`
	SortDirection string `json:"sort_direction"`
	Pattern       string `json:"pattern"`
	Filetype      string `json:"filetype"`
	GotoPath      string `json:"goto_path"`
	Additional    string `json:"additional"`
}
