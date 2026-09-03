# SpaceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spaces** | Pointer to [**[]SpaceItem**](SpaceItem.md) | Spaces are the caller org&#39;s spaces, oldest first as the store returns them. | [optional] 
**Total** | Pointer to **int64** | Total is how many spaces this org has. It equals len(spaces): the listing is not paged, because one bucket per (org, space) keeps an org&#39;s count small by construction, which is the whole reason a drive is a prefix and not a bucket. | [optional] 

## Methods

### NewSpaceList

`func NewSpaceList() *SpaceList`

NewSpaceList instantiates a new SpaceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceListWithDefaults

`func NewSpaceListWithDefaults() *SpaceList`

NewSpaceListWithDefaults instantiates a new SpaceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaces

`func (o *SpaceList) GetSpaces() []SpaceItem`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *SpaceList) GetSpacesOk() (*[]SpaceItem, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *SpaceList) SetSpaces(v []SpaceItem)`

SetSpaces sets Spaces field to given value.

### HasSpaces

`func (o *SpaceList) HasSpaces() bool`

HasSpaces returns a boolean if a field has been set.

### GetTotal

`func (o *SpaceList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SpaceList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SpaceList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SpaceList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


