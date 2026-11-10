# FlagOrgOptinIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to **bool** | Blocked turns the opt-in off for every member of the org when true, and back on when false. It never turns a beta on for anyone. | [optional] 

## Methods

### NewFlagOrgOptinIn

`func NewFlagOrgOptinIn() *FlagOrgOptinIn`

NewFlagOrgOptinIn instantiates a new FlagOrgOptinIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagOrgOptinInWithDefaults

`func NewFlagOrgOptinInWithDefaults() *FlagOrgOptinIn`

NewFlagOrgOptinInWithDefaults instantiates a new FlagOrgOptinIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *FlagOrgOptinIn) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *FlagOrgOptinIn) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *FlagOrgOptinIn) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *FlagOrgOptinIn) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


