# BenchmarkClaimRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **time.Time** | At is when a stored row was recorded. Zero for a seed row. | [optional] 
**Benchmark** | Pointer to **string** | Benchmark is the canonical test id the claim is about, from /catalog. | [optional] 
**By** | Pointer to **string** | By is the verified user who recorded a stored row. Empty for a seed row. | [optional] 
**Model** | Pointer to **string** | Model is the system the score is claimed for. | [optional] 
**Org** | Pointer to **string** | Org is the org that made the claim. \&quot;admin\&quot; is the platform&#39;s own curated reading; any other org is that org&#39;s claim, not the platform&#39;s. Empty on an unattributed row. | [optional] 
**Origin** | Pointer to **string** | Origin is \&quot;seed\&quot; for a compiled row, \&quot;stored\&quot; for one a verified caller wrote through this surface, and \&quot;unattributed\&quot; for a stored row no verified caller vouches for, which never reaches the leaderboard. | [optional] 
**Protocol** | Pointer to **string** | Protocol records HOW it was scored — provider-reported, agentic, third-party-leaderboard — so a provider card is never read as a measurement. | [optional] 
**Provider** | Pointer to **string** | Provider is who the claim belongs to — the lab or leaderboard whose number this is. | [optional] 
**Score** | Pointer to **float64** | Score is the reported aggregate, as a percentage. | [optional] 
**Source** | Pointer to **string** | Source is the citation the row was read from. | [optional] 
**Visibility** | Pointer to **string** | Visibility is \&quot;public\&quot;, readable by anyone, or \&quot;private\&quot;, readable by Org&#39;s own members only. Only the curator&#39;s public claims reach the leaderboard. | [optional] 

## Methods

### NewBenchmarkClaimRow

`func NewBenchmarkClaimRow() *BenchmarkClaimRow`

NewBenchmarkClaimRow instantiates a new BenchmarkClaimRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkClaimRowWithDefaults

`func NewBenchmarkClaimRowWithDefaults() *BenchmarkClaimRow`

NewBenchmarkClaimRowWithDefaults instantiates a new BenchmarkClaimRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *BenchmarkClaimRow) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *BenchmarkClaimRow) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *BenchmarkClaimRow) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *BenchmarkClaimRow) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBenchmark

`func (o *BenchmarkClaimRow) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *BenchmarkClaimRow) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *BenchmarkClaimRow) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *BenchmarkClaimRow) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetBy

`func (o *BenchmarkClaimRow) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *BenchmarkClaimRow) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *BenchmarkClaimRow) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *BenchmarkClaimRow) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetModel

`func (o *BenchmarkClaimRow) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BenchmarkClaimRow) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BenchmarkClaimRow) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BenchmarkClaimRow) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrg

`func (o *BenchmarkClaimRow) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BenchmarkClaimRow) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BenchmarkClaimRow) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BenchmarkClaimRow) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOrigin

`func (o *BenchmarkClaimRow) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *BenchmarkClaimRow) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *BenchmarkClaimRow) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *BenchmarkClaimRow) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetProtocol

`func (o *BenchmarkClaimRow) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BenchmarkClaimRow) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BenchmarkClaimRow) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BenchmarkClaimRow) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProvider

`func (o *BenchmarkClaimRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BenchmarkClaimRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BenchmarkClaimRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *BenchmarkClaimRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetScore

`func (o *BenchmarkClaimRow) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BenchmarkClaimRow) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BenchmarkClaimRow) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *BenchmarkClaimRow) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSource

`func (o *BenchmarkClaimRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BenchmarkClaimRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BenchmarkClaimRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BenchmarkClaimRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVisibility

`func (o *BenchmarkClaimRow) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BenchmarkClaimRow) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BenchmarkClaimRow) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BenchmarkClaimRow) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


