# Attempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** |  | [optional] 
**Benchmark** | Pointer to **string** |  | [optional] 
**Correct** | Pointer to **bool** |  | [optional] 
**Gold** | Pointer to **string** |  | [optional] 
**Item** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Response** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Ts** | Pointer to **int32** |  | [optional] 

## Methods

### NewAttempt

`func NewAttempt() *Attempt`

NewAttempt instantiates a new Attempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptWithDefaults

`func NewAttemptWithDefaults() *Attempt`

NewAttemptWithDefaults instantiates a new Attempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *Attempt) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *Attempt) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *Attempt) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *Attempt) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetBenchmark

`func (o *Attempt) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *Attempt) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *Attempt) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *Attempt) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetCorrect

`func (o *Attempt) GetCorrect() bool`

GetCorrect returns the Correct field if non-nil, zero value otherwise.

### GetCorrectOk

`func (o *Attempt) GetCorrectOk() (*bool, bool)`

GetCorrectOk returns a tuple with the Correct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrect

`func (o *Attempt) SetCorrect(v bool)`

SetCorrect sets Correct field to given value.

### HasCorrect

`func (o *Attempt) HasCorrect() bool`

HasCorrect returns a boolean if a field has been set.

### GetGold

`func (o *Attempt) GetGold() string`

GetGold returns the Gold field if non-nil, zero value otherwise.

### GetGoldOk

`func (o *Attempt) GetGoldOk() (*string, bool)`

GetGoldOk returns a tuple with the Gold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGold

`func (o *Attempt) SetGold(v string)`

SetGold sets Gold field to given value.

### HasGold

`func (o *Attempt) HasGold() bool`

HasGold returns a boolean if a field has been set.

### GetItem

`func (o *Attempt) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *Attempt) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *Attempt) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *Attempt) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetModel

`func (o *Attempt) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Attempt) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Attempt) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Attempt) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetResponse

`func (o *Attempt) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Attempt) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Attempt) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Attempt) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetRevision

`func (o *Attempt) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Attempt) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Attempt) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *Attempt) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSource

`func (o *Attempt) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Attempt) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Attempt) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Attempt) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *Attempt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Attempt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Attempt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Attempt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTs

`func (o *Attempt) GetTs() int32`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *Attempt) GetTsOk() (*int32, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *Attempt) SetTs(v int32)`

SetTs sets Ts field to given value.

### HasTs

`func (o *Attempt) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


