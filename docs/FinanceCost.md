# FinanceCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** |  | [optional] 
**Digitalocean** | Pointer to [**DoCost**](DoCost.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**TotalCents** | Pointer to **int32** |  | [optional] 
**Vendors** | Pointer to [**[]Vendor**](Vendor.md) |  | [optional] 

## Methods

### NewFinanceCost

`func NewFinanceCost() *FinanceCost`

NewFinanceCost instantiates a new FinanceCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinanceCostWithDefaults

`func NewFinanceCostWithDefaults() *FinanceCost`

NewFinanceCostWithDefaults instantiates a new FinanceCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *FinanceCost) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *FinanceCost) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *FinanceCost) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *FinanceCost) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetDigitalocean

`func (o *FinanceCost) GetDigitalocean() DoCost`

GetDigitalocean returns the Digitalocean field if non-nil, zero value otherwise.

### GetDigitaloceanOk

`func (o *FinanceCost) GetDigitaloceanOk() (*DoCost, bool)`

GetDigitaloceanOk returns a tuple with the Digitalocean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalocean

`func (o *FinanceCost) SetDigitalocean(v DoCost)`

SetDigitalocean sets Digitalocean field to given value.

### HasDigitalocean

`func (o *FinanceCost) HasDigitalocean() bool`

HasDigitalocean returns a boolean if a field has been set.

### GetError

`func (o *FinanceCost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FinanceCost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FinanceCost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FinanceCost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPeriod

`func (o *FinanceCost) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *FinanceCost) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *FinanceCost) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *FinanceCost) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetTotalCents

`func (o *FinanceCost) GetTotalCents() int32`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *FinanceCost) GetTotalCentsOk() (*int32, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *FinanceCost) SetTotalCents(v int32)`

SetTotalCents sets TotalCents field to given value.

### HasTotalCents

`func (o *FinanceCost) HasTotalCents() bool`

HasTotalCents returns a boolean if a field has been set.

### GetVendors

`func (o *FinanceCost) GetVendors() []Vendor`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *FinanceCost) GetVendorsOk() (*[]Vendor, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *FinanceCost) SetVendors(v []Vendor)`

SetVendors sets Vendors field to given value.

### HasVendors

`func (o *FinanceCost) HasVendors() bool`

HasVendors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


