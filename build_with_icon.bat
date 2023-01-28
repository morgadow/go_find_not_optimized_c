@echo off  & setlocal

go install github.com/tc-hib/go-winres@latest
go-winres simply --icon icon.png
go build