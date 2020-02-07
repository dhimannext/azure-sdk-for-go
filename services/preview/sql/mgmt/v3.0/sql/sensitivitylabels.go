package sql

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SensitivityLabelsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type SensitivityLabelsClient struct {
	BaseClient
}

// NewSensitivityLabelsClient creates an instance of the SensitivityLabelsClient client.
func NewSensitivityLabelsClient(subscriptionID string) SensitivityLabelsClient {
	return NewSensitivityLabelsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSensitivityLabelsClientWithBaseURI creates an instance of the SensitivityLabelsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewSensitivityLabelsClientWithBaseURI(baseURI string, subscriptionID string) SensitivityLabelsClient {
	return SensitivityLabelsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the sensitivity label of a given column
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
// parameters - the column sensitivity label resource.
func (client SensitivityLabelsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel) (result SensitivityLabel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SensitivityLabelsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SensitivityLabelsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "current"),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SensitivityLabelsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SensitivityLabelsClient) CreateOrUpdateResponder(resp *http.Response) (result SensitivityLabel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the sensitivity label of a given column
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
func (client SensitivityLabelsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SensitivityLabelsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SensitivityLabelsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "current"),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SensitivityLabelsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SensitivityLabelsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DisableRecommendation disables sensitivity recommendations on a given column
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
func (client SensitivityLabelsClient) DisableRecommendation(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SensitivityLabelsClient.DisableRecommendation")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DisableRecommendationPreparer(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "DisableRecommendation", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableRecommendationSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "DisableRecommendation", resp, "Failure sending request")
		return
	}

	result, err = client.DisableRecommendationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "DisableRecommendation", resp, "Failure responding to request")
	}

	return
}

// DisableRecommendationPreparer prepares the DisableRecommendation request.
func (client SensitivityLabelsClient) DisableRecommendationPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "recommended"),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/disable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableRecommendationSender sends the DisableRecommendation request. The method will close the
// http.Response Body if it receives an error.
func (client SensitivityLabelsClient) DisableRecommendationSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DisableRecommendationResponder handles the response to the DisableRecommendation request. The method always
// closes the http.Response Body.
func (client SensitivityLabelsClient) DisableRecommendationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// EnableRecommendation enables sensitivity recommendations on a given column (recommendations are enabled by default
// on all columns)
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
func (client SensitivityLabelsClient) EnableRecommendation(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SensitivityLabelsClient.EnableRecommendation")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnableRecommendationPreparer(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "EnableRecommendation", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableRecommendationSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "EnableRecommendation", resp, "Failure sending request")
		return
	}

	result, err = client.EnableRecommendationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "EnableRecommendation", resp, "Failure responding to request")
	}

	return
}

// EnableRecommendationPreparer prepares the EnableRecommendation request.
func (client SensitivityLabelsClient) EnableRecommendationPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "recommended"),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/enable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableRecommendationSender sends the EnableRecommendation request. The method will close the
// http.Response Body if it receives an error.
func (client SensitivityLabelsClient) EnableRecommendationSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// EnableRecommendationResponder handles the response to the EnableRecommendation request. The method always
// closes the http.Response Body.
func (client SensitivityLabelsClient) EnableRecommendationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the sensitivity label of a given column
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
// sensitivityLabelSource - the source of the sensitivity label.
func (client SensitivityLabelsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource) (result SensitivityLabel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SensitivityLabelsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName, sensitivityLabelSource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SensitivityLabelsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", sensitivityLabelSource),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SensitivityLabelsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SensitivityLabelsClient) GetResponder(resp *http.Response) (result SensitivityLabel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListCurrentByDatabase gets the sensitivity labels of a given database
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// filter - an OData filter expression that filters elements in the collection.
func (client SensitivityLabelsClient) ListCurrentByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string, filter string) (result SensitivityLabelListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SensitivityLabelsClient.ListCurrentByDatabase")
		defer func() {
			sc := -1
			if result.sllr.Response.Response != nil {
				sc = result.sllr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listCurrentByDatabaseNextResults
	req, err := client.ListCurrentByDatabasePreparer(ctx, resourceGroupName, serverName, databaseName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "ListCurrentByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCurrentByDatabaseSender(req)
	if err != nil {
		result.sllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "ListCurrentByDatabase", resp, "Failure sending request")
		return
	}

	result.sllr, err = client.ListCurrentByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "ListCurrentByDatabase", resp, "Failure responding to request")
	}

	return
}

// ListCurrentByDatabasePreparer prepares the ListCurrentByDatabase request.
func (client SensitivityLabelsClient) ListCurrentByDatabasePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/currentSensitivityLabels", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCurrentByDatabaseSender sends the ListCurrentByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client SensitivityLabelsClient) ListCurrentByDatabaseSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListCurrentByDatabaseResponder handles the response to the ListCurrentByDatabase request. The method always
// closes the http.Response Body.
func (client SensitivityLabelsClient) ListCurrentByDatabaseResponder(resp *http.Response) (result SensitivityLabelListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listCurrentByDatabaseNextResults retrieves the next set of results, if any.
func (client SensitivityLabelsClient) listCurrentByDatabaseNextResults(ctx context.Context, lastResults SensitivityLabelListResult) (result SensitivityLabelListResult, err error) {
	req, err := lastResults.sensitivityLabelListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "listCurrentByDatabaseNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListCurrentByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "listCurrentByDatabaseNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListCurrentByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "listCurrentByDatabaseNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListCurrentByDatabaseComplete enumerates all values, automatically crossing page boundaries as required.
func (client SensitivityLabelsClient) ListCurrentByDatabaseComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, filter string) (result SensitivityLabelListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SensitivityLabelsClient.ListCurrentByDatabase")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListCurrentByDatabase(ctx, resourceGroupName, serverName, databaseName, filter)
	return
}

// ListRecommendedByDatabase gets the sensitivity labels of a given database
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// includeDisabledRecommendations - specifies whether to include disabled recommendations or not.
// filter - an OData filter expression that filters elements in the collection.
func (client SensitivityLabelsClient) ListRecommendedByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string, includeDisabledRecommendations *bool, skipToken string, filter string) (result SensitivityLabelListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SensitivityLabelsClient.ListRecommendedByDatabase")
		defer func() {
			sc := -1
			if result.sllr.Response.Response != nil {
				sc = result.sllr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listRecommendedByDatabaseNextResults
	req, err := client.ListRecommendedByDatabasePreparer(ctx, resourceGroupName, serverName, databaseName, includeDisabledRecommendations, skipToken, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "ListRecommendedByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRecommendedByDatabaseSender(req)
	if err != nil {
		result.sllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "ListRecommendedByDatabase", resp, "Failure sending request")
		return
	}

	result.sllr, err = client.ListRecommendedByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "ListRecommendedByDatabase", resp, "Failure responding to request")
	}

	return
}

// ListRecommendedByDatabasePreparer prepares the ListRecommendedByDatabase request.
func (client SensitivityLabelsClient) ListRecommendedByDatabasePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, includeDisabledRecommendations *bool, skipToken string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if includeDisabledRecommendations != nil {
		queryParameters["includeDisabledRecommendations"] = autorest.Encode("query", *includeDisabledRecommendations)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/recommendedSensitivityLabels", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRecommendedByDatabaseSender sends the ListRecommendedByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client SensitivityLabelsClient) ListRecommendedByDatabaseSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListRecommendedByDatabaseResponder handles the response to the ListRecommendedByDatabase request. The method always
// closes the http.Response Body.
func (client SensitivityLabelsClient) ListRecommendedByDatabaseResponder(resp *http.Response) (result SensitivityLabelListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listRecommendedByDatabaseNextResults retrieves the next set of results, if any.
func (client SensitivityLabelsClient) listRecommendedByDatabaseNextResults(ctx context.Context, lastResults SensitivityLabelListResult) (result SensitivityLabelListResult, err error) {
	req, err := lastResults.sensitivityLabelListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "listRecommendedByDatabaseNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListRecommendedByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "listRecommendedByDatabaseNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListRecommendedByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SensitivityLabelsClient", "listRecommendedByDatabaseNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListRecommendedByDatabaseComplete enumerates all values, automatically crossing page boundaries as required.
func (client SensitivityLabelsClient) ListRecommendedByDatabaseComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, includeDisabledRecommendations *bool, skipToken string, filter string) (result SensitivityLabelListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SensitivityLabelsClient.ListRecommendedByDatabase")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListRecommendedByDatabase(ctx, resourceGroupName, serverName, databaseName, includeDisabledRecommendations, skipToken, filter)
	return
}