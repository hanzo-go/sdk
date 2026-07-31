# CloudFleetBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | Pointer to [**[]CloudFleetUnit**](CloudFleetUnit.md) | Units is the union across sources — agent run-targets, BYO workers, BYO clusters and Visor machines — each row naming the source it came from. | [optional] 

## Methods

### NewCloudFleetBoard

`func NewCloudFleetBoard() *CloudFleetBoard`

NewCloudFleetBoard instantiates a new CloudFleetBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFleetBoardWithDefaults

`func NewCloudFleetBoardWithDefaults() *CloudFleetBoard`

NewCloudFleetBoardWithDefaults instantiates a new CloudFleetBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *CloudFleetBoard) GetUnits() []CloudFleetUnit`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CloudFleetBoard) GetUnitsOk() (*[]CloudFleetUnit, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CloudFleetBoard) SetUnits(v []CloudFleetUnit)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *CloudFleetBoard) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


