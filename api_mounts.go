/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [PHP](https://github.com/dofusdude/dofusdude-php) - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 0.9.2
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


// MountsAPIService MountsAPI service
type MountsAPIService service

type ApiGetAllMountsListRequest struct {
	ctx context.Context
	ApiService *MountsAPIService
	language string
	game string
	filterFamilyName *string
	acceptEncoding *string
}

// only results with the translated family name
func (r ApiGetAllMountsListRequest) FilterFamilyName(filterFamilyName string) ApiGetAllMountsListRequest {
	r.filterFamilyName = &filterFamilyName
	return r
}

// optional compression for saving bandwidth
func (r ApiGetAllMountsListRequest) AcceptEncoding(acceptEncoding string) ApiGetAllMountsListRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}

func (r ApiGetAllMountsListRequest) Execute() (*MountsListPaged, *http.Response, error) {
	return r.ApiService.GetAllMountsListExecute(r)
}

/*
GetAllMountsList List All Mounts

Retrieve all mounts with one request. This endpoint is just an alias for the a list with disabled pagination (page[size]=-1) and all fields[type] set.

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
 @return ApiGetAllMountsListRequest
*/
func (a *MountsAPIService) GetAllMountsList(ctx context.Context, language string, game string) ApiGetAllMountsListRequest {
	return ApiGetAllMountsListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return MountsListPaged
func (a *MountsAPIService) GetAllMountsListExecute(r ApiGetAllMountsListRequest) (*MountsListPaged, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MountsListPaged
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MountsAPIService.GetAllMountsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/mounts/all"
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

	if r.filterFamilyName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[family_name]", r.filterFamilyName, "form", "")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Encoding", r.acceptEncoding, "simple", "")
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

type ApiGetMountsListRequest struct {
	ctx context.Context
	ApiService *MountsAPIService
	language string
	game string
	filterFamilyName *string
	pageSize *int32
	pageNumber *int32
	fieldsMount *[]string
}

// only results with the translated family name
func (r ApiGetMountsListRequest) FilterFamilyName(filterFamilyName string) ApiGetMountsListRequest {
	r.filterFamilyName = &filterFamilyName
	return r
}

// size of the results from the list. -1 disables pagination and gets all in one response.
func (r ApiGetMountsListRequest) PageSize(pageSize int32) ApiGetMountsListRequest {
	r.pageSize = &pageSize
	return r
}

// page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16.
func (r ApiGetMountsListRequest) PageNumber(pageNumber int32) ApiGetMountsListRequest {
	r.pageNumber = &pageNumber
	return r
}

// adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed.
func (r ApiGetMountsListRequest) FieldsMount(fieldsMount []string) ApiGetMountsListRequest {
	r.fieldsMount = &fieldsMount
	return r
}

func (r ApiGetMountsListRequest) Execute() (*MountsListPaged, *http.Response, error) {
	return r.ApiService.GetMountsListExecute(r)
}

/*
GetMountsList List Mounts

Retrieve a list of mounts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game
 @return ApiGetMountsListRequest
*/
func (a *MountsAPIService) GetMountsList(ctx context.Context, language string, game string) ApiGetMountsListRequest {
	return ApiGetMountsListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return MountsListPaged
func (a *MountsAPIService) GetMountsListExecute(r ApiGetMountsListRequest) (*MountsListPaged, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MountsListPaged
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MountsAPIService.GetMountsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/mounts"
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

	if r.filterFamilyName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[family_name]", r.filterFamilyName, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[size]", r.pageSize, "form", "")
	}
	if r.pageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[number]", r.pageNumber, "form", "")
	}
	if r.fieldsMount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[mount]", r.fieldsMount, "form", "csv")
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

type ApiGetMountsSearchRequest struct {
	ctx context.Context
	ApiService *MountsAPIService
	language string
	game string
	query *string
	filterFamilyName *string
	limit *int32
}

// case sensitive search query
func (r ApiGetMountsSearchRequest) Query(query string) ApiGetMountsSearchRequest {
	r.query = &query
	return r
}

// only results with the translated family name
func (r ApiGetMountsSearchRequest) FilterFamilyName(filterFamilyName string) ApiGetMountsSearchRequest {
	r.filterFamilyName = &filterFamilyName
	return r
}

// maximum number of returned results
func (r ApiGetMountsSearchRequest) Limit(limit int32) ApiGetMountsSearchRequest {
	r.limit = &limit
	return r
}

func (r ApiGetMountsSearchRequest) Execute() ([]MountListEntry, *http.Response, error) {
	return r.ApiService.GetMountsSearchExecute(r)
}

/*
GetMountsSearch Search Mounts

Search in all names and descriptions of mounts with a query.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game
 @return ApiGetMountsSearchRequest
*/
func (a *MountsAPIService) GetMountsSearch(ctx context.Context, language string, game string) ApiGetMountsSearchRequest {
	return ApiGetMountsSearchRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return []MountListEntry
func (a *MountsAPIService) GetMountsSearchExecute(r ApiGetMountsSearchRequest) ([]MountListEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MountListEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MountsAPIService.GetMountsSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/mounts/search"
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

	parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "form", "")
	if r.filterFamilyName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[family_name]", r.filterFamilyName, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
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

type ApiGetMountsSingleRequest struct {
	ctx context.Context
	ApiService *MountsAPIService
	language string
	ankamaId int32
	game string
}

func (r ApiGetMountsSingleRequest) Execute() (*Mount, *http.Response, error) {
	return r.ApiService.GetMountsSingleExecute(r)
}

/*
GetMountsSingle Single Mounts

Retrieve a specific mount with id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param ankamaId identifier
 @param game
 @return ApiGetMountsSingleRequest
*/
func (a *MountsAPIService) GetMountsSingle(ctx context.Context, language string, ankamaId int32, game string) ApiGetMountsSingleRequest {
	return ApiGetMountsSingleRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		ankamaId: ankamaId,
		game: game,
	}
}

// Execute executes the request
//  @return Mount
func (a *MountsAPIService) GetMountsSingleExecute(r ApiGetMountsSingleRequest) (*Mount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Mount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MountsAPIService.GetMountsSingle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/mounts/{ankama_id}"
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
