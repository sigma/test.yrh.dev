package main

import (
	"bytes"
	"context"
	"text/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	vanity string
	user   string
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	repo := request.Headers["X-Repo"]
	tpl := `
<!doctype html>
<html>
  <head>
    <meta charset=utf-8>
    <meta name=go-import content="{{.Vanity}}/{{.Repo}} git https://github.com/{{.User}}/{{.Repo}}">
	<meta name=go-source content="{{.Vanity}}/{{.Repo}} https://github.com/{{.User}}/{{.Repo}} https://github.com/{{.User}}/{{.Repo}}/tree/master{/dir} https://github.com/{{.User}}/{{.Repo}}/blob/master{/dir}/{file}#L{line}">
	<title>https://{{.Vanity}}/{{.Repo}}</title>
	<link rel="canonical" href="https://{{.Vanity}}/{{.Repo}}/"/>
	<meta http-equiv="content-type" content="text/html; charset=utf-8" />
	<meta http-equiv="refresh" content="0; url=https://{{.Vanity}}/{{.Repo}}/" />
  </head>
</html>
`

	buf := new(bytes.Buffer)
	obj := struct{ Vanity, User, Repo string }{vanity, user, repo}
	if err := template.Must(template.New("goget").Parse(tpl)).Execute(buf, obj); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       buf.String(),
		Headers: map[string]string{
			"Cache-Control": "public, max-age=31536000",
			"Vary":          "X-Repo",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
