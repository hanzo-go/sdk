# RiskRiskSearchRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candidates** | Pointer to **int64** | Candidates is how many model shapes will be tried. | [optional] 
**Events** | Pointer to **int64** | Events is how much of the organisation&#39;s own history the run will replay. | [optional] 
**Id** | Pointer to **string** | ID addresses the run. Read the result back with it. | [optional] 

## Methods

### NewRiskRiskSearchRun

`func NewRiskRiskSearchRun() *RiskRiskSearchRun`

NewRiskRiskSearchRun instantiates a new RiskRiskSearchRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskSearchRunWithDefaults

`func NewRiskRiskSearchRunWithDefaults() *RiskRiskSearchRun`

NewRiskRiskSearchRunWithDefaults instantiates a new RiskRiskSearchRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidates

`func (o *RiskRiskSearchRun) GetCandidates() int64`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *RiskRiskSearchRun) GetCandidatesOk() (*int64, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *RiskRiskSearchRun) SetCandidates(v int64)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *RiskRiskSearchRun) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.

### GetEvents

`func (o *RiskRiskSearchRun) GetEvents() int64`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RiskRiskSearchRun) GetEventsOk() (*int64, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RiskRiskSearchRun) SetEvents(v int64)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RiskRiskSearchRun) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *RiskRiskSearchRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskRiskSearchRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskRiskSearchRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskRiskSearchRun) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


