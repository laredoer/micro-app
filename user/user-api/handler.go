package main

import (
	"context"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
	"encoding/json"
	"strings"
)

type Handler struct {

}

func(h *Handler) Hello(ctx context.Context,req *api.Request,resp *api.Response) (error) {
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.user","no content")
	}
	resp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": "got you request" + strings.Join(name.Values," "),
	})

	resp.Body = string(b)
	return nil
}