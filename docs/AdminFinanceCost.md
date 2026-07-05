# AdminFinanceCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**TotalCents** | Pointer to **int64** |  | [optional] 
**Vendors** | Pointer to [**[]AdminVendorCost**](AdminVendorCost.md) |  | [optional] 
**Digitalocean** | Pointer to [**AdminDoCost**](AdminDoCost.md) |  | [optional] 

## Methods

### NewAdminFinanceCost

`func NewAdminFinanceCost() *AdminFinanceCost`

NewAdminFinanceCost instantiates a new AdminFinanceCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminFinanceCostWithDefaults

`func NewAdminFinanceCostWithDefaults() *AdminFinanceCost`

NewAdminFinanceCostWithDefaults instantiates a new AdminFinanceCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *AdminFinanceCost) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *AdminFinanceCost) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *AdminFinanceCost) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *AdminFinanceCost) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetError

`func (o *AdminFinanceCost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AdminFinanceCost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AdminFinanceCost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AdminFinanceCost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPeriod

`func (o *AdminFinanceCost) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AdminFinanceCost) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AdminFinanceCost) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AdminFinanceCost) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetTotalCents

`func (o *AdminFinanceCost) GetTotalCents() int64`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *AdminFinanceCost) GetTotalCentsOk() (*int64, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *AdminFinanceCost) SetTotalCents(v int64)`

SetTotalCents sets TotalCents field to given value.

### HasTotalCents

`func (o *AdminFinanceCost) HasTotalCents() bool`

HasTotalCents returns a boolean if a field has been set.

### GetVendors

`func (o *AdminFinanceCost) GetVendors() []AdminVendorCost`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *AdminFinanceCost) GetVendorsOk() (*[]AdminVendorCost, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *AdminFinanceCost) SetVendors(v []AdminVendorCost)`

SetVendors sets Vendors field to given value.

### HasVendors

`func (o *AdminFinanceCost) HasVendors() bool`

HasVendors returns a boolean if a field has been set.

### GetDigitalocean

`func (o *AdminFinanceCost) GetDigitalocean() AdminDoCost`

GetDigitalocean returns the Digitalocean field if non-nil, zero value otherwise.

### GetDigitaloceanOk

`func (o *AdminFinanceCost) GetDigitaloceanOk() (*AdminDoCost, bool)`

GetDigitaloceanOk returns a tuple with the Digitalocean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalocean

`func (o *AdminFinanceCost) SetDigitalocean(v AdminDoCost)`

SetDigitalocean sets Digitalocean field to given value.

### HasDigitalocean

`func (o *AdminFinanceCost) HasDigitalocean() bool`

HasDigitalocean returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


