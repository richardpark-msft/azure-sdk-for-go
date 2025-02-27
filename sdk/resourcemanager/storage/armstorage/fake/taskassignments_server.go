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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// TaskAssignmentsServer is a fake server for instances of the armstorage.TaskAssignmentsClient type.
type TaskAssignmentsServer struct {
	// BeginCreate is the fake for method TaskAssignmentsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, parameters armstorage.TaskAssignment, options *armstorage.TaskAssignmentsClientBeginCreateOptions) (resp azfake.PollerResponder[armstorage.TaskAssignmentsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method TaskAssignmentsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, options *armstorage.TaskAssignmentsClientBeginDeleteOptions) (resp azfake.PollerResponder[armstorage.TaskAssignmentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TaskAssignmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, options *armstorage.TaskAssignmentsClientGetOptions) (resp azfake.Responder[armstorage.TaskAssignmentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TaskAssignmentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armstorage.TaskAssignmentsClientListOptions) (resp azfake.PagerResponder[armstorage.TaskAssignmentsClientListResponse])

	// BeginUpdate is the fake for method TaskAssignmentsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, parameters armstorage.TaskAssignmentUpdateParameters, options *armstorage.TaskAssignmentsClientBeginUpdateOptions) (resp azfake.PollerResponder[armstorage.TaskAssignmentsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewTaskAssignmentsServerTransport creates a new instance of TaskAssignmentsServerTransport with the provided implementation.
// The returned TaskAssignmentsServerTransport instance is connected to an instance of armstorage.TaskAssignmentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTaskAssignmentsServerTransport(srv *TaskAssignmentsServer) *TaskAssignmentsServerTransport {
	return &TaskAssignmentsServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armstorage.TaskAssignmentsClientCreateResponse]](),
		beginDelete:  newTracker[azfake.PollerResponder[armstorage.TaskAssignmentsClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armstorage.TaskAssignmentsClientListResponse]](),
		beginUpdate:  newTracker[azfake.PollerResponder[armstorage.TaskAssignmentsClientUpdateResponse]](),
	}
}

// TaskAssignmentsServerTransport connects instances of armstorage.TaskAssignmentsClient to instances of TaskAssignmentsServer.
// Don't use this type directly, use NewTaskAssignmentsServerTransport instead.
type TaskAssignmentsServerTransport struct {
	srv          *TaskAssignmentsServer
	beginCreate  *tracker[azfake.PollerResponder[armstorage.TaskAssignmentsClientCreateResponse]]
	beginDelete  *tracker[azfake.PollerResponder[armstorage.TaskAssignmentsClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armstorage.TaskAssignmentsClientListResponse]]
	beginUpdate  *tracker[azfake.PollerResponder[armstorage.TaskAssignmentsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for TaskAssignmentsServerTransport.
func (t *TaskAssignmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TaskAssignmentsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if taskAssignmentsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = taskAssignmentsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TaskAssignmentsClient.BeginCreate":
				res.resp, res.err = t.dispatchBeginCreate(req)
			case "TaskAssignmentsClient.BeginDelete":
				res.resp, res.err = t.dispatchBeginDelete(req)
			case "TaskAssignmentsClient.Get":
				res.resp, res.err = t.dispatchGet(req)
			case "TaskAssignmentsClient.NewListPager":
				res.resp, res.err = t.dispatchNewListPager(req)
			case "TaskAssignmentsClient.BeginUpdate":
				res.resp, res.err = t.dispatchBeginUpdate(req)
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

func (t *TaskAssignmentsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if t.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := t.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageTaskAssignments/(?P<storageTaskAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorage.TaskAssignment](req)
		if err != nil {
			return nil, err
		}
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
		respr, errRespr := t.srv.BeginCreate(req.Context(), resourceGroupNameParam, accountNameParam, storageTaskAssignmentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		t.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		t.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		t.beginCreate.remove(req)
	}

	return resp, nil
}

func (t *TaskAssignmentsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if t.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := t.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageTaskAssignments/(?P<storageTaskAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
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
		respr, errRespr := t.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, storageTaskAssignmentNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		t.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		t.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		t.beginDelete.remove(req)
	}

	return resp, nil
}

func (t *TaskAssignmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageTaskAssignments/(?P<storageTaskAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
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
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, storageTaskAssignmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TaskAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TaskAssignmentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageTaskAssignments`
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
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
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
		var options *armstorage.TaskAssignmentsClientListOptions
		if maxpagesizeParam != nil {
			options = &armstorage.TaskAssignmentsClientListOptions{
				Maxpagesize: maxpagesizeParam,
			}
		}
		resp := t.srv.NewListPager(resourceGroupNameParam, accountNameParam, options)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armstorage.TaskAssignmentsClientListResponse, createLink func() string) {
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

func (t *TaskAssignmentsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := t.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageTaskAssignments/(?P<storageTaskAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorage.TaskAssignmentUpdateParameters](req)
		if err != nil {
			return nil, err
		}
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
		respr, errRespr := t.srv.BeginUpdate(req.Context(), resourceGroupNameParam, accountNameParam, storageTaskAssignmentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		t.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		t.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		t.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to TaskAssignmentsServerTransport
var taskAssignmentsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
