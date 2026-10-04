# GitKeyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GitKeyView**](GitKeyView.md) | Data holds the org&#39;s keys. | [optional] 

## Methods

### NewGitKeyList

`func NewGitKeyList() *GitKeyList`

NewGitKeyList instantiates a new GitKeyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitKeyListWithDefaults

`func NewGitKeyListWithDefaults() *GitKeyList`

NewGitKeyListWithDefaults instantiates a new GitKeyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GitKeyList) GetData() []GitKeyView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GitKeyList) GetDataOk() (*[]GitKeyView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GitKeyList) SetData(v []GitKeyView)`

SetData sets Data field to given value.

### HasData

`func (o *GitKeyList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


