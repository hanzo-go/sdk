# ConsoleGetMediaUploadUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | **string** |  | 
**ObservationId** | Pointer to **string** |  | [optional] 
**ContentType** | **string** |  | 
**ContentLength** | **int32** |  | 
**Sha256Hash** | **string** |  | 
**Field** | **string** | Trace/observation field: input, output, or metadata | 

## Methods

### NewConsoleGetMediaUploadUrlRequest

`func NewConsoleGetMediaUploadUrlRequest(traceId string, contentType string, contentLength int32, sha256Hash string, field string, ) *ConsoleGetMediaUploadUrlRequest`

NewConsoleGetMediaUploadUrlRequest instantiates a new ConsoleGetMediaUploadUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleGetMediaUploadUrlRequestWithDefaults

`func NewConsoleGetMediaUploadUrlRequestWithDefaults() *ConsoleGetMediaUploadUrlRequest`

NewConsoleGetMediaUploadUrlRequestWithDefaults instantiates a new ConsoleGetMediaUploadUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *ConsoleGetMediaUploadUrlRequest) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ConsoleGetMediaUploadUrlRequest) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ConsoleGetMediaUploadUrlRequest) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetObservationId

`func (o *ConsoleGetMediaUploadUrlRequest) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *ConsoleGetMediaUploadUrlRequest) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *ConsoleGetMediaUploadUrlRequest) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *ConsoleGetMediaUploadUrlRequest) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetContentType

`func (o *ConsoleGetMediaUploadUrlRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ConsoleGetMediaUploadUrlRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ConsoleGetMediaUploadUrlRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetContentLength

`func (o *ConsoleGetMediaUploadUrlRequest) GetContentLength() int32`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *ConsoleGetMediaUploadUrlRequest) GetContentLengthOk() (*int32, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *ConsoleGetMediaUploadUrlRequest) SetContentLength(v int32)`

SetContentLength sets ContentLength field to given value.


### GetSha256Hash

`func (o *ConsoleGetMediaUploadUrlRequest) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *ConsoleGetMediaUploadUrlRequest) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *ConsoleGetMediaUploadUrlRequest) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetField

`func (o *ConsoleGetMediaUploadUrlRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ConsoleGetMediaUploadUrlRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ConsoleGetMediaUploadUrlRequest) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


