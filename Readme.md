If you're having this error when importing gorilla/mux:
main.go:9:2: no required module provides package github.com/gorilla/mux: working directory is not part of a module 

run these two command in cmd
```cmd 
go mod init main
go mod tidy
```