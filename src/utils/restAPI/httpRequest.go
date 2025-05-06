package restAPI

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strings"
	"time"

	"omnenest-backend/src/constants"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/models"
	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/tracer"

	"github.com/bytedance/sonic"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	otelHttp "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// HttpRequest makes a HTTP request to the given base URL using the provided request payload.
func HttpRequest(ctx context.Context, baseUrl string, requestPayload *models.Request, header *http.Header, appName string) (*http.Response, error) {
	ctx, span := tracer.AddToSpan(ctx, "HttpRequest")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	var req *http.Request
	var err error
	applicationConfig := configs.GetApplicationConfig()
	enableSSL := applicationConfig.AppConfig.EnableSSL
	client := &http.Client{
		Transport: otelHttp.NewTransport(
			&http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: enableSSL},
			},
			otelHttp.WithClientTrace(func(ctx context.Context) *httptrace.ClientTrace {
				return otelhttptrace.NewClientTrace(ctx)
			}),
			otelHttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
				return r.URL.Path
			}),
			otelHttp.WithServerName(appName),
		),
		Timeout: time.Second * constants.OLTPHttpTimeoutInSec,
	}
	if requestPayload.Request != nil {
		requestData, err := sonic.Marshal(requestPayload.Request)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequestWithContext(ctx, requestPayload.Method, baseUrl, bytes.NewBuffer(requestData))
		if err != nil {
			return nil, err
		}

		header.Set(genericConstants.ContentType, genericConstants.ApplicationJSONTypeConfig)
	} else {
		req, err = http.NewRequestWithContext(ctx, requestPayload.Method, baseUrl, nil)
		if err != nil {
			return nil, err
		}
	}

	if header != nil {
		req.Header = *header
	}
	return client.Do(req)
}

// EncodeQueryParams encodes the given URL parameters into a query string.
func EncodeQueryParams(params url.Values) string {
	var encodedParams []string
	for key, values := range params {
		for _, value := range values {
			encodedParams = append(encodedParams, fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(value)))
		}
	}
	return strings.Join(encodedParams, "&")
}
