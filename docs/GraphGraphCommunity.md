# GraphGraphCommunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID numbers the community by size, largest first, then by its lowest key. The same graph at the same instant and level numbers it the same way every time; another instant may not. | [optional] 
**Members** | Pointer to **[]string** | Members are the entities in it, in key order. | [optional] 
**Size** | Pointer to **int64** | Size is how many entities belong to it, whether or not all are listed. | [optional] 

## Methods

### NewGraphGraphCommunity

`func NewGraphGraphCommunity() *GraphGraphCommunity`

NewGraphGraphCommunity instantiates a new GraphGraphCommunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphCommunityWithDefaults

`func NewGraphGraphCommunityWithDefaults() *GraphGraphCommunity`

NewGraphGraphCommunityWithDefaults instantiates a new GraphGraphCommunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphGraphCommunity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphGraphCommunity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphGraphCommunity) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GraphGraphCommunity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMembers

`func (o *GraphGraphCommunity) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GraphGraphCommunity) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GraphGraphCommunity) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GraphGraphCommunity) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetSize

`func (o *GraphGraphCommunity) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GraphGraphCommunity) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GraphGraphCommunity) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GraphGraphCommunity) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


