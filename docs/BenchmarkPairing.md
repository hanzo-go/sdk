# BenchmarkPairing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A** | Pointer to **string** | A is the first model id. | [optional] 
**ACorrect** | Pointer to **int64** | ACorrect is how many of those common items A got right. | [optional] 
**B** | Pointer to **string** | B is the second model id. | [optional] 
**BCorrect** | Pointer to **int64** | BCorrect is how many of those common items B got right. | [optional] 
**Benchmark** | Pointer to **string** | Benchmark is the catalog id the two arms were compared on. | [optional] 
**McnemarP** | Pointer to **float64** | McnemarP is the two-sided exact binomial p on the discordant pairs. It is 1 when nothing is discordant, which is \&quot;no evidence of a difference\&quot;, not an error. | [optional] 
**NCommon** | Pointer to **int64** | NCommon is how many items BOTH arms completed. It is the denominator, and the reason this comparison is valid where a raw accuracy difference is not. | [optional] 
**NetAMinusB** | Pointer to **int64** | NetAMinusB is the two rescue counts subtracted — A&#39;s advantage in items. | [optional] 
**RescueAOverB** | Pointer to **int64** | RescueAOverB is how many items A got right and B got wrong. | [optional] 
**RescueBOverA** | Pointer to **int64** | RescueBOverA is how many items B got right and A got wrong. | [optional] 

## Methods

### NewBenchmarkPairing

`func NewBenchmarkPairing() *BenchmarkPairing`

NewBenchmarkPairing instantiates a new BenchmarkPairing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkPairingWithDefaults

`func NewBenchmarkPairingWithDefaults() *BenchmarkPairing`

NewBenchmarkPairingWithDefaults instantiates a new BenchmarkPairing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *BenchmarkPairing) GetA() string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *BenchmarkPairing) GetAOk() (*string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *BenchmarkPairing) SetA(v string)`

SetA sets A field to given value.

### HasA

`func (o *BenchmarkPairing) HasA() bool`

HasA returns a boolean if a field has been set.

### GetACorrect

`func (o *BenchmarkPairing) GetACorrect() int64`

GetACorrect returns the ACorrect field if non-nil, zero value otherwise.

### GetACorrectOk

`func (o *BenchmarkPairing) GetACorrectOk() (*int64, bool)`

GetACorrectOk returns a tuple with the ACorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACorrect

`func (o *BenchmarkPairing) SetACorrect(v int64)`

SetACorrect sets ACorrect field to given value.

### HasACorrect

`func (o *BenchmarkPairing) HasACorrect() bool`

HasACorrect returns a boolean if a field has been set.

### GetB

`func (o *BenchmarkPairing) GetB() string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *BenchmarkPairing) GetBOk() (*string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *BenchmarkPairing) SetB(v string)`

SetB sets B field to given value.

### HasB

`func (o *BenchmarkPairing) HasB() bool`

HasB returns a boolean if a field has been set.

### GetBCorrect

`func (o *BenchmarkPairing) GetBCorrect() int64`

GetBCorrect returns the BCorrect field if non-nil, zero value otherwise.

### GetBCorrectOk

`func (o *BenchmarkPairing) GetBCorrectOk() (*int64, bool)`

GetBCorrectOk returns a tuple with the BCorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCorrect

`func (o *BenchmarkPairing) SetBCorrect(v int64)`

SetBCorrect sets BCorrect field to given value.

### HasBCorrect

`func (o *BenchmarkPairing) HasBCorrect() bool`

HasBCorrect returns a boolean if a field has been set.

### GetBenchmark

`func (o *BenchmarkPairing) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *BenchmarkPairing) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *BenchmarkPairing) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *BenchmarkPairing) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetMcnemarP

`func (o *BenchmarkPairing) GetMcnemarP() float64`

GetMcnemarP returns the McnemarP field if non-nil, zero value otherwise.

### GetMcnemarPOk

`func (o *BenchmarkPairing) GetMcnemarPOk() (*float64, bool)`

GetMcnemarPOk returns a tuple with the McnemarP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcnemarP

`func (o *BenchmarkPairing) SetMcnemarP(v float64)`

SetMcnemarP sets McnemarP field to given value.

### HasMcnemarP

`func (o *BenchmarkPairing) HasMcnemarP() bool`

HasMcnemarP returns a boolean if a field has been set.

### GetNCommon

`func (o *BenchmarkPairing) GetNCommon() int64`

GetNCommon returns the NCommon field if non-nil, zero value otherwise.

### GetNCommonOk

`func (o *BenchmarkPairing) GetNCommonOk() (*int64, bool)`

GetNCommonOk returns a tuple with the NCommon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCommon

`func (o *BenchmarkPairing) SetNCommon(v int64)`

SetNCommon sets NCommon field to given value.

### HasNCommon

`func (o *BenchmarkPairing) HasNCommon() bool`

HasNCommon returns a boolean if a field has been set.

### GetNetAMinusB

`func (o *BenchmarkPairing) GetNetAMinusB() int64`

GetNetAMinusB returns the NetAMinusB field if non-nil, zero value otherwise.

### GetNetAMinusBOk

`func (o *BenchmarkPairing) GetNetAMinusBOk() (*int64, bool)`

GetNetAMinusBOk returns a tuple with the NetAMinusB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAMinusB

`func (o *BenchmarkPairing) SetNetAMinusB(v int64)`

SetNetAMinusB sets NetAMinusB field to given value.

### HasNetAMinusB

`func (o *BenchmarkPairing) HasNetAMinusB() bool`

HasNetAMinusB returns a boolean if a field has been set.

### GetRescueAOverB

`func (o *BenchmarkPairing) GetRescueAOverB() int64`

GetRescueAOverB returns the RescueAOverB field if non-nil, zero value otherwise.

### GetRescueAOverBOk

`func (o *BenchmarkPairing) GetRescueAOverBOk() (*int64, bool)`

GetRescueAOverBOk returns a tuple with the RescueAOverB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueAOverB

`func (o *BenchmarkPairing) SetRescueAOverB(v int64)`

SetRescueAOverB sets RescueAOverB field to given value.

### HasRescueAOverB

`func (o *BenchmarkPairing) HasRescueAOverB() bool`

HasRescueAOverB returns a boolean if a field has been set.

### GetRescueBOverA

`func (o *BenchmarkPairing) GetRescueBOverA() int64`

GetRescueBOverA returns the RescueBOverA field if non-nil, zero value otherwise.

### GetRescueBOverAOk

`func (o *BenchmarkPairing) GetRescueBOverAOk() (*int64, bool)`

GetRescueBOverAOk returns a tuple with the RescueBOverA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueBOverA

`func (o *BenchmarkPairing) SetRescueBOverA(v int64)`

SetRescueBOverA sets RescueBOverA field to given value.

### HasRescueBOverA

`func (o *BenchmarkPairing) HasRescueBOverA() bool`

HasRescueBOverA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


