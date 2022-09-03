# wsui

## Features 
1. Prompt reconnect upon server temrinating connection
2. Bookmark menu to list connections

## Build sources 
Download dependencies 

`go mod download`

Build binary

`go build` || `go build -o wsui.x`

for Windows (make sure to include the flags or a termnal window will open alongside the app)

`go build -ldflags -H=windowsgui`

## Example
![](https://github.com/Lyca0n/wsui/blob/main/docs/Capture.PNG?raw=true)