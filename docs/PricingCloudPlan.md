# PricingCloudPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Vcpus** | Pointer to **int32** |  | [optional] 
**MemoryGB** | Pointer to **int32** |  | [optional] 
**DiskGB** | Pointer to **int32** |  | [optional] 
**CpuType** | Pointer to **string** |  | [optional] 
**MaxVMs** | Pointer to **int32** |  | [optional] 
**PriceMonthly** | Pointer to **float32** |  | [optional] 
**PriceHourly** | Pointer to **float32** |  | [optional] 
**FreeTier** | Pointer to **bool** |  | [optional] 
**Popular** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPricingCloudPlan

`func NewPricingCloudPlan() *PricingCloudPlan`

NewPricingCloudPlan instantiates a new PricingCloudPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingCloudPlanWithDefaults

`func NewPricingCloudPlanWithDefaults() *PricingCloudPlan`

NewPricingCloudPlanWithDefaults instantiates a new PricingCloudPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PricingCloudPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricingCloudPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricingCloudPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PricingCloudPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PricingCloudPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PricingCloudPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PricingCloudPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PricingCloudPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PricingCloudPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PricingCloudPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PricingCloudPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PricingCloudPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVcpus

`func (o *PricingCloudPlan) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *PricingCloudPlan) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *PricingCloudPlan) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *PricingCloudPlan) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetMemoryGB

`func (o *PricingCloudPlan) GetMemoryGB() int32`

GetMemoryGB returns the MemoryGB field if non-nil, zero value otherwise.

### GetMemoryGBOk

`func (o *PricingCloudPlan) GetMemoryGBOk() (*int32, bool)`

GetMemoryGBOk returns a tuple with the MemoryGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGB

`func (o *PricingCloudPlan) SetMemoryGB(v int32)`

SetMemoryGB sets MemoryGB field to given value.

### HasMemoryGB

`func (o *PricingCloudPlan) HasMemoryGB() bool`

HasMemoryGB returns a boolean if a field has been set.

### GetDiskGB

`func (o *PricingCloudPlan) GetDiskGB() int32`

GetDiskGB returns the DiskGB field if non-nil, zero value otherwise.

### GetDiskGBOk

`func (o *PricingCloudPlan) GetDiskGBOk() (*int32, bool)`

GetDiskGBOk returns a tuple with the DiskGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGB

`func (o *PricingCloudPlan) SetDiskGB(v int32)`

SetDiskGB sets DiskGB field to given value.

### HasDiskGB

`func (o *PricingCloudPlan) HasDiskGB() bool`

HasDiskGB returns a boolean if a field has been set.

### GetCpuType

`func (o *PricingCloudPlan) GetCpuType() string`

GetCpuType returns the CpuType field if non-nil, zero value otherwise.

### GetCpuTypeOk

`func (o *PricingCloudPlan) GetCpuTypeOk() (*string, bool)`

GetCpuTypeOk returns a tuple with the CpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuType

`func (o *PricingCloudPlan) SetCpuType(v string)`

SetCpuType sets CpuType field to given value.

### HasCpuType

`func (o *PricingCloudPlan) HasCpuType() bool`

HasCpuType returns a boolean if a field has been set.

### GetMaxVMs

`func (o *PricingCloudPlan) GetMaxVMs() int32`

GetMaxVMs returns the MaxVMs field if non-nil, zero value otherwise.

### GetMaxVMsOk

`func (o *PricingCloudPlan) GetMaxVMsOk() (*int32, bool)`

GetMaxVMsOk returns a tuple with the MaxVMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVMs

`func (o *PricingCloudPlan) SetMaxVMs(v int32)`

SetMaxVMs sets MaxVMs field to given value.

### HasMaxVMs

`func (o *PricingCloudPlan) HasMaxVMs() bool`

HasMaxVMs returns a boolean if a field has been set.

### GetPriceMonthly

`func (o *PricingCloudPlan) GetPriceMonthly() float32`

GetPriceMonthly returns the PriceMonthly field if non-nil, zero value otherwise.

### GetPriceMonthlyOk

`func (o *PricingCloudPlan) GetPriceMonthlyOk() (*float32, bool)`

GetPriceMonthlyOk returns a tuple with the PriceMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMonthly

`func (o *PricingCloudPlan) SetPriceMonthly(v float32)`

SetPriceMonthly sets PriceMonthly field to given value.

### HasPriceMonthly

`func (o *PricingCloudPlan) HasPriceMonthly() bool`

HasPriceMonthly returns a boolean if a field has been set.

### GetPriceHourly

`func (o *PricingCloudPlan) GetPriceHourly() float32`

GetPriceHourly returns the PriceHourly field if non-nil, zero value otherwise.

### GetPriceHourlyOk

`func (o *PricingCloudPlan) GetPriceHourlyOk() (*float32, bool)`

GetPriceHourlyOk returns a tuple with the PriceHourly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHourly

`func (o *PricingCloudPlan) SetPriceHourly(v float32)`

SetPriceHourly sets PriceHourly field to given value.

### HasPriceHourly

`func (o *PricingCloudPlan) HasPriceHourly() bool`

HasPriceHourly returns a boolean if a field has been set.

### GetFreeTier

`func (o *PricingCloudPlan) GetFreeTier() bool`

GetFreeTier returns the FreeTier field if non-nil, zero value otherwise.

### GetFreeTierOk

`func (o *PricingCloudPlan) GetFreeTierOk() (*bool, bool)`

GetFreeTierOk returns a tuple with the FreeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTier

`func (o *PricingCloudPlan) SetFreeTier(v bool)`

SetFreeTier sets FreeTier field to given value.

### HasFreeTier

`func (o *PricingCloudPlan) HasFreeTier() bool`

HasFreeTier returns a boolean if a field has been set.

### GetPopular

`func (o *PricingCloudPlan) GetPopular() bool`

GetPopular returns the Popular field if non-nil, zero value otherwise.

### GetPopularOk

`func (o *PricingCloudPlan) GetPopularOk() (*bool, bool)`

GetPopularOk returns a tuple with the Popular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopular

`func (o *PricingCloudPlan) SetPopular(v bool)`

SetPopular sets Popular field to given value.

### HasPopular

`func (o *PricingCloudPlan) HasPopular() bool`

HasPopular returns a boolean if a field has been set.

### GetFeatures

`func (o *PricingCloudPlan) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PricingCloudPlan) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PricingCloudPlan) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PricingCloudPlan) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


