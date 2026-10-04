# GraphGraphDiffOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asserted** | Pointer to [**[]GraphWireFact**](GraphWireFact.md) | Asserted is each statement holding at To that did not at From. | [optional] 
**Bound** | Pointer to **int64** | Bound is how many pairs this diff examined at most. | [optional] 
**From** | Pointer to **string** | From is the earlier instant, RFC 3339, echoed. | [optional] 
**FromKnown** | Pointer to **string** | FromKnown is the knowledge instant of the baseline, RFC 3339. | [optional] 
**Retracted** | Pointer to [**[]GraphGraphChange**](GraphGraphChange.md) | Retracted is each statement that held at From and no longer does at To, with what ended it as Now: the retraction, the later value, or the statement itself when its own until did. | [optional] 
**Superseded** | Pointer to [**[]GraphGraphChange**](GraphGraphChange.md) | Superseded is each property whose value at From was replaced by a different one at To, once per pair however many came in between. | [optional] 
**To** | Pointer to **string** | To is the later instant, RFC 3339: the one asked for, or the server&#39;s clock when none was. | [optional] 
**ToKnown** | Pointer to **string** | ToKnown is the knowledge instant of the later point, RFC 3339. | [optional] 
**Truncated** | Pointer to **bool** | Truncated says more pairs gathered assertions in the interval than Bound, or that one pair holds more than a single read returns. Narrow by entity or by a shorter interval to see the rest. | [optional] 

## Methods

### NewGraphGraphDiffOut

`func NewGraphGraphDiffOut() *GraphGraphDiffOut`

NewGraphGraphDiffOut instantiates a new GraphGraphDiffOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphDiffOutWithDefaults

`func NewGraphGraphDiffOutWithDefaults() *GraphGraphDiffOut`

NewGraphGraphDiffOutWithDefaults instantiates a new GraphGraphDiffOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsserted

`func (o *GraphGraphDiffOut) GetAsserted() []GraphWireFact`

GetAsserted returns the Asserted field if non-nil, zero value otherwise.

### GetAssertedOk

`func (o *GraphGraphDiffOut) GetAssertedOk() (*[]GraphWireFact, bool)`

GetAssertedOk returns a tuple with the Asserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsserted

`func (o *GraphGraphDiffOut) SetAsserted(v []GraphWireFact)`

SetAsserted sets Asserted field to given value.

### HasAsserted

`func (o *GraphGraphDiffOut) HasAsserted() bool`

HasAsserted returns a boolean if a field has been set.

### GetBound

`func (o *GraphGraphDiffOut) GetBound() int64`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *GraphGraphDiffOut) GetBoundOk() (*int64, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *GraphGraphDiffOut) SetBound(v int64)`

SetBound sets Bound field to given value.

### HasBound

`func (o *GraphGraphDiffOut) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetFrom

`func (o *GraphGraphDiffOut) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GraphGraphDiffOut) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GraphGraphDiffOut) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GraphGraphDiffOut) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFromKnown

`func (o *GraphGraphDiffOut) GetFromKnown() string`

GetFromKnown returns the FromKnown field if non-nil, zero value otherwise.

### GetFromKnownOk

`func (o *GraphGraphDiffOut) GetFromKnownOk() (*string, bool)`

GetFromKnownOk returns a tuple with the FromKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromKnown

`func (o *GraphGraphDiffOut) SetFromKnown(v string)`

SetFromKnown sets FromKnown field to given value.

### HasFromKnown

`func (o *GraphGraphDiffOut) HasFromKnown() bool`

HasFromKnown returns a boolean if a field has been set.

### GetRetracted

`func (o *GraphGraphDiffOut) GetRetracted() []GraphGraphChange`

GetRetracted returns the Retracted field if non-nil, zero value otherwise.

### GetRetractedOk

`func (o *GraphGraphDiffOut) GetRetractedOk() (*[]GraphGraphChange, bool)`

GetRetractedOk returns a tuple with the Retracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetracted

`func (o *GraphGraphDiffOut) SetRetracted(v []GraphGraphChange)`

SetRetracted sets Retracted field to given value.

### HasRetracted

`func (o *GraphGraphDiffOut) HasRetracted() bool`

HasRetracted returns a boolean if a field has been set.

### GetSuperseded

`func (o *GraphGraphDiffOut) GetSuperseded() []GraphGraphChange`

GetSuperseded returns the Superseded field if non-nil, zero value otherwise.

### GetSupersededOk

`func (o *GraphGraphDiffOut) GetSupersededOk() (*[]GraphGraphChange, bool)`

GetSupersededOk returns a tuple with the Superseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperseded

`func (o *GraphGraphDiffOut) SetSuperseded(v []GraphGraphChange)`

SetSuperseded sets Superseded field to given value.

### HasSuperseded

`func (o *GraphGraphDiffOut) HasSuperseded() bool`

HasSuperseded returns a boolean if a field has been set.

### GetTo

`func (o *GraphGraphDiffOut) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GraphGraphDiffOut) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GraphGraphDiffOut) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GraphGraphDiffOut) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetToKnown

`func (o *GraphGraphDiffOut) GetToKnown() string`

GetToKnown returns the ToKnown field if non-nil, zero value otherwise.

### GetToKnownOk

`func (o *GraphGraphDiffOut) GetToKnownOk() (*string, bool)`

GetToKnownOk returns a tuple with the ToKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToKnown

`func (o *GraphGraphDiffOut) SetToKnown(v string)`

SetToKnown sets ToKnown field to given value.

### HasToKnown

`func (o *GraphGraphDiffOut) HasToKnown() bool`

HasToKnown returns a boolean if a field has been set.

### GetTruncated

`func (o *GraphGraphDiffOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GraphGraphDiffOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GraphGraphDiffOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GraphGraphDiffOut) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


