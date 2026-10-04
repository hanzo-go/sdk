# DataroomDataroomRoomDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the room was created, in unix milliseconds. | [optional] 
**Description** | Pointer to **string** | Description is the room&#39;s description, null when none was given. | [optional] 
**Documents** | Pointer to [**[]DataroomDataroomMember**](DataroomDataroomMember.md) | Documents is every document in the room, in the order a visitor sees them. | [optional] 
**Id** | Pointer to **string** | ID is the room id. | [optional] 
**Name** | Pointer to **string** | Name is the room&#39;s display name. | [optional] 
**PId** | Pointer to **string** | PId is the room&#39;s short public identifier. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the room last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomDataroomRoomDetail

`func NewDataroomDataroomRoomDetail() *DataroomDataroomRoomDetail`

NewDataroomDataroomRoomDetail instantiates a new DataroomDataroomRoomDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomRoomDetailWithDefaults

`func NewDataroomDataroomRoomDetailWithDefaults() *DataroomDataroomRoomDetail`

NewDataroomDataroomRoomDetailWithDefaults instantiates a new DataroomDataroomRoomDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DataroomDataroomRoomDetail) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataroomDataroomRoomDetail) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataroomDataroomRoomDetail) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataroomDataroomRoomDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DataroomDataroomRoomDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataroomDataroomRoomDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataroomDataroomRoomDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataroomDataroomRoomDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocuments

`func (o *DataroomDataroomRoomDetail) GetDocuments() []DataroomDataroomMember`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DataroomDataroomRoomDetail) GetDocumentsOk() (*[]DataroomDataroomMember, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DataroomDataroomRoomDetail) SetDocuments(v []DataroomDataroomMember)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DataroomDataroomRoomDetail) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetId

`func (o *DataroomDataroomRoomDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomDataroomRoomDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomDataroomRoomDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomDataroomRoomDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataroomDataroomRoomDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomDataroomRoomDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomDataroomRoomDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomDataroomRoomDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPId

`func (o *DataroomDataroomRoomDetail) GetPId() string`

GetPId returns the PId field if non-nil, zero value otherwise.

### GetPIdOk

`func (o *DataroomDataroomRoomDetail) GetPIdOk() (*string, bool)`

GetPIdOk returns a tuple with the PId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPId

`func (o *DataroomDataroomRoomDetail) SetPId(v string)`

SetPId sets PId field to given value.

### HasPId

`func (o *DataroomDataroomRoomDetail) HasPId() bool`

HasPId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomDataroomRoomDetail) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomDataroomRoomDetail) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomDataroomRoomDetail) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomDataroomRoomDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


