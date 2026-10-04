# BotBotRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **string** | RunID is the run&#39;s id in the bot runtime, and the node id its live VNC session is registered under. | [optional] 
**SessionUrl** | Pointer to **string** | SessionURL is the live session the hanzo.app /vnc panel embeds to watch or attach to this run. Derived here from the run id, never sent by the runtime. | [optional] 
**StartedAt** | Pointer to **string** | StartedAt is when the run began, RFC 3339, as the runtime stamped it. | [optional] 
**Status** | Pointer to **string** | Status is the run&#39;s state as the runtime reports it; \&quot;running\&quot; when the runtime names none of its own. | [optional] 
**Surface** | Pointer to **string** | Surface is what the bot drives: the desktop or terminal sandbox it runs in. | [optional] 
**Task** | Pointer to **string** | Task is the instruction the bot is executing. | [optional] 

## Methods

### NewBotBotRun

`func NewBotBotRun() *BotBotRun`

NewBotBotRun instantiates a new BotBotRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotBotRunWithDefaults

`func NewBotBotRunWithDefaults() *BotBotRun`

NewBotBotRunWithDefaults instantiates a new BotBotRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *BotBotRun) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *BotBotRun) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *BotBotRun) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *BotBotRun) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetSessionUrl

`func (o *BotBotRun) GetSessionUrl() string`

GetSessionUrl returns the SessionUrl field if non-nil, zero value otherwise.

### GetSessionUrlOk

`func (o *BotBotRun) GetSessionUrlOk() (*string, bool)`

GetSessionUrlOk returns a tuple with the SessionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionUrl

`func (o *BotBotRun) SetSessionUrl(v string)`

SetSessionUrl sets SessionUrl field to given value.

### HasSessionUrl

`func (o *BotBotRun) HasSessionUrl() bool`

HasSessionUrl returns a boolean if a field has been set.

### GetStartedAt

`func (o *BotBotRun) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BotBotRun) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BotBotRun) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BotBotRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *BotBotRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BotBotRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BotBotRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BotBotRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSurface

`func (o *BotBotRun) GetSurface() string`

GetSurface returns the Surface field if non-nil, zero value otherwise.

### GetSurfaceOk

`func (o *BotBotRun) GetSurfaceOk() (*string, bool)`

GetSurfaceOk returns a tuple with the Surface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurface

`func (o *BotBotRun) SetSurface(v string)`

SetSurface sets Surface field to given value.

### HasSurface

`func (o *BotBotRun) HasSurface() bool`

HasSurface returns a boolean if a field has been set.

### GetTask

`func (o *BotBotRun) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *BotBotRun) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *BotBotRun) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *BotBotRun) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


