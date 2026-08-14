# VerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoUrl** | Pointer to **string** | RepoURL is what to claim: a repository (github.com/owner/name) or a whole OWNER (github.com/owner, no repository segment). gitlab.com is accepted too. | [optional] 

## Methods

### NewVerifyRequest

`func NewVerifyRequest() *VerifyRequest`

NewVerifyRequest instantiates a new VerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyRequestWithDefaults

`func NewVerifyRequestWithDefaults() *VerifyRequest`

NewVerifyRequestWithDefaults instantiates a new VerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoUrl

`func (o *VerifyRequest) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *VerifyRequest) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *VerifyRequest) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *VerifyRequest) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


