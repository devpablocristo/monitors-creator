package restclient

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/melisource/fury_go-core/pkg/rusty"
	"github.com/melisource/fury_go-core/pkg/transport/httpclient"
)

type EndpointType struct {
	URL           string
	RustyEndpoint *rusty.Endpoint
}

func NewEndpoint(url string) (*EndpointType, error) {
	// TODO TIRAR KEYS E FAZER HOOK PARA ADICIONAR HEADER
	apiKey := "894299b0b7abd078cfdc48d8c11aa090"
	appKey := "ac972bc28fe0ea7a9ca35b06f7699613bf9c1721"

	clientOpts := []httpclient.OptionRetryable{
		httpclient.WithRequestHook(logRequestHook),
		httpclient.WithResponseHook(logResponseHook),
	}
	httpClient := httpclient.NewRetryable(3, clientOpts...)

	opts := []rusty.EndpointOption{
		rusty.WithHeader("Content-Type", "application/json"),
		rusty.WithHeader("DD-API-KEY", apiKey),
		rusty.WithHeader("DD-APPLICATION-KEY", appKey),
	}
	endpoint, err := rusty.NewEndpoint(httpClient, url, opts...)
	if err != nil {
		return nil, err
	}

	return &EndpointType{
		URL:           url,
		RustyEndpoint: endpoint,
	}, nil
}

func logRequestHook(r *http.Request) error {
	// TODO: ENABLE LOGGING
	// mylog.Debug(r.Context(), "request dispatched",
	// 	mylog.String("scheme", r.URL.Scheme),
	// 	mylog.String("host", r.URL.Host),
	// 	mylog.String("path", r.URL.Path),
	// 	mylog.String("verb", r.Method),
	// 	mylog.String("query", r.URL.Query().Encode()))
	return nil
}

func logResponseHook(r *http.Request, w *http.Response, err error) {
	//TODO ENABLE LOGGING
	// status := "unknown"
	// if w != nil {
	// 	status = w.Status
	// }
	// mylog.Debug(r.Context(), "response received",
	// 	mylog.Bool("succeed", err == nil),
	// 	mylog.String("status", status))

}

func (e *EndpointType) Get(ctx context.Context, data any) (error) {
	r, err := e.RustyEndpoint.Get(ctx)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(r.Body, &data); err != nil {
		return err
	}

	return nil
}
