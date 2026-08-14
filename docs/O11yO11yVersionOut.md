# O11yO11yVersionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ee** | Pointer to **string** | EE says whether an enterprise edition is present; \&quot;N\&quot; in this build. | [optional] 
**SetupCompleted** | Pointer to **bool** | SetupCompleted says whether the first user has been created. | [optional] 
**Version** | Pointer to **string** | Version is the build version. | [optional] 

## Methods

### NewO11yO11yVersionOut

`func NewO11yO11yVersionOut() *O11yO11yVersionOut`

NewO11yO11yVersionOut instantiates a new O11yO11yVersionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yVersionOutWithDefaults

`func NewO11yO11yVersionOutWithDefaults() *O11yO11yVersionOut`

NewO11yO11yVersionOutWithDefaults instantiates a new O11yO11yVersionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEe

`func (o *O11yO11yVersionOut) GetEe() string`

GetEe returns the Ee field if non-nil, zero value otherwise.

### GetEeOk

`func (o *O11yO11yVersionOut) GetEeOk() (*string, bool)`

GetEeOk returns a tuple with the Ee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEe

`func (o *O11yO11yVersionOut) SetEe(v string)`

SetEe sets Ee field to given value.

### HasEe

`func (o *O11yO11yVersionOut) HasEe() bool`

HasEe returns a boolean if a field has been set.

### GetSetupCompleted

`func (o *O11yO11yVersionOut) GetSetupCompleted() bool`

GetSetupCompleted returns the SetupCompleted field if non-nil, zero value otherwise.

### GetSetupCompletedOk

`func (o *O11yO11yVersionOut) GetSetupCompletedOk() (*bool, bool)`

GetSetupCompletedOk returns a tuple with the SetupCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupCompleted

`func (o *O11yO11yVersionOut) SetSetupCompleted(v bool)`

SetSetupCompleted sets SetupCompleted field to given value.

### HasSetupCompleted

`func (o *O11yO11yVersionOut) HasSetupCompleted() bool`

HasSetupCompleted returns a boolean if a field has been set.

### GetVersion

`func (o *O11yO11yVersionOut) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *O11yO11yVersionOut) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *O11yO11yVersionOut) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *O11yO11yVersionOut) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


