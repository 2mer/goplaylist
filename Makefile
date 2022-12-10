FILENAME=gonga
BUILD_DIR=../dist
OUTFILE=$(BUILD_DIR)/$(FILENAME)


# build flavors
windows:
	cd src;\
	env GOOS=windows GOARCH=amd64 go build -o $(OUTFILE)-win64.exe

linux:
	cd src;\
	env GOOS=linux GOARCH=amd64 go build -o $(OUTFILE)


clean:
	rm -rf dist

build: clean windows

run:
	cd src;\
	go run .