# CloudVendorsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendors** | Pointer to [**[]CloudVendorRow**](CloudVendorRow.md) | Vendors is every vendor the org has recorded, canonical name ascending. | [optional] 

## Methods

### NewCloudVendorsOut

`func NewCloudVendorsOut() *CloudVendorsOut`

NewCloudVendorsOut instantiates a new CloudVendorsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVendorsOutWithDefaults

`func NewCloudVendorsOutWithDefaults() *CloudVendorsOut`

NewCloudVendorsOutWithDefaults instantiates a new CloudVendorsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendors

`func (o *CloudVendorsOut) GetVendors() []CloudVendorRow`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *CloudVendorsOut) GetVendorsOk() (*[]CloudVendorRow, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *CloudVendorsOut) SetVendors(v []CloudVendorRow)`

SetVendors sets Vendors field to given value.

### HasVendors

`func (o *CloudVendorsOut) HasVendors() bool`

HasVendors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


