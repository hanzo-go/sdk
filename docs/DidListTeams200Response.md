# DidListTeams200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Teams** | Pointer to [**[]DidTeam**](DidTeam.md) |  | [optional] 

## Methods

### NewDidListTeams200Response

`func NewDidListTeams200Response() *DidListTeams200Response`

NewDidListTeams200Response instantiates a new DidListTeams200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDidListTeams200ResponseWithDefaults

`func NewDidListTeams200ResponseWithDefaults() *DidListTeams200Response`

NewDidListTeams200ResponseWithDefaults instantiates a new DidListTeams200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeams

`func (o *DidListTeams200Response) GetTeams() []DidTeam`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *DidListTeams200Response) GetTeamsOk() (*[]DidTeam, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *DidListTeams200Response) SetTeams(v []DidTeam)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *DidListTeams200Response) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


