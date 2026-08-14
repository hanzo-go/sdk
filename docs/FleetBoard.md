# FleetBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | Pointer to [**[]FleetUnit**](FleetUnit.md) | Units is the union across sources — agent run-targets, BYO workers, BYO clusters and Visor machines — each row naming the source it came from. | [optional] 

## Methods

### NewFleetBoard

`func NewFleetBoard() *FleetBoard`

NewFleetBoard instantiates a new FleetBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetBoardWithDefaults

`func NewFleetBoardWithDefaults() *FleetBoard`

NewFleetBoardWithDefaults instantiates a new FleetBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *FleetBoard) GetUnits() []FleetUnit`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *FleetBoard) GetUnitsOk() (*[]FleetUnit, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *FleetBoard) SetUnits(v []FleetUnit)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *FleetBoard) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


