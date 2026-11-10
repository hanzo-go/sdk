# FlagOptinOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to **bool** | Blocked is true when the org the caller acts in turned the opt-in off for its members. Their own choice is kept and applies again once unblocked. | [optional] 
**On** | Pointer to **bool** | On is true when the caller opted in: every public beta, and each one promoted later, is on for them unless Blocked. | [optional] 
**Public** | Pointer to **[]string** | Public is the id of every beta anyone may opt in to (beta.&lt;id&gt;). The rest are private: a SuperAdmin grants them, person by person. | [optional] 

## Methods

### NewFlagOptinOut

`func NewFlagOptinOut() *FlagOptinOut`

NewFlagOptinOut instantiates a new FlagOptinOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagOptinOutWithDefaults

`func NewFlagOptinOutWithDefaults() *FlagOptinOut`

NewFlagOptinOutWithDefaults instantiates a new FlagOptinOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *FlagOptinOut) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *FlagOptinOut) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *FlagOptinOut) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *FlagOptinOut) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetOn

`func (o *FlagOptinOut) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *FlagOptinOut) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *FlagOptinOut) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *FlagOptinOut) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetPublic

`func (o *FlagOptinOut) GetPublic() []string`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *FlagOptinOut) GetPublicOk() (*[]string, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *FlagOptinOut) SetPublic(v []string)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *FlagOptinOut) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


