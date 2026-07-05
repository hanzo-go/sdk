# DidCreateTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ParentTeam** | Pointer to **string** |  | [optional] 

## Methods

### NewDidCreateTeamRequest

`func NewDidCreateTeamRequest(name string, ) *DidCreateTeamRequest`

NewDidCreateTeamRequest instantiates a new DidCreateTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDidCreateTeamRequestWithDefaults

`func NewDidCreateTeamRequestWithDefaults() *DidCreateTeamRequest`

NewDidCreateTeamRequestWithDefaults instantiates a new DidCreateTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DidCreateTeamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DidCreateTeamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DidCreateTeamRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DidCreateTeamRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DidCreateTeamRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DidCreateTeamRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DidCreateTeamRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentTeam

`func (o *DidCreateTeamRequest) GetParentTeam() string`

GetParentTeam returns the ParentTeam field if non-nil, zero value otherwise.

### GetParentTeamOk

`func (o *DidCreateTeamRequest) GetParentTeamOk() (*string, bool)`

GetParentTeamOk returns a tuple with the ParentTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTeam

`func (o *DidCreateTeamRequest) SetParentTeam(v string)`

SetParentTeam sets ParentTeam field to given value.

### HasParentTeam

`func (o *DidCreateTeamRequest) HasParentTeam() bool`

HasParentTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


