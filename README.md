# Raincheck
## Abstract
Yoinked weatherapi to make a simple CLI app

<p align="center"><img src="https://github.com/The-Golang-Way/raincheck/blob/main/demo/demo.png" width="80%"></p>


## Things to note:
If the color module is not working (`go: go.mod file not found in current directory or any parent directory` or for some reason, it was not imported in the code), run this command `go env -w GO111MODULE=off` in the project directory. If it still does not work, manually download module by `go get github.com/fatih/color`

## Configuration:
- Make sure you have Golang installed (`go version`). If not, run `sudo apt install golang-go` for debian based systems. If you have a diffrent os, google it.
- cd into project directory and run `go run main.go`
- if you're on a linux distro and want to run the program anywhere in your file system, do `go build` in the project directory, then `sudo mv raincheck /usr/local/bin` and finally run `raincheck`.
