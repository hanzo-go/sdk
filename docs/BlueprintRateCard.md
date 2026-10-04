# BlueprintRateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis names where the rates come from, so a published price can be explained rather than merely asserted. | [optional] 
**MicroUsdPerGbHour** | Pointer to **int64** | MicroUSDPerGBHour is the price of one GiB of memory for one hour, in millionths of a US dollar. | [optional] 
**MicroUsdPerVcpuHour** | Pointer to **int64** | MicroUSDPerVCPUHour is the price of one vCPU for one hour, in millionths of a US dollar. | [optional] 

## Methods

### NewBlueprintRateCard

`func NewBlueprintRateCard() *BlueprintRateCard`

NewBlueprintRateCard instantiates a new BlueprintRateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintRateCardWithDefaults

`func NewBlueprintRateCardWithDefaults() *BlueprintRateCard`

NewBlueprintRateCardWithDefaults instantiates a new BlueprintRateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *BlueprintRateCard) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *BlueprintRateCard) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *BlueprintRateCard) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *BlueprintRateCard) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetMicroUsdPerGbHour

`func (o *BlueprintRateCard) GetMicroUsdPerGbHour() int64`

GetMicroUsdPerGbHour returns the MicroUsdPerGbHour field if non-nil, zero value otherwise.

### GetMicroUsdPerGbHourOk

`func (o *BlueprintRateCard) GetMicroUsdPerGbHourOk() (*int64, bool)`

GetMicroUsdPerGbHourOk returns a tuple with the MicroUsdPerGbHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroUsdPerGbHour

`func (o *BlueprintRateCard) SetMicroUsdPerGbHour(v int64)`

SetMicroUsdPerGbHour sets MicroUsdPerGbHour field to given value.

### HasMicroUsdPerGbHour

`func (o *BlueprintRateCard) HasMicroUsdPerGbHour() bool`

HasMicroUsdPerGbHour returns a boolean if a field has been set.

### GetMicroUsdPerVcpuHour

`func (o *BlueprintRateCard) GetMicroUsdPerVcpuHour() int64`

GetMicroUsdPerVcpuHour returns the MicroUsdPerVcpuHour field if non-nil, zero value otherwise.

### GetMicroUsdPerVcpuHourOk

`func (o *BlueprintRateCard) GetMicroUsdPerVcpuHourOk() (*int64, bool)`

GetMicroUsdPerVcpuHourOk returns a tuple with the MicroUsdPerVcpuHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroUsdPerVcpuHour

`func (o *BlueprintRateCard) SetMicroUsdPerVcpuHour(v int64)`

SetMicroUsdPerVcpuHour sets MicroUsdPerVcpuHour field to given value.

### HasMicroUsdPerVcpuHour

`func (o *BlueprintRateCard) HasMicroUsdPerVcpuHour() bool`

HasMicroUsdPerVcpuHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


