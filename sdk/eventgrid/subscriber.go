package eventgrid

import (
	"errors"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
)

// Event allows for easy processing of Event Grid Events.
//
// External documentation on Event Grid Events can be found here:
// https://docs.microsoft.com/en-us/azure/event-grid/event-schema
type Event struct {
	ID              string      `json:"id"`
	Topic           string      `json:"topic"`
	Subject         string      `json:"subject"`
	Data            interface{} `json:"data"`
	EventType       string      `json:"eventType"`
	EventTime       time.Time   `json:"eventTime"`
	MetadataVersion string      `json:"metadataVersion"`
	DataVersion     string      `json:"dataVersion"`
}

// EventHandler extends the definition of buffalo.Handler to include
// an `Event`.
type EventHandler func(buffalo.Context, Event) error

// Subscriber allows for quick implementation of RESTful actions while
// working with Event Grid events.
type Subscriber interface {
	List(buffalo.Context) error
	New(buffalo.Context) error
	Receive(buffalo.Context) error
	Show(buffalo.Context) error
}

// BaseSubscriber will always respond to request by returning an HTTP 404 status.
type BaseSubscriber struct{}

// List responds with an HTTP NotFound Status Code.
func (s BaseSubscriber) List(c buffalo.Context) error {
	return c.Error(http.StatusNotFound, errors.New("subscriber not implemented"))
}

// New responds with an HTTP NotFound Status Code.
func (s BaseSubscriber) New(c buffalo.Context) error {
	return c.Error(http.StatusNotFound, errors.New("subscriber not implemented"))
}

// Receive responds with an HTTP NotFound Status Code.
func (s BaseSubscriber) Receive(c buffalo.Context) error {
	return c.Error(http.StatusNotFound, errors.New("subscriber not implemented"))
}

// Show responds with an HTTP NotFound Status Code.
func (s BaseSubscriber) Show(c buffalo.Context) error {
	return c.Error(http.StatusNotFound, errors.New("subscriber not implemented"))
}