package priceplans

import (
	"context"
	"net/url"
	"strconv"

	//	"fmt"

	"github.com/go-kit/kit/endpoint"

	"joi-energy-golang/http/contextkeys"
)

func makeCompareAllPricePlansEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		err := validateSmartMeterId(req)
		if err != nil {
			return nil, err
		}
		r, err := s.CompareAllPricePlans(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

func makeRecommendPricePlansEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//fmt.Println("this is ur.values from test!!",contextkeys.QueryValues)
		//context.Background.WithValue(type string, val 10)
		limitString := ctx.Value(contextkeys.QueryValues).(url.Values).Get("limit")
		//fmt.Println("this is our data------->>",ctx.Value(contextkeys.QueryValues))
		limit, err := strconv.ParseUint(limitString, 10, 64)
		if limitString != "" && err != nil {
			return nil, err
		}
		req := request.(string)
		err = validateSmartMeterId(req)
		if err != nil {
			return nil, err
		}
		r, err := s.RecommendPricePlans(req, limit)
		if err != nil {
			return nil, err
		}
		return r, nil // here r is endpoint.Endpoint,
	}
}

func makeRecommendPricePlansWeeklyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//fmt.Println("this is ur.values from test!!",contextkeys.QueryValues)
		//context.Background.WithValue(type string, val 10)
		limitString := ctx.Value(contextkeys.QueryValues).(url.Values).Get("limit")
		//fmt.Println("this is our data------->>",ctx.Value(contextkeys.QueryValues))
		limit, err := strconv.ParseUint(limitString, 10, 64)
		if limitString != "" && err != nil {
			return nil, err
		}
		req := request.(string)
		err = validateSmartMeterId(req)
		if err != nil {
			return nil, err
		}
		r, err := s.RecommendPricePlansWeekly(req, limit)
		if err != nil {
			return nil, err
		}
		return r, nil // here r is endpoint.Endpoint,
	}
}
