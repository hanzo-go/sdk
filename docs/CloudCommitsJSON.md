# CloudCommitsJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commits** | Pointer to [**[]CloudCommitJSON**](CloudCommitJSON.md) | Commits are newest first. | [optional] 

## Methods

### NewCloudCommitsJSON

`func NewCloudCommitsJSON() *CloudCommitsJSON`

NewCloudCommitsJSON instantiates a new CloudCommitsJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCommitsJSONWithDefaults

`func NewCloudCommitsJSONWithDefaults() *CloudCommitsJSON`

NewCloudCommitsJSONWithDefaults instantiates a new CloudCommitsJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommits

`func (o *CloudCommitsJSON) GetCommits() []CloudCommitJSON`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *CloudCommitsJSON) GetCommitsOk() (*[]CloudCommitJSON, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *CloudCommitsJSON) SetCommits(v []CloudCommitJSON)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *CloudCommitsJSON) HasCommits() bool`

HasCommits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


