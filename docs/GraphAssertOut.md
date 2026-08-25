# GraphAssertOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplicate** | Pointer to **int32** | Duplicate is how many members this plane already held. A redelivery collides on its content address and is counted here, not refused: it is the success a retrying caller depends on. | [optional] 
**Reasons** | Pointer to **[]string** | Reasons names why each refused member was refused, in the order sent. | [optional] 
**Recorded** | Pointer to **int32** | Recorded is how many members became new rows. | [optional] 
**Refused** | Pointer to **int32** | Refused is how many members were turned away on arrival, before the store was touched — a missing entity, a timestamp that is not RFC 3339, a confidence outside [0,1]. The rest of the batch was still recorded. | [optional] 

## Methods

### NewGraphAssertOut

`func NewGraphAssertOut() *GraphAssertOut`

NewGraphAssertOut instantiates a new GraphAssertOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphAssertOutWithDefaults

`func NewGraphAssertOutWithDefaults() *GraphAssertOut`

NewGraphAssertOutWithDefaults instantiates a new GraphAssertOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplicate

`func (o *GraphAssertOut) GetDuplicate() int32`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *GraphAssertOut) GetDuplicateOk() (*int32, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *GraphAssertOut) SetDuplicate(v int32)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *GraphAssertOut) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetReasons

`func (o *GraphAssertOut) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *GraphAssertOut) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *GraphAssertOut) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *GraphAssertOut) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetRecorded

`func (o *GraphAssertOut) GetRecorded() int32`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *GraphAssertOut) GetRecordedOk() (*int32, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *GraphAssertOut) SetRecorded(v int32)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *GraphAssertOut) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetRefused

`func (o *GraphAssertOut) GetRefused() int32`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *GraphAssertOut) GetRefusedOk() (*int32, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *GraphAssertOut) SetRefused(v int32)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *GraphAssertOut) HasRefused() bool`

HasRefused returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


