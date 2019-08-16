package transport

import (
	"cargo-routing/datastruct"
	dt "cargo-routing/datastruct"
	svc "cargo-routing/service"
	"context"
	"encoding/json"
	"log"

	//"kit/endpoint"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type AphService interface {
	GetStatusDeliveryService(context.Context, dt.Delivery) []dt.Delivery
}

type aphService struct {
}

func (aphService) GetStatusDeliveryService(_ context.Context, del dt.Delivery) []dt.Delivery {
	return call_ServiceGetStatusDeliveryService(del)
}

func call_ServiceGetStatusDeliveryService(del dt.Delivery) []dt.Delivery {
	retDel := svc.GetStatusDelivery(del)

	return retDel
}

func makeGetStatusDeliveryEndpoint(aph AphService) endpoint.Endpoint {
	log.Println("Process")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.GetStatusDeliveryRequest)
		paramDel := dt.Delivery{}
		paramDel.ID_ROUTE = req.ROUTE_ID
		paramDel.ID_ITENARY = req.ITENARY_ID
		aph.GetStatusDeliveryService(ctx, paramDel)
		return datastruct.GetStatusDeliveryResponse{
			ID_DELIVERY:         1,
			ROUTE_ID:            1,
			ITENARY_ID:          1,
			ROUTING_STATUS:      "unload",
			TRANSPORT_STATUS:    "inport",
			LAST_KNOWN_LOCATION: "surabaya",
			RESPONSE_CODE:       200,
			RESPONSE_DESC:       "Success",
		}, nil
	}
}

func decodeGetDeliveryStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.GetStatusDeliveryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHttpsServicesAndStartListener() {
	aph := aphService{}

	GetStatusDeliveryHandler := httptransport.NewServer(
		makeGetStatusDeliveryEndpoint(aph),
		decodeGetDeliveryStatusRequest,
		encodeResponse,
	)
	//url path of our API service
	http.Handle("/GetStatusDelivery", GetStatusDeliveryHandler)

}
