# TeamTeamTeamspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | Archived reports that the teamspace is closed: its documents still read, and none can be created, moved into it, renamed or commented on. | [optional] 
**Id** | Pointer to **string** | ID is the teamspace document&#39;s id, the value teamDoc.teamspace names. | [optional] 
**Name** | Pointer to **string** | Name is what a person sees in a sidebar. | [optional] 
**Private** | Pointer to **bool** | Private reports that the teamspace is open to its members only. | [optional] 

## Methods

### NewTeamTeamTeamspace

`func NewTeamTeamTeamspace() *TeamTeamTeamspace`

NewTeamTeamTeamspace instantiates a new TeamTeamTeamspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamTeamspaceWithDefaults

`func NewTeamTeamTeamspaceWithDefaults() *TeamTeamTeamspace`

NewTeamTeamTeamspaceWithDefaults instantiates a new TeamTeamTeamspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TeamTeamTeamspace) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TeamTeamTeamspace) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TeamTeamTeamspace) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *TeamTeamTeamspace) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamTeamspace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamTeamspace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamTeamspace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamTeamspace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TeamTeamTeamspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamTeamspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamTeamspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamTeamspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivate

`func (o *TeamTeamTeamspace) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TeamTeamTeamspace) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TeamTeamTeamspace) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *TeamTeamTeamspace) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


