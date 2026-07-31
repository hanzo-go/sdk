# CloudVpcList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpcs** | Pointer to [**[]CloudVpcView**](CloudVpcView.md) | VPCs are the caller org&#39;s VPCs under their friendly names. | [optional] 

## Methods

### NewCloudVpcList

`func NewCloudVpcList() *CloudVpcList`

NewCloudVpcList instantiates a new CloudVpcList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVpcListWithDefaults

`func NewCloudVpcListWithDefaults() *CloudVpcList`

NewCloudVpcListWithDefaults instantiates a new CloudVpcList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcs

`func (o *CloudVpcList) GetVpcs() []CloudVpcView`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *CloudVpcList) GetVpcsOk() (*[]CloudVpcView, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *CloudVpcList) SetVpcs(v []CloudVpcView)`

SetVpcs sets Vpcs field to given value.

### HasVpcs

`func (o *CloudVpcList) HasVpcs() bool`

HasVpcs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


