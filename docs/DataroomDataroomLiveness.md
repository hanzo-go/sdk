# DataroomDataroomLiveness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Service names the subsystem answering, so a probe response is attributable when several are collected together. | [optional] 
**Status** | Pointer to **string** | Status is &#x60;ok&#x60;. This probe has no degraded answer by design: it reports process liveness and nothing that could be false while the process serves. | [optional] 

## Methods

### NewDataroomDataroomLiveness

`func NewDataroomDataroomLiveness() *DataroomDataroomLiveness`

NewDataroomDataroomLiveness instantiates a new DataroomDataroomLiveness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomLivenessWithDefaults

`func NewDataroomDataroomLivenessWithDefaults() *DataroomDataroomLiveness`

NewDataroomDataroomLivenessWithDefaults instantiates a new DataroomDataroomLiveness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *DataroomDataroomLiveness) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DataroomDataroomLiveness) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DataroomDataroomLiveness) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DataroomDataroomLiveness) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *DataroomDataroomLiveness) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataroomDataroomLiveness) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataroomDataroomLiveness) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataroomDataroomLiveness) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


