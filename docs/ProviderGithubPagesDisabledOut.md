# ProviderGithubPagesDisabledOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Disabled is always true — a failure is an HTTP error, never this shape. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository whose site was deleted. | [optional] 

## Methods

### NewProviderGithubPagesDisabledOut

`func NewProviderGithubPagesDisabledOut() *ProviderGithubPagesDisabledOut`

NewProviderGithubPagesDisabledOut instantiates a new ProviderGithubPagesDisabledOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubPagesDisabledOutWithDefaults

`func NewProviderGithubPagesDisabledOutWithDefaults() *ProviderGithubPagesDisabledOut`

NewProviderGithubPagesDisabledOutWithDefaults instantiates a new ProviderGithubPagesDisabledOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *ProviderGithubPagesDisabledOut) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ProviderGithubPagesDisabledOut) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ProviderGithubPagesDisabledOut) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ProviderGithubPagesDisabledOut) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetRepo

`func (o *ProviderGithubPagesDisabledOut) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProviderGithubPagesDisabledOut) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProviderGithubPagesDisabledOut) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProviderGithubPagesDisabledOut) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


