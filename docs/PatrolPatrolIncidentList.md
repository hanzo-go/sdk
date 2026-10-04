# PatrolPatrolIncidentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PatrolPatrolIncident**](PatrolPatrolIncident.md) | Data is the matching incidents, without their trails. | [optional] 

## Methods

### NewPatrolPatrolIncidentList

`func NewPatrolPatrolIncidentList() *PatrolPatrolIncidentList`

NewPatrolPatrolIncidentList instantiates a new PatrolPatrolIncidentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolIncidentListWithDefaults

`func NewPatrolPatrolIncidentListWithDefaults() *PatrolPatrolIncidentList`

NewPatrolPatrolIncidentListWithDefaults instantiates a new PatrolPatrolIncidentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PatrolPatrolIncidentList) GetData() []PatrolPatrolIncident`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatrolPatrolIncidentList) GetDataOk() (*[]PatrolPatrolIncident, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatrolPatrolIncidentList) SetData(v []PatrolPatrolIncident)`

SetData sets Data field to given value.

### HasData

`func (o *PatrolPatrolIncidentList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


