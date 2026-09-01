# DriveList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drives** | Pointer to [**[]DriveItem**](DriveItem.md) | Drives are the drives at the space&#39;s root. | [optional] 
**Space** | Pointer to **string** | Space is the space that was listed. | [optional] 
**Total** | Pointer to **int32** | Total is how many drives came back. The listing is BOUNDED, so it is what came back and not a count of what the space holds. | [optional] 

## Methods

### NewDriveList

`func NewDriveList() *DriveList`

NewDriveList instantiates a new DriveList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveListWithDefaults

`func NewDriveListWithDefaults() *DriveList`

NewDriveListWithDefaults instantiates a new DriveList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrives

`func (o *DriveList) GetDrives() []DriveItem`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *DriveList) GetDrivesOk() (*[]DriveItem, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *DriveList) SetDrives(v []DriveItem)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *DriveList) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetSpace

`func (o *DriveList) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *DriveList) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *DriveList) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *DriveList) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTotal

`func (o *DriveList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DriveList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DriveList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DriveList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


