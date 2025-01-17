package api

import (
	"GoSDK-SynologyFileStation/internal/data"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// API Reference:
// https://global.download.synology.com/download/Document/Software/DeveloperGuide/Package/FileStation/All/enu/Synology_File_Station_API_Guide.pdf

/*
Implemented Methods:
	SYNO.API.Auth.login
	SYNO.API.Auth.logout
	SYNO.FileStation.List.list
	SYNO.FileStation.CopyMove.start
	SYNO.FileStation.CopyMove.status
	SYNO.FileStation.CopyMove.stop
*/

type FileStation struct {
	sid     string
	baseUrl string
	user    string
	pass    string
	client  *http.Client
}

func (api *FileStation) Close() error {
	if err := api.logout(); err != nil {
		return err
	}
	return nil
}

func NewFileStationAPIConn(host string, user string, password string) (*FileStation, error) {
	fs := FileStation{}
	fs.baseUrl = fmt.Sprintf("http://%s:5000/webapi/", host)
	fs.user = user
	fs.pass = password
	fs.client = &http.Client{}
	if err := fs.login(); err != nil {
		return nil, err
	}
	return &fs, nil
}

func (api *FileStation) buildStandardRequest(apiName string, version int, method string, params map[string]string) (*http.Request, error) {
	reqUrl := fmt.Sprintf(
		"%sentry.cgi?api=SYNO.FileStation.%s&version=%d&method=%s&sid=%s",
		api.baseUrl,
		apiName,
		version,
		method,
		api.sid,
	)
	for key, value := range params {
		reqUrl += fmt.Sprintf("&%s=%s", key, value)
	}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cookie", fmt.Sprintf("id=%s", api.sid))

	return req, nil
}

func (api *FileStation) decodeResponse(resp *http.Response, out interface{}) error {

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	// Decode JSON from the body
	if err := json.Unmarshal(bodyBytes, out); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	return nil
}

func (api *FileStation) error(req *http.Request, err data.Error) error {
	params := req.URL.Query()
	apiName := params["api"]
	method := params["method"]
	errDesc := "UNKNOWN ERROR"
	if desc, ok := data.ErrorCodes[err.Code]; ok {
		errDesc = desc
	}
	return fmt.Errorf("error encountered calling %s.%s: Code %d\nURL: %s\nParams:%s\nError Code Desc: %s", apiName, method, err.Code, req.URL.String(), params, errDesc)
}

func (api *FileStation) buildParamsFromStruct(v interface{}) (map[string]string, error) {
	result := make(map[string]string)

	jsonData, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var intermediate map[string]interface{}
	err = json.Unmarshal(jsonData, &intermediate)
	if err != nil {
		return nil, err
	}

	// Iterate through the map and convert values to string, filtering out nil values
	for key, value := range intermediate {
		if value == nil {
			continue
		}

		// Convert to string (if possible)
		strValue, ok := value.(string)
		if !ok {
			strValue = fmt.Sprintf("%v", value)
		}

		if strValue != "" {
			result[key] = strValue
		}
	}

	return result, nil
}

func (api *FileStation) doRequest(
	apiName string,
	version int,
	method string,
	request interface{},
	out interface{},
) (*http.Request, interface{}, error) {
	params, err := api.buildParamsFromStruct(request)
	if err != nil {
		return nil, nil, err
	}

	req, err := api.buildStandardRequest(
		apiName,
		version,
		method,
		params,
	)
	if err != nil {
		return nil, nil, err
	}
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if err = api.decodeResponse(resp, &out); err != nil {
		return nil, nil, err
	}

	return req, out, nil
}
