# GitLinked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Held** | Pointer to **[]string** | Held are the repositories that hold their own history here, which a link would hide. | [optional] 
**Linked** | Pointer to [**[]GitRepoView**](GitRepoView.md) | Linked are the repositories it linked, each to its namesake on GitHub. | [optional] 
**Unmatched** | Pointer to [**[]GitUnmatched**](GitUnmatched.md) | Unmatched are the repositories it left as they were, each with what GitHub answered: no namesake readable as the caller is never replaced by a guess. | [optional] 

## Methods

### NewGitLinked

`func NewGitLinked() *GitLinked`

NewGitLinked instantiates a new GitLinked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitLinkedWithDefaults

`func NewGitLinkedWithDefaults() *GitLinked`

NewGitLinkedWithDefaults instantiates a new GitLinked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeld

`func (o *GitLinked) GetHeld() []string`

GetHeld returns the Held field if non-nil, zero value otherwise.

### GetHeldOk

`func (o *GitLinked) GetHeldOk() (*[]string, bool)`

GetHeldOk returns a tuple with the Held field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeld

`func (o *GitLinked) SetHeld(v []string)`

SetHeld sets Held field to given value.

### HasHeld

`func (o *GitLinked) HasHeld() bool`

HasHeld returns a boolean if a field has been set.

### GetLinked

`func (o *GitLinked) GetLinked() []GitRepoView`

GetLinked returns the Linked field if non-nil, zero value otherwise.

### GetLinkedOk

`func (o *GitLinked) GetLinkedOk() (*[]GitRepoView, bool)`

GetLinkedOk returns a tuple with the Linked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinked

`func (o *GitLinked) SetLinked(v []GitRepoView)`

SetLinked sets Linked field to given value.

### HasLinked

`func (o *GitLinked) HasLinked() bool`

HasLinked returns a boolean if a field has been set.

### GetUnmatched

`func (o *GitLinked) GetUnmatched() []GitUnmatched`

GetUnmatched returns the Unmatched field if non-nil, zero value otherwise.

### GetUnmatchedOk

`func (o *GitLinked) GetUnmatchedOk() (*[]GitUnmatched, bool)`

GetUnmatchedOk returns a tuple with the Unmatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatched

`func (o *GitLinked) SetUnmatched(v []GitUnmatched)`

SetUnmatched sets Unmatched field to given value.

### HasUnmatched

`func (o *GitLinked) HasUnmatched() bool`

HasUnmatched returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


