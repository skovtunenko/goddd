package handling

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"

	shipping "github.com/marcusolsson/goddd"
)

type registerIncidentRequest struct {
	ID             shipping.TrackingID
	Location       shipping.UNLocode
	Voyage         shipping.VoyageNumber
	EventType      shipping.HandlingEventType
	CompletionTime time.Time
}

type registerIncidentResponse struct {
	Err error `json:"error,omitempty"`
}

func (r registerIncidentResponse) error() error { return r.Err }

func makeRegisterIncidentEndpoint(hs Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerIncidentRequest)
		err := hs.RegisterHandlingEvent(req.CompletionTime, req.ID, req.Voyage, req.Location, req.EventType)
		return registerIncidentResponse{Err: err}, nil
	}
}
