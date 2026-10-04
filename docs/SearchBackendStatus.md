# SearchBackendStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capped** | Pointer to **bool** | Capped is true when the leg answered from only PART of its corpus because it stopped at its scan cap — the code leg compares a query with at most 50000 stored vectors — so a match past the cap was never seen. The status is still ok: what it returned is right, it is just not the whole. Absent otherwise. | [optional] 
**Error** | Pointer to **string** | Error is the failure text from a leg whose status is degraded — the reason a configured backend could not answer. Absent otherwise. | [optional] 
**Hits** | Pointer to **int64** | Hits is how many results this leg returned, counted BEFORE fusion, so it is not the number that survived into Fusion.Hits — fusion merges what both legs found and the caller&#39;s limit and offset then page it. 0 for a leg that did not run. | [optional] 
**Name** | Pointer to **string** | Name is which leg this reports: \&quot;index\&quot;, the lexical store, \&quot;vector\&quot;, the semantic one, \&quot;code\&quot;, the org&#39;s own repositories, or \&quot;rerank\&quot;, the relevance pass over the fused window. Match.Backend uses the same names. | [optional] 
**Status** | Pointer to **string** | Status is one of ok, degraded, disabled, skipped — four distinct operational facts that are never collapsed. It ran and answered; it is configured and FAILED (Error says how, and only this one is a fault); this deployment never provisioned it; or the request&#39;s mode excluded it. | [optional] 
**TookMs** | Pointer to **int64** | TookMS is how long this leg took, in milliseconds, timed around its own call and excluding fusion. 0 for a leg that was skipped or is disabled, since nothing was called. | [optional] 

## Methods

### NewSearchBackendStatus

`func NewSearchBackendStatus() *SearchBackendStatus`

NewSearchBackendStatus instantiates a new SearchBackendStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBackendStatusWithDefaults

`func NewSearchBackendStatusWithDefaults() *SearchBackendStatus`

NewSearchBackendStatusWithDefaults instantiates a new SearchBackendStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapped

`func (o *SearchBackendStatus) GetCapped() bool`

GetCapped returns the Capped field if non-nil, zero value otherwise.

### GetCappedOk

`func (o *SearchBackendStatus) GetCappedOk() (*bool, bool)`

GetCappedOk returns a tuple with the Capped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapped

`func (o *SearchBackendStatus) SetCapped(v bool)`

SetCapped sets Capped field to given value.

### HasCapped

`func (o *SearchBackendStatus) HasCapped() bool`

HasCapped returns a boolean if a field has been set.

### GetError

`func (o *SearchBackendStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchBackendStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchBackendStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SearchBackendStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHits

`func (o *SearchBackendStatus) GetHits() int64`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchBackendStatus) GetHitsOk() (*int64, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchBackendStatus) SetHits(v int64)`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchBackendStatus) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetName

`func (o *SearchBackendStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchBackendStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchBackendStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchBackendStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *SearchBackendStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchBackendStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchBackendStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchBackendStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTookMs

`func (o *SearchBackendStatus) GetTookMs() int64`

GetTookMs returns the TookMs field if non-nil, zero value otherwise.

### GetTookMsOk

`func (o *SearchBackendStatus) GetTookMsOk() (*int64, bool)`

GetTookMsOk returns a tuple with the TookMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTookMs

`func (o *SearchBackendStatus) SetTookMs(v int64)`

SetTookMs sets TookMs field to given value.

### HasTookMs

`func (o *SearchBackendStatus) HasTookMs() bool`

HasTookMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


