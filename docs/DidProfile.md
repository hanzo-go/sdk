# DidProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**LinkedIdentities** | Pointer to [**[]DidLinkedIdentity**](DidLinkedIdentity.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDidProfile

`func NewDidProfile() *DidProfile`

NewDidProfile instantiates a new DidProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDidProfileWithDefaults

`func NewDidProfileWithDefaults() *DidProfile`

NewDidProfileWithDefaults instantiates a new DidProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DidProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DidProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DidProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DidProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DidProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DidProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DidProfile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DidProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DidProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DidProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DidProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DidProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *DidProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DidProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DidProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DidProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatar

`func (o *DidProfile) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *DidProfile) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *DidProfile) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *DidProfile) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetOrganization

`func (o *DidProfile) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DidProfile) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DidProfile) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DidProfile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetTeams

`func (o *DidProfile) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *DidProfile) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *DidProfile) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *DidProfile) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetMetadata

`func (o *DidProfile) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DidProfile) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DidProfile) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DidProfile) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLinkedIdentities

`func (o *DidProfile) GetLinkedIdentities() []DidLinkedIdentity`

GetLinkedIdentities returns the LinkedIdentities field if non-nil, zero value otherwise.

### GetLinkedIdentitiesOk

`func (o *DidProfile) GetLinkedIdentitiesOk() (*[]DidLinkedIdentity, bool)`

GetLinkedIdentitiesOk returns a tuple with the LinkedIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIdentities

`func (o *DidProfile) SetLinkedIdentities(v []DidLinkedIdentity)`

SetLinkedIdentities sets LinkedIdentities field to given value.

### HasLinkedIdentities

`func (o *DidProfile) HasLinkedIdentities() bool`

HasLinkedIdentities returns a boolean if a field has been set.

### GetStatus

`func (o *DidProfile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DidProfile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DidProfile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DidProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DidProfile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DidProfile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DidProfile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DidProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DidProfile) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DidProfile) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DidProfile) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DidProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


