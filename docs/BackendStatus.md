# BackendStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Hits** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TookMs** | Pointer to **int32** |  | [optional] 

## Methods

### NewBackendStatus

`func NewBackendStatus() *BackendStatus`

NewBackendStatus instantiates a new BackendStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackendStatusWithDefaults

`func NewBackendStatusWithDefaults() *BackendStatus`

NewBackendStatusWithDefaults instantiates a new BackendStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *BackendStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BackendStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BackendStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BackendStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHits

`func (o *BackendStatus) GetHits() int32`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *BackendStatus) GetHitsOk() (*int32, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *BackendStatus) SetHits(v int32)`

SetHits sets Hits field to given value.

### HasHits

`func (o *BackendStatus) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetName

`func (o *BackendStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackendStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackendStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackendStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *BackendStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackendStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackendStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackendStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTookMs

`func (o *BackendStatus) GetTookMs() int32`

GetTookMs returns the TookMs field if non-nil, zero value otherwise.

### GetTookMsOk

`func (o *BackendStatus) GetTookMsOk() (*int32, bool)`

GetTookMsOk returns a tuple with the TookMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTookMs

`func (o *BackendStatus) SetTookMs(v int32)`

SetTookMs sets TookMs field to given value.

### HasTookMs

`func (o *BackendStatus) HasTookMs() bool`

HasTookMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


