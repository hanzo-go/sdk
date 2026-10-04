# GraphGraphAssertOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplicate** | Pointer to **int64** | Duplicate is how many members this plane already held. A redelivery collides on its content address and is counted here, not refused: it is the success a retrying caller depends on. | [optional] 
**Reasons** | Pointer to **[]string** | Reasons names why each refused member was refused, in the order sent. | [optional] 
**Recorded** | Pointer to **int64** | Recorded is how many members became new rows. | [optional] 
**Refused** | Pointer to **int64** | Refused is how many members were turned away — a missing entity, a timestamp that is not RFC 3339, a confidence outside [0,1], a relation whose declared domain or range the entity&#39;s type does not meet, or a declaration from a caller who is not an admin of the organization. The rest of the batch was still recorded; a batch refused whole answers 400, or 403 when a member of it was refused for want of authority. | [optional] 

## Methods

### NewGraphGraphAssertOut

`func NewGraphGraphAssertOut() *GraphGraphAssertOut`

NewGraphGraphAssertOut instantiates a new GraphGraphAssertOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphAssertOutWithDefaults

`func NewGraphGraphAssertOutWithDefaults() *GraphGraphAssertOut`

NewGraphGraphAssertOutWithDefaults instantiates a new GraphGraphAssertOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplicate

`func (o *GraphGraphAssertOut) GetDuplicate() int64`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *GraphGraphAssertOut) GetDuplicateOk() (*int64, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *GraphGraphAssertOut) SetDuplicate(v int64)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *GraphGraphAssertOut) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetReasons

`func (o *GraphGraphAssertOut) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *GraphGraphAssertOut) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *GraphGraphAssertOut) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *GraphGraphAssertOut) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetRecorded

`func (o *GraphGraphAssertOut) GetRecorded() int64`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *GraphGraphAssertOut) GetRecordedOk() (*int64, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *GraphGraphAssertOut) SetRecorded(v int64)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *GraphGraphAssertOut) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetRefused

`func (o *GraphGraphAssertOut) GetRefused() int64`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *GraphGraphAssertOut) GetRefusedOk() (*int64, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *GraphGraphAssertOut) SetRefused(v int64)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *GraphGraphAssertOut) HasRefused() bool`

HasRefused returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


