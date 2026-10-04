# ReferenceReferenceOverrideIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the member: a domain, a CIDR or address, an issuer prefix, a device digest. It is matched the same way the baseline is, so a deny on tempbox.example also covers mail.tempbox.example. | [optional] 
**Note** | Pointer to **string** | Note is why, in your own words. Optional, bounded to 512 bytes. | [optional] 
**Verdict** | Pointer to **string** | Verdict is allow or deny, and nothing else. An override is a decision — unlike a baseline entry, which states facts and leaves the decision to your policy — because your organisation is the only party entitled to say \&quot;for us, this one is fine\&quot;. | [optional] 

## Methods

### NewReferenceReferenceOverrideIn

`func NewReferenceReferenceOverrideIn() *ReferenceReferenceOverrideIn`

NewReferenceReferenceOverrideIn instantiates a new ReferenceReferenceOverrideIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceReferenceOverrideInWithDefaults

`func NewReferenceReferenceOverrideInWithDefaults() *ReferenceReferenceOverrideIn`

NewReferenceReferenceOverrideInWithDefaults instantiates a new ReferenceReferenceOverrideIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ReferenceReferenceOverrideIn) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ReferenceReferenceOverrideIn) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ReferenceReferenceOverrideIn) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ReferenceReferenceOverrideIn) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNote

`func (o *ReferenceReferenceOverrideIn) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ReferenceReferenceOverrideIn) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ReferenceReferenceOverrideIn) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ReferenceReferenceOverrideIn) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetVerdict

`func (o *ReferenceReferenceOverrideIn) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ReferenceReferenceOverrideIn) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ReferenceReferenceOverrideIn) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *ReferenceReferenceOverrideIn) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


