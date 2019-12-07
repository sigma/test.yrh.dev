build:
	mkdir -p functions
	go get ./src/...
	# change static/_redirects when increasing version
	go build -o functions/vanity_v1 -ldflags "-X main.user=sigma -X main.vanity=test.yrh.dev" ./src/vanity/vanity.go
	hugo --gc --minify -b $(URL)

