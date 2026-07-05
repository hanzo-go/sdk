# DidUpdateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Teams** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDidUpdateProfileRequest

`func NewDidUpdateProfileRequest() *DidUpdateProfileRequest`

NewDidUpdateProfileRequest instantiates a new DidUpdateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDidUpdateProfileRequestWithDefaults

`func NewDidUpdateProfileRequestWithDefaults() *DidUpdateProfileRequest`

NewDidUpdateProfileRequestWithDefaults instantiates a new DidUpdateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DidUpdateProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DidUpdateProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DidUpdateProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DidUpdateProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatar

`func (o *DidUpdateProfileRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *DidUpdateProfileRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *DidUpdateProfileRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *DidUpdateProfileRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetMetadata

`func (o *DidUpdateProfileRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DidUpdateProfileRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DidUpdateProfileRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DidUpdateProfileRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTeams

`func (o *DidUpdateProfileRequest) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *DidUpdateProfileRequest) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *DidUpdateProfileRequest) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *DidUpdateProfileRequest) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


