# BenchmarkClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **string** | Benchmark is the canonical test id the claim is about, from /catalog. | [optional] 
**Evidence** | Pointer to **[]string** | Evidence are the evidence uris the claim rests on, at most 64. | [optional] 
**Model** | Pointer to **string** | Model is the system the score is claimed for. | [optional] 
**Protocol** | Pointer to **string** | Protocol records HOW it was scored — provider-reported, agentic, third-party-leaderboard — so a provider card is never read as a measurement. | [optional] 
**Provider** | Pointer to **string** | Provider is who the claim belongs to — the lab or leaderboard whose number this is. | [optional] 
**Runs** | Pointer to **[]string** | Runs are the research run ids the claim rests on, at most 64. | [optional] 
**Score** | Pointer to **float64** | Score is the reported aggregate, as a percentage. | [optional] 
**Source** | Pointer to **string** | Source is the citation the row was read from. Required. | [optional] 

## Methods

### NewBenchmarkClaim

`func NewBenchmarkClaim() *BenchmarkClaim`

NewBenchmarkClaim instantiates a new BenchmarkClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkClaimWithDefaults

`func NewBenchmarkClaimWithDefaults() *BenchmarkClaim`

NewBenchmarkClaimWithDefaults instantiates a new BenchmarkClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *BenchmarkClaim) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *BenchmarkClaim) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *BenchmarkClaim) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *BenchmarkClaim) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetEvidence

`func (o *BenchmarkClaim) GetEvidence() []string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *BenchmarkClaim) GetEvidenceOk() (*[]string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *BenchmarkClaim) SetEvidence(v []string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *BenchmarkClaim) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetModel

`func (o *BenchmarkClaim) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BenchmarkClaim) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BenchmarkClaim) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BenchmarkClaim) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProtocol

`func (o *BenchmarkClaim) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BenchmarkClaim) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BenchmarkClaim) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BenchmarkClaim) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProvider

`func (o *BenchmarkClaim) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BenchmarkClaim) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BenchmarkClaim) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *BenchmarkClaim) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRuns

`func (o *BenchmarkClaim) GetRuns() []string`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *BenchmarkClaim) GetRunsOk() (*[]string, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *BenchmarkClaim) SetRuns(v []string)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *BenchmarkClaim) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetScore

`func (o *BenchmarkClaim) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BenchmarkClaim) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BenchmarkClaim) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *BenchmarkClaim) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSource

`func (o *BenchmarkClaim) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BenchmarkClaim) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BenchmarkClaim) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BenchmarkClaim) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


