package main

import (
	"context"
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type EndPoints struct {
	ShortenUrlEndPoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) EndPoints {
	return EndPoints{
		ShortenUrlEndPoint: MakeShortenUrlEndPoint(s),
	}
}

func MakeClientEndpoints(instance string) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	options := []httptransport.ClientOption{}

	return Endpoints{
		ShortenUrlEndPoint: httptransport.NewClient("POST", tgt, encodePostProfileRequest, decodePostProfileRequest, options...).Endpoint(),
	}, nil
}

func (e Endpoints) ShortenUrl(ctx context.Context, url string) error {
	request := postUrlRequest{URL: url}
	response, err := e.ShortenUrlEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(string)
	return resp.Err
}

type postUrlRequest struct {
	Profile Profile
}
