# GitRefJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the short ref name (\&quot;main\&quot;, \&quot;v1.2.0\&quot;), not the full refs/… path. | [optional] 
**Sha** | Pointer to **string** | SHA is the full commit hash the ref resolves to. | [optional] 

## Methods

### NewGitRefJSON

`func NewGitRefJSON() *GitRefJSON`

NewGitRefJSON instantiates a new GitRefJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRefJSONWithDefaults

`func NewGitRefJSONWithDefaults() *GitRefJSON`

NewGitRefJSONWithDefaults instantiates a new GitRefJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GitRefJSON) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitRefJSON) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitRefJSON) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitRefJSON) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSha

`func (o *GitRefJSON) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitRefJSON) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitRefJSON) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitRefJSON) HasSha() bool`

HasSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


