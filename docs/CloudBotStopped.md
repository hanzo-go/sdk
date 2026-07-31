# CloudBotStopped

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **string** | RunID is the run that was stopped. | [optional] 
**Status** | Pointer to **string** | Status is the run&#39;s terminal state: \&quot;stopped\&quot;. | [optional] 

## Methods

### NewCloudBotStopped

`func NewCloudBotStopped() *CloudBotStopped`

NewCloudBotStopped instantiates a new CloudBotStopped object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBotStoppedWithDefaults

`func NewCloudBotStoppedWithDefaults() *CloudBotStopped`

NewCloudBotStoppedWithDefaults instantiates a new CloudBotStopped object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *CloudBotStopped) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *CloudBotStopped) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *CloudBotStopped) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *CloudBotStopped) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetStatus

`func (o *CloudBotStopped) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudBotStopped) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudBotStopped) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudBotStopped) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


