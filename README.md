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

## Installation
`go get github.com/MSIFLA/GoSDK-SynologyFileStation`

## Getting Started
```go
package main

import (
  "github.com/MSIFLA/GoSDK-SynologyFileStation/pkg/api"
)

const (
	host     = "YOUR_HOST_ADDRESS"
	user     = "YOUR_API_USER"
	password = "YOUR_USER_PASSWORD"
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

## FileStation: List
**Listing a Directory, Sorted by Name Descending with Size & Time Info:**
```go
resp, err := fs.List(&api.ListDirRequest{
  Path:          "/folder/path",
  SortBy:        "name",
  SortDirection: "desc",
  Additional:    "[\"size\", \"time\"]",
})
if err != nil {
  panic(err)
}
```

## FileStation: CopyMove
**Starting a CopyMove Operation:**
```go
// Directories only used as examples
copyStart, err := fs.CopyMoveStart(&api.CopyMoveStartRequest{
  Path:           "/folder/path",
  DestFolderPath: "/usbshare1/destination",
})
if err != nil {
  panic(err)
}
```

**Checking the Status of a CopyMove Operation:**
```go
status, err := fs.CopyMoveStatus(&api.CopyMoveStatusRequest{TaskId: copyStart.Data.TaskID})
if err != nil {
  panic(err)
}
```

**Stopping a CopyMove Operation Early:**
```go
_, err = fs.CopyMoveStop(&api.CopyMoveStopRequest{TaskId: copyStart.Data.TaskID})
if err != nil {
  panic(err)
}
```