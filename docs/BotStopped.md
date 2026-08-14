# BotStopped

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **string** | RunID is the run that was stopped. | [optional] 
**Status** | Pointer to **string** | Status is the run&#39;s terminal state: \&quot;stopped\&quot;. | [optional] 

## Methods

### NewBotStopped

`func NewBotStopped() *BotStopped`

NewBotStopped instantiates a new BotStopped object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotStoppedWithDefaults

`func NewBotStoppedWithDefaults() *BotStopped`

NewBotStoppedWithDefaults instantiates a new BotStopped object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *BotStopped) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *BotStopped) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *BotStopped) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *BotStopped) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetStatus

`func (o *BotStopped) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BotStopped) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BotStopped) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BotStopped) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


