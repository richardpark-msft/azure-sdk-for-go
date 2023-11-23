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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v3"
	"net/http"
	"net/url"
	"regexp"
)

// SitesServer is a fake server for instances of the armmobilenetwork.SitesClient type.
type SitesServer struct {
	// BeginCreateOrUpdate is the fake for method SitesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters armmobilenetwork.Site, options *armmobilenetwork.SitesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmobilenetwork.SitesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SitesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, options *armmobilenetwork.SitesClientBeginDeleteOptions) (resp azfake.PollerResponder[armmobilenetwork.SitesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDeletePacketCore is the fake for method SitesClient.BeginDeletePacketCore
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDeletePacketCore func(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters armmobilenetwork.SiteDeletePacketCore, options *armmobilenetwork.SitesClientBeginDeletePacketCoreOptions) (resp azfake.PollerResponder[armmobilenetwork.SitesClientDeletePacketCoreResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SitesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, options *armmobilenetwork.SitesClientGetOptions) (resp azfake.Responder[armmobilenetwork.SitesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByMobileNetworkPager is the fake for method SitesClient.NewListByMobileNetworkPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByMobileNetworkPager func(resourceGroupName string, mobileNetworkName string, options *armmobilenetwork.SitesClientListByMobileNetworkOptions) (resp azfake.PagerResponder[armmobilenetwork.SitesClientListByMobileNetworkResponse])

	// UpdateTags is the fake for method SitesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters armmobilenetwork.TagsObject, options *armmobilenetwork.SitesClientUpdateTagsOptions) (resp azfake.Responder[armmobilenetwork.SitesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewSitesServerTransport creates a new instance of SitesServerTransport with the provided implementation.
// The returned SitesServerTransport instance is connected to an instance of armmobilenetwork.SitesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSitesServerTransport(srv *SitesServer) *SitesServerTransport {
	return &SitesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armmobilenetwork.SitesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armmobilenetwork.SitesClientDeleteResponse]](),
		beginDeletePacketCore:       newTracker[azfake.PollerResponder[armmobilenetwork.SitesClientDeletePacketCoreResponse]](),
		newListByMobileNetworkPager: newTracker[azfake.PagerResponder[armmobilenetwork.SitesClientListByMobileNetworkResponse]](),
	}
}

// SitesServerTransport connects instances of armmobilenetwork.SitesClient to instances of SitesServer.
// Don't use this type directly, use NewSitesServerTransport instead.
type SitesServerTransport struct {
	srv                         *SitesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armmobilenetwork.SitesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armmobilenetwork.SitesClientDeleteResponse]]
	beginDeletePacketCore       *tracker[azfake.PollerResponder[armmobilenetwork.SitesClientDeletePacketCoreResponse]]
	newListByMobileNetworkPager *tracker[azfake.PagerResponder[armmobilenetwork.SitesClientListByMobileNetworkResponse]]
}

// Do implements the policy.Transporter interface for SitesServerTransport.
func (s *SitesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SitesClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "SitesClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SitesClient.BeginDeletePacketCore":
		resp, err = s.dispatchBeginDeletePacketCore(req)
	case "SitesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SitesClient.NewListByMobileNetworkPager":
		resp, err = s.dispatchNewListByMobileNetworkPager(req)
	case "SitesClient.UpdateTags":
		resp, err = s.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SitesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobileNetwork/mobileNetworks/(?P<mobileNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmobilenetwork.Site](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mobileNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mobileNetworkName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, mobileNetworkNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *SitesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobileNetwork/mobileNetworks/(?P<mobileNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mobileNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mobileNetworkName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, mobileNetworkNameParam, siteNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SitesServerTransport) dispatchBeginDeletePacketCore(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDeletePacketCore == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeletePacketCore not implemented")}
	}
	beginDeletePacketCore := s.beginDeletePacketCore.get(req)
	if beginDeletePacketCore == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobileNetwork/mobileNetworks/(?P<mobileNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletePacketCore`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmobilenetwork.SiteDeletePacketCore](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mobileNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mobileNetworkName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDeletePacketCore(req.Context(), resourceGroupNameParam, mobileNetworkNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeletePacketCore = &respr
		s.beginDeletePacketCore.add(req, beginDeletePacketCore)
	}

	resp, err := server.PollerResponderNext(beginDeletePacketCore, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginDeletePacketCore.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeletePacketCore) {
		s.beginDeletePacketCore.remove(req)
	}

	return resp, nil
}

func (s *SitesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobileNetwork/mobileNetworks/(?P<mobileNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	mobileNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mobileNetworkName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, mobileNetworkNameParam, siteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Site, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SitesServerTransport) dispatchNewListByMobileNetworkPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByMobileNetworkPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByMobileNetworkPager not implemented")}
	}
	newListByMobileNetworkPager := s.newListByMobileNetworkPager.get(req)
	if newListByMobileNetworkPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobileNetwork/mobileNetworks/(?P<mobileNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mobileNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mobileNetworkName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByMobileNetworkPager(resourceGroupNameParam, mobileNetworkNameParam, nil)
		newListByMobileNetworkPager = &resp
		s.newListByMobileNetworkPager.add(req, newListByMobileNetworkPager)
		server.PagerResponderInjectNextLinks(newListByMobileNetworkPager, req, func(page *armmobilenetwork.SitesClientListByMobileNetworkResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByMobileNetworkPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByMobileNetworkPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByMobileNetworkPager) {
		s.newListByMobileNetworkPager.remove(req)
	}
	return resp, nil
}

func (s *SitesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if s.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobileNetwork/mobileNetworks/(?P<mobileNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmobilenetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	mobileNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mobileNetworkName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.UpdateTags(req.Context(), resourceGroupNameParam, mobileNetworkNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Site, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}