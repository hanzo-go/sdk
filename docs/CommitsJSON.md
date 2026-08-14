# CommitsJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commits** | Pointer to [**[]CommitJSON**](CommitJSON.md) | Commits are newest first. | [optional] 

## Methods

### NewCommitsJSON

`func NewCommitsJSON() *CommitsJSON`

NewCommitsJSON instantiates a new CommitsJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitsJSONWithDefaults

`func NewCommitsJSONWithDefaults() *CommitsJSON`

NewCommitsJSONWithDefaults instantiates a new CommitsJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommits

`func (o *CommitsJSON) GetCommits() []CommitJSON`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *CommitsJSON) GetCommitsOk() (*[]CommitJSON, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *CommitsJSON) SetCommits(v []CommitJSON)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *CommitsJSON) HasCommits() bool`

HasCommits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


