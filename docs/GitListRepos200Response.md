# GitListRepos200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GitRepo**](GitRepo.md) |  | [optional] 

## Methods

### NewGitListRepos200Response

`func NewGitListRepos200Response() *GitListRepos200Response`

NewGitListRepos200Response instantiates a new GitListRepos200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitListRepos200ResponseWithDefaults

`func NewGitListRepos200ResponseWithDefaults() *GitListRepos200Response`

NewGitListRepos200ResponseWithDefaults instantiates a new GitListRepos200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GitListRepos200Response) GetData() []GitRepo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GitListRepos200Response) GetDataOk() (*[]GitRepo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GitListRepos200Response) SetData(v []GitRepo)`

SetData sets Data field to given value.

### HasData

`func (o *GitListRepos200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


