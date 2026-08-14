# VendorsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendors** | Pointer to [**[]VendorRow**](VendorRow.md) | Vendors is every vendor the org has recorded, canonical name ascending. | [optional] 

## Methods

### NewVendorsOut

`func NewVendorsOut() *VendorsOut`

NewVendorsOut instantiates a new VendorsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorsOutWithDefaults

`func NewVendorsOutWithDefaults() *VendorsOut`

NewVendorsOutWithDefaults instantiates a new VendorsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendors

`func (o *VendorsOut) GetVendors() []VendorRow`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *VendorsOut) GetVendorsOk() (*[]VendorRow, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *VendorsOut) SetVendors(v []VendorRow)`

SetVendors sets Vendors field to given value.

### HasVendors

`func (o *VendorsOut) HasVendors() bool`

HasVendors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


