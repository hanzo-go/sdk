# ResearchAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | **string** |  | 
**Item** | **string** |  | 
**Model** | **string** |  | 
**Revision** | Pointer to **string** |  | [optional] [default to "original"]
**Status** | Pointer to **string** | faulted retains a negative result | [optional] [default to "complete"]
**Gold** | Pointer to **string** |  | [optional] 
**Answer** | Pointer to **string** |  | [optional] 
**Correct** | Pointer to **bool** |  | [optional] 
**Response** | Pointer to **string** | raw model output (raw-artifact retention class) | [optional] 
**Source** | Pointer to **string** |  | [optional] [default to "hanzo-measured"]
**Ts** | Pointer to **int64** |  | [optional] 

## Methods

### NewResearchAttempt

`func NewResearchAttempt(benchmark string, item string, model string, ) *ResearchAttempt`

NewResearchAttempt instantiates a new ResearchAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchAttemptWithDefaults

`func NewResearchAttemptWithDefaults() *ResearchAttempt`

NewResearchAttemptWithDefaults instantiates a new ResearchAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *ResearchAttempt) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *ResearchAttempt) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *ResearchAttempt) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.


### GetItem

`func (o *ResearchAttempt) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ResearchAttempt) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ResearchAttempt) SetItem(v string)`

SetItem sets Item field to given value.


### GetModel

`func (o *ResearchAttempt) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ResearchAttempt) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ResearchAttempt) SetModel(v string)`

SetModel sets Model field to given value.


### GetRevision

`func (o *ResearchAttempt) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ResearchAttempt) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ResearchAttempt) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ResearchAttempt) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStatus

`func (o *ResearchAttempt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResearchAttempt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResearchAttempt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResearchAttempt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGold

`func (o *ResearchAttempt) GetGold() string`

GetGold returns the Gold field if non-nil, zero value otherwise.

### GetGoldOk

`func (o *ResearchAttempt) GetGoldOk() (*string, bool)`

GetGoldOk returns a tuple with the Gold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGold

`func (o *ResearchAttempt) SetGold(v string)`

SetGold sets Gold field to given value.

### HasGold

`func (o *ResearchAttempt) HasGold() bool`

HasGold returns a boolean if a field has been set.

### GetAnswer

`func (o *ResearchAttempt) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *ResearchAttempt) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *ResearchAttempt) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *ResearchAttempt) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetCorrect

`func (o *ResearchAttempt) GetCorrect() bool`

GetCorrect returns the Correct field if non-nil, zero value otherwise.

### GetCorrectOk

`func (o *ResearchAttempt) GetCorrectOk() (*bool, bool)`

GetCorrectOk returns a tuple with the Correct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrect

`func (o *ResearchAttempt) SetCorrect(v bool)`

SetCorrect sets Correct field to given value.

### HasCorrect

`func (o *ResearchAttempt) HasCorrect() bool`

HasCorrect returns a boolean if a field has been set.

### GetResponse

`func (o *ResearchAttempt) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ResearchAttempt) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ResearchAttempt) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ResearchAttempt) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSource

`func (o *ResearchAttempt) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ResearchAttempt) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ResearchAttempt) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ResearchAttempt) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTs

`func (o *ResearchAttempt) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ResearchAttempt) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ResearchAttempt) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ResearchAttempt) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


