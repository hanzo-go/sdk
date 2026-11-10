# GraphGraphConvoOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertions** | Pointer to [**[]GraphGraphFact**](GraphGraphFact.md) | Assertions is the assertions extracted from the turns, in turn order. | [optional] 
**Duplicate** | Pointer to **int64** | Duplicate is how many extracted assertions the store already held. | [optional] 
**Reasons** | Pointer to **[]string** | Reasons names why each refused assertion was refused, in order. | [optional] 
**Recorded** | Pointer to **int64** | Recorded is how many extracted assertions became new rows in the store. Zero when DryRun is true. | [optional] 
**Refused** | Pointer to **int64** | Refused is how many extracted assertions were refused by admission or schema checks. | [optional] 

## Methods

### NewGraphGraphConvoOut

`func NewGraphGraphConvoOut() *GraphGraphConvoOut`

NewGraphGraphConvoOut instantiates a new GraphGraphConvoOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphConvoOutWithDefaults

`func NewGraphGraphConvoOutWithDefaults() *GraphGraphConvoOut`

NewGraphGraphConvoOutWithDefaults instantiates a new GraphGraphConvoOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertions

`func (o *GraphGraphConvoOut) GetAssertions() []GraphGraphFact`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *GraphGraphConvoOut) GetAssertionsOk() (*[]GraphGraphFact, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *GraphGraphConvoOut) SetAssertions(v []GraphGraphFact)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *GraphGraphConvoOut) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.

### GetDuplicate

`func (o *GraphGraphConvoOut) GetDuplicate() int64`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *GraphGraphConvoOut) GetDuplicateOk() (*int64, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *GraphGraphConvoOut) SetDuplicate(v int64)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *GraphGraphConvoOut) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetReasons

`func (o *GraphGraphConvoOut) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *GraphGraphConvoOut) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *GraphGraphConvoOut) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *GraphGraphConvoOut) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetRecorded

`func (o *GraphGraphConvoOut) GetRecorded() int64`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *GraphGraphConvoOut) GetRecordedOk() (*int64, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *GraphGraphConvoOut) SetRecorded(v int64)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *GraphGraphConvoOut) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetRefused

`func (o *GraphGraphConvoOut) GetRefused() int64`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *GraphGraphConvoOut) GetRefusedOk() (*int64, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *GraphGraphConvoOut) SetRefused(v int64)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *GraphGraphConvoOut) HasRefused() bool`

HasRefused returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


