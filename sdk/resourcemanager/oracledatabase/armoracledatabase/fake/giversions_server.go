// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
	"net/http"
	"net/url"
	"regexp"
)

// GiVersionsServer is a fake server for instances of the armoracledatabase.GiVersionsClient type.
type GiVersionsServer struct {
	// Get is the fake for method GiVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, giversionname string, options *armoracledatabase.GiVersionsClientGetOptions) (resp azfake.Responder[armoracledatabase.GiVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByLocationPager is the fake for method GiVersionsClient.NewListByLocationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByLocationPager func(location string, options *armoracledatabase.GiVersionsClientListByLocationOptions) (resp azfake.PagerResponder[armoracledatabase.GiVersionsClientListByLocationResponse])
}

// NewGiVersionsServerTransport creates a new instance of GiVersionsServerTransport with the provided implementation.
// The returned GiVersionsServerTransport instance is connected to an instance of armoracledatabase.GiVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGiVersionsServerTransport(srv *GiVersionsServer) *GiVersionsServerTransport {
	return &GiVersionsServerTransport{
		srv:                    srv,
		newListByLocationPager: newTracker[azfake.PagerResponder[armoracledatabase.GiVersionsClientListByLocationResponse]](),
	}
}

// GiVersionsServerTransport connects instances of armoracledatabase.GiVersionsClient to instances of GiVersionsServer.
// Don't use this type directly, use NewGiVersionsServerTransport instead.
type GiVersionsServerTransport struct {
	srv                    *GiVersionsServer
	newListByLocationPager *tracker[azfake.PagerResponder[armoracledatabase.GiVersionsClientListByLocationResponse]]
}

// Do implements the policy.Transporter interface for GiVersionsServerTransport.
func (g *GiVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GiVersionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if giVersionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = giVersionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "GiVersionsClient.Get":
				res.resp, res.err = g.dispatchGet(req)
			case "GiVersionsClient.NewListByLocationPager":
				res.resp, res.err = g.dispatchNewListByLocationPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (g *GiVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/giVersions/(?P<giversionname>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	giversionnameParam, err := url.PathUnescape(matches[regex.SubexpIndex("giversionname")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), locationParam, giversionnameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GiVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GiVersionsServerTransport) dispatchNewListByLocationPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByLocationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByLocationPager not implemented")}
	}
	newListByLocationPager := g.newListByLocationPager.get(req)
	if newListByLocationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/giVersions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		shapeUnescaped, err := url.QueryUnescape(qp.Get("shape"))
		if err != nil {
			return nil, err
		}
		shapeParam := getOptional(armoracledatabase.SystemShapes(shapeUnescaped))
		zoneUnescaped, err := url.QueryUnescape(qp.Get("zone"))
		if err != nil {
			return nil, err
		}
		zoneParam := getOptional(zoneUnescaped)
		var options *armoracledatabase.GiVersionsClientListByLocationOptions
		if shapeParam != nil || zoneParam != nil {
			options = &armoracledatabase.GiVersionsClientListByLocationOptions{
				Shape: shapeParam,
				Zone:  zoneParam,
			}
		}
		resp := g.srv.NewListByLocationPager(locationParam, options)
		newListByLocationPager = &resp
		g.newListByLocationPager.add(req, newListByLocationPager)
		server.PagerResponderInjectNextLinks(newListByLocationPager, req, func(page *armoracledatabase.GiVersionsClientListByLocationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByLocationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByLocationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByLocationPager) {
		g.newListByLocationPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to GiVersionsServerTransport
var giVersionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
