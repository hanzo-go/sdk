# EnvironmentEnvSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the environment variable the run exports the value under. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository&#39;s name in the caller&#39;s org. | [optional] 
**Value** | Pointer to **string** | Value is sealed in KMS and never answered, logged or echoed. It is never read from a query string, because a URL is logged in more places than a body is. | [optional] 

## Methods

### NewEnvironmentEnvSecret

`func NewEnvironmentEnvSecret() *EnvironmentEnvSecret`

NewEnvironmentEnvSecret instantiates a new EnvironmentEnvSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEnvSecretWithDefaults

`func NewEnvironmentEnvSecretWithDefaults() *EnvironmentEnvSecret`

NewEnvironmentEnvSecretWithDefaults instantiates a new EnvironmentEnvSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentEnvSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentEnvSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentEnvSecret) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentEnvSecret) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepo

`func (o *EnvironmentEnvSecret) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *EnvironmentEnvSecret) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *EnvironmentEnvSecret) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *EnvironmentEnvSecret) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetValue

`func (o *EnvironmentEnvSecret) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironmentEnvSecret) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironmentEnvSecret) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvironmentEnvSecret) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


