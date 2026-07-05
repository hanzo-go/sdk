# DidDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to [**[]DidTeam**](DidTeam.md) |  | [optional] 
**MemberCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewDidDirectory

`func NewDidDirectory() *DidDirectory`

NewDidDirectory instantiates a new DidDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDidDirectoryWithDefaults

`func NewDidDirectoryWithDefaults() *DidDirectory`

NewDidDirectoryWithDefaults instantiates a new DidDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *DidDirectory) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DidDirectory) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DidDirectory) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DidDirectory) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetTeams

`func (o *DidDirectory) GetTeams() []DidTeam`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *DidDirectory) GetTeamsOk() (*[]DidTeam, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *DidDirectory) SetTeams(v []DidTeam)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *DidDirectory) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetMemberCount

`func (o *DidDirectory) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *DidDirectory) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *DidDirectory) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *DidDirectory) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


