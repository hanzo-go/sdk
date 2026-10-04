# ReferenceReferenceAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | Pointer to **string** | Age is how old that is, as a duration. | [optional] 
**AsOf** | Pointer to **string** | AsOf is when the oldest contributing publisher was current, RFC 3339. | [optional] 
**From** | Pointer to **string** | From is override or baseline — which plane answered. | [optional] 
**Hit** | Pointer to **bool** | Hit is whether the key is a member. It is meaningful ONLY when Refusal is empty: false with a refusal means the set could not be consulted, which is not the same as the key being clean. | [optional] 
**Key** | Pointer to **string** | Key is the key as asked. | [optional] 
**Matched** | Pointer to **string** | Matched is the member that covered the key, which for a domain or a network is the enclosing entry rather than the key itself. | [optional] 
**Refusal** | Pointer to **string** | Refusal is why the set could not be consulted, when it could not: never loaded, held elsewhere, or a source we hold no licence for. Non-empty means Hit must not be read as an answer. | [optional] 
**Score** | Pointer to **float64** | Score is the published risk weight where the source expresses one. | [optional] 
**Set** | Pointer to **string** | Set is the set consulted. | [optional] 
**Stale** | Pointer to **bool** | Stale is whether the set is past its freshness bound. A stale set still answers — yesterday&#39;s list beats none — and this is how a decision knows it leaned on one. | [optional] 
**Value** | Pointer to **map[string]string** | Value is what the publisher says about the member — class, operator, scheme, region. | [optional] 
**Verdict** | Pointer to **string** | Verdict is the tenant&#39;s own allow or deny, present only for an override. The baseline never carries one: it states facts and leaves the decision to the caller&#39;s policy. | [optional] 
**Version** | Pointer to **string** | Version is the exact baseline version consulted, composed of each contributing publisher and its content digest. It is what makes a decision reproducible: an auditor takes this string and knows precisely what was consulted. | [optional] 

## Methods

### NewReferenceReferenceAnswer

`func NewReferenceReferenceAnswer() *ReferenceReferenceAnswer`

NewReferenceReferenceAnswer instantiates a new ReferenceReferenceAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceReferenceAnswerWithDefaults

`func NewReferenceReferenceAnswerWithDefaults() *ReferenceReferenceAnswer`

NewReferenceReferenceAnswerWithDefaults instantiates a new ReferenceReferenceAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *ReferenceReferenceAnswer) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ReferenceReferenceAnswer) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ReferenceReferenceAnswer) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *ReferenceReferenceAnswer) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAsOf

`func (o *ReferenceReferenceAnswer) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ReferenceReferenceAnswer) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ReferenceReferenceAnswer) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *ReferenceReferenceAnswer) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetFrom

`func (o *ReferenceReferenceAnswer) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReferenceReferenceAnswer) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReferenceReferenceAnswer) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ReferenceReferenceAnswer) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHit

`func (o *ReferenceReferenceAnswer) GetHit() bool`

GetHit returns the Hit field if non-nil, zero value otherwise.

### GetHitOk

`func (o *ReferenceReferenceAnswer) GetHitOk() (*bool, bool)`

GetHitOk returns a tuple with the Hit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHit

`func (o *ReferenceReferenceAnswer) SetHit(v bool)`

SetHit sets Hit field to given value.

### HasHit

`func (o *ReferenceReferenceAnswer) HasHit() bool`

HasHit returns a boolean if a field has been set.

### GetKey

`func (o *ReferenceReferenceAnswer) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ReferenceReferenceAnswer) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ReferenceReferenceAnswer) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ReferenceReferenceAnswer) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMatched

`func (o *ReferenceReferenceAnswer) GetMatched() string`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *ReferenceReferenceAnswer) GetMatchedOk() (*string, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *ReferenceReferenceAnswer) SetMatched(v string)`

SetMatched sets Matched field to given value.

### HasMatched

`func (o *ReferenceReferenceAnswer) HasMatched() bool`

HasMatched returns a boolean if a field has been set.

### GetRefusal

`func (o *ReferenceReferenceAnswer) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *ReferenceReferenceAnswer) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *ReferenceReferenceAnswer) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *ReferenceReferenceAnswer) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetScore

`func (o *ReferenceReferenceAnswer) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ReferenceReferenceAnswer) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ReferenceReferenceAnswer) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *ReferenceReferenceAnswer) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSet

`func (o *ReferenceReferenceAnswer) GetSet() string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *ReferenceReferenceAnswer) GetSetOk() (*string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *ReferenceReferenceAnswer) SetSet(v string)`

SetSet sets Set field to given value.

### HasSet

`func (o *ReferenceReferenceAnswer) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetStale

`func (o *ReferenceReferenceAnswer) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ReferenceReferenceAnswer) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ReferenceReferenceAnswer) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ReferenceReferenceAnswer) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetValue

`func (o *ReferenceReferenceAnswer) GetValue() map[string]string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReferenceReferenceAnswer) GetValueOk() (*map[string]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReferenceReferenceAnswer) SetValue(v map[string]string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ReferenceReferenceAnswer) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVerdict

`func (o *ReferenceReferenceAnswer) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ReferenceReferenceAnswer) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ReferenceReferenceAnswer) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *ReferenceReferenceAnswer) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### GetVersion

`func (o *ReferenceReferenceAnswer) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReferenceReferenceAnswer) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReferenceReferenceAnswer) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReferenceReferenceAnswer) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


