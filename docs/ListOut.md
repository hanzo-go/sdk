# ListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]ProviderView**](ProviderView.md) | Providers is the whole catalog. Never null; [] when nothing is registered. | [optional] 

## Methods

### NewListOut

`func NewListOut() *ListOut`

NewListOut instantiates a new ListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOutWithDefaults

`func NewListOutWithDefaults() *ListOut`

NewListOutWithDefaults instantiates a new ListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *ListOut) GetProviders() []ProviderView`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ListOut) GetProvidersOk() (*[]ProviderView, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ListOut) SetProviders(v []ProviderView)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ListOut) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


