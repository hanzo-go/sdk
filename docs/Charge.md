# Charge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is what this line costs. | [optional] 
**AsOf** | Pointer to **string** | AsOf is when a pass-through amount was last checked against its source. | [optional] 
**Code** | Pointer to **string** | Code names the line so a caller can branch on it without reading prose. | [optional] 
**Label** | Pointer to **string** | Label is what the payer sees on the invoice. | [optional] 
**PassThrough** | Pointer to **bool** | PassThrough marks money we collect and remit rather than keep — the state&#39;s fee is not our revenue, and a quote that hides that is a quote that reads as a bigger margin than it is. | [optional] 
**Recurring** | Pointer to **string** | Recurring marks a line that repeats. An agent of record is billed every year for as long as the entity stands, and a payer agreeing to a one-time total is not agreeing to that. | [optional] 
**Source** | Pointer to **string** | Source names who publishes this amount, for a line we merely pass through. Empty for a price of ours, which needs no external authority. | [optional] 
**Stale** | Pointer to **bool** | Stale reports that AsOf is older than the review window — the figure may have moved and nobody has looked. It does not block; it tells. | [optional] 

## Methods

### NewCharge

`func NewCharge() *Charge`

NewCharge instantiates a new Charge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeWithDefaults

`func NewChargeWithDefaults() *Charge`

NewChargeWithDefaults instantiates a new Charge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *Charge) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *Charge) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *Charge) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *Charge) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAsOf

`func (o *Charge) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *Charge) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *Charge) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *Charge) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetCode

`func (o *Charge) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Charge) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Charge) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Charge) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *Charge) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Charge) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Charge) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Charge) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPassThrough

`func (o *Charge) GetPassThrough() bool`

GetPassThrough returns the PassThrough field if non-nil, zero value otherwise.

### GetPassThroughOk

`func (o *Charge) GetPassThroughOk() (*bool, bool)`

GetPassThroughOk returns a tuple with the PassThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThrough

`func (o *Charge) SetPassThrough(v bool)`

SetPassThrough sets PassThrough field to given value.

### HasPassThrough

`func (o *Charge) HasPassThrough() bool`

HasPassThrough returns a boolean if a field has been set.

### GetRecurring

`func (o *Charge) GetRecurring() string`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *Charge) GetRecurringOk() (*string, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *Charge) SetRecurring(v string)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *Charge) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetSource

`func (o *Charge) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Charge) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Charge) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Charge) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStale

`func (o *Charge) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *Charge) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *Charge) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *Charge) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


