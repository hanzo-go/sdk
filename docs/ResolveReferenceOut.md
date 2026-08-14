# ResolveReferenceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**[]ReferenceAnswer**](ReferenceAnswer.md) | Answers is one entry per (set, key) consulted. | [optional] 
**Consulted** | Pointer to [**[]ReferenceVersion**](ReferenceVersion.md) | Consulted names the version of every set that took part, so a decision can record precisely what it leaned on. Record this with the decision: it is what makes the decision reproducible a year later. | [optional] 
**Refused** | Pointer to **[]string** | Refused names the consulted sets that could not answer at all. A key that missed in one of these is UNKNOWN, not clean. | [optional] 
**Stale** | Pointer to **[]string** | Stale names the consulted sets past their freshness bound. Staleness is itself a risk signal — a decision taken against a three-week-old list is a weaker decision, and this is how it knows. | [optional] 

## Methods

### NewResolveReferenceOut

`func NewResolveReferenceOut() *ResolveReferenceOut`

NewResolveReferenceOut instantiates a new ResolveReferenceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolveReferenceOutWithDefaults

`func NewResolveReferenceOutWithDefaults() *ResolveReferenceOut`

NewResolveReferenceOutWithDefaults instantiates a new ResolveReferenceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *ResolveReferenceOut) GetAnswers() []ReferenceAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *ResolveReferenceOut) GetAnswersOk() (*[]ReferenceAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *ResolveReferenceOut) SetAnswers(v []ReferenceAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *ResolveReferenceOut) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetConsulted

`func (o *ResolveReferenceOut) GetConsulted() []ReferenceVersion`

GetConsulted returns the Consulted field if non-nil, zero value otherwise.

### GetConsultedOk

`func (o *ResolveReferenceOut) GetConsultedOk() (*[]ReferenceVersion, bool)`

GetConsultedOk returns a tuple with the Consulted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsulted

`func (o *ResolveReferenceOut) SetConsulted(v []ReferenceVersion)`

SetConsulted sets Consulted field to given value.

### HasConsulted

`func (o *ResolveReferenceOut) HasConsulted() bool`

HasConsulted returns a boolean if a field has been set.

### GetRefused

`func (o *ResolveReferenceOut) GetRefused() []string`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *ResolveReferenceOut) GetRefusedOk() (*[]string, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *ResolveReferenceOut) SetRefused(v []string)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *ResolveReferenceOut) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetStale

`func (o *ResolveReferenceOut) GetStale() []string`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ResolveReferenceOut) GetStaleOk() (*[]string, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ResolveReferenceOut) SetStale(v []string)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ResolveReferenceOut) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


