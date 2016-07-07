# About

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
