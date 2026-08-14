# ReferenceOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when it was written, RFC 3339. | [optional] 
**By** | Pointer to **string** | By is who wrote it. | [optional] 
**Key** | Pointer to **string** | Key is the member this organisation is speaking about. | [optional] 
**Note** | Pointer to **string** | Note is why, in the operator&#39;s own words. Optional, and bounded. | [optional] 
**Verdict** | Pointer to **string** | Verdict is allow or deny. | [optional] 

## Methods

### NewReferenceOverride

`func NewReferenceOverride() *ReferenceOverride`

NewReferenceOverride instantiates a new ReferenceOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceOverrideWithDefaults

`func NewReferenceOverrideWithDefaults() *ReferenceOverride`

NewReferenceOverrideWithDefaults instantiates a new ReferenceOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *ReferenceOverride) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *ReferenceOverride) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *ReferenceOverride) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *ReferenceOverride) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *ReferenceOverride) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *ReferenceOverride) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *ReferenceOverride) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *ReferenceOverride) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetKey

`func (o *ReferenceOverride) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ReferenceOverride) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ReferenceOverride) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ReferenceOverride) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNote

`func (o *ReferenceOverride) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ReferenceOverride) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ReferenceOverride) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ReferenceOverride) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetVerdict

`func (o *ReferenceOverride) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ReferenceOverride) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ReferenceOverride) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *ReferenceOverride) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


