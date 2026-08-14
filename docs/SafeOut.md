# SafeOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsignRef** | Pointer to **string** | EsignRef is the provider&#39;s reference for the signature request. | [optional] 
**Provider** | Pointer to **string** | Provider is the wired e-signature provider&#39;s name. | [optional] 

## Methods

### NewSafeOut

`func NewSafeOut() *SafeOut`

NewSafeOut instantiates a new SafeOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeOutWithDefaults

`func NewSafeOutWithDefaults() *SafeOut`

NewSafeOutWithDefaults instantiates a new SafeOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsignRef

`func (o *SafeOut) GetEsignRef() string`

GetEsignRef returns the EsignRef field if non-nil, zero value otherwise.

### GetEsignRefOk

`func (o *SafeOut) GetEsignRefOk() (*string, bool)`

GetEsignRefOk returns a tuple with the EsignRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsignRef

`func (o *SafeOut) SetEsignRef(v string)`

SetEsignRef sets EsignRef field to given value.

### HasEsignRef

`func (o *SafeOut) HasEsignRef() bool`

HasEsignRef returns a boolean if a field has been set.

### GetProvider

`func (o *SafeOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SafeOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SafeOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SafeOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


