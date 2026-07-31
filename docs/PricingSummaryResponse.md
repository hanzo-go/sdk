# PricingSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | Pointer to **time.Time** |  | [optional] 
**ZenModels** | Pointer to **int32** |  | [optional] 
**ThirdPartyModels** | Pointer to **int32** |  | [optional] 
**TotalModels** | Pointer to **int32** |  | [optional] 
**Providers** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPricingSummaryResponse

`func NewPricingSummaryResponse() *PricingSummaryResponse`

NewPricingSummaryResponse instantiates a new PricingSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingSummaryResponseWithDefaults

`func NewPricingSummaryResponseWithDefaults() *PricingSummaryResponse`

NewPricingSummaryResponseWithDefaults instantiates a new PricingSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *PricingSummaryResponse) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PricingSummaryResponse) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PricingSummaryResponse) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PricingSummaryResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetZenModels

`func (o *PricingSummaryResponse) GetZenModels() int32`

GetZenModels returns the ZenModels field if non-nil, zero value otherwise.

### GetZenModelsOk

`func (o *PricingSummaryResponse) GetZenModelsOk() (*int32, bool)`

GetZenModelsOk returns a tuple with the ZenModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZenModels

`func (o *PricingSummaryResponse) SetZenModels(v int32)`

SetZenModels sets ZenModels field to given value.

### HasZenModels

`func (o *PricingSummaryResponse) HasZenModels() bool`

HasZenModels returns a boolean if a field has been set.

### GetThirdPartyModels

`func (o *PricingSummaryResponse) GetThirdPartyModels() int32`

GetThirdPartyModels returns the ThirdPartyModels field if non-nil, zero value otherwise.

### GetThirdPartyModelsOk

`func (o *PricingSummaryResponse) GetThirdPartyModelsOk() (*int32, bool)`

GetThirdPartyModelsOk returns a tuple with the ThirdPartyModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyModels

`func (o *PricingSummaryResponse) SetThirdPartyModels(v int32)`

SetThirdPartyModels sets ThirdPartyModels field to given value.

### HasThirdPartyModels

`func (o *PricingSummaryResponse) HasThirdPartyModels() bool`

HasThirdPartyModels returns a boolean if a field has been set.

### GetTotalModels

`func (o *PricingSummaryResponse) GetTotalModels() int32`

GetTotalModels returns the TotalModels field if non-nil, zero value otherwise.

### GetTotalModelsOk

`func (o *PricingSummaryResponse) GetTotalModelsOk() (*int32, bool)`

GetTotalModelsOk returns a tuple with the TotalModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalModels

`func (o *PricingSummaryResponse) SetTotalModels(v int32)`

SetTotalModels sets TotalModels field to given value.

### HasTotalModels

`func (o *PricingSummaryResponse) HasTotalModels() bool`

HasTotalModels returns a boolean if a field has been set.

### GetProviders

`func (o *PricingSummaryResponse) GetProviders() map[string]interface{}`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *PricingSummaryResponse) GetProvidersOk() (*map[string]interface{}, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *PricingSummaryResponse) SetProviders(v map[string]interface{})`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *PricingSummaryResponse) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


