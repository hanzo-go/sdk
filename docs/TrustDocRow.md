# TrustDocRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attested** | Pointer to **bool** | Attested reports whether somebody OUTSIDE this organization put their name to it. Those are the artifacts a reviewer asks for, and they are released through a grant rather than published — there is no field that can say otherwise. | [optional] 
**Href** | Pointer to **string** | Href is where to read it, present only when this reader may. | [optional] 
**Id** | Pointer to **string** | ID is the document&#39;s id within this organization&#39;s centre. | [optional] 
**Kind** | Pointer to **string** | Kind is the artifact type — soc2, iso, pentest, letter, caiq, sig, vsa, questionnaire, policy or other. | [optional] 
**Label** | Pointer to **string** | Label is the artifact type in words, for rendering. | [optional] 
**Note** | Pointer to **string** | Note is anything the organization says about this artifact. | [optional] 
**Released** | Pointer to **bool** | Released reports whether THIS reader may read it. False means the artifact exists and is available on request. | [optional] 
**Tier** | Pointer to **string** | Tier is \&quot;public\&quot; or \&quot;gated\&quot;. It defaults to gated, so a new artifact is closed until somebody opens it deliberately. | [optional] 
**Title** | Pointer to **string** | Title is what the document is called. | [optional] 
**Updated** | Pointer to **int64** | Updated is when the record last changed, unix milliseconds. | [optional] 

## Methods

### NewTrustDocRow

`func NewTrustDocRow() *TrustDocRow`

NewTrustDocRow instantiates a new TrustDocRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustDocRowWithDefaults

`func NewTrustDocRowWithDefaults() *TrustDocRow`

NewTrustDocRowWithDefaults instantiates a new TrustDocRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttested

`func (o *TrustDocRow) GetAttested() bool`

GetAttested returns the Attested field if non-nil, zero value otherwise.

### GetAttestedOk

`func (o *TrustDocRow) GetAttestedOk() (*bool, bool)`

GetAttestedOk returns a tuple with the Attested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttested

`func (o *TrustDocRow) SetAttested(v bool)`

SetAttested sets Attested field to given value.

### HasAttested

`func (o *TrustDocRow) HasAttested() bool`

HasAttested returns a boolean if a field has been set.

### GetHref

`func (o *TrustDocRow) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TrustDocRow) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TrustDocRow) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *TrustDocRow) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *TrustDocRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustDocRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustDocRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrustDocRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TrustDocRow) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrustDocRow) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrustDocRow) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrustDocRow) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *TrustDocRow) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TrustDocRow) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TrustDocRow) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TrustDocRow) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetNote

`func (o *TrustDocRow) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *TrustDocRow) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *TrustDocRow) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *TrustDocRow) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetReleased

`func (o *TrustDocRow) GetReleased() bool`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *TrustDocRow) GetReleasedOk() (*bool, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *TrustDocRow) SetReleased(v bool)`

SetReleased sets Released field to given value.

### HasReleased

`func (o *TrustDocRow) HasReleased() bool`

HasReleased returns a boolean if a field has been set.

### GetTier

`func (o *TrustDocRow) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *TrustDocRow) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *TrustDocRow) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *TrustDocRow) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetTitle

`func (o *TrustDocRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrustDocRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrustDocRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrustDocRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *TrustDocRow) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TrustDocRow) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TrustDocRow) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *TrustDocRow) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


