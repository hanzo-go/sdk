# Fusion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backends** | Pointer to [**[]BackendStatus**](BackendStatus.md) | Backends is the per-leg report. Always populated. | [optional] 
**Hits** | Pointer to [**[]Hit**](Hit.md) | Hits is the fused, ranked result set. | [optional] 
**Mode** | Pointer to **string** | Mode is the mode actually used after &#x60;auto&#x60; resolution. | [optional] 
**Status** | Pointer to **string** | Status is the query&#39;s overall honesty signal:   ok          every consulted leg answered.   partial     at least one leg failed; Hits holds the survivors&#39; results.   unavailable every consulted leg failed; Hits is empty AND that is stated. | [optional] 
**TookMs** | Pointer to **int64** | TookMS is the whole query&#39;s wall time in milliseconds — every leg it consulted, plus fusion and paging. Each leg&#39;s own share is in Backends[].TookMS; the legs run in sequence, so this is at least their sum. | [optional] 

## Methods

### NewFusion

`func NewFusion() *Fusion`

NewFusion instantiates a new Fusion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFusionWithDefaults

`func NewFusionWithDefaults() *Fusion`

NewFusionWithDefaults instantiates a new Fusion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackends

`func (o *Fusion) GetBackends() []BackendStatus`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *Fusion) GetBackendsOk() (*[]BackendStatus, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *Fusion) SetBackends(v []BackendStatus)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *Fusion) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetHits

`func (o *Fusion) GetHits() []Hit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *Fusion) GetHitsOk() (*[]Hit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *Fusion) SetHits(v []Hit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *Fusion) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetMode

`func (o *Fusion) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Fusion) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Fusion) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Fusion) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetStatus

`func (o *Fusion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Fusion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Fusion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Fusion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTookMs

`func (o *Fusion) GetTookMs() int64`

GetTookMs returns the TookMs field if non-nil, zero value otherwise.

### GetTookMsOk

`func (o *Fusion) GetTookMsOk() (*int64, bool)`

GetTookMsOk returns a tuple with the TookMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTookMs

`func (o *Fusion) SetTookMs(v int64)`

SetTookMs sets TookMs field to given value.

### HasTookMs

`func (o *Fusion) HasTookMs() bool`

HasTookMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


