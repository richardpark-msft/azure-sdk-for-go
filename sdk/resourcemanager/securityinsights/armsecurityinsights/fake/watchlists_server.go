//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
	"net/http"
	"net/url"
	"regexp"
)

// WatchlistsServer is a fake server for instances of the armsecurityinsights.WatchlistsClient type.
type WatchlistsServer struct {
	// CreateOrUpdate is the fake for method WatchlistsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlist armsecurityinsights.Watchlist, options *armsecurityinsights.WatchlistsClientCreateOrUpdateOptions) (resp azfake.Responder[armsecurityinsights.WatchlistsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method WatchlistsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, options *armsecurityinsights.WatchlistsClientDeleteOptions) (resp azfake.Responder[armsecurityinsights.WatchlistsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WatchlistsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, options *armsecurityinsights.WatchlistsClientGetOptions) (resp azfake.Responder[armsecurityinsights.WatchlistsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WatchlistsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armsecurityinsights.WatchlistsClientListOptions) (resp azfake.PagerResponder[armsecurityinsights.WatchlistsClientListResponse])
}

// NewWatchlistsServerTransport creates a new instance of WatchlistsServerTransport with the provided implementation.
// The returned WatchlistsServerTransport instance is connected to an instance of armsecurityinsights.WatchlistsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWatchlistsServerTransport(srv *WatchlistsServer) *WatchlistsServerTransport {
	return &WatchlistsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurityinsights.WatchlistsClientListResponse]](),
	}
}

// WatchlistsServerTransport connects instances of armsecurityinsights.WatchlistsClient to instances of WatchlistsServer.
// Don't use this type directly, use NewWatchlistsServerTransport instead.
type WatchlistsServerTransport struct {
	srv          *WatchlistsServer
	newListPager *tracker[azfake.PagerResponder[armsecurityinsights.WatchlistsClientListResponse]]
}

// Do implements the policy.Transporter interface for WatchlistsServerTransport.
func (w *WatchlistsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WatchlistsClient.CreateOrUpdate":
		resp, err = w.dispatchCreateOrUpdate(req)
	case "WatchlistsClient.Delete":
		resp, err = w.dispatchDelete(req)
	case "WatchlistsClient.Get":
		resp, err = w.dispatchGet(req)
	case "WatchlistsClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WatchlistsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/watchlists/(?P<watchlistAlias>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.Watchlist](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	watchlistAliasParam, err := url.PathUnescape(matches[regex.SubexpIndex("watchlistAlias")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, watchlistAliasParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Watchlist, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).AzureAsyncOperation; val != nil {
		resp.Header.Set("Azure-AsyncOperation", *val)
	}
	return resp, nil
}

func (w *WatchlistsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if w.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/watchlists/(?P<watchlistAlias>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	watchlistAliasParam, err := url.PathUnescape(matches[regex.SubexpIndex("watchlistAlias")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, watchlistAliasParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).AzureAsyncOperation; val != nil {
		resp.Header.Set("Azure-AsyncOperation", *val)
	}
	return resp, nil
}

func (w *WatchlistsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/watchlists/(?P<watchlistAlias>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	watchlistAliasParam, err := url.PathUnescape(matches[regex.SubexpIndex("watchlistAlias")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, watchlistAliasParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Watchlist, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WatchlistsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/watchlists`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armsecurityinsights.WatchlistsClientListOptions
		if skipTokenParam != nil {
			options = &armsecurityinsights.WatchlistsClientListOptions{
				SkipToken: skipTokenParam,
			}
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, options)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurityinsights.WatchlistsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}
