/*
Whisparr

Whisparr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package whisparr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// HistoryApiService HistoryApi service
type HistoryApiService service
type ApiCreateHistoryFailedByIdRequest struct {
	ctx context.Context
	ApiService *HistoryApiService
	id int32
}

func (r ApiCreateHistoryFailedByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateHistoryFailedByIdExecute(r)
}

/*
CreateHistoryFailedById Method for CreateHistoryFailedById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCreateHistoryFailedByIdRequest
*/
func (a *HistoryApiService) CreateHistoryFailedById(ctx context.Context, id int32) ApiCreateHistoryFailedByIdRequest {
	return ApiCreateHistoryFailedByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *HistoryApiService) CreateHistoryFailedByIdExecute(r ApiCreateHistoryFailedByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.CreateHistoryFailedById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/history/failed/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type ApiGetHistoryRequest struct {
	ctx context.Context
	ApiService *HistoryApiService
	includeSeries *bool
	includeEpisode *bool
}

func (r ApiGetHistoryRequest) IncludeSeries(includeSeries bool) ApiGetHistoryRequest {
	r.includeSeries = &includeSeries
	return r
}

func (r ApiGetHistoryRequest) IncludeEpisode(includeEpisode bool) ApiGetHistoryRequest {
	r.includeEpisode = &includeEpisode
	return r
}

func (r ApiGetHistoryRequest) Execute() (*HistoryResourcePagingResource, *http.Response, error) {
	return r.ApiService.GetHistoryExecute(r)
}

/*
GetHistory Method for GetHistory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHistoryRequest
*/
func (a *HistoryApiService) GetHistory(ctx context.Context) ApiGetHistoryRequest {
	return ApiGetHistoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistoryResourcePagingResource
func (a *HistoryApiService) GetHistoryExecute(r ApiGetHistoryRequest) (*HistoryResourcePagingResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistoryResourcePagingResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.GetHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeSeries != nil {
		localVarQueryParams.Add("includeSeries", parameterToString(*r.includeSeries, ""))
	}
	if r.includeEpisode != nil {
		localVarQueryParams.Add("includeEpisode", parameterToString(*r.includeEpisode, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
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
type ApiListHistorySeriesRequest struct {
	ctx context.Context
	ApiService *HistoryApiService
	seriesId *int32
	seasonNumber *int32
	eventType *EpisodeHistoryEventType
	includeSeries *bool
	includeEpisode *bool
}

func (r ApiListHistorySeriesRequest) SeriesId(seriesId int32) ApiListHistorySeriesRequest {
	r.seriesId = &seriesId
	return r
}

func (r ApiListHistorySeriesRequest) SeasonNumber(seasonNumber int32) ApiListHistorySeriesRequest {
	r.seasonNumber = &seasonNumber
	return r
}

func (r ApiListHistorySeriesRequest) EventType(eventType EpisodeHistoryEventType) ApiListHistorySeriesRequest {
	r.eventType = &eventType
	return r
}

func (r ApiListHistorySeriesRequest) IncludeSeries(includeSeries bool) ApiListHistorySeriesRequest {
	r.includeSeries = &includeSeries
	return r
}

func (r ApiListHistorySeriesRequest) IncludeEpisode(includeEpisode bool) ApiListHistorySeriesRequest {
	r.includeEpisode = &includeEpisode
	return r
}

func (r ApiListHistorySeriesRequest) Execute() ([]*HistoryResource, *http.Response, error) {
	return r.ApiService.ListHistorySeriesExecute(r)
}

/*
ListHistorySeries Method for ListHistorySeries

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHistorySeriesRequest
*/
func (a *HistoryApiService) ListHistorySeries(ctx context.Context) ApiListHistorySeriesRequest {
	return ApiListHistorySeriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []HistoryResource
func (a *HistoryApiService) ListHistorySeriesExecute(r ApiListHistorySeriesRequest) ([]*HistoryResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []*HistoryResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.ListHistorySeries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/history/series"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.seriesId != nil {
		localVarQueryParams.Add("seriesId", parameterToString(*r.seriesId, ""))
	}
	if r.seasonNumber != nil {
		localVarQueryParams.Add("seasonNumber", parameterToString(*r.seasonNumber, ""))
	}
	if r.eventType != nil {
		localVarQueryParams.Add("eventType", parameterToString(*r.eventType, ""))
	}
	if r.includeSeries != nil {
		localVarQueryParams.Add("includeSeries", parameterToString(*r.includeSeries, ""))
	}
	if r.includeEpisode != nil {
		localVarQueryParams.Add("includeEpisode", parameterToString(*r.includeEpisode, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
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
type ApiListHistorySinceRequest struct {
	ctx context.Context
	ApiService *HistoryApiService
	date *time.Time
	eventType *EpisodeHistoryEventType
	includeSeries *bool
	includeEpisode *bool
}

func (r ApiListHistorySinceRequest) Date(date time.Time) ApiListHistorySinceRequest {
	r.date = &date
	return r
}

func (r ApiListHistorySinceRequest) EventType(eventType EpisodeHistoryEventType) ApiListHistorySinceRequest {
	r.eventType = &eventType
	return r
}

func (r ApiListHistorySinceRequest) IncludeSeries(includeSeries bool) ApiListHistorySinceRequest {
	r.includeSeries = &includeSeries
	return r
}

func (r ApiListHistorySinceRequest) IncludeEpisode(includeEpisode bool) ApiListHistorySinceRequest {
	r.includeEpisode = &includeEpisode
	return r
}

func (r ApiListHistorySinceRequest) Execute() ([]*HistoryResource, *http.Response, error) {
	return r.ApiService.ListHistorySinceExecute(r)
}

/*
ListHistorySince Method for ListHistorySince

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHistorySinceRequest
*/
func (a *HistoryApiService) ListHistorySince(ctx context.Context) ApiListHistorySinceRequest {
	return ApiListHistorySinceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []HistoryResource
func (a *HistoryApiService) ListHistorySinceExecute(r ApiListHistorySinceRequest) ([]*HistoryResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []*HistoryResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.ListHistorySince")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/history/since"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		localVarQueryParams.Add("date", parameterToString(*r.date, ""))
	}
	if r.eventType != nil {
		localVarQueryParams.Add("eventType", parameterToString(*r.eventType, ""))
	}
	if r.includeSeries != nil {
		localVarQueryParams.Add("includeSeries", parameterToString(*r.includeSeries, ""))
	}
	if r.includeEpisode != nil {
		localVarQueryParams.Add("includeEpisode", parameterToString(*r.includeEpisode, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
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
