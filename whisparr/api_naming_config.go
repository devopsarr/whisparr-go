/*
Radarr

Radarr API docs

API version: b08981dee068e1ed23e4f45a0d8fe70ef7bf7703
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
)


// NamingConfigAPIService NamingConfigAPI service
type NamingConfigAPIService service
type ApiGetNamingConfigRequest struct {
	ctx context.Context
	ApiService *NamingConfigAPIService
}

func (r ApiGetNamingConfigRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.GetNamingConfigExecute(r)
}

/*
GetNamingConfig Method for GetNamingConfig

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNamingConfigRequest
*/
func (a *NamingConfigAPIService) GetNamingConfig(ctx context.Context) ApiGetNamingConfigRequest {
	return ApiGetNamingConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigAPIService) GetNamingConfigExecute(r ApiGetNamingConfigRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigAPIService.GetNamingConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming"

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
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
type ApiGetNamingConfigByIdRequest struct {
	ctx context.Context
	ApiService *NamingConfigAPIService
	id int32
}

func (r ApiGetNamingConfigByIdRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.GetNamingConfigByIdExecute(r)
}

/*
GetNamingConfigById Method for GetNamingConfigById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetNamingConfigByIdRequest
*/
func (a *NamingConfigAPIService) GetNamingConfigById(ctx context.Context, id int32) ApiGetNamingConfigByIdRequest {
	return ApiGetNamingConfigByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigAPIService) GetNamingConfigByIdExecute(r ApiGetNamingConfigByIdRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigAPIService.GetNamingConfigById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/{id}"
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
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
type ApiGetNamingConfigExamplesRequest struct {
	ctx context.Context
	ApiService *NamingConfigAPIService
	renameMovies *bool
	replaceIllegalCharacters *bool
	colonReplacementFormat *ColonReplacementFormat
	standardMovieFormat *string
	movieFolderFormat *string
	includeQuality *bool
	replaceSpaces *bool
	separator *string
	numberStyle *string
	id *int32
	resourceName *string
}

func (r ApiGetNamingConfigExamplesRequest) RenameMovies(renameMovies bool) ApiGetNamingConfigExamplesRequest {
	r.renameMovies = &renameMovies
	return r
}

func (r ApiGetNamingConfigExamplesRequest) ReplaceIllegalCharacters(replaceIllegalCharacters bool) ApiGetNamingConfigExamplesRequest {
	r.replaceIllegalCharacters = &replaceIllegalCharacters
	return r
}

func (r ApiGetNamingConfigExamplesRequest) ColonReplacementFormat(colonReplacementFormat ColonReplacementFormat) ApiGetNamingConfigExamplesRequest {
	r.colonReplacementFormat = &colonReplacementFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) StandardMovieFormat(standardMovieFormat string) ApiGetNamingConfigExamplesRequest {
	r.standardMovieFormat = &standardMovieFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) MovieFolderFormat(movieFolderFormat string) ApiGetNamingConfigExamplesRequest {
	r.movieFolderFormat = &movieFolderFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) IncludeQuality(includeQuality bool) ApiGetNamingConfigExamplesRequest {
	r.includeQuality = &includeQuality
	return r
}

func (r ApiGetNamingConfigExamplesRequest) ReplaceSpaces(replaceSpaces bool) ApiGetNamingConfigExamplesRequest {
	r.replaceSpaces = &replaceSpaces
	return r
}

func (r ApiGetNamingConfigExamplesRequest) Separator(separator string) ApiGetNamingConfigExamplesRequest {
	r.separator = &separator
	return r
}

func (r ApiGetNamingConfigExamplesRequest) NumberStyle(numberStyle string) ApiGetNamingConfigExamplesRequest {
	r.numberStyle = &numberStyle
	return r
}

func (r ApiGetNamingConfigExamplesRequest) Id(id int32) ApiGetNamingConfigExamplesRequest {
	r.id = &id
	return r
}

func (r ApiGetNamingConfigExamplesRequest) ResourceName(resourceName string) ApiGetNamingConfigExamplesRequest {
	r.resourceName = &resourceName
	return r
}

func (r ApiGetNamingConfigExamplesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetNamingConfigExamplesExecute(r)
}

/*
GetNamingConfigExamples Method for GetNamingConfigExamples

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNamingConfigExamplesRequest
*/
func (a *NamingConfigAPIService) GetNamingConfigExamples(ctx context.Context) ApiGetNamingConfigExamplesRequest {
	return ApiGetNamingConfigExamplesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NamingConfigAPIService) GetNamingConfigExamplesExecute(r ApiGetNamingConfigExamplesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigAPIService.GetNamingConfigExamples")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/examples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.renameMovies != nil {
		localVarQueryParams.Add("renameMovies", parameterToString(*r.renameMovies, ""))
	}
	if r.replaceIllegalCharacters != nil {
		localVarQueryParams.Add("replaceIllegalCharacters", parameterToString(*r.replaceIllegalCharacters, ""))
	}
	if r.colonReplacementFormat != nil {
		localVarQueryParams.Add("colonReplacementFormat", parameterToString(*r.colonReplacementFormat, ""))
	}
	if r.standardMovieFormat != nil {
		localVarQueryParams.Add("standardMovieFormat", parameterToString(*r.standardMovieFormat, ""))
	}
	if r.movieFolderFormat != nil {
		localVarQueryParams.Add("movieFolderFormat", parameterToString(*r.movieFolderFormat, ""))
	}
	if r.includeQuality != nil {
		localVarQueryParams.Add("includeQuality", parameterToString(*r.includeQuality, ""))
	}
	if r.replaceSpaces != nil {
		localVarQueryParams.Add("replaceSpaces", parameterToString(*r.replaceSpaces, ""))
	}
	if r.separator != nil {
		localVarQueryParams.Add("separator", parameterToString(*r.separator, ""))
	}
	if r.numberStyle != nil {
		localVarQueryParams.Add("numberStyle", parameterToString(*r.numberStyle, ""))
	}
	if r.id != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id, ""))
	}
	if r.resourceName != nil {
		localVarQueryParams.Add("resourceName", parameterToString(*r.resourceName, ""))
	}
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
type ApiUpdateNamingConfigRequest struct {
	ctx context.Context
	ApiService *NamingConfigAPIService
	id string
	namingConfigResource *NamingConfigResource
}

func (r ApiUpdateNamingConfigRequest) NamingConfigResource(namingConfigResource NamingConfigResource) ApiUpdateNamingConfigRequest {
	r.namingConfigResource = &namingConfigResource
	return r
}

func (r ApiUpdateNamingConfigRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.UpdateNamingConfigExecute(r)
}

/*
UpdateNamingConfig Method for UpdateNamingConfig

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateNamingConfigRequest
*/
func (a *NamingConfigAPIService) UpdateNamingConfig(ctx context.Context, id string) ApiUpdateNamingConfigRequest {
	return ApiUpdateNamingConfigRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigAPIService) UpdateNamingConfigExecute(r ApiUpdateNamingConfigRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigAPIService.UpdateNamingConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.namingConfigResource
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
