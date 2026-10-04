# GraphGraphDeriveOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is the knowledge instant it was taken at, RFC 3339, the same way. | [optional] 
**AsOf** | Pointer to **string** | AsOf is the instant the graph was read at, RFC 3339: the one asked for, or the server&#39;s clock when none was. | [optional] 
**Bound** | Pointer to **int64** | Bound is the most tuples one evaluation derives. | [optional] 
**Cut** | Pointer to **string** | Cut names the bound that stopped it: &#x60;tuples&#x60;, at Bound tuples derived, or &#x60;steps&#x60;, at Steps rows examined. Absent exactly when Truncated is false. | [optional] 
**Derived** | Pointer to [**[]GraphGraphConclusion**](GraphGraphConclusion.md) | Derived is every tuple of the predicates asked for that the rules derive, ordered by predicate, subject and object. A tuple an assertion already states is asserted, not derived, and is not here. | [optional] 
**Filed** | Pointer to [**GraphGraphAssertOut**](GraphGraphAssertOut.md) | Filed is what filing Derived did, member by member as an assertion batch reports it. Absent unless file was asked. | [optional] 
**Gained** | Pointer to [**[]GraphGraphAtom**](GraphGraphAtom.md) | Gained is each that holds only without them — a negation those assertions had been refuting. Empty when Without is absent. | [optional] 
**Lemmas** | Pointer to [**[]GraphGraphConclusion**](GraphGraphConclusion.md) | Lemmas is every derived tuple of any other predicate a proof in Derived cites, in the same order, so every support that is not an assertion can be followed to a proof. | [optional] 
**Lost** | Pointer to [**[]GraphGraphAtom**](GraphGraphAtom.md) | Lost is each tuple of the predicates asked for, asserted or derived, that holds with everything and not without the assertions Without names. Empty when Without is absent. | [optional] 
**Steps** | Pointer to **int64** | Steps is the most rows one evaluation&#39;s joins examine. | [optional] 
**Truncated** | Pointer to **bool** | Truncated says a bound stopped the evaluation before its fixpoint. Every tuple listed still holds; some that hold are missing. | [optional] 

## Methods

### NewGraphGraphDeriveOut

`func NewGraphGraphDeriveOut() *GraphGraphDeriveOut`

NewGraphGraphDeriveOut instantiates a new GraphGraphDeriveOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphDeriveOutWithDefaults

`func NewGraphGraphDeriveOutWithDefaults() *GraphGraphDeriveOut`

NewGraphGraphDeriveOutWithDefaults instantiates a new GraphGraphDeriveOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphDeriveOut) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphDeriveOut) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphDeriveOut) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphDeriveOut) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphDeriveOut) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphDeriveOut) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphDeriveOut) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphDeriveOut) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetBound

`func (o *GraphGraphDeriveOut) GetBound() int64`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *GraphGraphDeriveOut) GetBoundOk() (*int64, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *GraphGraphDeriveOut) SetBound(v int64)`

SetBound sets Bound field to given value.

### HasBound

`func (o *GraphGraphDeriveOut) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetCut

`func (o *GraphGraphDeriveOut) GetCut() string`

GetCut returns the Cut field if non-nil, zero value otherwise.

### GetCutOk

`func (o *GraphGraphDeriveOut) GetCutOk() (*string, bool)`

GetCutOk returns a tuple with the Cut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCut

`func (o *GraphGraphDeriveOut) SetCut(v string)`

SetCut sets Cut field to given value.

### HasCut

`func (o *GraphGraphDeriveOut) HasCut() bool`

HasCut returns a boolean if a field has been set.

### GetDerived

`func (o *GraphGraphDeriveOut) GetDerived() []GraphGraphConclusion`

GetDerived returns the Derived field if non-nil, zero value otherwise.

### GetDerivedOk

`func (o *GraphGraphDeriveOut) GetDerivedOk() (*[]GraphGraphConclusion, bool)`

GetDerivedOk returns a tuple with the Derived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerived

`func (o *GraphGraphDeriveOut) SetDerived(v []GraphGraphConclusion)`

SetDerived sets Derived field to given value.

### HasDerived

`func (o *GraphGraphDeriveOut) HasDerived() bool`

HasDerived returns a boolean if a field has been set.

### GetFiled

`func (o *GraphGraphDeriveOut) GetFiled() GraphGraphAssertOut`

GetFiled returns the Filed field if non-nil, zero value otherwise.

### GetFiledOk

`func (o *GraphGraphDeriveOut) GetFiledOk() (*GraphGraphAssertOut, bool)`

GetFiledOk returns a tuple with the Filed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiled

`func (o *GraphGraphDeriveOut) SetFiled(v GraphGraphAssertOut)`

SetFiled sets Filed field to given value.

### HasFiled

`func (o *GraphGraphDeriveOut) HasFiled() bool`

HasFiled returns a boolean if a field has been set.

### GetGained

`func (o *GraphGraphDeriveOut) GetGained() []GraphGraphAtom`

GetGained returns the Gained field if non-nil, zero value otherwise.

### GetGainedOk

`func (o *GraphGraphDeriveOut) GetGainedOk() (*[]GraphGraphAtom, bool)`

GetGainedOk returns a tuple with the Gained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGained

`func (o *GraphGraphDeriveOut) SetGained(v []GraphGraphAtom)`

SetGained sets Gained field to given value.

### HasGained

`func (o *GraphGraphDeriveOut) HasGained() bool`

HasGained returns a boolean if a field has been set.

### GetLemmas

`func (o *GraphGraphDeriveOut) GetLemmas() []GraphGraphConclusion`

GetLemmas returns the Lemmas field if non-nil, zero value otherwise.

### GetLemmasOk

`func (o *GraphGraphDeriveOut) GetLemmasOk() (*[]GraphGraphConclusion, bool)`

GetLemmasOk returns a tuple with the Lemmas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLemmas

`func (o *GraphGraphDeriveOut) SetLemmas(v []GraphGraphConclusion)`

SetLemmas sets Lemmas field to given value.

### HasLemmas

`func (o *GraphGraphDeriveOut) HasLemmas() bool`

HasLemmas returns a boolean if a field has been set.

### GetLost

`func (o *GraphGraphDeriveOut) GetLost() []GraphGraphAtom`

GetLost returns the Lost field if non-nil, zero value otherwise.

### GetLostOk

`func (o *GraphGraphDeriveOut) GetLostOk() (*[]GraphGraphAtom, bool)`

GetLostOk returns a tuple with the Lost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLost

`func (o *GraphGraphDeriveOut) SetLost(v []GraphGraphAtom)`

SetLost sets Lost field to given value.

### HasLost

`func (o *GraphGraphDeriveOut) HasLost() bool`

HasLost returns a boolean if a field has been set.

### GetSteps

`func (o *GraphGraphDeriveOut) GetSteps() int64`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *GraphGraphDeriveOut) GetStepsOk() (*int64, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *GraphGraphDeriveOut) SetSteps(v int64)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *GraphGraphDeriveOut) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTruncated

`func (o *GraphGraphDeriveOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GraphGraphDeriveOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GraphGraphDeriveOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GraphGraphDeriveOut) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


