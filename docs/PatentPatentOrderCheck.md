# PatentPatentOrderCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GapSeq** | Pointer to **int64** | GapSeq is how many appends separate them, negative when the implementation precedes the idea. It is a distance in the LOG rather than in seconds, because seconds are comparable only within one server&#39;s clock. | [optional] 
**Idea** | Pointer to **string** | Idea is the conception it cites. | [optional] 
**Implementation** | Pointer to **string** | Implementation is the artifact being examined. | [optional] 
**Verdict** | Pointer to **string** | Verdict is before, after or unproven. | [optional] 
**Why** | Pointer to **string** | Why states, in one sentence, what the verdict is resting on. A verdict without its basis is the thing this whole surface exists to avoid. | [optional] 

## Methods

### NewPatentPatentOrderCheck

`func NewPatentPatentOrderCheck() *PatentPatentOrderCheck`

NewPatentPatentOrderCheck instantiates a new PatentPatentOrderCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentOrderCheckWithDefaults

`func NewPatentPatentOrderCheckWithDefaults() *PatentPatentOrderCheck`

NewPatentPatentOrderCheckWithDefaults instantiates a new PatentPatentOrderCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGapSeq

`func (o *PatentPatentOrderCheck) GetGapSeq() int64`

GetGapSeq returns the GapSeq field if non-nil, zero value otherwise.

### GetGapSeqOk

`func (o *PatentPatentOrderCheck) GetGapSeqOk() (*int64, bool)`

GetGapSeqOk returns a tuple with the GapSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGapSeq

`func (o *PatentPatentOrderCheck) SetGapSeq(v int64)`

SetGapSeq sets GapSeq field to given value.

### HasGapSeq

`func (o *PatentPatentOrderCheck) HasGapSeq() bool`

HasGapSeq returns a boolean if a field has been set.

### GetIdea

`func (o *PatentPatentOrderCheck) GetIdea() string`

GetIdea returns the Idea field if non-nil, zero value otherwise.

### GetIdeaOk

`func (o *PatentPatentOrderCheck) GetIdeaOk() (*string, bool)`

GetIdeaOk returns a tuple with the Idea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdea

`func (o *PatentPatentOrderCheck) SetIdea(v string)`

SetIdea sets Idea field to given value.

### HasIdea

`func (o *PatentPatentOrderCheck) HasIdea() bool`

HasIdea returns a boolean if a field has been set.

### GetImplementation

`func (o *PatentPatentOrderCheck) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *PatentPatentOrderCheck) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *PatentPatentOrderCheck) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *PatentPatentOrderCheck) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### GetVerdict

`func (o *PatentPatentOrderCheck) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *PatentPatentOrderCheck) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *PatentPatentOrderCheck) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *PatentPatentOrderCheck) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### GetWhy

`func (o *PatentPatentOrderCheck) GetWhy() string`

GetWhy returns the Why field if non-nil, zero value otherwise.

### GetWhyOk

`func (o *PatentPatentOrderCheck) GetWhyOk() (*string, bool)`

GetWhyOk returns a tuple with the Why field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhy

`func (o *PatentPatentOrderCheck) SetWhy(v string)`

SetWhy sets Why field to given value.

### HasWhy

`func (o *PatentPatentOrderCheck) HasWhy() bool`

HasWhy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


