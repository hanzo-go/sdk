# FormationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formation** | Pointer to [**Formation**](Formation.md) | Formation is the org&#39;s one incorporation record. | [optional] 
**NextStages** | Pointer to **[]string** | NextStages are the stages reachable from the formation&#39;s current stage, whether or not their guards are satisfied yet. | [optional] 

## Methods

### NewFormationView

`func NewFormationView() *FormationView`

NewFormationView instantiates a new FormationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormationViewWithDefaults

`func NewFormationViewWithDefaults() *FormationView`

NewFormationViewWithDefaults instantiates a new FormationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormation

`func (o *FormationView) GetFormation() Formation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *FormationView) GetFormationOk() (*Formation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *FormationView) SetFormation(v Formation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *FormationView) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetNextStages

`func (o *FormationView) GetNextStages() []string`

GetNextStages returns the NextStages field if non-nil, zero value otherwise.

### GetNextStagesOk

`func (o *FormationView) GetNextStagesOk() (*[]string, bool)`

GetNextStagesOk returns a tuple with the NextStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStages

`func (o *FormationView) SetNextStages(v []string)`

SetNextStages sets NextStages field to given value.

### HasNextStages

`func (o *FormationView) HasNextStages() bool`

HasNextStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


