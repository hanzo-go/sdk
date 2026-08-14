# RateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis names where the rates come from, so a published price can be explained rather than merely asserted. | [optional] 
**MicroUsdPerGbHour** | Pointer to **int32** | MicroUSDPerGBHour is the price of one GiB of memory for one hour, in millionths of a US dollar. | [optional] 
**MicroUsdPerVcpuHour** | Pointer to **int32** | MicroUSDPerVCPUHour is the price of one vCPU for one hour, in millionths of a US dollar. | [optional] 

## Methods

### NewRateCard

`func NewRateCard() *RateCard`

NewRateCard instantiates a new RateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardWithDefaults

`func NewRateCardWithDefaults() *RateCard`

NewRateCardWithDefaults instantiates a new RateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *RateCard) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *RateCard) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *RateCard) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *RateCard) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetMicroUsdPerGbHour

`func (o *RateCard) GetMicroUsdPerGbHour() int32`

GetMicroUsdPerGbHour returns the MicroUsdPerGbHour field if non-nil, zero value otherwise.

### GetMicroUsdPerGbHourOk

`func (o *RateCard) GetMicroUsdPerGbHourOk() (*int32, bool)`

GetMicroUsdPerGbHourOk returns a tuple with the MicroUsdPerGbHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroUsdPerGbHour

`func (o *RateCard) SetMicroUsdPerGbHour(v int32)`

SetMicroUsdPerGbHour sets MicroUsdPerGbHour field to given value.

### HasMicroUsdPerGbHour

`func (o *RateCard) HasMicroUsdPerGbHour() bool`

HasMicroUsdPerGbHour returns a boolean if a field has been set.

### GetMicroUsdPerVcpuHour

`func (o *RateCard) GetMicroUsdPerVcpuHour() int32`

GetMicroUsdPerVcpuHour returns the MicroUsdPerVcpuHour field if non-nil, zero value otherwise.

### GetMicroUsdPerVcpuHourOk

`func (o *RateCard) GetMicroUsdPerVcpuHourOk() (*int32, bool)`

GetMicroUsdPerVcpuHourOk returns a tuple with the MicroUsdPerVcpuHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroUsdPerVcpuHour

`func (o *RateCard) SetMicroUsdPerVcpuHour(v int32)`

SetMicroUsdPerVcpuHour sets MicroUsdPerVcpuHour field to given value.

### HasMicroUsdPerVcpuHour

`func (o *RateCard) HasMicroUsdPerVcpuHour() bool`

HasMicroUsdPerVcpuHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


