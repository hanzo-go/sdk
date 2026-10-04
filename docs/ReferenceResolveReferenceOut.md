# ReferenceResolveReferenceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**[]ReferenceReferenceAnswer**](ReferenceReferenceAnswer.md) | Answers is one entry per (set, key) consulted. | [optional] 
**Consulted** | Pointer to [**[]ReferenceReferenceVersion**](ReferenceReferenceVersion.md) | Consulted names the version of every set that took part, so a decision can record precisely what it leaned on. Record this with the decision: it is what makes the decision reproducible a year later. | [optional] 
**Refused** | Pointer to **[]string** | Refused names the consulted sets that could not answer at all. A key that missed in one of these is UNKNOWN, not clean. | [optional] 
**Stale** | Pointer to **[]string** | Stale names the consulted sets past their freshness bound. Staleness is itself a risk signal — a decision taken against a three-week-old list is a weaker decision, and this is how it knows. | [optional] 

## Methods

### NewReferenceResolveReferenceOut

`func NewReferenceResolveReferenceOut() *ReferenceResolveReferenceOut`

NewReferenceResolveReferenceOut instantiates a new ReferenceResolveReferenceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceResolveReferenceOutWithDefaults

`func NewReferenceResolveReferenceOutWithDefaults() *ReferenceResolveReferenceOut`

NewReferenceResolveReferenceOutWithDefaults instantiates a new ReferenceResolveReferenceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *ReferenceResolveReferenceOut) GetAnswers() []ReferenceReferenceAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *ReferenceResolveReferenceOut) GetAnswersOk() (*[]ReferenceReferenceAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *ReferenceResolveReferenceOut) SetAnswers(v []ReferenceReferenceAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *ReferenceResolveReferenceOut) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetConsulted

`func (o *ReferenceResolveReferenceOut) GetConsulted() []ReferenceReferenceVersion`

GetConsulted returns the Consulted field if non-nil, zero value otherwise.

### GetConsultedOk

`func (o *ReferenceResolveReferenceOut) GetConsultedOk() (*[]ReferenceReferenceVersion, bool)`

GetConsultedOk returns a tuple with the Consulted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsulted

`func (o *ReferenceResolveReferenceOut) SetConsulted(v []ReferenceReferenceVersion)`

SetConsulted sets Consulted field to given value.

### HasConsulted

`func (o *ReferenceResolveReferenceOut) HasConsulted() bool`

HasConsulted returns a boolean if a field has been set.

### GetRefused

`func (o *ReferenceResolveReferenceOut) GetRefused() []string`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *ReferenceResolveReferenceOut) GetRefusedOk() (*[]string, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *ReferenceResolveReferenceOut) SetRefused(v []string)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *ReferenceResolveReferenceOut) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetStale

`func (o *ReferenceResolveReferenceOut) GetStale() []string`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ReferenceResolveReferenceOut) GetStaleOk() (*[]string, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ReferenceResolveReferenceOut) SetStale(v []string)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ReferenceResolveReferenceOut) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


