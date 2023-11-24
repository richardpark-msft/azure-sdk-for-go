//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// MonitorServer is a fake server for instances of the armelastic.MonitorClient type.
type MonitorServer struct {
	// BeginUpgrade is the fake for method MonitorClient.BeginUpgrade
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginUpgrade func(ctx context.Context, resourceGroupName string, monitorName string, options *armelastic.MonitorClientBeginUpgradeOptions) (resp azfake.PollerResponder[armelastic.MonitorClientUpgradeResponse], errResp azfake.ErrorResponder)
}

// NewMonitorServerTransport creates a new instance of MonitorServerTransport with the provided implementation.
// The returned MonitorServerTransport instance is connected to an instance of armelastic.MonitorClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMonitorServerTransport(srv *MonitorServer) *MonitorServerTransport {
	return &MonitorServerTransport{
		srv:          srv,
		beginUpgrade: newTracker[azfake.PollerResponder[armelastic.MonitorClientUpgradeResponse]](),
	}
}

// MonitorServerTransport connects instances of armelastic.MonitorClient to instances of MonitorServer.
// Don't use this type directly, use NewMonitorServerTransport instead.
type MonitorServerTransport struct {
	srv          *MonitorServer
	beginUpgrade *tracker[azfake.PollerResponder[armelastic.MonitorClientUpgradeResponse]]
}

// Do implements the policy.Transporter interface for MonitorServerTransport.
func (m *MonitorServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "MonitorClient.BeginUpgrade":
		resp, err = m.dispatchBeginUpgrade(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *MonitorServerTransport) dispatchBeginUpgrade(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpgrade == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpgrade not implemented")}
	}
	beginUpgrade := m.beginUpgrade.get(req)
	if beginUpgrade == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Elastic/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/upgrade`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armelastic.MonitorUpgrade](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
		if err != nil {
			return nil, err
		}
		var options *armelastic.MonitorClientBeginUpgradeOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armelastic.MonitorClientBeginUpgradeOptions{
				Body: &body,
			}
		}
		respr, errRespr := m.srv.BeginUpgrade(req.Context(), resourceGroupNameParam, monitorNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpgrade = &respr
		m.beginUpgrade.add(req, beginUpgrade)
	}

	resp, err := server.PollerResponderNext(beginUpgrade, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		m.beginUpgrade.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpgrade) {
		m.beginUpgrade.remove(req)
	}

	return resp, nil
}
