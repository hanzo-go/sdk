# DriftBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]AppView**](AppView.md) | Apps are the service rows, ordered by org, then app, then env. | [optional] 
**Summary** | Pointer to [**FleetSummary**](FleetSummary.md) | Summary counts the board by drift severity. | [optional] 

## Methods

### NewDriftBoard

`func NewDriftBoard() *DriftBoard`

NewDriftBoard instantiates a new DriftBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriftBoardWithDefaults

`func NewDriftBoardWithDefaults() *DriftBoard`

NewDriftBoardWithDefaults instantiates a new DriftBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *DriftBoard) GetApps() []AppView`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *DriftBoard) GetAppsOk() (*[]AppView, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *DriftBoard) SetApps(v []AppView)`

SetApps sets Apps field to given value.

### HasApps

`func (o *DriftBoard) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetSummary

`func (o *DriftBoard) GetSummary() FleetSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DriftBoard) GetSummaryOk() (*FleetSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DriftBoard) SetSummary(v FleetSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *DriftBoard) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


