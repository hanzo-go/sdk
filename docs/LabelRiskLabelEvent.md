# LabelRiskLabelEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is the event&#39;s own instant, RFC 3339. It is part of the event&#39;s IDENTITY and not a filter: it is matched exactly, to the second, against the &#x60;at&#x60; the assertions were filed under, so an instant a second off names a different event and resolves to nothing. It is also what this event&#39;s as-of is measured from — At plus the horizon. | [optional] 
**Kind** | Pointer to **string** | Kind is the judged entity&#39;s type, from the closed set: account, agent, merchant, payout, person, session or transaction. One outside it is refused rather than answered &#x60;unlabelled&#x60;, because it could only ever match nothing and the caller would read a real absence into a typo. | [optional] 
**Subject** | Pointer to **string** | Subject is the entity id in the tenant&#39;s own namespace, at most 512 bytes. It is matched EXACTLY against what was recorded — this is a lookup, not a search, and no prefix, pattern or normalisation is applied. | [optional] 

## Methods

### NewLabelRiskLabelEvent

`func NewLabelRiskLabelEvent() *LabelRiskLabelEvent`

NewLabelRiskLabelEvent instantiates a new LabelRiskLabelEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelRiskLabelEventWithDefaults

`func NewLabelRiskLabelEventWithDefaults() *LabelRiskLabelEvent`

NewLabelRiskLabelEventWithDefaults instantiates a new LabelRiskLabelEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *LabelRiskLabelEvent) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *LabelRiskLabelEvent) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *LabelRiskLabelEvent) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *LabelRiskLabelEvent) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetKind

`func (o *LabelRiskLabelEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LabelRiskLabelEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LabelRiskLabelEvent) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LabelRiskLabelEvent) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSubject

`func (o *LabelRiskLabelEvent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *LabelRiskLabelEvent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *LabelRiskLabelEvent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *LabelRiskLabelEvent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


