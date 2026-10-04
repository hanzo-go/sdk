# TeamTeamDocNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | Pointer to **string** | Parent nests the new document under another one. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid. Optional for a caller in exactly one space. | [optional] 
**Teamspace** | Pointer to **string** | Teamspace is where the document goes. Optional: a document nested under a parent goes in the parent&#39;s teamspace, and a space with exactly one teamspace the caller can write needs no name — with none, the first document opens a public \&quot;General\&quot; teamspace for the whole space. | [optional] 
**Title** | Pointer to **string** | Title is the document&#39;s title. Empty reads \&quot;Untitled\&quot;, as in the Team client. | [optional] 

## Methods

### NewTeamTeamDocNew

`func NewTeamTeamDocNew() *TeamTeamDocNew`

NewTeamTeamDocNew instantiates a new TeamTeamDocNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamDocNewWithDefaults

`func NewTeamTeamDocNewWithDefaults() *TeamTeamDocNew`

NewTeamTeamDocNewWithDefaults instantiates a new TeamTeamDocNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *TeamTeamDocNew) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TeamTeamDocNew) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TeamTeamDocNew) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TeamTeamDocNew) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamDocNew) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamDocNew) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamDocNew) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamDocNew) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTeamspace

`func (o *TeamTeamDocNew) GetTeamspace() string`

GetTeamspace returns the Teamspace field if non-nil, zero value otherwise.

### GetTeamspaceOk

`func (o *TeamTeamDocNew) GetTeamspaceOk() (*string, bool)`

GetTeamspaceOk returns a tuple with the Teamspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamspace

`func (o *TeamTeamDocNew) SetTeamspace(v string)`

SetTeamspace sets Teamspace field to given value.

### HasTeamspace

`func (o *TeamTeamDocNew) HasTeamspace() bool`

HasTeamspace returns a boolean if a field has been set.

### GetTitle

`func (o *TeamTeamDocNew) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TeamTeamDocNew) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TeamTeamDocNew) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TeamTeamDocNew) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


