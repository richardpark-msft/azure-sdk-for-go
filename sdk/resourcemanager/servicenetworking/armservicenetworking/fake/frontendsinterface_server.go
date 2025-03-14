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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
	"net/http"
	"net/url"
	"regexp"
)

// FrontendsInterfaceServer is a fake server for instances of the armservicenetworking.FrontendsInterfaceClient type.
type FrontendsInterfaceServer struct {
	// BeginCreateOrUpdate is the fake for method FrontendsInterfaceClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, trafficControllerName string, frontendName string, resource armservicenetworking.Frontend, options *armservicenetworking.FrontendsInterfaceClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armservicenetworking.FrontendsInterfaceClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FrontendsInterfaceClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, trafficControllerName string, frontendName string, options *armservicenetworking.FrontendsInterfaceClientBeginDeleteOptions) (resp azfake.PollerResponder[armservicenetworking.FrontendsInterfaceClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FrontendsInterfaceClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, trafficControllerName string, frontendName string, options *armservicenetworking.FrontendsInterfaceClientGetOptions) (resp azfake.Responder[armservicenetworking.FrontendsInterfaceClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByTrafficControllerPager is the fake for method FrontendsInterfaceClient.NewListByTrafficControllerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByTrafficControllerPager func(resourceGroupName string, trafficControllerName string, options *armservicenetworking.FrontendsInterfaceClientListByTrafficControllerOptions) (resp azfake.PagerResponder[armservicenetworking.FrontendsInterfaceClientListByTrafficControllerResponse])

	// Update is the fake for method FrontendsInterfaceClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, trafficControllerName string, frontendName string, properties armservicenetworking.FrontendUpdate, options *armservicenetworking.FrontendsInterfaceClientUpdateOptions) (resp azfake.Responder[armservicenetworking.FrontendsInterfaceClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewFrontendsInterfaceServerTransport creates a new instance of FrontendsInterfaceServerTransport with the provided implementation.
// The returned FrontendsInterfaceServerTransport instance is connected to an instance of armservicenetworking.FrontendsInterfaceClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFrontendsInterfaceServerTransport(srv *FrontendsInterfaceServer) *FrontendsInterfaceServerTransport {
	return &FrontendsInterfaceServerTransport{
		srv:                             srv,
		beginCreateOrUpdate:             newTracker[azfake.PollerResponder[armservicenetworking.FrontendsInterfaceClientCreateOrUpdateResponse]](),
		beginDelete:                     newTracker[azfake.PollerResponder[armservicenetworking.FrontendsInterfaceClientDeleteResponse]](),
		newListByTrafficControllerPager: newTracker[azfake.PagerResponder[armservicenetworking.FrontendsInterfaceClientListByTrafficControllerResponse]](),
	}
}

// FrontendsInterfaceServerTransport connects instances of armservicenetworking.FrontendsInterfaceClient to instances of FrontendsInterfaceServer.
// Don't use this type directly, use NewFrontendsInterfaceServerTransport instead.
type FrontendsInterfaceServerTransport struct {
	srv                             *FrontendsInterfaceServer
	beginCreateOrUpdate             *tracker[azfake.PollerResponder[armservicenetworking.FrontendsInterfaceClientCreateOrUpdateResponse]]
	beginDelete                     *tracker[azfake.PollerResponder[armservicenetworking.FrontendsInterfaceClientDeleteResponse]]
	newListByTrafficControllerPager *tracker[azfake.PagerResponder[armservicenetworking.FrontendsInterfaceClientListByTrafficControllerResponse]]
}

// Do implements the policy.Transporter interface for FrontendsInterfaceServerTransport.
func (f *FrontendsInterfaceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FrontendsInterfaceServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if frontendsInterfaceServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = frontendsInterfaceServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "FrontendsInterfaceClient.BeginCreateOrUpdate":
				res.resp, res.err = f.dispatchBeginCreateOrUpdate(req)
			case "FrontendsInterfaceClient.BeginDelete":
				res.resp, res.err = f.dispatchBeginDelete(req)
			case "FrontendsInterfaceClient.Get":
				res.resp, res.err = f.dispatchGet(req)
			case "FrontendsInterfaceClient.NewListByTrafficControllerPager":
				res.resp, res.err = f.dispatchNewListByTrafficControllerPager(req)
			case "FrontendsInterfaceClient.Update":
				res.resp, res.err = f.dispatchUpdate(req)
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

func (f *FrontendsInterfaceServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := f.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceNetworking/trafficControllers/(?P<trafficControllerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontends/(?P<frontendName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armservicenetworking.Frontend](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		trafficControllerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("trafficControllerName")])
		if err != nil {
			return nil, err
		}
		frontendNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontendName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, trafficControllerNameParam, frontendNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		f.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		f.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		f.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (f *FrontendsInterfaceServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceNetworking/trafficControllers/(?P<trafficControllerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontends/(?P<frontendName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		trafficControllerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("trafficControllerName")])
		if err != nil {
			return nil, err
		}
		frontendNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontendName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameParam, trafficControllerNameParam, frontendNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		f.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FrontendsInterfaceServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceNetworking/trafficControllers/(?P<trafficControllerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontends/(?P<frontendName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	trafficControllerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("trafficControllerName")])
	if err != nil {
		return nil, err
	}
	frontendNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontendName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, trafficControllerNameParam, frontendNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Frontend, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FrontendsInterfaceServerTransport) dispatchNewListByTrafficControllerPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListByTrafficControllerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByTrafficControllerPager not implemented")}
	}
	newListByTrafficControllerPager := f.newListByTrafficControllerPager.get(req)
	if newListByTrafficControllerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceNetworking/trafficControllers/(?P<trafficControllerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontends`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		trafficControllerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("trafficControllerName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListByTrafficControllerPager(resourceGroupNameParam, trafficControllerNameParam, nil)
		newListByTrafficControllerPager = &resp
		f.newListByTrafficControllerPager.add(req, newListByTrafficControllerPager)
		server.PagerResponderInjectNextLinks(newListByTrafficControllerPager, req, func(page *armservicenetworking.FrontendsInterfaceClientListByTrafficControllerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByTrafficControllerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListByTrafficControllerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByTrafficControllerPager) {
		f.newListByTrafficControllerPager.remove(req)
	}
	return resp, nil
}

func (f *FrontendsInterfaceServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceNetworking/trafficControllers/(?P<trafficControllerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontends/(?P<frontendName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armservicenetworking.FrontendUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	trafficControllerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("trafficControllerName")])
	if err != nil {
		return nil, err
	}
	frontendNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontendName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Update(req.Context(), resourceGroupNameParam, trafficControllerNameParam, frontendNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Frontend, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to FrontendsInterfaceServerTransport
var frontendsInterfaceServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
