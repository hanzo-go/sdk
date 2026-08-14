# DataroomRoomDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when the room was created, in unix milliseconds. | [optional] 
**Description** | Pointer to **string** | Description is the room&#39;s description, null when none was given. | [optional] 
**Documents** | Pointer to [**[]DataroomMember**](DataroomMember.md) | Documents is every document in the room, in the order a visitor sees them. | [optional] 
**Id** | Pointer to **string** | ID is the room id. | [optional] 
**Name** | Pointer to **string** | Name is the room&#39;s display name. | [optional] 
**PId** | Pointer to **string** | PId is the room&#39;s short public identifier. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the room last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomRoomDetail

`func NewDataroomRoomDetail() *DataroomRoomDetail`

NewDataroomRoomDetail instantiates a new DataroomRoomDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomRoomDetailWithDefaults

`func NewDataroomRoomDetailWithDefaults() *DataroomRoomDetail`

NewDataroomRoomDetailWithDefaults instantiates a new DataroomRoomDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DataroomRoomDetail) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataroomRoomDetail) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataroomRoomDetail) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataroomRoomDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DataroomRoomDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataroomRoomDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataroomRoomDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataroomRoomDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocuments

`func (o *DataroomRoomDetail) GetDocuments() []DataroomMember`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DataroomRoomDetail) GetDocumentsOk() (*[]DataroomMember, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DataroomRoomDetail) SetDocuments(v []DataroomMember)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DataroomRoomDetail) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetId

`func (o *DataroomRoomDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomRoomDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomRoomDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomRoomDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataroomRoomDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomRoomDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomRoomDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomRoomDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPId

`func (o *DataroomRoomDetail) GetPId() string`

GetPId returns the PId field if non-nil, zero value otherwise.

### GetPIdOk

`func (o *DataroomRoomDetail) GetPIdOk() (*string, bool)`

GetPIdOk returns a tuple with the PId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPId

`func (o *DataroomRoomDetail) SetPId(v string)`

SetPId sets PId field to given value.

### HasPId

`func (o *DataroomRoomDetail) HasPId() bool`

HasPId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomRoomDetail) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomRoomDetail) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomRoomDetail) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomRoomDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


