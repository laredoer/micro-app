package main

import (
	"encoding/json"
	"log"

	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
	"context"
)

type Foo struct{}

// Foo.Bar is a method which will be served by http request /example/foo/bar
// Because Foo is not the same as the service name it is mapped beyond /example/
func (f *Foo) Bar(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Foo.Bar request")

	// check method
	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.example", "require post")
	}

	// let's make sure we get json
	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "need content-type")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.example", "expect application/json")
	}

	// parse body
	var body map[string]interface{}
	json.Unmarshal([]byte(req.Body), &body)

	// do something with parsed body

	return nil
}