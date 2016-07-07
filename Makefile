.PHONY: www

www:
	gopherjs build -v -o www/gol.js cmd/gol/main.go
	go run cmd/gol_http/main.go
