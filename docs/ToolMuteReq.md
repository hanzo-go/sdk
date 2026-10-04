# ToolMuteReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mute** | Pointer to **[]string** | Mute takes these out of the caller&#39;s runs: a skill (skill_&lt;name&gt;), a server&#39;s tool (&lt;server&gt;_&lt;tool&gt;) or a whole server (&lt;server&gt;), each one the org&#39;s kit carries. | [optional] 
**Unmute** | Pointer to **[]string** | Unmute gives these back, as far as the org&#39;s kit carries them. | [optional] 

## Methods

### NewToolMuteReq

`func NewToolMuteReq() *ToolMuteReq`

NewToolMuteReq instantiates a new ToolMuteReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolMuteReqWithDefaults

`func NewToolMuteReqWithDefaults() *ToolMuteReq`

NewToolMuteReqWithDefaults instantiates a new ToolMuteReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMute

`func (o *ToolMuteReq) GetMute() []string`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *ToolMuteReq) GetMuteOk() (*[]string, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *ToolMuteReq) SetMute(v []string)`

SetMute sets Mute field to given value.

### HasMute

`func (o *ToolMuteReq) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetUnmute

`func (o *ToolMuteReq) GetUnmute() []string`

GetUnmute returns the Unmute field if non-nil, zero value otherwise.

### GetUnmuteOk

`func (o *ToolMuteReq) GetUnmuteOk() (*[]string, bool)`

GetUnmuteOk returns a tuple with the Unmute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmute

`func (o *ToolMuteReq) SetUnmute(v []string)`

SetUnmute sets Unmute field to given value.

### HasUnmute

`func (o *ToolMuteReq) HasUnmute() bool`

HasUnmute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


