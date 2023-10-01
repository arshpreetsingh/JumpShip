package urlshortner

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Endpoints struct {
	ShortenUrlEndpoint endpoint.Endpoint
	GetLongUrlEndpoint *http.Handler
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service. Useful in a profilesvc
// server.
func MakeServerEndpoints(s ServiceURL) Endpoints {
	return Endpoints{
		ShortenUrlEndpoint: MakeShortenUrlEndpoint(s),
		GetLongUrlEndpoint: MakeLongUrlEndpoint(s),
	}
}

// MakeClientEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the remote instance, via a transport/http.Client.

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

	// Note that the request encoders need to modify the request URL, changing
	// the path. That's fine: we simply need to provide specific encoders for
	// each endpoint.

	return Endpoints{
		ShortenUrlEndpoint: httptransport.NewClient("POST", tgt, encodeShortenUrlRequest, decodeShortenUrlResponse, options...).Endpoint(),
	}, nil
}

// ShortenUrl implements Service. Primarily useful in a client.
func (e Endpoints) ShortenUrl(ctx context.Context, model UrlModel) error {
	request := shortenUrlRequest{UrlModel: model}
	response, err := e.ShortenUrlEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(shortenUrlResponse)
	return resp.Err
}

func MakeShortenUrlEndpoint(s ServiceURL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(shortenUrlRequest)
		fmt.Println("hellloooooooooowwwwwwwww")
		response, e := s.ShortenUrl(ctx, req.UrlModel)
		//convert response{} to string
		return shortenUrlResponse{GENUrl: response.(string), Err: e}, nil
	}
}

//type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

func MakeLongUrlEndpoint(s ServiceURL) *http.Handler {
	return func(ctx context.Context, request interface{}) *http.Handler {
		req := request.(getProfileRequest)
		fmt.Println("shorten Url request--->>", req)
		//fmt.Println(mux.Vars())
		//response, e := s.GetLongUrl(ctx, req.UrlModel)
		//convert response{} to string
		resp := "hello"
		err = errors.New("ok Error")
		// that part is fixed, Just need to fix the URL-Redirect method here!!
		http.Redirect(ctx, request, "http://www.google.com", 200)
		//return http.Redirect("http://google.com", http.StatusSeeOther)
		return resp, err
		//return shortenUrlResponse{GENUrl: resp, Err: err}, nil
		//return shortenUrlResponse{GENUrl: response.(string), Err: e}, nil
	}
}

type shortenUrlRequest struct {
	UrlModel UrlModel
}

type getProfileRequest struct {
	ID string
}

type shortenUrlResponse struct {
	GENUrl string
	Err    error `json:"err,omitempty"`
}

func (r shortenUrlResponse) error() error { return r.Err }
