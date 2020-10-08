package automanage

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

// ConfigurationProfilePreferencesClient is the automanage Client
type ConfigurationProfilePreferencesClient struct {
	BaseClient
}

// NewConfigurationProfilePreferencesClient creates an instance of the ConfigurationProfilePreferencesClient client.
func NewConfigurationProfilePreferencesClient(subscriptionID string) ConfigurationProfilePreferencesClient {
	return NewConfigurationProfilePreferencesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationProfilePreferencesClientWithBaseURI creates an instance of the ConfigurationProfilePreferencesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewConfigurationProfilePreferencesClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationProfilePreferencesClient {
	return ConfigurationProfilePreferencesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a configuration profile preference
// Parameters:
// configurationProfilePreferenceName - name of the configuration profile preference.
// resourceGroupName - the resource group name.
// parameters - parameters supplied to create or update configuration profile preference.
func (client ConfigurationProfilePreferencesClient) CreateOrUpdate(ctx context.Context, configurationProfilePreferenceName string, resourceGroupName string, parameters ConfigurationProfilePreference) (result ConfigurationProfilePreference, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfilePreferencesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfilePreferencesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, configurationProfilePreferenceName, resourceGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConfigurationProfilePreferencesClient) CreateOrUpdatePreparer(ctx context.Context, configurationProfilePreferenceName string, resourceGroupName string, parameters ConfigurationProfilePreference) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationProfilePreferenceName": autorest.Encode("path", configurationProfilePreferenceName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfilePreferences/{configurationProfilePreferenceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfilePreferencesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConfigurationProfilePreferencesClient) CreateOrUpdateResponder(resp *http.Response) (result ConfigurationProfilePreference, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a configuration profile preference
// Parameters:
// resourceGroupName - the resource group name.
// configurationProfilePreferenceName - name of the configuration profile preference
func (client ConfigurationProfilePreferencesClient) Delete(ctx context.Context, resourceGroupName string, configurationProfilePreferenceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfilePreferencesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfilePreferencesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, configurationProfilePreferenceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConfigurationProfilePreferencesClient) DeletePreparer(ctx context.Context, resourceGroupName string, configurationProfilePreferenceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationProfilePreferenceName": autorest.Encode("path", configurationProfilePreferenceName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfilePreferences/{configurationProfilePreferenceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfilePreferencesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConfigurationProfilePreferencesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get information about a configuration profile preference
// Parameters:
// configurationProfilePreferenceName - the configuration profile preference name.
// resourceGroupName - the resource group name.
func (client ConfigurationProfilePreferencesClient) Get(ctx context.Context, configurationProfilePreferenceName string, resourceGroupName string) (result ConfigurationProfilePreference, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfilePreferencesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfilePreferencesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, configurationProfilePreferenceName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigurationProfilePreferencesClient) GetPreparer(ctx context.Context, configurationProfilePreferenceName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationProfilePreferenceName": autorest.Encode("path", configurationProfilePreferenceName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfilePreferences/{configurationProfilePreferenceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfilePreferencesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigurationProfilePreferencesClient) GetResponder(resp *http.Response) (result ConfigurationProfilePreference, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup retrieve a list of configuration profile preferences within a given resource group
// Parameters:
// resourceGroupName - the resource group name.
func (client ConfigurationProfilePreferencesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ConfigurationProfilePreferenceList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfilePreferencesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfilePreferencesClient", "ListByResourceGroup", err.Error())
	}

	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ConfigurationProfilePreferencesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfilePreferences", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfilePreferencesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ConfigurationProfilePreferencesClient) ListByResourceGroupResponder(resp *http.Response) (result ConfigurationProfilePreferenceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscription retrieve a list of configuration profile preferences within a subscription
func (client ConfigurationProfilePreferencesClient) ListBySubscription(ctx context.Context) (result ConfigurationProfilePreferenceList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfilePreferencesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ConfigurationProfilePreferencesClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Automanage/configurationProfilePreferences", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfilePreferencesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ConfigurationProfilePreferencesClient) ListBySubscriptionResponder(resp *http.Response) (result ConfigurationProfilePreferenceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a configuration profile preference
// Parameters:
// configurationProfilePreferenceName - name of the configuration profile preference.
// resourceGroupName - the resource group name.
// parameters - parameters supplied to create or update configuration profile preference.
func (client ConfigurationProfilePreferencesClient) Update(ctx context.Context, configurationProfilePreferenceName string, resourceGroupName string, parameters ConfigurationProfilePreference) (result ConfigurationProfilePreference, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfilePreferencesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfilePreferencesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, configurationProfilePreferenceName, resourceGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilePreferencesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ConfigurationProfilePreferencesClient) UpdatePreparer(ctx context.Context, configurationProfilePreferenceName string, resourceGroupName string, parameters ConfigurationProfilePreference) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationProfilePreferenceName": autorest.Encode("path", configurationProfilePreferenceName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfilePreferences/{configurationProfilePreferenceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfilePreferencesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ConfigurationProfilePreferencesClient) UpdateResponder(resp *http.Response) (result ConfigurationProfilePreference, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}