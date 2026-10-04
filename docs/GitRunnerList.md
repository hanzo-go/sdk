# GitRunnerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GitRunnerView**](GitRunnerView.md) | Data is the daemons registered into this org&#39;s pools, newest first. | [optional] 

## Methods

### NewGitRunnerList

`func NewGitRunnerList() *GitRunnerList`

NewGitRunnerList instantiates a new GitRunnerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRunnerListWithDefaults

`func NewGitRunnerListWithDefaults() *GitRunnerList`

NewGitRunnerListWithDefaults instantiates a new GitRunnerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GitRunnerList) GetData() []GitRunnerView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GitRunnerList) GetDataOk() (*[]GitRunnerView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GitRunnerList) SetData(v []GitRunnerView)`

SetData sets Data field to given value.

### HasData

`func (o *GitRunnerList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


