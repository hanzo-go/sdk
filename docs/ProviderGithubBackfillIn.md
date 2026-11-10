# ProviderGithubBackfillIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repo** | Pointer to **string** | Repo narrows the pass to ONE granted repository, named &#x60;owner/name&#x60; or bare &#x60;name&#x60; (case-insensitive). Empty walks every granted repository. A project page syncs its own issues with this; the org-wide pass stays the default. A name the connection does not grant is a 404, never a silent empty pass that reads as \&quot;this repository has no issues\&quot;. | [optional] 
**State** | Pointer to **string** | State is the GitHub issue state to walk: \&quot;open\&quot; (the default), \&quot;closed\&quot; or \&quot;all\&quot;. Anything else is a 400. | [optional] 

## Methods

### NewProviderGithubBackfillIn

`func NewProviderGithubBackfillIn() *ProviderGithubBackfillIn`

NewProviderGithubBackfillIn instantiates a new ProviderGithubBackfillIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubBackfillInWithDefaults

`func NewProviderGithubBackfillInWithDefaults() *ProviderGithubBackfillIn`

NewProviderGithubBackfillInWithDefaults instantiates a new ProviderGithubBackfillIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepo

`func (o *ProviderGithubBackfillIn) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProviderGithubBackfillIn) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProviderGithubBackfillIn) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProviderGithubBackfillIn) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetState

`func (o *ProviderGithubBackfillIn) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProviderGithubBackfillIn) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProviderGithubBackfillIn) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProviderGithubBackfillIn) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


