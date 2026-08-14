# CommitJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorEmail** | Pointer to **string** | AuthorEmail is the commit author&#39;s email. | [optional] 
**AuthorName** | Pointer to **string** | AuthorName is the commit author&#39;s name. | [optional] 
**Date** | Pointer to **string** | Date is the author date, RFC 3339 UTC. | [optional] 
**Message** | Pointer to **string** | Message is the commit&#39;s SUBJECT — its first line only. | [optional] 
**Sha** | Pointer to **string** | SHA is the full commit hash. | [optional] 
**ShortSha** | Pointer to **string** | ShortSHA is the abbreviated hash a UI displays. | [optional] 

## Methods

### NewCommitJSON

`func NewCommitJSON() *CommitJSON`

NewCommitJSON instantiates a new CommitJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitJSONWithDefaults

`func NewCommitJSONWithDefaults() *CommitJSON`

NewCommitJSONWithDefaults instantiates a new CommitJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorEmail

`func (o *CommitJSON) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *CommitJSON) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *CommitJSON) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.

### HasAuthorEmail

`func (o *CommitJSON) HasAuthorEmail() bool`

HasAuthorEmail returns a boolean if a field has been set.

### GetAuthorName

`func (o *CommitJSON) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *CommitJSON) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *CommitJSON) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *CommitJSON) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetDate

`func (o *CommitJSON) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CommitJSON) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CommitJSON) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CommitJSON) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMessage

`func (o *CommitJSON) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommitJSON) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommitJSON) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CommitJSON) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSha

`func (o *CommitJSON) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CommitJSON) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CommitJSON) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *CommitJSON) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetShortSha

`func (o *CommitJSON) GetShortSha() string`

GetShortSha returns the ShortSha field if non-nil, zero value otherwise.

### GetShortShaOk

`func (o *CommitJSON) GetShortShaOk() (*string, bool)`

GetShortShaOk returns a tuple with the ShortSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortSha

`func (o *CommitJSON) SetShortSha(v string)`

SetShortSha sets ShortSha field to given value.

### HasShortSha

`func (o *CommitJSON) HasShortSha() bool`

HasShortSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


