# PublishedClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **string** | Benchmark is the canonical test id the claim is about, from /catalog. | [optional] 
**Model** | Pointer to **string** | Model is the system the score is claimed for. | [optional] 
**Protocol** | Pointer to **string** | Protocol records HOW it was scored — provider-reported, agentic, third-party-leaderboard — because a provider card and a third party running its own harness are different kinds of number and must not be blended. | [optional] 
**Provider** | Pointer to **string** | Provider is who the claim belongs to — the lab or leaderboard whose number this is. It joins a claim to the attempts measured for that same model. | [optional] 
**Score** | Pointer to **float64** | Score is the reported aggregate, as a percentage. | [optional] 
**Source** | Pointer to **string** | Source is the citation the row was read from. A claim without one is a number nobody can check, so every write requires it. | [optional] 

## Methods

### NewPublishedClaim

`func NewPublishedClaim() *PublishedClaim`

NewPublishedClaim instantiates a new PublishedClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedClaimWithDefaults

`func NewPublishedClaimWithDefaults() *PublishedClaim`

NewPublishedClaimWithDefaults instantiates a new PublishedClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *PublishedClaim) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *PublishedClaim) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *PublishedClaim) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *PublishedClaim) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetModel

`func (o *PublishedClaim) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PublishedClaim) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PublishedClaim) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PublishedClaim) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProtocol

`func (o *PublishedClaim) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PublishedClaim) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PublishedClaim) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PublishedClaim) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProvider

`func (o *PublishedClaim) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PublishedClaim) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PublishedClaim) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PublishedClaim) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetScore

`func (o *PublishedClaim) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *PublishedClaim) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *PublishedClaim) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *PublishedClaim) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSource

`func (o *PublishedClaim) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PublishedClaim) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PublishedClaim) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PublishedClaim) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


