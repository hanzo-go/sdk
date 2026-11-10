# FormImported

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**[]FormAnswer**](FormAnswer.md) | Answers are its answers, each read against the form&#39;s field as a submission&#39;s is; a required field may be unanswered. | [optional] 
**At** | Pointer to **string** | At is when it was made, RFC 3339. Required. | [optional] 
**Email** | Pointer to **string** | Email is the address it was from, when the source knew one. On a subscribe form the address is the response&#39;s key: one address is one subscriber, however many times and places it subscribed. | [optional] 
**External** | Pointer to **string** | External is its id where it was made: at most 200 characters, and the response&#39;s key, so importing it again is the same response. | [optional] 
**Hidden** | Pointer to **map[string]string** | Hidden are values it carried, by the names the form declares. | [optional] 
**Source** | Pointer to **string** | Source is where it was made, at most 100 characters: commerce, a list&#39;s name. | [optional] 

## Methods

### NewFormImported

`func NewFormImported() *FormImported`

NewFormImported instantiates a new FormImported object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormImportedWithDefaults

`func NewFormImportedWithDefaults() *FormImported`

NewFormImportedWithDefaults instantiates a new FormImported object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *FormImported) GetAnswers() []FormAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *FormImported) GetAnswersOk() (*[]FormAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *FormImported) SetAnswers(v []FormAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *FormImported) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetAt

`func (o *FormImported) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *FormImported) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *FormImported) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *FormImported) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetEmail

`func (o *FormImported) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FormImported) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FormImported) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FormImported) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternal

`func (o *FormImported) GetExternal() string`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *FormImported) GetExternalOk() (*string, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *FormImported) SetExternal(v string)`

SetExternal sets External field to given value.

### HasExternal

`func (o *FormImported) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetHidden

`func (o *FormImported) GetHidden() map[string]string`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FormImported) GetHiddenOk() (*map[string]string, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FormImported) SetHidden(v map[string]string)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FormImported) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetSource

`func (o *FormImported) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FormImported) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FormImported) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FormImported) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


