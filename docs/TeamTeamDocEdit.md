# TeamTeamDocEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the document, from the path. | [optional] 
**Parent** | Pointer to **string** | Parent moves the document under another one; \&quot;\&quot; moves it to the top of its teamspace. It goes after its new siblings. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid holding it. Body-only. | [optional] 
**Teamspace** | Pointer to **string** | Teamspace moves the document — with everything nested under it — to another teamspace, at the top unless Parent names a document there. | [optional] 
**Title** | Pointer to **string** | Title renames the document. Empty reads \&quot;Untitled\&quot;. | [optional] 

## Methods

### NewTeamTeamDocEdit

`func NewTeamTeamDocEdit() *TeamTeamDocEdit`

NewTeamTeamDocEdit instantiates a new TeamTeamDocEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamDocEditWithDefaults

`func NewTeamTeamDocEditWithDefaults() *TeamTeamDocEdit`

NewTeamTeamDocEditWithDefaults instantiates a new TeamTeamDocEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamTeamDocEdit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamDocEdit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamDocEdit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamDocEdit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParent

`func (o *TeamTeamDocEdit) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TeamTeamDocEdit) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TeamTeamDocEdit) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TeamTeamDocEdit) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamDocEdit) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamDocEdit) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamDocEdit) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamDocEdit) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTeamspace

`func (o *TeamTeamDocEdit) GetTeamspace() string`

GetTeamspace returns the Teamspace field if non-nil, zero value otherwise.

### GetTeamspaceOk

`func (o *TeamTeamDocEdit) GetTeamspaceOk() (*string, bool)`

GetTeamspaceOk returns a tuple with the Teamspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamspace

`func (o *TeamTeamDocEdit) SetTeamspace(v string)`

SetTeamspace sets Teamspace field to given value.

### HasTeamspace

`func (o *TeamTeamDocEdit) HasTeamspace() bool`

HasTeamspace returns a boolean if a field has been set.

### GetTitle

`func (o *TeamTeamDocEdit) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TeamTeamDocEdit) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TeamTeamDocEdit) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TeamTeamDocEdit) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


