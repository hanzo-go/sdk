# IamTombstone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Founder** | Pointer to **string** | Founder carries the deleted org&#39;s founder, so an org counted against its founder&#39;s cap stays counted after it is deleted. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | the organization&#39;s name | [optional] 
**Owner** | Pointer to **string** | the admin registry, as for organizations | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIamTombstone

`func NewIamTombstone() *IamTombstone`

NewIamTombstone instantiates a new IamTombstone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamTombstoneWithDefaults

`func NewIamTombstoneWithDefaults() *IamTombstone`

NewIamTombstoneWithDefaults instantiates a new IamTombstone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IamTombstone) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamTombstone) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamTombstone) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamTombstone) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamTombstone) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamTombstone) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamTombstone) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamTombstone) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeleted

`func (o *IamTombstone) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamTombstone) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamTombstone) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamTombstone) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFounder

`func (o *IamTombstone) GetFounder() string`

GetFounder returns the Founder field if non-nil, zero value otherwise.

### GetFounderOk

`func (o *IamTombstone) GetFounderOk() (*string, bool)`

GetFounderOk returns a tuple with the Founder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFounder

`func (o *IamTombstone) SetFounder(v string)`

SetFounder sets Founder field to given value.

### HasFounder

`func (o *IamTombstone) HasFounder() bool`

HasFounder returns a boolean if a field has been set.

### GetId

`func (o *IamTombstone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamTombstone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamTombstone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamTombstone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IamTombstone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamTombstone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamTombstone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamTombstone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamTombstone) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamTombstone) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamTombstone) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamTombstone) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamTombstone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamTombstone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamTombstone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamTombstone) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


