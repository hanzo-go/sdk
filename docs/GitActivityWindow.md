# GitActivityWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** | By is the commit date counted: \&quot;committer\&quot;, when the work landed. | [optional] 
**From** | Pointer to **string** | From is the first UTC day counted, YYYY-MM-DD. | [optional] 
**To** | Pointer to **string** | To is the last UTC day counted, YYYY-MM-DD. | [optional] 

## Methods

### NewGitActivityWindow

`func NewGitActivityWindow() *GitActivityWindow`

NewGitActivityWindow instantiates a new GitActivityWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitActivityWindowWithDefaults

`func NewGitActivityWindowWithDefaults() *GitActivityWindow`

NewGitActivityWindowWithDefaults instantiates a new GitActivityWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *GitActivityWindow) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *GitActivityWindow) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *GitActivityWindow) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *GitActivityWindow) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetFrom

`func (o *GitActivityWindow) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GitActivityWindow) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GitActivityWindow) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GitActivityWindow) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GitActivityWindow) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GitActivityWindow) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GitActivityWindow) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GitActivityWindow) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


