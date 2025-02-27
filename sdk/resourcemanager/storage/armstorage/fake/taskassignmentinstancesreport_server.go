// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// TaskAssignmentInstancesReportServer is a fake server for instances of the armstorage.TaskAssignmentInstancesReportClient type.
type TaskAssignmentInstancesReportServer struct {
	// NewListPager is the fake for method TaskAssignmentInstancesReportClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, storageTaskAssignmentName string, options *armstorage.TaskAssignmentInstancesReportClientListOptions) (resp azfake.PagerResponder[armstorage.TaskAssignmentInstancesReportClientListResponse])
}

// NewTaskAssignmentInstancesReportServerTransport creates a new instance of TaskAssignmentInstancesReportServerTransport with the provided implementation.
// The returned TaskAssignmentInstancesReportServerTransport instance is connected to an instance of armstorage.TaskAssignmentInstancesReportClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTaskAssignmentInstancesReportServerTransport(srv *TaskAssignmentInstancesReportServer) *TaskAssignmentInstancesReportServerTransport {
	return &TaskAssignmentInstancesReportServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armstorage.TaskAssignmentInstancesReportClientListResponse]](),
	}
}

// TaskAssignmentInstancesReportServerTransport connects instances of armstorage.TaskAssignmentInstancesReportClient to instances of TaskAssignmentInstancesReportServer.
// Don't use this type directly, use NewTaskAssignmentInstancesReportServerTransport instead.
type TaskAssignmentInstancesReportServerTransport struct {
	srv          *TaskAssignmentInstancesReportServer
	newListPager *tracker[azfake.PagerResponder[armstorage.TaskAssignmentInstancesReportClientListResponse]]
}

// Do implements the policy.Transporter interface for TaskAssignmentInstancesReportServerTransport.
func (t *TaskAssignmentInstancesReportServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TaskAssignmentInstancesReportServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if taskAssignmentInstancesReportServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = taskAssignmentInstancesReportServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TaskAssignmentInstancesReportClient.NewListPager":
				res.resp, res.err = t.dispatchNewListPager(req)
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

func (t *TaskAssignmentInstancesReportServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageTaskAssignments/(?P<storageTaskAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		storageTaskAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageTaskAssignmentName")])
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("$maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam, err := parseOptional(maxpagesizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armstorage.TaskAssignmentInstancesReportClientListOptions
		if maxpagesizeParam != nil || filterParam != nil {
			options = &armstorage.TaskAssignmentInstancesReportClientListOptions{
				Maxpagesize: maxpagesizeParam,
				Filter:      filterParam,
			}
		}
		resp := t.srv.NewListPager(resourceGroupNameParam, accountNameParam, storageTaskAssignmentNameParam, options)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armstorage.TaskAssignmentInstancesReportClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		t.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to TaskAssignmentInstancesReportServerTransport
var taskAssignmentInstancesReportServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
