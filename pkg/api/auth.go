package api

import (
	"GoSDK-SynologyFileStation/internal/data"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (api *FileStation) login() error {
	reqUrl := fmt.Sprintf(
		"%sauth.cgi?api=SYNO.API.Auth&version=3&method=login&account=%s&passwd=%s&session=FileStation",
		api.baseUrl,
		api.user,
		api.pass,
	)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	var decoded data.AuthLoginResponse
	if err = api.decodeResponse(resp, &decoded); err != nil {
		return err
	}
	if !decoded.Success {
		return api.authError(reqUrl, decoded.Error)
	}
	api.sid = decoded.Data.SID
	return nil
}

func (api *FileStation) logout() error {
	reqUrl := fmt.Sprintf(
		"%sauth.cgi?api=SYNO.API.Auth&version=1&method=logout&session=FileStation",
		api.baseUrl,
	)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	var decoded data.GenericResponse
	if err = api.decodeResponse(resp, &decoded); err != nil {
		return err
	}
	if !decoded.Success {
		return api.authError(reqUrl, decoded.Error)
	}
	return nil
}

func (api *FileStation) authError(reqUrl string, authErr data.Error) error {
	parsedUrl, err := url.Parse(reqUrl)
	if err != nil {
		return err
	}

	params := parsedUrl.Query()
	apiName := params["api"]
	method := params["method"]
	errDesc := "UNKNOWN ERROR"
	if desc, ok := data.AuthErrorCodes[authErr.Code]; ok {
		errDesc = desc
	}
	return fmt.Errorf("auth error encountered calling %s.%s: Code %d\nURL: %s\nParams:%s\nError Code Desc: %s", apiName, method, authErr.Code, parsedUrl.String(), params, errDesc)
}
