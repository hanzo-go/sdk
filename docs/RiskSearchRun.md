# RiskSearchRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candidates** | Pointer to **int32** | Candidates is how many model shapes will be tried. | [optional] 
**Events** | Pointer to **int32** | Events is how much of the organisation&#39;s own history the run will replay. | [optional] 
**Id** | Pointer to **string** | ID addresses the run. Read the result back with it. | [optional] 

## Methods

### NewRiskSearchRun

`func NewRiskSearchRun() *RiskSearchRun`

NewRiskSearchRun instantiates a new RiskSearchRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSearchRunWithDefaults

`func NewRiskSearchRunWithDefaults() *RiskSearchRun`

NewRiskSearchRunWithDefaults instantiates a new RiskSearchRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidates

`func (o *RiskSearchRun) GetCandidates() int32`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *RiskSearchRun) GetCandidatesOk() (*int32, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *RiskSearchRun) SetCandidates(v int32)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *RiskSearchRun) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.

### GetEvents

`func (o *RiskSearchRun) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RiskSearchRun) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RiskSearchRun) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RiskSearchRun) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *RiskSearchRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskSearchRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskSearchRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskSearchRun) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


