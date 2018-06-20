.PHONY: all clean ensure clean-vendor

all:
	rm -rf vendor
	dep ensure -v
	GOOS=linux go build -o shop github.com/rancaka/shop/src
