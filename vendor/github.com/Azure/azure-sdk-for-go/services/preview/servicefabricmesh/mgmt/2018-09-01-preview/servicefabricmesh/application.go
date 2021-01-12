package servicefabricmesh

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ApplicationClient is the service Fabric Mesh Management Client
type ApplicationClient struct {
	BaseClient
}

// NewApplicationClient creates an instance of the ApplicationClient client.
func NewApplicationClient(subscriptionID string) ApplicationClient {
	return NewApplicationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationClientWithBaseURI creates an instance of the ApplicationClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewApplicationClientWithBaseURI(baseURI string, subscriptionID string) ApplicationClient {
	return ApplicationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates an application resource with the specified name, description and properties. If an application
// resource with the same name exists, then it is updated with the specified description and properties.
// Parameters:
// resourceGroupName - azure resource group name
// applicationResourceName - the identity of the application.
// applicationResourceDescription - description for creating a Application resource.
func (client ApplicationClient) Create(ctx context.Context, resourceGroupName string, applicationResourceName string, applicationResourceDescription ApplicationResourceDescription) (result ApplicationResourceDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: applicationResourceDescription,
			Constraints: []validation.Constraint{{Target: "applicationResourceDescription.ApplicationResourceProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicefabricmesh.ApplicationClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, applicationResourceName, applicationResourceDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ApplicationClient) CreatePreparer(ctx context.Context, resourceGroupName string, applicationResourceName string, applicationResourceDescription ApplicationResourceDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationResourceName": applicationResourceName,
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}", pathParameters),
		autorest.WithJSON(applicationResourceDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ApplicationClient) CreateResponder(resp *http.Response) (result ApplicationResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the application resource identified by the name.
// Parameters:
// resourceGroupName - azure resource group name
// applicationResourceName - the identity of the application.
func (client ApplicationClient) Delete(ctx context.Context, resourceGroupName string, applicationResourceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, applicationResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationClient) DeletePreparer(ctx context.Context, resourceGroupName string, applicationResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationResourceName": applicationResourceName,
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the information about the application resource with the given name. The information include the description
// and other properties of the application.
// Parameters:
// resourceGroupName - azure resource group name
// applicationResourceName - the identity of the application.
func (client ApplicationClient) Get(ctx context.Context, resourceGroupName string, applicationResourceName string) (result ApplicationResourceDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, applicationResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationClient) GetPreparer(ctx context.Context, resourceGroupName string, applicationResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationResourceName": applicationResourceName,
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationClient) GetResponder(resp *http.Response) (result ApplicationResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets the information about all application resources in a given resource group. The information
// include the description and other properties of the Application.
// Parameters:
// resourceGroupName - azure resource group name
func (client ApplicationClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ApplicationResourceDescriptionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.ardl.Response.Response != nil {
				sc = result.ardl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.ardl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.ardl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.ardl.hasNextLink() && result.ardl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ApplicationClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ApplicationClient) ListByResourceGroupResponder(resp *http.Response) (result ApplicationResourceDescriptionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ApplicationClient) listByResourceGroupNextResults(ctx context.Context, lastResults ApplicationResourceDescriptionList) (result ApplicationResourceDescriptionList, err error) {
	req, err := lastResults.applicationResourceDescriptionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ApplicationResourceDescriptionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySubscription gets the information about all application resources in a given resource group. The information
// include the description and other properties of the application.
func (client ApplicationClient) ListBySubscription(ctx context.Context) (result ApplicationResourceDescriptionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.ardl.Response.Response != nil {
				sc = result.ardl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.ardl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.ardl, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.ardl.hasNextLink() && result.ardl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ApplicationClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/applications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ApplicationClient) ListBySubscriptionResponder(resp *http.Response) (result ApplicationResourceDescriptionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client ApplicationClient) listBySubscriptionNextResults(ctx context.Context, lastResults ApplicationResourceDescriptionList) (result ApplicationResourceDescriptionList, err error) {
	req, err := lastResults.applicationResourceDescriptionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ApplicationClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationClient) ListBySubscriptionComplete(ctx context.Context) (result ApplicationResourceDescriptionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}
