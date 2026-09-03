# O11yO11yOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCount** | Pointer to **int32** | ErrorCount is how many of those runs errored. | [optional] 
**Name** | Pointer to **string** | Name is the operation (span name). | [optional] 
**NumCalls** | Pointer to **int32** | NumCalls is how many times it ran in the window. | [optional] 
**P50** | Pointer to **float64** | P50 is its median latency, nanoseconds. | [optional] 
**P95** | Pointer to **float64** | P95 is its p95 latency, nanoseconds. | [optional] 
**P99** | Pointer to **float64** | P99 is its p99 latency, nanoseconds. | [optional] 

## Methods

### NewO11yO11yOperation

`func NewO11yO11yOperation() *O11yO11yOperation`

NewO11yO11yOperation instantiates a new O11yO11yOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yOperationWithDefaults

`func NewO11yO11yOperationWithDefaults() *O11yO11yOperation`

NewO11yO11yOperationWithDefaults instantiates a new O11yO11yOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCount

`func (o *O11yO11yOperation) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *O11yO11yOperation) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *O11yO11yOperation) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *O11yO11yOperation) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yOperation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yOperation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumCalls

`func (o *O11yO11yOperation) GetNumCalls() int32`

GetNumCalls returns the NumCalls field if non-nil, zero value otherwise.

### GetNumCallsOk

`func (o *O11yO11yOperation) GetNumCallsOk() (*int32, bool)`

GetNumCallsOk returns a tuple with the NumCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCalls

`func (o *O11yO11yOperation) SetNumCalls(v int32)`

SetNumCalls sets NumCalls field to given value.

### HasNumCalls

`func (o *O11yO11yOperation) HasNumCalls() bool`

HasNumCalls returns a boolean if a field has been set.

### GetP50

`func (o *O11yO11yOperation) GetP50() float64`

GetP50 returns the P50 field if non-nil, zero value otherwise.

### GetP50Ok

`func (o *O11yO11yOperation) GetP50Ok() (*float64, bool)`

GetP50Ok returns a tuple with the P50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP50

`func (o *O11yO11yOperation) SetP50(v float64)`

SetP50 sets P50 field to given value.

### HasP50

`func (o *O11yO11yOperation) HasP50() bool`

HasP50 returns a boolean if a field has been set.

### GetP95

`func (o *O11yO11yOperation) GetP95() float64`

GetP95 returns the P95 field if non-nil, zero value otherwise.

### GetP95Ok

`func (o *O11yO11yOperation) GetP95Ok() (*float64, bool)`

GetP95Ok returns a tuple with the P95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95

`func (o *O11yO11yOperation) SetP95(v float64)`

SetP95 sets P95 field to given value.

### HasP95

`func (o *O11yO11yOperation) HasP95() bool`

HasP95 returns a boolean if a field has been set.

### GetP99

`func (o *O11yO11yOperation) GetP99() float64`

GetP99 returns the P99 field if non-nil, zero value otherwise.

### GetP99Ok

`func (o *O11yO11yOperation) GetP99Ok() (*float64, bool)`

GetP99Ok returns a tuple with the P99 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99

`func (o *O11yO11yOperation) SetP99(v float64)`

SetP99 sets P99 field to given value.

### HasP99

`func (o *O11yO11yOperation) HasP99() bool`

HasP99 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


