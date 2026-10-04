# SpaceDriveList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drives** | Pointer to [**[]SpaceDriveItem**](SpaceDriveItem.md) | Drives are the drives at the space&#39;s root. | [optional] 
**Space** | Pointer to **string** | Space is the space that was listed. | [optional] 
**Total** | Pointer to **int64** | Total is how many drives came back. The listing is BOUNDED, so it is what came back and not a count of what the space holds. | [optional] 

## Methods

### NewSpaceDriveList

`func NewSpaceDriveList() *SpaceDriveList`

NewSpaceDriveList instantiates a new SpaceDriveList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceDriveListWithDefaults

`func NewSpaceDriveListWithDefaults() *SpaceDriveList`

NewSpaceDriveListWithDefaults instantiates a new SpaceDriveList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrives

`func (o *SpaceDriveList) GetDrives() []SpaceDriveItem`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *SpaceDriveList) GetDrivesOk() (*[]SpaceDriveItem, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *SpaceDriveList) SetDrives(v []SpaceDriveItem)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *SpaceDriveList) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetSpace

`func (o *SpaceDriveList) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *SpaceDriveList) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *SpaceDriveList) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *SpaceDriveList) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTotal

`func (o *SpaceDriveList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SpaceDriveList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SpaceDriveList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SpaceDriveList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


