# SeoRateOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rates** | Pointer to [**[]SeoCharge**](SeoCharge.md) | Rates is one row per op on this surface. | [optional] 

## Methods

### NewSeoRateOut

`func NewSeoRateOut() *SeoRateOut`

NewSeoRateOut instantiates a new SeoRateOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoRateOutWithDefaults

`func NewSeoRateOutWithDefaults() *SeoRateOut`

NewSeoRateOutWithDefaults instantiates a new SeoRateOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRates

`func (o *SeoRateOut) GetRates() []SeoCharge`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *SeoRateOut) GetRatesOk() (*[]SeoCharge, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *SeoRateOut) SetRates(v []SeoCharge)`

SetRates sets Rates field to given value.

### HasRates

`func (o *SeoRateOut) HasRates() bool`

HasRates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


