# GitPullList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GitPullView**](GitPullView.md) | Data holds the repo&#39;s pull requests, newest number first. | [optional] 

## Methods

### NewGitPullList

`func NewGitPullList() *GitPullList`

NewGitPullList instantiates a new GitPullList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitPullListWithDefaults

`func NewGitPullListWithDefaults() *GitPullList`

NewGitPullListWithDefaults instantiates a new GitPullList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GitPullList) GetData() []GitPullView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GitPullList) GetDataOk() (*[]GitPullView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GitPullList) SetData(v []GitPullView)`

SetData sets Data field to given value.

### HasData

`func (o *GitPullList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


