# ClaimRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **time.Time** | At is when a stored row was recorded. Zero for a seed row. | [optional] 
**Benchmark** | Pointer to **string** | Benchmark is the canonical test id the claim is about, from /catalog. | [optional] 
**By** | Pointer to **string** | By is who recorded it, when the caller said. | [optional] 
**Model** | Pointer to **string** | Model is the system the score is claimed for. | [optional] 
**Origin** | Pointer to **string** | Origin is \&quot;seed\&quot; for a compiled row and \&quot;stored\&quot; for one written through this surface. It is the difference between what we shipped and what an operator has since corrected. | [optional] 
**Protocol** | Pointer to **string** | Protocol records HOW it was scored — provider-reported, agentic, third-party-leaderboard — so a provider card is never read as a measurement. | [optional] 
**Provider** | Pointer to **string** | Provider is who the claim belongs to — the lab or leaderboard whose number this is. | [optional] 
**Score** | Pointer to **float64** | Score is the reported aggregate, as a percentage. | [optional] 
**Source** | Pointer to **string** | Source is the citation the row was read from. | [optional] 

## Methods

### NewClaimRow

`func NewClaimRow() *ClaimRow`

NewClaimRow instantiates a new ClaimRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimRowWithDefaults

`func NewClaimRowWithDefaults() *ClaimRow`

NewClaimRowWithDefaults instantiates a new ClaimRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *ClaimRow) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *ClaimRow) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *ClaimRow) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *ClaimRow) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBenchmark

`func (o *ClaimRow) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *ClaimRow) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *ClaimRow) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *ClaimRow) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetBy

`func (o *ClaimRow) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *ClaimRow) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *ClaimRow) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *ClaimRow) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetModel

`func (o *ClaimRow) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClaimRow) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClaimRow) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ClaimRow) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrigin

`func (o *ClaimRow) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ClaimRow) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ClaimRow) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ClaimRow) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetProtocol

`func (o *ClaimRow) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ClaimRow) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ClaimRow) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ClaimRow) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProvider

`func (o *ClaimRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClaimRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClaimRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ClaimRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetScore

`func (o *ClaimRow) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ClaimRow) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ClaimRow) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *ClaimRow) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSource

`func (o *ClaimRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ClaimRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ClaimRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ClaimRow) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


