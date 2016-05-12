// Package httpexpect helps to write nice tests for your HTTP API.
//
// Usage example
//
// See example directory:
//  - https://godoc.org/github.com/gavv/httpexpect/example
//  - https://github.com/gavv/httpexpect/tree/master/example
//
// Value equality
//
// Whenever values are checked for equality in httpexpect, they are converted
// to "canonical form":
//  - type aliases are removed
//  - numeric types are converted to float64
//  - non-nil interfaces pointing to nil slices and maps are replaced with nil interfaces
//  - structs are converted to map[string]interface{}
//
// This is equivalent to subsequently json.Marshal() and json.Unmarshal() the value.
//
// Failure handling
//
// When some check fails, failure is reported. If non-fatal failures are used
// (see Reporter interface), execution is continued and instance that was checked
// is marked as failed.
//
// If specific instance is marked as failed, all subsequent checks are ignored
// for this instance and for any child instances retrieved after failure.
//
// Example:
//  array := NewArray(NewAssertReporter(t), []interface{}{"foo", 123})
//
//  e0 := array.Element(0)  // success
//  e1 := array.Element(1)  // success
//
//  s0 := e0.String()       // success
//  s1 := e1.String()       // failure; e1 and s1 are marked as failed, e0 and s0 are not
//
//  s0.Equal("foo")         // success
//  s1.Equal("bar")         // this check is ignored because s1 is marked as failed
package httpexpect

import (
	"net/http"
	"testing"
)

// Expect is a toplevel object that contains user Config and allows
// to construct Request objects.
type Expect struct {
	config Config
}

// Config contains various settings.
type Config struct {
	// BaseURL is a URL to prepended to all request. My be empty. If
	// non-empty, trailing slash is allowed but not required and is
	// appended automatically.
	BaseURL string

	// Client is used to send http.Request and receive http.Response.
	// Should not be nil.
	//
	// You can use http.DefaultClient or http.Client, or provide
	// custom implementation.
	Client Client

	// Printer is used to print requests and responses.
	// May be nil.
	//
	// You can use CompactPrinter or DebugPrinter, or provide custom
	// implementation.
	//
	// You can also use CompactPrinter or DebugPrinter with alternative
	// Logger if you're happy with their format, but want to send logs
	// somewhere else instead of testing.T.
	Printer Printer

	// Reporter is used to report failures.
	// Should not be nil.
	//
	// You can use AssertReporter, RequireReporter (they use testify),
	// or testing.T, or provide custom implementation.
	Reporter Reporter
}

// Client is used to send http.Request and receive http.Response.
// Note that http.Client implements this interface.
type Client interface {
	// Do sends request and returns response.
	Do(*http.Request) (*http.Response, error)
}

// Printer is used to print requests and responses.
// CompactPrinter and DebugPrinter implement this interface.
type Printer interface {
	// Request is called before request is sent.
	Request(*http.Request)

	// Response is called after response is received.
	Response(*http.Response)
}

// Logger is used as output backend for Printer.
// testing.T implements this interface.
type Logger interface {
	// Logf writes message to log.
	Logf(fmt string, args ...interface{})
}

// Reporter is used to report failures.
// testing.T implements this interface. AssertReporter or RequireReporter,
// also implement this interface using testify.
type Reporter interface {
	// Errorf reports failure.
	// Allowed to return normally or terminate test using t.FailNow().
	Errorf(message string, args ...interface{})
}

// New returns a new Expect object.
//
// baseURL specifies URL to prepended to all request. My be empty. If non-empty,
// trailing slash is allowed but not required and is appended automatically.
//
// New is shorthand for WithConfig. It uses:
//  - http.DefaultClient as Client
//  - CompactPrinter as Printer with testing.T as Logger
//  - AssertReporter as Reporter
//
// Example:
//  func TestAPI(t *testing.T) {
//      e := httpexpect.New(t, "http://example.org/")
//      e.GET("/path").Expect().Status(http.StatusOK)
//  }
func New(t *testing.T, baseURL string) *Expect {
	return WithConfig(Config{
		BaseURL:  baseURL,
		Reporter: NewAssertReporter(t),
		Printer:  NewCompactPrinter(t),
	})
}

// WithConfig returns a new Expect object with given config.
//
// If Config.Client is nil, http.DefaultClient is used.
//
// Example:
//  func TestAPI(t *testing.T) {
//      e := httpexpect.WithConfig(httpexpect.Config{
//          BaseURL:  "http://example.org/",
//          Client:   http.DefaultClient,
//          Printer:  httpexpect.NewDebugPrinter(t),
//          Reporter: httpexpect.NewAssertReporter(t),
//      })
//      e.GET("/path").Expect().Status(http.StatusOK)
//  }
func WithConfig(config Config) *Expect {
	if config.Client == nil {
		config.Client = http.DefaultClient
	}
	if config.Reporter == nil {
		panic("config.Reporter is nil")
	}
	return &Expect{config}
}

// Request is a shorthand for NewRequest(config, method, url, args...).
func (e *Expect) Request(method, url string, args ...interface{}) *Request {
	return NewRequest(e.config, method, url, args...)
}

// OPTIONS is a shorthand for NewRequest(config, "OPTIONS", url, args...).
func (e *Expect) OPTIONS(url string, args ...interface{}) *Request {
	return NewRequest(e.config, "OPTIONS", url, args...)
}

// HEAD is a shorthand for NewRequest(config, "HEAD", url, args...).
func (e *Expect) HEAD(url string, args ...interface{}) *Request {
	return NewRequest(e.config, "HEAD", url, args...)
}

// GET is a shorthand for NewRequest(config, "GET", url, args...).
func (e *Expect) GET(url string, args ...interface{}) *Request {
	return NewRequest(e.config, "GET", url, args...)
}

// POST is a shorthand for NewRequest(config, "POST", url, args...).
func (e *Expect) POST(url string, args ...interface{}) *Request {
	return NewRequest(e.config, "POST", url, args...)
}

// PUT is a shorthand for NewRequest(config, "PUT", url, args...).
func (e *Expect) PUT(url string, args ...interface{}) *Request {
	return NewRequest(e.config, "PUT", url, args...)
}

// PATCH is a shorthand for NewRequest(config, "PATCH", url, args...).
func (e *Expect) PATCH(url string, args ...interface{}) *Request {
	return NewRequest(e.config, "PATCH", url, args...)
}

// DELETE is a shorthand for NewRequest(config, "DELETE", url, args...).
func (e *Expect) DELETE(url string, args ...interface{}) *Request {
	return NewRequest(e.config, "DELETE", url, args...)
}

// Value is a shorthand for NewValue(Config.Reporter, value).
func (e *Expect) Value(value interface{}) *Value {
	return NewValue(e.config.Reporter, value)
}

// Object is a shorthand for NewObject(Config.Reporter, value).
func (e *Expect) Object(value map[string]interface{}) *Object {
	return NewObject(e.config.Reporter, value)
}

// Array is a shorthand for NewArray(Config.Reporter, value).
func (e *Expect) Array(value []interface{}) *Array {
	return NewArray(e.config.Reporter, value)
}

// String is a shorthand for NewString(Config.Reporter, value).
func (e *Expect) String(value string) *String {
	return NewString(e.config.Reporter, value)
}

// Number is a shorthand for NewNumber(Config.Reporter, value).
func (e *Expect) Number(value float64) *Number {
	return NewNumber(e.config.Reporter, value)
}

// Boolean is a shorthand for NewBoolean(Config.Reporter, value).
func (e *Expect) Boolean(value bool) *Boolean {
	return NewBoolean(e.config.Reporter, value)
}
