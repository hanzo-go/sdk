# DocRow

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

### NewDocRow

`func NewDocRow() *DocRow`

NewDocRow instantiates a new DocRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocRowWithDefaults

`func NewDocRowWithDefaults() *DocRow`

NewDocRowWithDefaults instantiates a new DocRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttested

`func (o *DocRow) GetAttested() bool`

GetAttested returns the Attested field if non-nil, zero value otherwise.

### GetAttestedOk

`func (o *DocRow) GetAttestedOk() (*bool, bool)`

GetAttestedOk returns a tuple with the Attested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttested

`func (o *DocRow) SetAttested(v bool)`

SetAttested sets Attested field to given value.

### HasAttested

`func (o *DocRow) HasAttested() bool`

HasAttested returns a boolean if a field has been set.

### GetHref

`func (o *DocRow) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DocRow) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DocRow) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DocRow) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *DocRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *DocRow) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DocRow) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DocRow) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DocRow) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *DocRow) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DocRow) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DocRow) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DocRow) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetNote

`func (o *DocRow) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *DocRow) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *DocRow) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *DocRow) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetReleased

`func (o *DocRow) GetReleased() bool`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *DocRow) GetReleasedOk() (*bool, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *DocRow) SetReleased(v bool)`

SetReleased sets Released field to given value.

### HasReleased

`func (o *DocRow) HasReleased() bool`

HasReleased returns a boolean if a field has been set.

### GetTier

`func (o *DocRow) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *DocRow) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *DocRow) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *DocRow) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetTitle

`func (o *DocRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *DocRow) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DocRow) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DocRow) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *DocRow) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


