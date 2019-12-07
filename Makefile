build:
	mkdir -p functions
	go get ./src/...
	go build -o functions/app -ldflags "-X main.user=sigma -X main.vanity=yrh.dev" ./src/app.go
	hugo --gc --minify -b $(URL)

