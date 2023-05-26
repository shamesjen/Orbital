package main

import (
	"context"
	api "github.com/shamesjen/Orbital/kitex_gen/api"
)

// TestImpl implements the last service interface defined in the IDL.
type TestImpl struct{}

// Reply implements the TestImpl interface.
func (s *TestImpl) Reply(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...

	if (req.Message == "james") {
		return &api.Response{Message: "Hello there, " + req.Message}, nil
	} else {
		return &api.Response{Message: "You are not James"}, nil
	}
}

