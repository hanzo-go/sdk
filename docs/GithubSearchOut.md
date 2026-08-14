# GithubSearchOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Repos** | Pointer to [**[]GithubSearchHit**](GithubSearchHit.md) |  | [optional] 

## Methods

### NewGithubSearchOut

`func NewGithubSearchOut() *GithubSearchOut`

NewGithubSearchOut instantiates a new GithubSearchOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubSearchOutWithDefaults

`func NewGithubSearchOutWithDefaults() *GithubSearchOut`

NewGithubSearchOutWithDefaults instantiates a new GithubSearchOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GithubSearchOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GithubSearchOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GithubSearchOut) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GithubSearchOut) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRepos

`func (o *GithubSearchOut) GetRepos() []GithubSearchHit`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *GithubSearchOut) GetReposOk() (*[]GithubSearchHit, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *GithubSearchOut) SetRepos(v []GithubSearchHit)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *GithubSearchOut) HasRepos() bool`

HasRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


