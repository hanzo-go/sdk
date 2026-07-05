# DidTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Members** | Pointer to **int32** |  | [optional] 
**ParentTeam** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDidTeam

`func NewDidTeam() *DidTeam`

NewDidTeam instantiates a new DidTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDidTeamWithDefaults

`func NewDidTeamWithDefaults() *DidTeam`

NewDidTeamWithDefaults instantiates a new DidTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DidTeam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DidTeam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DidTeam) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DidTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DidTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DidTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DidTeam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DidTeam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *DidTeam) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DidTeam) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DidTeam) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DidTeam) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetDescription

`func (o *DidTeam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DidTeam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DidTeam) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DidTeam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMembers

`func (o *DidTeam) GetMembers() int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DidTeam) GetMembersOk() (*int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DidTeam) SetMembers(v int32)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *DidTeam) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetParentTeam

`func (o *DidTeam) GetParentTeam() string`

GetParentTeam returns the ParentTeam field if non-nil, zero value otherwise.

### GetParentTeamOk

`func (o *DidTeam) GetParentTeamOk() (*string, bool)`

GetParentTeamOk returns a tuple with the ParentTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTeam

`func (o *DidTeam) SetParentTeam(v string)`

SetParentTeam sets ParentTeam field to given value.

### HasParentTeam

`func (o *DidTeam) HasParentTeam() bool`

HasParentTeam returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DidTeam) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DidTeam) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DidTeam) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DidTeam) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


