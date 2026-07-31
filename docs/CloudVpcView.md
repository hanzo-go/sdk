# CloudVpcView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subnets** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudVpcView

`func NewCloudVpcView() *CloudVpcView`

NewCloudVpcView instantiates a new CloudVpcView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVpcViewWithDefaults

`func NewCloudVpcViewWithDefaults() *CloudVpcView`

NewCloudVpcViewWithDefaults instantiates a new CloudVpcView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *CloudVpcView) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CloudVpcView) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CloudVpcView) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CloudVpcView) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetId

`func (o *CloudVpcView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudVpcView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudVpcView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudVpcView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudVpcView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudVpcView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudVpcView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudVpcView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CloudVpcView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudVpcView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudVpcView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudVpcView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *CloudVpcView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudVpcView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudVpcView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudVpcView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubnets

`func (o *CloudVpcView) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CloudVpcView) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CloudVpcView) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *CloudVpcView) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


