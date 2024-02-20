package component

import (
	"errors"

	"github.com/leslie-qiwa/flat"
	"github.com/projectdiscovery/retryablehttp-go"
)

// ErrSetValue is a error raised when a value cannot be set
var ErrSetValue = errors.New("could not set value")

// Component is a component for a request
type Component interface {
	// Name returns the name of the component
	Name() string
	// Parse parses the component and returns the
	// parsed component
	Parse(req *retryablehttp.Request) (bool, error)
	// Iterate iterates over all values of a component
	// ex in case of query component, it will iterate over each query parameter
	// depending on the rule if mode is single
	// request is rebuilt for each value in this callback
	// and in case of multiple, request will be rebuilt after iteration of all values
	Iterate(func(key string, value interface{}))
	// SetValue sets a value in the component
	// for a key
	//
	// After calling setValue for mutation, the value must be
	// called again so as to reset the body to its original state.
	SetValue(key string, value string) error
	// Rebuild returns a new request with the
	// component rebuilt
	Rebuild() (*retryablehttp.Request, error)
}

const (
	// RequestBodyComponent is the name of the request body component
	RequestBodyComponent = "body"
	// RequestQueryComponent is the name of the request query component
	RequestQueryComponent = "query"
	// RequestPathComponent is the name of the request url component
	RequestPathComponent = "path"
	// RequestHeaderComponent is the name of the request header component
	RequestHeaderComponent = "headers"
	// RequestCookieComponent is the name of the request cookie component
	RequestCookieComponent = "cookie"
)

// Components is a list of all available components
var Components = []string{
	RequestBodyComponent,
	RequestQueryComponent,
	RequestPathComponent,
	RequestHeaderComponent,
	RequestCookieComponent,
}

// New creates a new component for a componentType
func New(componentType string) Component {
	switch componentType {
	case "body":
		return NewBody()
	case "query":
		return NewQuery()
	case "path":
		return NewPath()
	case "headers":
		return NewHeader()
	case "cookie":
		return NewCookie()
	}
	return nil
}

var (
	flatOpts = &flat.Options{
		Safe:      true,
		Delimiter: "~",
	}
)