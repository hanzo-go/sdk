# ReferenceTaken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **int64** | Keys is how many members it carries. | [optional] 
**Refusal** | Pointer to **string** | Refusal is why this publisher contributed nothing, if it did not. The set keeps its previous version of this source rather than shrinking. | [optional] 
**Resumed** | Pointer to **bool** | Resumed is true when this run continued a version a previous run left half-landed. | [optional] 
**Source** | Pointer to **string** | Source is the publisher. | [optional] 
**Unchanged** | Pointer to **bool** | Unchanged is true when the publisher&#39;s data was byte-for-byte the set we already held. | [optional] 
**Version** | Pointer to **string** | Version is the content digest that landed. | [optional] 
**Wrote** | Pointer to **int64** | Wrote is how many rows this run actually wrote. Zero with Unchanged means the publisher served the same set again. | [optional] 

## Methods

### NewReferenceTaken

`func NewReferenceTaken() *ReferenceTaken`

NewReferenceTaken instantiates a new ReferenceTaken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceTakenWithDefaults

`func NewReferenceTakenWithDefaults() *ReferenceTaken`

NewReferenceTakenWithDefaults instantiates a new ReferenceTaken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ReferenceTaken) GetKeys() int64`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ReferenceTaken) GetKeysOk() (*int64, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ReferenceTaken) SetKeys(v int64)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ReferenceTaken) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetRefusal

`func (o *ReferenceTaken) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *ReferenceTaken) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *ReferenceTaken) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *ReferenceTaken) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetResumed

`func (o *ReferenceTaken) GetResumed() bool`

GetResumed returns the Resumed field if non-nil, zero value otherwise.

### GetResumedOk

`func (o *ReferenceTaken) GetResumedOk() (*bool, bool)`

GetResumedOk returns a tuple with the Resumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumed

`func (o *ReferenceTaken) SetResumed(v bool)`

SetResumed sets Resumed field to given value.

### HasResumed

`func (o *ReferenceTaken) HasResumed() bool`

HasResumed returns a boolean if a field has been set.

### GetSource

`func (o *ReferenceTaken) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReferenceTaken) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReferenceTaken) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReferenceTaken) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUnchanged

`func (o *ReferenceTaken) GetUnchanged() bool`

GetUnchanged returns the Unchanged field if non-nil, zero value otherwise.

### GetUnchangedOk

`func (o *ReferenceTaken) GetUnchangedOk() (*bool, bool)`

GetUnchangedOk returns a tuple with the Unchanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnchanged

`func (o *ReferenceTaken) SetUnchanged(v bool)`

SetUnchanged sets Unchanged field to given value.

### HasUnchanged

`func (o *ReferenceTaken) HasUnchanged() bool`

HasUnchanged returns a boolean if a field has been set.

### GetVersion

`func (o *ReferenceTaken) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReferenceTaken) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReferenceTaken) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReferenceTaken) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWrote

`func (o *ReferenceTaken) GetWrote() int64`

GetWrote returns the Wrote field if non-nil, zero value otherwise.

### GetWroteOk

`func (o *ReferenceTaken) GetWroteOk() (*int64, bool)`

GetWroteOk returns a tuple with the Wrote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrote

`func (o *ReferenceTaken) SetWrote(v int64)`

SetWrote sets Wrote field to given value.

### HasWrote

`func (o *ReferenceTaken) HasWrote() bool`

HasWrote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


