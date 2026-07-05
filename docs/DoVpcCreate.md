# DoVpcCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Friendly name; must match ^[a-z0-9]([a-z0-9-]{0,38}[a-z0-9])?$ | 
**Region** | **string** | DO region slug | 
**IpRange** | Pointer to **string** | CIDR (empty → DO auto-assigns) | [optional] 

## Methods

### NewDoVpcCreate

`func NewDoVpcCreate(name string, region string, ) *DoVpcCreate`

NewDoVpcCreate instantiates a new DoVpcCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoVpcCreateWithDefaults

`func NewDoVpcCreateWithDefaults() *DoVpcCreate`

NewDoVpcCreateWithDefaults instantiates a new DoVpcCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DoVpcCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DoVpcCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DoVpcCreate) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *DoVpcCreate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DoVpcCreate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DoVpcCreate) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetIpRange

`func (o *DoVpcCreate) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *DoVpcCreate) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *DoVpcCreate) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *DoVpcCreate) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


