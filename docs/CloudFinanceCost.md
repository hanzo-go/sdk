# CloudFinanceCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** |  | [optional] 
**Digitalocean** | Pointer to [**CloudDoCost**](CloudDoCost.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**TotalCents** | Pointer to **int32** |  | [optional] 
**Vendors** | Pointer to [**[]CloudVendor**](CloudVendor.md) |  | [optional] 

## Methods

### NewCloudFinanceCost

`func NewCloudFinanceCost() *CloudFinanceCost`

NewCloudFinanceCost instantiates a new CloudFinanceCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFinanceCostWithDefaults

`func NewCloudFinanceCostWithDefaults() *CloudFinanceCost`

NewCloudFinanceCostWithDefaults instantiates a new CloudFinanceCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *CloudFinanceCost) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *CloudFinanceCost) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *CloudFinanceCost) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *CloudFinanceCost) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetDigitalocean

`func (o *CloudFinanceCost) GetDigitalocean() CloudDoCost`

GetDigitalocean returns the Digitalocean field if non-nil, zero value otherwise.

### GetDigitaloceanOk

`func (o *CloudFinanceCost) GetDigitaloceanOk() (*CloudDoCost, bool)`

GetDigitaloceanOk returns a tuple with the Digitalocean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalocean

`func (o *CloudFinanceCost) SetDigitalocean(v CloudDoCost)`

SetDigitalocean sets Digitalocean field to given value.

### HasDigitalocean

`func (o *CloudFinanceCost) HasDigitalocean() bool`

HasDigitalocean returns a boolean if a field has been set.

### GetError

`func (o *CloudFinanceCost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudFinanceCost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudFinanceCost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudFinanceCost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPeriod

`func (o *CloudFinanceCost) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CloudFinanceCost) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CloudFinanceCost) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CloudFinanceCost) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetTotalCents

`func (o *CloudFinanceCost) GetTotalCents() int32`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *CloudFinanceCost) GetTotalCentsOk() (*int32, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *CloudFinanceCost) SetTotalCents(v int32)`

SetTotalCents sets TotalCents field to given value.

### HasTotalCents

`func (o *CloudFinanceCost) HasTotalCents() bool`

HasTotalCents returns a boolean if a field has been set.

### GetVendors

`func (o *CloudFinanceCost) GetVendors() []CloudVendor`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *CloudFinanceCost) GetVendorsOk() (*[]CloudVendor, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *CloudFinanceCost) SetVendors(v []CloudVendor)`

SetVendors sets Vendors field to given value.

### HasVendors

`func (o *CloudFinanceCost) HasVendors() bool`

HasVendors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


