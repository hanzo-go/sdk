# Obligation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** | AmountCents is what it costs each period. | [optional] 
**AsOf** | Pointer to **string** | AsOf is when that amount was last checked against its source, RFC 3339 date. | [optional] 
**Code** | Pointer to **string** | Code names the obligation so a caller can branch without reading prose. | [optional] 
**Every** | Pointer to **string** | Every is how often it falls due — \&quot;yearly\&quot; for every obligation known here. | [optional] 
**Label** | Pointer to **string** | Label is what the payer sees. | [optional] 
**Minimum** | Pointer to **bool** | Minimum marks a floor rather than a fixed price — a franchise tax that scales with shares or assets is quoted at its minimum, and an entity past the threshold owes more. Saying so is the difference between a quote and a number someone later disputes. | [optional] 
**PassThrough** | Pointer to **bool** | PassThrough marks money we collect and remit to the state rather than keep. | [optional] 
**Source** | Pointer to **string** | Source names the authority that publishes a pass-through amount, so the figure can be checked without first working out who would know. | [optional] 
**Stale** | Pointer to **bool** | Stale reports that AsOf is older than the review window. It tells; it does not block. | [optional] 

## Methods

### NewObligation

`func NewObligation() *Obligation`

NewObligation instantiates a new Obligation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObligationWithDefaults

`func NewObligationWithDefaults() *Obligation`

NewObligationWithDefaults instantiates a new Obligation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *Obligation) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *Obligation) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *Obligation) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *Obligation) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAsOf

`func (o *Obligation) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *Obligation) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *Obligation) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *Obligation) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetCode

`func (o *Obligation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Obligation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Obligation) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Obligation) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEvery

`func (o *Obligation) GetEvery() string`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *Obligation) GetEveryOk() (*string, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *Obligation) SetEvery(v string)`

SetEvery sets Every field to given value.

### HasEvery

`func (o *Obligation) HasEvery() bool`

HasEvery returns a boolean if a field has been set.

### GetLabel

`func (o *Obligation) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Obligation) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Obligation) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Obligation) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMinimum

`func (o *Obligation) GetMinimum() bool`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *Obligation) GetMinimumOk() (*bool, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *Obligation) SetMinimum(v bool)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *Obligation) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetPassThrough

`func (o *Obligation) GetPassThrough() bool`

GetPassThrough returns the PassThrough field if non-nil, zero value otherwise.

### GetPassThroughOk

`func (o *Obligation) GetPassThroughOk() (*bool, bool)`

GetPassThroughOk returns a tuple with the PassThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThrough

`func (o *Obligation) SetPassThrough(v bool)`

SetPassThrough sets PassThrough field to given value.

### HasPassThrough

`func (o *Obligation) HasPassThrough() bool`

HasPassThrough returns a boolean if a field has been set.

### GetSource

`func (o *Obligation) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Obligation) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Obligation) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Obligation) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStale

`func (o *Obligation) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *Obligation) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *Obligation) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *Obligation) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


