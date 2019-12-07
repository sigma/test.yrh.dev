package main

import (
	"bytes"
	"context"
	"log"
	"strings"
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
	log.Println(">>>", path)
	repo := strings.Split(path, "/")[0]
	tpl := `
<!doctype html>
<html>
  <head>
    <meta charset=utf-8>
    <meta name=go-import content="{{.Vanity}}/{{.Repo}} git https://github.com/{{.User}}/{{.Repo}}">
    <meta name=go-source content="{{.Vanity}}/{{.Repo}} https://github.com/{{.User}}/{{.Repo}} https://github.com/{{.User}}/{{.Repo}}/tree/master{/dir} https://github.com/{{.User}}/{{.Repo}}/blob/master{/dir}/{file}#L{line}">
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
	}, nil
}

func main() {
	lambda.Start(handler)
}
