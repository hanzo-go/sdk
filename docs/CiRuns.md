# CiRuns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | Ref is what was asked for; HEAD when the request named none. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository, as owner/name. | [optional] 
**Runs** | Pointer to [**[]CiRun**](CiRun.md) | Runs is every run on that commit, newest first. Empty, never null, when nothing ran. | [optional] 
**Sha** | Pointer to **string** | SHA is the full commit the ref resolved to when it was read. A branch moves and this does not, so it is what the runs below belong to. | [optional] 

## Methods

### NewCiRuns

`func NewCiRuns() *CiRuns`

NewCiRuns instantiates a new CiRuns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiRunsWithDefaults

`func NewCiRunsWithDefaults() *CiRuns`

NewCiRunsWithDefaults instantiates a new CiRuns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *CiRuns) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CiRuns) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CiRuns) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CiRuns) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRepo

`func (o *CiRuns) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CiRuns) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CiRuns) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CiRuns) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRuns

`func (o *CiRuns) GetRuns() []CiRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *CiRuns) GetRunsOk() (*[]CiRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *CiRuns) SetRuns(v []CiRun)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *CiRuns) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetSha

`func (o *CiRuns) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CiRuns) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CiRuns) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *CiRuns) HasSha() bool`

HasSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


