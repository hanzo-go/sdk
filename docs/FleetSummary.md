# FleetSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByDrift** | Pointer to [**DriftTally**](DriftTally.md) | ByDrift counts those rows green, yellow and red. | [optional] 
**Total** | Pointer to **int64** | Total is how many rows the board returned, after filtering. | [optional] 

## Methods

### NewFleetSummary

`func NewFleetSummary() *FleetSummary`

NewFleetSummary instantiates a new FleetSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetSummaryWithDefaults

`func NewFleetSummaryWithDefaults() *FleetSummary`

NewFleetSummaryWithDefaults instantiates a new FleetSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByDrift

`func (o *FleetSummary) GetByDrift() DriftTally`

GetByDrift returns the ByDrift field if non-nil, zero value otherwise.

### GetByDriftOk

`func (o *FleetSummary) GetByDriftOk() (*DriftTally, bool)`

GetByDriftOk returns a tuple with the ByDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByDrift

`func (o *FleetSummary) SetByDrift(v DriftTally)`

SetByDrift sets ByDrift field to given value.

### HasByDrift

`func (o *FleetSummary) HasByDrift() bool`

HasByDrift returns a boolean if a field has been set.

### GetTotal

`func (o *FleetSummary) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FleetSummary) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FleetSummary) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FleetSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


