# DirectoryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**DirectoryData**](DirectoryData.md) | Data is the affiliate directory and its tally. | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewDirectoryOut

`func NewDirectoryOut() *DirectoryOut`

NewDirectoryOut instantiates a new DirectoryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryOutWithDefaults

`func NewDirectoryOutWithDefaults() *DirectoryOut`

NewDirectoryOutWithDefaults instantiates a new DirectoryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DirectoryOut) GetData() DirectoryData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DirectoryOut) GetDataOk() (*DirectoryData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DirectoryOut) SetData(v DirectoryData)`

SetData sets Data field to given value.

### HasData

`func (o *DirectoryOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *DirectoryOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *DirectoryOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *DirectoryOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *DirectoryOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *DirectoryOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DirectoryOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DirectoryOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DirectoryOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


