package fixture

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

// HandlerTestCase describes a successful HTTP handler scenario.
type HandlerTestCase[TService, THandler any] struct {
	Name   string
	Method string
	Path   string
	Body   any
	Expect func(TService)
	Invoke func(THandler, *gin.Context)
	Want   any
}

// NewFormHTTPContext creates a Gin test context with form-urlencoded data.
func NewFormHTTPContext(t *testing.T, method, path string, form url.Values) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	request := httptest.NewRequest(method, path, strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = request

	return ctx, recorder
}

// HandlerErrorTestCase describes a service-error HTTP handler scenario.
type HandlerErrorTestCase[TService, THandler any] struct {
	Name   string
	Expect func(TService)
	Invoke func(THandler, *gin.Context)
}

// HandlerInvalidInputTestCase describes an invalid transport input scenario.
type HandlerInvalidInputTestCase[THandler any] struct {
	Name   string
	Invoke func(THandler, *gin.Context)
}

// NewMockSUT creates a GoMock service and passes it to the handler constructor.
func NewMockSUT[TService, THandler any](
	t *testing.T,
	newService func(*gomock.Controller) TService,
	newHandler func(TService) THandler,
) (TService, THandler) {
	t.Helper()
	service := newService(gomock.NewController(t))
	return service, newHandler(service)
}

// NewHTTPContext creates a Gin test context with an optional JSON body.
func NewHTTPContext(t *testing.T, method, path string, body any) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	var data []byte
	if body != nil {
		var err error
		data, err = json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
	}

	request := httptest.NewRequest(method, path, bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = request

	return ctx, recorder
}

// NewRawHTTPContext creates a Gin test context with an unprocessed request body.
func NewRawHTTPContext(t *testing.T, method, path, contentType, body string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	request := httptest.NewRequest(method, path, strings.NewReader(body))
	request.Header.Set("Content-Type", contentType)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = request

	return ctx, recorder
}

// AssertHTTPResponse compares the response status and JSON body.
func AssertHTTPResponse(t *testing.T, recorder *httptest.ResponseRecorder, status int, want any) {
	t.Helper()
	if recorder.Code != status {
		t.Fatalf("status: want %d, got %d; body=%s", status, recorder.Code, recorder.Body.String())
	}

	if want == nil {
		if recorder.Body.String() != "null" && recorder.Body.String() != "" {
			t.Fatalf("body: want null or empty, got %s", recorder.Body.String())
		}
		return
	}

	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if recorder.Body.String() != string(wantJSON) {
		t.Fatalf("body: want %s, got %s", wantJSON, recorder.Body.String())
	}
}
