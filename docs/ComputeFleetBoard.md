# ComputeFleetBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | Pointer to [**[]ComputeFleetUnit**](ComputeFleetUnit.md) | Units is the union across sources — agent run-targets, BYO workers, BYO clusters and Visor machines — each row naming the source it came from. | [optional] 

## Methods

### NewComputeFleetBoard

`func NewComputeFleetBoard() *ComputeFleetBoard`

NewComputeFleetBoard instantiates a new ComputeFleetBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeFleetBoardWithDefaults

`func NewComputeFleetBoardWithDefaults() *ComputeFleetBoard`

NewComputeFleetBoardWithDefaults instantiates a new ComputeFleetBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *ComputeFleetBoard) GetUnits() []ComputeFleetUnit`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ComputeFleetBoard) GetUnitsOk() (*[]ComputeFleetUnit, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ComputeFleetBoard) SetUnits(v []ComputeFleetUnit)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ComputeFleetBoard) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


