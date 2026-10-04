# TeamTeamDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | Author is the account uuid that created it. | [optional] 
**Collaborator** | Pointer to **string** | Collaborator is the documentId to open the body with on the /v1/team/collaborator Y.js socket: \&quot;&lt;space&gt;|document:class:Document|&lt;id&gt;|content\&quot;. | [optional] 
**Comments** | Pointer to **int64** | Comments is how many comments it has. | [optional] 
**CreatedOn** | Pointer to **int64** | CreatedOn is when it was created, unix milliseconds. | [optional] 
**Id** | Pointer to **string** | ID is the document&#39;s own id. | [optional] 
**ModifiedOn** | Pointer to **int64** | ModifiedOn is when its title or place last changed, unix milliseconds. Edits to the body happen on the collaborator socket and do not move it. | [optional] 
**Parent** | Pointer to **string** | Parent is the document this one is nested under. Empty for a document at the top of its teamspace. | [optional] 
**Rank** | Pointer to **string** | Rank orders siblings: compare as plain strings, ascending. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid holding it. | [optional] 
**Teamspace** | Pointer to **string** | Teamspace is the teamspace it belongs to. | [optional] 
**Title** | Pointer to **string** | Title is the document&#39;s title. | [optional] 

## Methods

### NewTeamTeamDoc

`func NewTeamTeamDoc() *TeamTeamDoc`

NewTeamTeamDoc instantiates a new TeamTeamDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamDocWithDefaults

`func NewTeamTeamDocWithDefaults() *TeamTeamDoc`

NewTeamTeamDocWithDefaults instantiates a new TeamTeamDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *TeamTeamDoc) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TeamTeamDoc) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TeamTeamDoc) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *TeamTeamDoc) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCollaborator

`func (o *TeamTeamDoc) GetCollaborator() string`

GetCollaborator returns the Collaborator field if non-nil, zero value otherwise.

### GetCollaboratorOk

`func (o *TeamTeamDoc) GetCollaboratorOk() (*string, bool)`

GetCollaboratorOk returns a tuple with the Collaborator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborator

`func (o *TeamTeamDoc) SetCollaborator(v string)`

SetCollaborator sets Collaborator field to given value.

### HasCollaborator

`func (o *TeamTeamDoc) HasCollaborator() bool`

HasCollaborator returns a boolean if a field has been set.

### GetComments

`func (o *TeamTeamDoc) GetComments() int64`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *TeamTeamDoc) GetCommentsOk() (*int64, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *TeamTeamDoc) SetComments(v int64)`

SetComments sets Comments field to given value.

### HasComments

`func (o *TeamTeamDoc) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedOn

`func (o *TeamTeamDoc) GetCreatedOn() int64`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *TeamTeamDoc) GetCreatedOnOk() (*int64, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *TeamTeamDoc) SetCreatedOn(v int64)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *TeamTeamDoc) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamDoc) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamDoc) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamDoc) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamDoc) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedOn

`func (o *TeamTeamDoc) GetModifiedOn() int64`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *TeamTeamDoc) GetModifiedOnOk() (*int64, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *TeamTeamDoc) SetModifiedOn(v int64)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *TeamTeamDoc) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetParent

`func (o *TeamTeamDoc) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TeamTeamDoc) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TeamTeamDoc) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TeamTeamDoc) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRank

`func (o *TeamTeamDoc) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *TeamTeamDoc) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *TeamTeamDoc) SetRank(v string)`

SetRank sets Rank field to given value.

### HasRank

`func (o *TeamTeamDoc) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamDoc) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamDoc) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamDoc) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamDoc) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTeamspace

`func (o *TeamTeamDoc) GetTeamspace() string`

GetTeamspace returns the Teamspace field if non-nil, zero value otherwise.

### GetTeamspaceOk

`func (o *TeamTeamDoc) GetTeamspaceOk() (*string, bool)`

GetTeamspaceOk returns a tuple with the Teamspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamspace

`func (o *TeamTeamDoc) SetTeamspace(v string)`

SetTeamspace sets Teamspace field to given value.

### HasTeamspace

`func (o *TeamTeamDoc) HasTeamspace() bool`

HasTeamspace returns a boolean if a field has been set.

### GetTitle

`func (o *TeamTeamDoc) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TeamTeamDoc) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TeamTeamDoc) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TeamTeamDoc) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


