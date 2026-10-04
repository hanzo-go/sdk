# PatrolPatrolNearest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eta** | Pointer to **int64** | Eta is the minutes the drive would take at the planning speed. | [optional] 
**Km** | Pointer to **float64** | Km is the great-circle distance to the site in kilometres. | [optional] 
**Unit** | Pointer to **string** | Unit is the unit&#39;s call sign. | [optional] 

## Methods

### NewPatrolPatrolNearest

`func NewPatrolPatrolNearest() *PatrolPatrolNearest`

NewPatrolPatrolNearest instantiates a new PatrolPatrolNearest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolNearestWithDefaults

`func NewPatrolPatrolNearestWithDefaults() *PatrolPatrolNearest`

NewPatrolPatrolNearestWithDefaults instantiates a new PatrolPatrolNearest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEta

`func (o *PatrolPatrolNearest) GetEta() int64`

GetEta returns the Eta field if non-nil, zero value otherwise.

### GetEtaOk

`func (o *PatrolPatrolNearest) GetEtaOk() (*int64, bool)`

GetEtaOk returns a tuple with the Eta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEta

`func (o *PatrolPatrolNearest) SetEta(v int64)`

SetEta sets Eta field to given value.

### HasEta

`func (o *PatrolPatrolNearest) HasEta() bool`

HasEta returns a boolean if a field has been set.

### GetKm

`func (o *PatrolPatrolNearest) GetKm() float64`

GetKm returns the Km field if non-nil, zero value otherwise.

### GetKmOk

`func (o *PatrolPatrolNearest) GetKmOk() (*float64, bool)`

GetKmOk returns a tuple with the Km field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKm

`func (o *PatrolPatrolNearest) SetKm(v float64)`

SetKm sets Km field to given value.

### HasKm

`func (o *PatrolPatrolNearest) HasKm() bool`

HasKm returns a boolean if a field has been set.

### GetUnit

`func (o *PatrolPatrolNearest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PatrolPatrolNearest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PatrolPatrolNearest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PatrolPatrolNearest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


