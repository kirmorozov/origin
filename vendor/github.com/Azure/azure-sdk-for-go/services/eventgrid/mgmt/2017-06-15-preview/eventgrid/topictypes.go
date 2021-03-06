package eventgrid

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// TopicTypesClient is the azure EventGrid Management Client
type TopicTypesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// NewTopicTypesClient creates an instance of the TopicTypesClient client.
func NewTopicTypesClient(subscriptionID string) TopicTypesClient {
	return NewTopicTypesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// NewTopicTypesClientWithBaseURI creates an instance of the TopicTypesClient client.
func NewTopicTypesClientWithBaseURI(baseURI string, subscriptionID string) TopicTypesClient {
	return TopicTypesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// Get get information about a topic type
//
// topicTypeName is name of the topic type
func (client TopicTypesClient) Get(ctx context.Context, topicTypeName string) (result TopicTypeInfo, err error) {
	req, err := client.GetPreparer(ctx, topicTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// GetPreparer prepares the Get request.
func (client TopicTypesClient) GetPreparer(ctx context.Context, topicTypeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"topicTypeName": autorest.Encode("path", topicTypeName),
	}

	const APIVersion = "2017-06-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TopicTypesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TopicTypesClient) GetResponder(resp *http.Response) (result TopicTypeInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// List list all registered topic types
func (client TopicTypesClient) List(ctx context.Context) (result TopicTypesListResult, err error) {
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// ListPreparer prepares the List request.
func (client TopicTypesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2017-06-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.EventGrid/topicTypes"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TopicTypesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TopicTypesClient) ListResponder(resp *http.Response) (result TopicTypesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// ListEventTypes list event types for a topic type
//
// topicTypeName is name of the topic type
func (client TopicTypesClient) ListEventTypes(ctx context.Context, topicTypeName string) (result EventTypesListResult, err error) {
	req, err := client.ListEventTypesPreparer(ctx, topicTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "ListEventTypes", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListEventTypesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "ListEventTypes", resp, "Failure sending request")
		return
	}

	result, err = client.ListEventTypesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.TopicTypesClient", "ListEventTypes", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// ListEventTypesPreparer prepares the ListEventTypes request.
func (client TopicTypesClient) ListEventTypesPreparer(ctx context.Context, topicTypeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"topicTypeName": autorest.Encode("path", topicTypeName),
	}

	const APIVersion = "2017-06-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// ListEventTypesSender sends the ListEventTypes request. The method will close the
// http.Response Body if it receives an error.
func (client TopicTypesClient) ListEventTypesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2017-06-15-preview/eventgrid instead.
// ListEventTypesResponder handles the response to the ListEventTypes request. The method always
// closes the http.Response Body.
func (client TopicTypesClient) ListEventTypesResponder(resp *http.Response) (result EventTypesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
