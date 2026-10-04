# PlatformFleetSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByDrift** | Pointer to [**PlatformDriftTally**](PlatformDriftTally.md) | ByDrift counts those rows green, yellow and red. | [optional] 
**Total** | Pointer to **int64** | Total is how many rows the board returned, after filtering. | [optional] 

## Methods

### NewPlatformFleetSummary

`func NewPlatformFleetSummary() *PlatformFleetSummary`

NewPlatformFleetSummary instantiates a new PlatformFleetSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformFleetSummaryWithDefaults

`func NewPlatformFleetSummaryWithDefaults() *PlatformFleetSummary`

NewPlatformFleetSummaryWithDefaults instantiates a new PlatformFleetSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByDrift

`func (o *PlatformFleetSummary) GetByDrift() PlatformDriftTally`

GetByDrift returns the ByDrift field if non-nil, zero value otherwise.

### GetByDriftOk

`func (o *PlatformFleetSummary) GetByDriftOk() (*PlatformDriftTally, bool)`

GetByDriftOk returns a tuple with the ByDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByDrift

`func (o *PlatformFleetSummary) SetByDrift(v PlatformDriftTally)`

SetByDrift sets ByDrift field to given value.

### HasByDrift

`func (o *PlatformFleetSummary) HasByDrift() bool`

HasByDrift returns a boolean if a field has been set.

### GetTotal

`func (o *PlatformFleetSummary) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PlatformFleetSummary) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PlatformFleetSummary) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PlatformFleetSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


