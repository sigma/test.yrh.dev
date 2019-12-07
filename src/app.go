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
	path := request.QueryStringParameters["path"]
	tpl := `
<!doctype html>
<html>
  <head>
    <meta charset=utf-8>
    <meta name=go-import content="{{.Vanity}}/{{.Path}} git https://github.com/{{.User}}/{{.Path}}">
    <meta name=go-source content="{{.Vanity}}/{{.Path}} https://github.com/{{.User}}/{{.Path}} https://github.com/{{.User}}/{{.Path}}/tree/master{/dir} https://github.com/{{.User}}/{{.Path}}/blob/master{/dir}/{file}#L{line}">
    <meta http-equiv=refresh content="0; url=https://github.com/{{.User}}/{{.Path}}">
  </head>
</html>
`

	buf := new(bytes.Buffer)
	obj := struct{ Vanity, User, Path string }{vanity, user, path}
	if err := template.Must(template.New("goget").Parse(tpl)).Execute(buf, obj); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       buf.String(),
	}, nil
}

func main() {
	lambda.Start(handler)
}
