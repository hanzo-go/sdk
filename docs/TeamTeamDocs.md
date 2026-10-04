# TeamTeamDocs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Docs** | Pointer to [**[]TeamTeamDoc**](TeamTeamDoc.md) | Docs are the documents, grouped by teamspace, then parent, then the order the Team client shows siblings in. | [optional] 
**Teamspaces** | Pointer to [**[]TeamTeamTeamspace**](TeamTeamTeamspace.md) | Teamspaces are the teamspaces the caller may see, by name — the groups a client draws the tree under. | [optional] 

## Methods

### NewTeamTeamDocs

`func NewTeamTeamDocs() *TeamTeamDocs`

NewTeamTeamDocs instantiates a new TeamTeamDocs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamDocsWithDefaults

`func NewTeamTeamDocsWithDefaults() *TeamTeamDocs`

NewTeamTeamDocsWithDefaults instantiates a new TeamTeamDocs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocs

`func (o *TeamTeamDocs) GetDocs() []TeamTeamDoc`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *TeamTeamDocs) GetDocsOk() (*[]TeamTeamDoc, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *TeamTeamDocs) SetDocs(v []TeamTeamDoc)`

SetDocs sets Docs field to given value.

### HasDocs

`func (o *TeamTeamDocs) HasDocs() bool`

HasDocs returns a boolean if a field has been set.

### GetTeamspaces

`func (o *TeamTeamDocs) GetTeamspaces() []TeamTeamTeamspace`

GetTeamspaces returns the Teamspaces field if non-nil, zero value otherwise.

### GetTeamspacesOk

`func (o *TeamTeamDocs) GetTeamspacesOk() (*[]TeamTeamTeamspace, bool)`

GetTeamspacesOk returns a tuple with the Teamspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamspaces

`func (o *TeamTeamDocs) SetTeamspaces(v []TeamTeamTeamspace)`

SetTeamspaces sets Teamspaces field to given value.

### HasTeamspaces

`func (o *TeamTeamDocs) HasTeamspaces() bool`

HasTeamspaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


