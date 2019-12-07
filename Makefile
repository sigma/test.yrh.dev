build:
	mkdir -p functions
	go get ./src/...
	go build -o functions/vanity -ldflags "-X main.user=sigma -X main.vanity=yrh.dev" ./src/vanity/vanity.go
	hugo --gc --minify -b $(URL)

