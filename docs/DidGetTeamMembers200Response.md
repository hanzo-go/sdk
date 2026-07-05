# DidGetTeamMembers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | Pointer to [**DidTeam**](DidTeam.md) |  | [optional] 
**Members** | Pointer to [**[]DidProfile**](DidProfile.md) |  | [optional] 

## Methods

### NewDidGetTeamMembers200Response

`func NewDidGetTeamMembers200Response() *DidGetTeamMembers200Response`

NewDidGetTeamMembers200Response instantiates a new DidGetTeamMembers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDidGetTeamMembers200ResponseWithDefaults

`func NewDidGetTeamMembers200ResponseWithDefaults() *DidGetTeamMembers200Response`

NewDidGetTeamMembers200ResponseWithDefaults instantiates a new DidGetTeamMembers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeam

`func (o *DidGetTeamMembers200Response) GetTeam() DidTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *DidGetTeamMembers200Response) GetTeamOk() (*DidTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *DidGetTeamMembers200Response) SetTeam(v DidTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *DidGetTeamMembers200Response) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetMembers

`func (o *DidGetTeamMembers200Response) GetMembers() []DidProfile`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DidGetTeamMembers200Response) GetMembersOk() (*[]DidProfile, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DidGetTeamMembers200Response) SetMembers(v []DidProfile)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *DidGetTeamMembers200Response) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


