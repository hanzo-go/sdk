# PatrolPatrolTour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | Date is the day the round belongs to. | [optional] 
**End** | Pointer to **string** | End is when it finished. | [optional] 
**Name** | Pointer to **string** | Name is the tour&#39;s document name. | [optional] 
**Points** | Pointer to [**[]PatrolPatrolCheckpoint**](PatrolPatrolCheckpoint.md) | Points is the round&#39;s checkpoints in store order. | [optional] 
**Start** | Pointer to **string** | Start is when the round began. | [optional] 
**Unit** | Pointer to **string** | Unit is the call sign running the round. | [optional] 

## Methods

### NewPatrolPatrolTour

`func NewPatrolPatrolTour() *PatrolPatrolTour`

NewPatrolPatrolTour instantiates a new PatrolPatrolTour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolTourWithDefaults

`func NewPatrolPatrolTourWithDefaults() *PatrolPatrolTour`

NewPatrolPatrolTourWithDefaults instantiates a new PatrolPatrolTour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PatrolPatrolTour) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PatrolPatrolTour) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PatrolPatrolTour) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PatrolPatrolTour) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEnd

`func (o *PatrolPatrolTour) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PatrolPatrolTour) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PatrolPatrolTour) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *PatrolPatrolTour) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolTour) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolTour) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolTour) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolTour) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoints

`func (o *PatrolPatrolTour) GetPoints() []PatrolPatrolCheckpoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *PatrolPatrolTour) GetPointsOk() (*[]PatrolPatrolCheckpoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *PatrolPatrolTour) SetPoints(v []PatrolPatrolCheckpoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *PatrolPatrolTour) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetStart

`func (o *PatrolPatrolTour) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PatrolPatrolTour) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PatrolPatrolTour) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *PatrolPatrolTour) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUnit

`func (o *PatrolPatrolTour) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PatrolPatrolTour) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PatrolPatrolTour) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PatrolPatrolTour) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


