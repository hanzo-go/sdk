# DataroomRoom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when the room was created, in unix milliseconds. | [optional] 
**Description** | Pointer to **string** | Description is the room&#39;s description, null when none was given. | [optional] 
**Id** | Pointer to **string** | ID is the room id, which is what other dataroom calls address it by. | [optional] 
**Name** | Pointer to **string** | Name is the room&#39;s display name. | [optional] 
**PId** | Pointer to **string** | PId is the room&#39;s short public identifier, unique within the tenant. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the room last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomRoom

`func NewDataroomRoom() *DataroomRoom`

NewDataroomRoom instantiates a new DataroomRoom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomRoomWithDefaults

`func NewDataroomRoomWithDefaults() *DataroomRoom`

NewDataroomRoomWithDefaults instantiates a new DataroomRoom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DataroomRoom) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataroomRoom) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataroomRoom) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataroomRoom) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DataroomRoom) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataroomRoom) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataroomRoom) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataroomRoom) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DataroomRoom) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomRoom) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomRoom) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomRoom) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataroomRoom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomRoom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomRoom) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomRoom) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPId

`func (o *DataroomRoom) GetPId() string`

GetPId returns the PId field if non-nil, zero value otherwise.

### GetPIdOk

`func (o *DataroomRoom) GetPIdOk() (*string, bool)`

GetPIdOk returns a tuple with the PId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPId

`func (o *DataroomRoom) SetPId(v string)`

SetPId sets PId field to given value.

### HasPId

`func (o *DataroomRoom) HasPId() bool`

HasPId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomRoom) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomRoom) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomRoom) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomRoom) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


