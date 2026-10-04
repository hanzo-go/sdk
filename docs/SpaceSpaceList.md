# SpaceSpaceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spaces** | Pointer to [**[]SpaceSpaceItem**](SpaceSpaceItem.md) | Spaces are the caller org&#39;s spaces, oldest first as the store returns them. | [optional] 
**Total** | Pointer to **int64** | Total is how many spaces this org has. It equals len(spaces): the listing is not paged, because one bucket per (org, space) keeps an org&#39;s count small by construction, which is the whole reason a drive is a prefix and not a bucket. | [optional] 

## Methods

### NewSpaceSpaceList

`func NewSpaceSpaceList() *SpaceSpaceList`

NewSpaceSpaceList instantiates a new SpaceSpaceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceSpaceListWithDefaults

`func NewSpaceSpaceListWithDefaults() *SpaceSpaceList`

NewSpaceSpaceListWithDefaults instantiates a new SpaceSpaceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaces

`func (o *SpaceSpaceList) GetSpaces() []SpaceSpaceItem`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *SpaceSpaceList) GetSpacesOk() (*[]SpaceSpaceItem, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *SpaceSpaceList) SetSpaces(v []SpaceSpaceItem)`

SetSpaces sets Spaces field to given value.

### HasSpaces

`func (o *SpaceSpaceList) HasSpaces() bool`

HasSpaces returns a boolean if a field has been set.

### GetTotal

`func (o *SpaceSpaceList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SpaceSpaceList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SpaceSpaceList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SpaceSpaceList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


