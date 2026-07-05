# PricingBlockStoragePricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PricePerGBPerMonth** | Pointer to **float32** |  | [optional] 
**MinSizeGB** | Pointer to **int32** |  | [optional] 
**MaxSizeGB** | Pointer to **int32** |  | [optional] 
**Snapshots** | Pointer to [**PricingBlockStoragePricingSnapshots**](PricingBlockStoragePricingSnapshots.md) |  | [optional] 

## Methods

### NewPricingBlockStoragePricing

`func NewPricingBlockStoragePricing() *PricingBlockStoragePricing`

NewPricingBlockStoragePricing instantiates a new PricingBlockStoragePricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingBlockStoragePricingWithDefaults

`func NewPricingBlockStoragePricingWithDefaults() *PricingBlockStoragePricing`

NewPricingBlockStoragePricingWithDefaults instantiates a new PricingBlockStoragePricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPricePerGBPerMonth

`func (o *PricingBlockStoragePricing) GetPricePerGBPerMonth() float32`

GetPricePerGBPerMonth returns the PricePerGBPerMonth field if non-nil, zero value otherwise.

### GetPricePerGBPerMonthOk

`func (o *PricingBlockStoragePricing) GetPricePerGBPerMonthOk() (*float32, bool)`

GetPricePerGBPerMonthOk returns a tuple with the PricePerGBPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerGBPerMonth

`func (o *PricingBlockStoragePricing) SetPricePerGBPerMonth(v float32)`

SetPricePerGBPerMonth sets PricePerGBPerMonth field to given value.

### HasPricePerGBPerMonth

`func (o *PricingBlockStoragePricing) HasPricePerGBPerMonth() bool`

HasPricePerGBPerMonth returns a boolean if a field has been set.

### GetMinSizeGB

`func (o *PricingBlockStoragePricing) GetMinSizeGB() int32`

GetMinSizeGB returns the MinSizeGB field if non-nil, zero value otherwise.

### GetMinSizeGBOk

`func (o *PricingBlockStoragePricing) GetMinSizeGBOk() (*int32, bool)`

GetMinSizeGBOk returns a tuple with the MinSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSizeGB

`func (o *PricingBlockStoragePricing) SetMinSizeGB(v int32)`

SetMinSizeGB sets MinSizeGB field to given value.

### HasMinSizeGB

`func (o *PricingBlockStoragePricing) HasMinSizeGB() bool`

HasMinSizeGB returns a boolean if a field has been set.

### GetMaxSizeGB

`func (o *PricingBlockStoragePricing) GetMaxSizeGB() int32`

GetMaxSizeGB returns the MaxSizeGB field if non-nil, zero value otherwise.

### GetMaxSizeGBOk

`func (o *PricingBlockStoragePricing) GetMaxSizeGBOk() (*int32, bool)`

GetMaxSizeGBOk returns a tuple with the MaxSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSizeGB

`func (o *PricingBlockStoragePricing) SetMaxSizeGB(v int32)`

SetMaxSizeGB sets MaxSizeGB field to given value.

### HasMaxSizeGB

`func (o *PricingBlockStoragePricing) HasMaxSizeGB() bool`

HasMaxSizeGB returns a boolean if a field has been set.

### GetSnapshots

`func (o *PricingBlockStoragePricing) GetSnapshots() PricingBlockStoragePricingSnapshots`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *PricingBlockStoragePricing) GetSnapshotsOk() (*PricingBlockStoragePricingSnapshots, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *PricingBlockStoragePricing) SetSnapshots(v PricingBlockStoragePricingSnapshots)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *PricingBlockStoragePricing) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


