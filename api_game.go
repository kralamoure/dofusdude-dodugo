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


// GameAPIService GameAPI service
type GameAPIService service

type ApiGetGameSearchRequest struct {
	ctx context.Context
	ApiService *GameAPIService
	language string
	game string
	query *string
	filterType *[]string
	limit *int32
	fieldsItem *[]string
}

// search query
func (r ApiGetGameSearchRequest) Query(query string) ApiGetGameSearchRequest {
	r.query = &query
	return r
}

// only results with all specific type
func (r ApiGetGameSearchRequest) FilterType(filterType []string) ApiGetGameSearchRequest {
	r.filterType = &filterType
	return r
}

// maximum number of returned results
func (r ApiGetGameSearchRequest) Limit(limit int32) ApiGetGameSearchRequest {
	r.limit = &limit
	return r
}

// adds fields from the item search to the list entries if the hit is a item. Multiple comma separated values allowed.
func (r ApiGetGameSearchRequest) FieldsItem(fieldsItem []string) ApiGetGameSearchRequest {
	r.fieldsItem = &fieldsItem
	return r
}

func (r ApiGetGameSearchRequest) Execute() ([]GetGameSearch200ResponseInner, *http.Response, error) {
	return r.ApiService.GetGameSearchExecute(r)
}

/*
GetGameSearch Game Search

Search in all names and descriptions of all supported types in the game. For the list of supported types see the endpoint /dofus2/meta/search/types.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game
 @return ApiGetGameSearchRequest
*/
func (a *GameAPIService) GetGameSearch(ctx context.Context, language string, game string) ApiGetGameSearchRequest {
	return ApiGetGameSearchRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return []GetGameSearch200ResponseInner
func (a *GameAPIService) GetGameSearchExecute(r ApiGetGameSearchRequest) ([]GetGameSearch200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetGameSearch200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GameAPIService.GetGameSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/search"
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
	if r.filterType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type]", r.filterType, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 8
		r.limit = &defaultValue
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

type ApiGetItemsAllSearchRequest struct {
	ctx context.Context
	ApiService *GameAPIService
	language string
	game string
	query *string
	filterTypeName *string
	filterMinLevel *int32
	filterMaxLevel *int32
	limit *int32
}

// case sensitive search query
func (r ApiGetItemsAllSearchRequest) Query(query string) ApiGetItemsAllSearchRequest {
	r.query = &query
	return r
}

// only results with the translated type name across all item_subtypes
func (r ApiGetItemsAllSearchRequest) FilterTypeName(filterTypeName string) ApiGetItemsAllSearchRequest {
	r.filterTypeName = &filterTypeName
	return r
}

// only results which level is equal or above this value
func (r ApiGetItemsAllSearchRequest) FilterMinLevel(filterMinLevel int32) ApiGetItemsAllSearchRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetItemsAllSearchRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetItemsAllSearchRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// maximum number of returned results
func (r ApiGetItemsAllSearchRequest) Limit(limit int32) ApiGetItemsAllSearchRequest {
	r.limit = &limit
	return r
}

func (r ApiGetItemsAllSearchRequest) Execute() ([]ItemsListEntryTyped, *http.Response, error) {
	return r.ApiService.GetItemsAllSearchExecute(r)
}

/*
GetItemsAllSearch Search All Items

Search in all names and descriptions of Dofus items (including all subtypes) with a query.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game
 @return ApiGetItemsAllSearchRequest
*/
func (a *GameAPIService) GetItemsAllSearch(ctx context.Context, language string, game string) ApiGetItemsAllSearchRequest {
	return ApiGetItemsAllSearchRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return []ItemsListEntryTyped
func (a *GameAPIService) GetItemsAllSearchExecute(r ApiGetItemsAllSearchRequest) ([]ItemsListEntryTyped, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ItemsListEntryTyped
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GameAPIService.GetItemsAllSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/items/search"
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
