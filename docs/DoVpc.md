# DoVpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Subnets** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewDoVpc

`func NewDoVpc() *DoVpc`

NewDoVpc instantiates a new DoVpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoVpcWithDefaults

`func NewDoVpcWithDefaults() *DoVpc`

NewDoVpcWithDefaults instantiates a new DoVpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DoVpc) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DoVpc) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DoVpc) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DoVpc) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DoVpc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DoVpc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DoVpc) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DoVpc) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCidr

`func (o *DoVpc) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *DoVpc) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *DoVpc) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *DoVpc) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetRegion

`func (o *DoVpc) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DoVpc) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DoVpc) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DoVpc) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSubnets

`func (o *DoVpc) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *DoVpc) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *DoVpc) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *DoVpc) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetStatus

`func (o *DoVpc) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DoVpc) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DoVpc) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DoVpc) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


