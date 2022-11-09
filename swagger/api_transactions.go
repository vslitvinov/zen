
/*
 * ZEN eCommerce Public API
 *
 * ZEN allows you to accept multiple global and local payment methods (including credit cards, instant bank transfers, and many more). The following document will guide you through the integration process and introduce you to the main features of the Public API..<br><br><p><strong><span data-contrast=\"auto\">1. Allowed HTTPs requests:</span></strong><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><ul><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">PU</span></strong><strong><span data-contrast=\"auto\">T</span></strong><strong><span data-contrast=\"auto\">:</span></strong><span data-contrast=\"auto\">&nbsp;To create a resource&nbsp;</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">POST:</span></strong><span data-contrast=\"auto\">&nbsp;To update ra esource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">GET:</span></strong><span data-contrast=\"auto\">&nbsp;To get a resource or a list of resources</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">DELETE:</span></strong><span data-contrast=\"auto\">&nbsp;To delete a resource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></li></ul><p><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span><strong><span data-contrast=\"auto\">2. Description&nbsp;</span></strong><strong><span data-contrast=\"auto\">o</span></strong><strong><span data-contrast=\"auto\">f&nbsp;</span></strong><strong><span data-contrast=\"auto\">u</span></strong><strong><span data-contrast=\"auto\">sual&nbsp;</span></strong><strong><span data-contrast=\"auto\">s</span></strong><strong><span data-contrast=\"auto\">erver&nbsp;</span></strong><strong><span data-contrast=\"auto\">r</span></strong><strong><span data-contrast=\"auto\">esponses:</span></strong><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><ul><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">200</span></strong><strong><span data-contrast=\"auto\">&nbsp;</span></strong><strong><span data-contrast=\"auto\">--</span></strong><strong><span data-contrast=\"auto\">&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">OK</span></strong><span data-contrast=\"auto\"> - the request was successful (some API calls may return 201 instead)</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">201&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">  Created</span></strong><span data-contrast=\"auto\"> - the request was successful and a resource was created</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"3\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">204  </span></strong><strong><span data-contrast=\"auto\">--&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">No Content</span></strong><span data-contrast=\"auto\"> - the request was successful, but there is no representation to return (</span><span data-contrast=\"auto\">i</span><span data-contrast=\"auto\">.e. the response is empty)</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"4\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">400&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">Bad Request</span></strong><span data-contrast=\"auto\"> - the request could not be understood or was missing required parameters</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"5\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">401&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">  Unauthorized</span></strong><span data-contrast=\"auto\"> - authentication failed or user does not have permissions for requested operation</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">403&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><span data-contrast=\"auto\">&nbsp;</span><strong><span data-contrast=\"auto\">Forbidden </span></strong><span data-contrast=\"auto\"> - access denied</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">404&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><span data-contrast=\"auto\"> </span><strong><span data-contrast=\"auto\">Not Found</span></strong><span data-contrast=\"auto\"> - the resource was not found</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"3\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">405  </span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">&nbsp;</span></strong><strong><span data-contrast=\"auto\">Method Not Allowed</span></strong><span data-contrast=\"auto\"> - requested method is not supported for the resource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li></ul><p><strong><span data-contrast=\"auto\">3.&nbsp;</span></strong><strong><span data-contrast=\"auto\">API Integration Guidelines</span></strong><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><p><span data-contrast=\"auto\">To start using ZEN eCommerce Public API please follow these instructions:&nbsp;</span><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:6,&quot;335551620&quot;:6,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#before-you-begin%20\">https://www.zen.com/developer/api-integration#before-you-begin&nbsp;</a><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559737&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></li></ol><p><strong><span data-contrast=\"auto\">4. IPN API Secrets</span></strong><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><p><p><span data-contrast=\"auto\">Instant Payment Notification (IPN) informs you about transaction status updates, e.g., about changing the status from pending to rejected. </span><br /><span data-contrast=\"auto\">Information abut IPN is available in the following link:&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#ipn\"><span data-contrast=\"auto\">https://www.zen.com/developer/api-integration#ipn</span></a></p> <p><strong><span data-contrast=\"auto\">4. Error codes</span></strong><p><span data-contrast=\"auto\">The full list of error codes can be found here:&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#rejection-codes\">https://www.zen.com/developer/api-integration#rejection-codes</a></p>  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type TransactionsApiService service
/*
TransactionsApiService Process transaction refunds using initial transaction&#x27;s ID, received in POST /transaction
Processes transaction refunds using initial transaction&#x27;s ID, received in POST /transaction
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param requestId Unique identifier generated by requesting client
 * @param terminalId ID of the ZEN terminal
@return DetailsAboutTransaction
*/
func (a *TransactionsApiService) CreateRefund(ctx context.Context, body ActionRefundObject, requestId string, terminalId string) (DetailsAboutTransaction, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue DetailsAboutTransaction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/transactions/refund"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(requestId) < 38 {
		return localVarReturnValue, nil, reportError("requestId must have at least 38 elements")
	}
	if strlen(requestId) > 1024 {
		return localVarReturnValue, nil, reportError("requestId must have less than 1024 elements")
	}
	if strlen(terminalId) < 3 {
		return localVarReturnValue, nil, reportError("terminalId must have at least 3 elements")
	}
	if strlen(terminalId) > 128 {
		return localVarReturnValue, nil, reportError("terminalId must have less than 128 elements")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["request-id"] = parameterToString(requestId, "")
	localVarHeaderParams["terminal-id"] = parameterToString(terminalId, "")
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v DetailsAboutTransaction
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TransactionsApiService Create transaction
Initiates transaction creation by sending transaction parameteres to ZEN. Important! For some payment methods, a returnUrl will be returned to you. To proceed with the transaction, the customer needs to be redirected to that URL.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param requestId Unique identifier generated by requesting client
 * @param terminalId ID of the ZEN terminal
@return DetailsAboutTransaction
*/
func (a *TransactionsApiService) CreateTransaction(ctx context.Context, body ActionCreateObject, requestId string, terminalId string) (DetailsAboutTransaction, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue DetailsAboutTransaction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(requestId) < 38 {
		return localVarReturnValue, nil, reportError("requestId must have at least 38 elements")
	}
	if strlen(requestId) > 1024 {
		return localVarReturnValue, nil, reportError("requestId must have less than 1024 elements")
	}
	if strlen(terminalId) < 3 {
		return localVarReturnValue, nil, reportError("terminalId must have at least 3 elements")
	}
	if strlen(terminalId) > 128 {
		return localVarReturnValue, nil, reportError("terminalId must have less than 128 elements")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["request-id"] = parameterToString(requestId, "")
	localVarHeaderParams["terminal-id"] = parameterToString(terminalId, "")
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v DetailsAboutTransaction
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TransactionsApiService Get transaction details using merchant transaction ID
Returns transaction details using Merchant transaction ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param merchantTransactionId Unique merchant transaction identifier
 * @param requestId Unique identifier generated by requesting client
 * @param terminalId ID of the ZEN terminal
@return DetailsAboutTransaction
*/
func (a *TransactionsApiService) GetMerchantTransaction(ctx context.Context, merchantTransactionId string, requestId string, terminalId string) (DetailsAboutTransaction, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue DetailsAboutTransaction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/transactions/merchant/{merchantTransactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"merchantTransactionId"+"}", fmt.Sprintf("%v", merchantTransactionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(merchantTransactionId) < 1 {
		return localVarReturnValue, nil, reportError("merchantTransactionId must have at least 1 elements")
	}
	if strlen(merchantTransactionId) > 128 {
		return localVarReturnValue, nil, reportError("merchantTransactionId must have less than 128 elements")
	}
	if strlen(requestId) < 38 {
		return localVarReturnValue, nil, reportError("requestId must have at least 38 elements")
	}
	if strlen(requestId) > 1024 {
		return localVarReturnValue, nil, reportError("requestId must have less than 1024 elements")
	}
	if strlen(terminalId) < 3 {
		return localVarReturnValue, nil, reportError("terminalId must have at least 3 elements")
	}
	if strlen(terminalId) > 128 {
		return localVarReturnValue, nil, reportError("terminalId must have less than 128 elements")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["request-id"] = parameterToString(requestId, "")
	localVarHeaderParams["terminal-id"] = parameterToString(terminalId, "")
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v DetailsAboutTransaction
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TransactionsApiService Get transaction details using ZEN ID
Returns transaction details using ZEN ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Unique identifier
 * @param requestId Unique identifier generated by requesting client
 * @param terminalId ID of the ZEN terminal
@return DetailsAboutTransaction
*/
func (a *TransactionsApiService) GetTransaction(ctx context.Context, id string, requestId string, terminalId string) (DetailsAboutTransaction, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue DetailsAboutTransaction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/transactions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(requestId) < 38 {
		return localVarReturnValue, nil, reportError("requestId must have at least 38 elements")
	}
	if strlen(requestId) > 1024 {
		return localVarReturnValue, nil, reportError("requestId must have less than 1024 elements")
	}
	if strlen(terminalId) < 3 {
		return localVarReturnValue, nil, reportError("terminalId must have at least 3 elements")
	}
	if strlen(terminalId) > 128 {
		return localVarReturnValue, nil, reportError("terminalId must have less than 128 elements")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["request-id"] = parameterToString(requestId, "")
	localVarHeaderParams["terminal-id"] = parameterToString(terminalId, "")
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v DetailsAboutTransaction
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TransactionsApiService Create new authorization attempt on the same TRX transaction in case of expired or rejected previous authorization
Creates new authorization attempt on the same TRX transaction in case of expired or rejected previous authorization
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param requestId Unique identifier generated by requesting client
 * @param terminalId ID of the ZEN terminal
@return DetailsAboutTransaction
*/
func (a *TransactionsApiService) RenewAuthorization(ctx context.Context, body ActionRenewAuthorizationObject, requestId string, terminalId string) (DetailsAboutTransaction, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue DetailsAboutTransaction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/transactions/renewAuthorization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(requestId) < 38 {
		return localVarReturnValue, nil, reportError("requestId must have at least 38 elements")
	}
	if strlen(requestId) > 1024 {
		return localVarReturnValue, nil, reportError("requestId must have less than 1024 elements")
	}
	if strlen(terminalId) < 3 {
		return localVarReturnValue, nil, reportError("terminalId must have at least 3 elements")
	}
	if strlen(terminalId) > 128 {
		return localVarReturnValue, nil, reportError("terminalId must have less than 128 elements")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["request-id"] = parameterToString(requestId, "")
	localVarHeaderParams["terminal-id"] = parameterToString(terminalId, "")
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v DetailsAboutTransaction
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
