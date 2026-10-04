# DataroomDataroomRoom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the room was created, in unix milliseconds. | [optional] 
**Description** | Pointer to **string** | Description is the room&#39;s description, null when none was given. | [optional] 
**Id** | Pointer to **string** | ID is the room id, which is what other dataroom calls address it by. | [optional] 
**Name** | Pointer to **string** | Name is the room&#39;s display name. | [optional] 
**PId** | Pointer to **string** | PId is the room&#39;s short public identifier, unique within the tenant. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the room last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomDataroomRoom

`func NewDataroomDataroomRoom() *DataroomDataroomRoom`

NewDataroomDataroomRoom instantiates a new DataroomDataroomRoom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomRoomWithDefaults

`func NewDataroomDataroomRoomWithDefaults() *DataroomDataroomRoom`

NewDataroomDataroomRoomWithDefaults instantiates a new DataroomDataroomRoom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DataroomDataroomRoom) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataroomDataroomRoom) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataroomDataroomRoom) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataroomDataroomRoom) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DataroomDataroomRoom) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataroomDataroomRoom) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataroomDataroomRoom) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataroomDataroomRoom) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DataroomDataroomRoom) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomDataroomRoom) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomDataroomRoom) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomDataroomRoom) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataroomDataroomRoom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomDataroomRoom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomDataroomRoom) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomDataroomRoom) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPId

`func (o *DataroomDataroomRoom) GetPId() string`

GetPId returns the PId field if non-nil, zero value otherwise.

### GetPIdOk

`func (o *DataroomDataroomRoom) GetPIdOk() (*string, bool)`

GetPIdOk returns a tuple with the PId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPId

`func (o *DataroomDataroomRoom) SetPId(v string)`

SetPId sets PId field to given value.

### HasPId

`func (o *DataroomDataroomRoom) HasPId() bool`

HasPId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomDataroomRoom) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomDataroomRoom) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomDataroomRoom) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomDataroomRoom) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


