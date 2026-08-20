# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backends** | Pointer to [**[]BackendStatus**](BackendStatus.md) | Backends is the per-leg report. Always populated. | [optional] 
**Hits** | Pointer to [**[]Hit**](Hit.md) | Hits is the fused, ranked result set. | [optional] 
**Mode** | Pointer to **string** | Mode is the mode actually used after &#x60;auto&#x60; resolution. | [optional] 
**Status** | Pointer to **string** | Status is the query&#39;s overall honesty signal:   ok          every consulted leg answered.   partial     at least one leg failed; Hits holds the survivors&#39; results.   unavailable every consulted leg failed; Hits is empty AND that is stated. | [optional] 
**TookMs** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponse

`func NewResponse() *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackends

`func (o *Response) GetBackends() []BackendStatus`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *Response) GetBackendsOk() (*[]BackendStatus, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *Response) SetBackends(v []BackendStatus)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *Response) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetHits

`func (o *Response) GetHits() []Hit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *Response) GetHitsOk() (*[]Hit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *Response) SetHits(v []Hit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *Response) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetMode

`func (o *Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetStatus

`func (o *Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTookMs

`func (o *Response) GetTookMs() int32`

GetTookMs returns the TookMs field if non-nil, zero value otherwise.

### GetTookMsOk

`func (o *Response) GetTookMsOk() (*int32, bool)`

GetTookMsOk returns a tuple with the TookMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTookMs

`func (o *Response) SetTookMs(v int32)`

SetTookMs sets TookMs field to given value.

### HasTookMs

`func (o *Response) HasTookMs() bool`

HasTookMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


