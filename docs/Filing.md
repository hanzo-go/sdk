# Filing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int32** | At is the unix second the filing record was written. | [optional] 
**Note** | Pointer to **string** | Note explains a filing Hanzo did not perform itself: what remains to be done and by whom. | [optional] 
**Provider** | Pointer to **string** | Provider is the filing partner that performed the filing, or \&quot;manual\&quot; when no partner is wired. | [optional] 
**Ref** | Pointer to **string** | Ref is the partner&#39;s or the state&#39;s filing reference. Empty when nothing was actually filed — no filing id is ever fabricated. | [optional] 
**Status** | Pointer to **string** | Status is manual (no partner wired — a registered agent files out-of-band), submitted (the partner accepted it, awaiting the state), filed (the state accepted it) or rejected. | [optional] 

## Methods

### NewFiling

`func NewFiling() *Filing`

NewFiling instantiates a new Filing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilingWithDefaults

`func NewFilingWithDefaults() *Filing`

NewFilingWithDefaults instantiates a new Filing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *Filing) GetAt() int32`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *Filing) GetAtOk() (*int32, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *Filing) SetAt(v int32)`

SetAt sets At field to given value.

### HasAt

`func (o *Filing) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetNote

`func (o *Filing) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Filing) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Filing) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Filing) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetProvider

`func (o *Filing) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Filing) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Filing) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Filing) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRef

`func (o *Filing) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Filing) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Filing) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Filing) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetStatus

`func (o *Filing) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Filing) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Filing) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Filing) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


