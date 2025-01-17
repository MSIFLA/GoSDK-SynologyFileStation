# GoSDK-SynologyFileStation
Golang SDK for the Synology File Station API

> This SDK is an unofficial client library for Synology's File Station REST API. It is not affiliated with, endorsed by, or maintained by Synology. Use of this SDK is subject to Synology's API terms of use.

[**Synology API Reference**](https://global.download.synology.com/download/Document/Software/DeveloperGuide/Package/FileStation/All/enu/Synology_File_Station_API_Guide.pdf)

## Currently Implimented Methods
```
SYNO.API.Auth.login
SYNO.API.Auth.logout
SYNO.FileStation.List.list
SYNO.FileStation.CopyMove.start
SYNO.FileStation.CopyMove.status
SYNO.FileStation.CopyMove.stop
```

## Getting Started
```go
package main

import (
  ""
)

func main() {
  fs, err := api.NewFileStationAPIConn(
    host,
    user,
    password,
  )
  if err != nil {
    panic(err)
  }
  defer func(fs *api.FileStation) {
    _ = fs.Close()
  }(fs)
}
```
