# PricingPricingModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | Pointer to **time.Time** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Models** | Pointer to [**[]PricingModel**](PricingModel.md) |  | [optional] 

## Methods

### NewPricingPricingModelsResponse

`func NewPricingPricingModelsResponse() *PricingPricingModelsResponse`

NewPricingPricingModelsResponse instantiates a new PricingPricingModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingPricingModelsResponseWithDefaults

`func NewPricingPricingModelsResponseWithDefaults() *PricingPricingModelsResponse`

NewPricingPricingModelsResponseWithDefaults instantiates a new PricingPricingModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *PricingPricingModelsResponse) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PricingPricingModelsResponse) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PricingPricingModelsResponse) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PricingPricingModelsResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetTotal

`func (o *PricingPricingModelsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PricingPricingModelsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PricingPricingModelsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PricingPricingModelsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetModels

`func (o *PricingPricingModelsResponse) GetModels() []PricingModel`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *PricingPricingModelsResponse) GetModelsOk() (*[]PricingModel, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *PricingPricingModelsResponse) SetModels(v []PricingModel)`

SetModels sets Models field to given value.

### HasModels

`func (o *PricingPricingModelsResponse) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


