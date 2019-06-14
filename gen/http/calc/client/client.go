// Code generated by goa v3.0.2, DO NOT EDIT.
//
// calc client HTTP transport
//
// Command:
// $ goa gen goa-ddb/design

package client

import (
	"context"
	calc "goa-ddb/gen/calc"
	"mime/multipart"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the calc service endpoint HTTP clients.
type Client struct {
	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// Addresume Doer is the HTTP client used to make requests to the addresume
	// endpoint.
	AddresumeDoer goahttp.Doer

	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// CalcAddresumeEncoderFunc is the type to encode multipart request for the
// "calc" service "addresume" endpoint.
type CalcAddresumeEncoderFunc func(*multipart.Writer, []*calc.Resume) error

// NewClient instantiates HTTP clients for all the calc service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AddDoer:             doer,
		AddresumeDoer:       doer,
		ListDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Add returns an endpoint that makes HTTP requests to the calc service add
// server.
func (c *Client) Add() goa.Endpoint {
	var (
		decodeResponse = DecodeAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "add", err)
		}
		return decodeResponse(resp)
	}
}

// Addresume returns an endpoint that makes HTTP requests to the calc service
// addresume server.
func (c *Client) Addresume(calcAddresumeEncoderFn CalcAddresumeEncoderFunc) goa.Endpoint {
	var (
		encodeRequest  = EncodeAddresumeRequest(NewCalcAddresumeEncoder(calcAddresumeEncoderFn))
		decodeResponse = DecodeAddresumeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddresumeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddresumeDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "addresume", err)
		}
		return decodeResponse(resp)
	}
}

// List returns an endpoint that makes HTTP requests to the calc service list
// server.
func (c *Client) List() goa.Endpoint {
	var (
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "list", err)
		}
		return decodeResponse(resp)
	}
}