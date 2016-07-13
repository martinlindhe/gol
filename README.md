# About

[![Travis-CI](https://api.travis-ci.org/martinlindhe/gol.svg)](https://travis-ci.org/martinlindhe/gol)
[![codecov.io](https://codecov.io/github/martinlindhe/gol/coverage.svg?branch=master)](https://codecov.io/github/martinlindhe/gol?branch=master)
[![GoDoc](https://godoc.org/github.com/martinlindhe/gol?status.svg)](https://godoc.org/github.com/martinlindhe/gol)
[![Go Report Card](https://goreportcard.com/badge/github.com/martinlindhe/gol)](https://goreportcard.com/report/github.com/martinlindhe/gol)


game of life

![screenshot](screen.png)


# Try it

to launch desktop version:

    go get -u github.com/martinlindhe/gol/cmd/gol
    gol

to launch browser version:

    cd $GOPATH/src/github.com/martinlindhe/gol
    gopherjs build -v -o www/gol.js cmd/gol/main.go
	go run cmd/gol_http/main.go


## License

Under [MIT](LICENSE)
