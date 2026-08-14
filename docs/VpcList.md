# VpcList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpcs** | Pointer to [**[]VpcView**](VpcView.md) | VPCs are the caller org&#39;s VPCs under their friendly names. | [optional] 

## Methods

### NewVpcList

`func NewVpcList() *VpcList`

NewVpcList instantiates a new VpcList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcListWithDefaults

`func NewVpcListWithDefaults() *VpcList`

NewVpcListWithDefaults instantiates a new VpcList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcs

`func (o *VpcList) GetVpcs() []VpcView`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *VpcList) GetVpcsOk() (*[]VpcView, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *VpcList) SetVpcs(v []VpcView)`

SetVpcs sets Vpcs field to given value.

### HasVpcs

`func (o *VpcList) HasVpcs() bool`

HasVpcs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


