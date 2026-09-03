# Pairing

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

### NewPairing

`func NewPairing() *Pairing`

NewPairing instantiates a new Pairing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPairingWithDefaults

`func NewPairingWithDefaults() *Pairing`

NewPairingWithDefaults instantiates a new Pairing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *Pairing) GetA() string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *Pairing) GetAOk() (*string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *Pairing) SetA(v string)`

SetA sets A field to given value.

### HasA

`func (o *Pairing) HasA() bool`

HasA returns a boolean if a field has been set.

### GetACorrect

`func (o *Pairing) GetACorrect() int64`

GetACorrect returns the ACorrect field if non-nil, zero value otherwise.

### GetACorrectOk

`func (o *Pairing) GetACorrectOk() (*int64, bool)`

GetACorrectOk returns a tuple with the ACorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACorrect

`func (o *Pairing) SetACorrect(v int64)`

SetACorrect sets ACorrect field to given value.

### HasACorrect

`func (o *Pairing) HasACorrect() bool`

HasACorrect returns a boolean if a field has been set.

### GetB

`func (o *Pairing) GetB() string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *Pairing) GetBOk() (*string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *Pairing) SetB(v string)`

SetB sets B field to given value.

### HasB

`func (o *Pairing) HasB() bool`

HasB returns a boolean if a field has been set.

### GetBCorrect

`func (o *Pairing) GetBCorrect() int64`

GetBCorrect returns the BCorrect field if non-nil, zero value otherwise.

### GetBCorrectOk

`func (o *Pairing) GetBCorrectOk() (*int64, bool)`

GetBCorrectOk returns a tuple with the BCorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCorrect

`func (o *Pairing) SetBCorrect(v int64)`

SetBCorrect sets BCorrect field to given value.

### HasBCorrect

`func (o *Pairing) HasBCorrect() bool`

HasBCorrect returns a boolean if a field has been set.

### GetBenchmark

`func (o *Pairing) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *Pairing) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *Pairing) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *Pairing) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetMcnemarP

`func (o *Pairing) GetMcnemarP() float64`

GetMcnemarP returns the McnemarP field if non-nil, zero value otherwise.

### GetMcnemarPOk

`func (o *Pairing) GetMcnemarPOk() (*float64, bool)`

GetMcnemarPOk returns a tuple with the McnemarP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcnemarP

`func (o *Pairing) SetMcnemarP(v float64)`

SetMcnemarP sets McnemarP field to given value.

### HasMcnemarP

`func (o *Pairing) HasMcnemarP() bool`

HasMcnemarP returns a boolean if a field has been set.

### GetNCommon

`func (o *Pairing) GetNCommon() int64`

GetNCommon returns the NCommon field if non-nil, zero value otherwise.

### GetNCommonOk

`func (o *Pairing) GetNCommonOk() (*int64, bool)`

GetNCommonOk returns a tuple with the NCommon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCommon

`func (o *Pairing) SetNCommon(v int64)`

SetNCommon sets NCommon field to given value.

### HasNCommon

`func (o *Pairing) HasNCommon() bool`

HasNCommon returns a boolean if a field has been set.

### GetNetAMinusB

`func (o *Pairing) GetNetAMinusB() int64`

GetNetAMinusB returns the NetAMinusB field if non-nil, zero value otherwise.

### GetNetAMinusBOk

`func (o *Pairing) GetNetAMinusBOk() (*int64, bool)`

GetNetAMinusBOk returns a tuple with the NetAMinusB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAMinusB

`func (o *Pairing) SetNetAMinusB(v int64)`

SetNetAMinusB sets NetAMinusB field to given value.

### HasNetAMinusB

`func (o *Pairing) HasNetAMinusB() bool`

HasNetAMinusB returns a boolean if a field has been set.

### GetRescueAOverB

`func (o *Pairing) GetRescueAOverB() int64`

GetRescueAOverB returns the RescueAOverB field if non-nil, zero value otherwise.

### GetRescueAOverBOk

`func (o *Pairing) GetRescueAOverBOk() (*int64, bool)`

GetRescueAOverBOk returns a tuple with the RescueAOverB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueAOverB

`func (o *Pairing) SetRescueAOverB(v int64)`

SetRescueAOverB sets RescueAOverB field to given value.

### HasRescueAOverB

`func (o *Pairing) HasRescueAOverB() bool`

HasRescueAOverB returns a boolean if a field has been set.

### GetRescueBOverA

`func (o *Pairing) GetRescueBOverA() int64`

GetRescueBOverA returns the RescueBOverA field if non-nil, zero value otherwise.

### GetRescueBOverAOk

`func (o *Pairing) GetRescueBOverAOk() (*int64, bool)`

GetRescueBOverAOk returns a tuple with the RescueBOverA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueBOverA

`func (o *Pairing) SetRescueBOverA(v int64)`

SetRescueBOverA sets RescueBOverA field to given value.

### HasRescueBOverA

`func (o *Pairing) HasRescueBOverA() bool`

HasRescueBOverA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


