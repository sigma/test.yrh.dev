build:
	mkdir -p functions
	go get ./src/...
	# change netlify.toml when increasing version
	go build -o functions/vanity_v3 -ldflags "-X main.user=sigma -X main.vanity=test.yrh.dev" ./src/vanity/vanity.go
	hugo --gc --minify -b $(URL)

