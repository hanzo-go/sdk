# TrafficGlobe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Points** | Pointer to [**[]TrafficPoint**](TrafficPoint.md) |  | [optional] 
**Totals** | Pointer to [**TrafficTotals**](TrafficTotals.md) |  | [optional] 
**Window** | Pointer to [**TrafficWindow**](TrafficWindow.md) |  | [optional] 

## Methods

### NewTrafficGlobe

`func NewTrafficGlobe() *TrafficGlobe`

NewTrafficGlobe instantiates a new TrafficGlobe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficGlobeWithDefaults

`func NewTrafficGlobeWithDefaults() *TrafficGlobe`

NewTrafficGlobeWithDefaults instantiates a new TrafficGlobe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoints

`func (o *TrafficGlobe) GetPoints() []TrafficPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *TrafficGlobe) GetPointsOk() (*[]TrafficPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *TrafficGlobe) SetPoints(v []TrafficPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *TrafficGlobe) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetTotals

`func (o *TrafficGlobe) GetTotals() TrafficTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *TrafficGlobe) GetTotalsOk() (*TrafficTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *TrafficGlobe) SetTotals(v TrafficTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *TrafficGlobe) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetWindow

`func (o *TrafficGlobe) GetWindow() TrafficWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *TrafficGlobe) GetWindowOk() (*TrafficWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *TrafficGlobe) SetWindow(v TrafficWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *TrafficGlobe) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


