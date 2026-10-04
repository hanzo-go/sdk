# PlatformDriftBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]PlatformAppView**](PlatformAppView.md) | Apps are the service rows, ordered by org, then app, then env. | [optional] 
**Summary** | Pointer to [**PlatformFleetSummary**](PlatformFleetSummary.md) | Summary counts the board by drift severity. | [optional] 

## Methods

### NewPlatformDriftBoard

`func NewPlatformDriftBoard() *PlatformDriftBoard`

NewPlatformDriftBoard instantiates a new PlatformDriftBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDriftBoardWithDefaults

`func NewPlatformDriftBoardWithDefaults() *PlatformDriftBoard`

NewPlatformDriftBoardWithDefaults instantiates a new PlatformDriftBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *PlatformDriftBoard) GetApps() []PlatformAppView`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *PlatformDriftBoard) GetAppsOk() (*[]PlatformAppView, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *PlatformDriftBoard) SetApps(v []PlatformAppView)`

SetApps sets Apps field to given value.

### HasApps

`func (o *PlatformDriftBoard) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetSummary

`func (o *PlatformDriftBoard) GetSummary() PlatformFleetSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PlatformDriftBoard) GetSummaryOk() (*PlatformFleetSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PlatformDriftBoard) SetSummary(v PlatformFleetSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PlatformDriftBoard) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


