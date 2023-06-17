# raincheck
Yoinked weatherapi to make a simple CLI app

Things to note:
If the color module is not working (`go: go.mod file not found in current directory or any parent directory`), run this command `go env -w GO111MODULE=off` in the project directory. If it still does not work, manually download module by `go get github.com/fatih/color`
