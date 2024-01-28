/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save - [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save - [Go](https://github.com/dofusdude/dodugo) go get -u github.com/dofusdude/dodugo - [Python](https://github.com/dofusdude/dofusdude-py) pip install dofusdude - [PHP](https://github.com/dofusdude/dofusdude-php)  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my Discord.   ## Deploy now. Use forever. Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility.  There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.  ## Thank you! I highly welcome everyone on my [Discord](https://discord.gg/3EtHskZD8h) to just talk about projects and use cases or give feedback of any kind.  The servers have a fixed monthly cost to provide very fast responses. If you want to help me keeping them running or simply donate to that cause, consider becoming a [GitHub Sponsor](https://github.com/sponsors/dofusdude).

API version: 0.8.2
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ConsumablesAPIService ConsumablesAPI service
type ConsumablesAPIService service

type ApiGetAllItemsConsumablesListRequest struct {
	ctx context.Context
	ApiService *ConsumablesAPIService
	language string
	game string
	sortLevel *string
	filterTypeName *string
	filterMinLevel *int32
	filterMaxLevel *int32
	acceptEncoding *string
}

// sort the resulting list by level, default unsorted
func (r ApiGetAllItemsConsumablesListRequest) SortLevel(sortLevel string) ApiGetAllItemsConsumablesListRequest {
	r.sortLevel = &sortLevel
	return r
}

// only results with the translated type name
func (r ApiGetAllItemsConsumablesListRequest) FilterTypeName(filterTypeName string) ApiGetAllItemsConsumablesListRequest {
	r.filterTypeName = &filterTypeName
	return r
}

// only results which level is equal or above this value
func (r ApiGetAllItemsConsumablesListRequest) FilterMinLevel(filterMinLevel int32) ApiGetAllItemsConsumablesListRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetAllItemsConsumablesListRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetAllItemsConsumablesListRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// optional compression for saving bandwidth
func (r ApiGetAllItemsConsumablesListRequest) AcceptEncoding(acceptEncoding string) ApiGetAllItemsConsumablesListRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}

func (r ApiGetAllItemsConsumablesListRequest) Execute() (*ItemsListPaged, *http.Response, error) {
	return r.ApiService.GetAllItemsConsumablesListExecute(r)
}

/*
GetAllItemsConsumablesList List All Consumables

Retrieve all consumable items with one request. This endpoint is just an alias for the a list with disabled pagination (page[size]=-1) and all fields[type] set.

If you want everything unfiltered, delete the other query parameters.

Be careful with testing or (god forbid) using /all in your browser, the returned json is huge and will slow down the browser!

Tip: set the HTTP Header 'Accept-Encoding: gzip' for saving bandwidth. You will need to uncompress it on your end.
Example with cURL:
```
curl -sH 'Accept-Encoding: gzip' <api-endpoint> | gunzip -
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game
 @return ApiGetAllItemsConsumablesListRequest
*/
func (a *ConsumablesAPIService) GetAllItemsConsumablesList(ctx context.Context, language string, game string) ApiGetAllItemsConsumablesListRequest {
	return ApiGetAllItemsConsumablesListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return ItemsListPaged
func (a *ConsumablesAPIService) GetAllItemsConsumablesListExecute(r ApiGetAllItemsConsumablesListRequest) (*ItemsListPaged, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ItemsListPaged
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConsumablesAPIService.GetAllItemsConsumablesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/items/consumables/all"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}

	if r.sortLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort[level]", r.sortLevel, "")
	}
	if r.filterTypeName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type_name]", r.filterTypeName, "")
	}
	if r.filterMinLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_level]", r.filterMinLevel, "")
	}
	if r.filterMaxLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_level]", r.filterMaxLevel, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptEncoding != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Encoding", r.acceptEncoding, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetItemsConsumablesListRequest struct {
	ctx context.Context
	ApiService *ConsumablesAPIService
	language string
	game string
	sortLevel *string
	filterTypeName *string
	filterMinLevel *int32
	filterMaxLevel *int32
	pageSize *int32
	pageNumber *int32
	fieldsItem *[]string
}

// sort the resulting list by level, default unsorted
func (r ApiGetItemsConsumablesListRequest) SortLevel(sortLevel string) ApiGetItemsConsumablesListRequest {
	r.sortLevel = &sortLevel
	return r
}

// only results with the translated type name
func (r ApiGetItemsConsumablesListRequest) FilterTypeName(filterTypeName string) ApiGetItemsConsumablesListRequest {
	r.filterTypeName = &filterTypeName
	return r
}

// only results which level is equal or above this value
func (r ApiGetItemsConsumablesListRequest) FilterMinLevel(filterMinLevel int32) ApiGetItemsConsumablesListRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetItemsConsumablesListRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetItemsConsumablesListRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// size of the results from the list. -1 disables pagination and gets all in one response.
func (r ApiGetItemsConsumablesListRequest) PageSize(pageSize int32) ApiGetItemsConsumablesListRequest {
	r.pageSize = &pageSize
	return r
}

// page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16.
func (r ApiGetItemsConsumablesListRequest) PageNumber(pageNumber int32) ApiGetItemsConsumablesListRequest {
	r.pageNumber = &pageNumber
	return r
}

// adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed.
func (r ApiGetItemsConsumablesListRequest) FieldsItem(fieldsItem []string) ApiGetItemsConsumablesListRequest {
	r.fieldsItem = &fieldsItem
	return r
}

func (r ApiGetItemsConsumablesListRequest) Execute() (*ItemsListPaged, *http.Response, error) {
	return r.ApiService.GetItemsConsumablesListExecute(r)
}

/*
GetItemsConsumablesList List Consumables

Retrieve a list of consumable items.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game
 @return ApiGetItemsConsumablesListRequest
*/
func (a *ConsumablesAPIService) GetItemsConsumablesList(ctx context.Context, language string, game string) ApiGetItemsConsumablesListRequest {
	return ApiGetItemsConsumablesListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return ItemsListPaged
func (a *ConsumablesAPIService) GetItemsConsumablesListExecute(r ApiGetItemsConsumablesListRequest) (*ItemsListPaged, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ItemsListPaged
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConsumablesAPIService.GetItemsConsumablesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/items/consumables"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}

	if r.sortLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort[level]", r.sortLevel, "")
	}
	if r.filterTypeName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type_name]", r.filterTypeName, "")
	}
	if r.filterMinLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_level]", r.filterMinLevel, "")
	}
	if r.filterMaxLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_level]", r.filterMaxLevel, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[size]", r.pageSize, "")
	}
	if r.pageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[number]", r.pageNumber, "")
	}
	if r.fieldsItem != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[item]", r.fieldsItem, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetItemsConsumablesSearchRequest struct {
	ctx context.Context
	ApiService *ConsumablesAPIService
	language string
	game string
	query *string
	filterTypeName *string
	filterMinLevel *int32
	filterMaxLevel *int32
	limit *int32
}

// case sensitive search query
func (r ApiGetItemsConsumablesSearchRequest) Query(query string) ApiGetItemsConsumablesSearchRequest {
	r.query = &query
	return r
}

// only results with the translated type name
func (r ApiGetItemsConsumablesSearchRequest) FilterTypeName(filterTypeName string) ApiGetItemsConsumablesSearchRequest {
	r.filterTypeName = &filterTypeName
	return r
}

// only results which level is equal or above this value
func (r ApiGetItemsConsumablesSearchRequest) FilterMinLevel(filterMinLevel int32) ApiGetItemsConsumablesSearchRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetItemsConsumablesSearchRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetItemsConsumablesSearchRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// maximum number of returned results
func (r ApiGetItemsConsumablesSearchRequest) Limit(limit int32) ApiGetItemsConsumablesSearchRequest {
	r.limit = &limit
	return r
}

func (r ApiGetItemsConsumablesSearchRequest) Execute() ([]ItemListEntry, *http.Response, error) {
	return r.ApiService.GetItemsConsumablesSearchExecute(r)
}

/*
GetItemsConsumablesSearch Search Consumables

Search in all names and descriptions of consumable items with a query.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game
 @return ApiGetItemsConsumablesSearchRequest
*/
func (a *ConsumablesAPIService) GetItemsConsumablesSearch(ctx context.Context, language string, game string) ApiGetItemsConsumablesSearchRequest {
	return ApiGetItemsConsumablesSearchRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return []ItemListEntry
func (a *ConsumablesAPIService) GetItemsConsumablesSearchExecute(r ApiGetItemsConsumablesSearchRequest) ([]ItemListEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ItemListEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConsumablesAPIService.GetItemsConsumablesSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/items/consumables/search"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}
	if r.query == nil {
		return localVarReturnValue, nil, reportError("query is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "")
	if r.filterTypeName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type_name]", r.filterTypeName, "")
	}
	if r.filterMinLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_level]", r.filterMinLevel, "")
	}
	if r.filterMaxLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_level]", r.filterMaxLevel, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 8
		r.limit = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetItemsConsumablesSingleRequest struct {
	ctx context.Context
	ApiService *ConsumablesAPIService
	language string
	ankamaId int32
	game string
}

func (r ApiGetItemsConsumablesSingleRequest) Execute() (*Resource, *http.Response, error) {
	return r.ApiService.GetItemsConsumablesSingleExecute(r)
}

/*
GetItemsConsumablesSingle Single Consumables

Retrieve a specific consumable item with id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param ankamaId identifier
 @param game
 @return ApiGetItemsConsumablesSingleRequest
*/
func (a *ConsumablesAPIService) GetItemsConsumablesSingle(ctx context.Context, language string, ankamaId int32, game string) ApiGetItemsConsumablesSingleRequest {
	return ApiGetItemsConsumablesSingleRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		ankamaId: ankamaId,
		game: game,
	}
}

// Execute executes the request
//  @return Resource
func (a *ConsumablesAPIService) GetItemsConsumablesSingleExecute(r ApiGetItemsConsumablesSingleRequest) (*Resource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Resource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConsumablesAPIService.GetItemsConsumablesSingle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/items/consumables/{ankama_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ankama_id"+"}", url.PathEscape(parameterValueToString(r.ankamaId, "ankamaId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
