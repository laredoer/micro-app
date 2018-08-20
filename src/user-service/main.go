package main

import (
	"encoding/json"
	"log"
	"strings"

	proto "github.com/micro/examples/api/default/proto"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"

	"context"
)

type Example struct{}


// Example.Call is a method which will be served by http request /example/call
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /example/call goes to go.micro.api.example Example.Call
func (e *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Example.Call request")

	// parse values from the get request
	name, ok := req.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}

	// set response status
	rsp.StatusCode = 200

	// respond with some json
	b, _ := json.Marshal(map[string]string{
		"message": "got your request " + strings.Join(name.Values, " "),
	})

	// set json body
	rsp.Body = string(b)

	return nil
}




func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)

	service.Init()

	// register example handler
	proto.RegisterExampleHandler(service.Server(), new(Example))

	// register foo handler
	proto.RegisterFooHandler(service.Server(), new(Foo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}