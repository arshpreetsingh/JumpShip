package readings

import (
	"context"
//   "fmt"
	"github.com/go-kit/kit/endpoint"

	"joi-energy-golang/domain"
	//"reflect"
)

// All Functions in EndPoint will Return Endpoint Handler!
func makeStoreReadingsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// converting interfce to Struct ni golang. aka type-Casting
		req := request.(domain.StoreReadings)
		s.StoreReadings(req.SmartMeterId, req.ElectricityReadings)
		return req, nil
	}
}

func makeGetReadingsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		err := validateSmartMeterId(req)
		if err != nil {
			return nil, err
		}
		res := s.GetReadings(req)
		return domain.StoreReadings{
			SmartMeterId: req,
			ElectricityReadings: res,
		}, nil
	}
}
