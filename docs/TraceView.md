# TraceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyHash** | Pointer to **string** | APIKeyHash is the non-reversible credential ref (never a plaintext key), so a trace correlates to the key that drove it without the store holding a secret. | [optional] 
**DatasetItemId** | Pointer to **string** | ItemID is the example the call answered. | [optional] 
**DatasetName** | Pointer to **string** | Dataset is the set the graded example came from. | [optional] 
**EndTime** | Pointer to **string** | EndTime is when it returned. | [optional] 
**Id** | Pointer to **string** | ID is the trace&#39;s handle, the value a score points at. | [optional] 
**Input** | Pointer to **map[string]interface{}** | Input is what the model was given. | [optional] 
**LatencyMs** | Pointer to **float32** | LatencyMs is EndTime-StartTime in milliseconds, nil when the trace carries no timing (so the console renders \&quot;—\&quot;, never a fabricated 0). | [optional] 
**Model** | Pointer to **string** | Model is the model that answered. | [optional] 
**Name** | Pointer to **string** | Name is the trace&#39;s label, \&quot;eval:&lt;run&gt;\&quot; for a call a run made. | [optional] 
**Output** | Pointer to **string** | Output is what it answered. | [optional] 
**ProjectId** | Pointer to **string** | ProjectID is the sub-scope within the org the call was made under. | [optional] 
**RunName** | Pointer to **string** | RunName is the run the call belongs to. | [optional] 
**SessionId** | Pointer to **string** | SessionID groups the calls of one run. | [optional] 
**StartTime** | Pointer to **string** | StartTime is when the call began. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp is the trace&#39;s own clock, equal to StartTime for a timed call. | [optional] 

## Methods

### NewTraceView

`func NewTraceView() *TraceView`

NewTraceView instantiates a new TraceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceViewWithDefaults

`func NewTraceViewWithDefaults() *TraceView`

NewTraceViewWithDefaults instantiates a new TraceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyHash

`func (o *TraceView) GetApiKeyHash() string`

GetApiKeyHash returns the ApiKeyHash field if non-nil, zero value otherwise.

### GetApiKeyHashOk

`func (o *TraceView) GetApiKeyHashOk() (*string, bool)`

GetApiKeyHashOk returns a tuple with the ApiKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyHash

`func (o *TraceView) SetApiKeyHash(v string)`

SetApiKeyHash sets ApiKeyHash field to given value.

### HasApiKeyHash

`func (o *TraceView) HasApiKeyHash() bool`

HasApiKeyHash returns a boolean if a field has been set.

### GetDatasetItemId

`func (o *TraceView) GetDatasetItemId() string`

GetDatasetItemId returns the DatasetItemId field if non-nil, zero value otherwise.

### GetDatasetItemIdOk

`func (o *TraceView) GetDatasetItemIdOk() (*string, bool)`

GetDatasetItemIdOk returns a tuple with the DatasetItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetItemId

`func (o *TraceView) SetDatasetItemId(v string)`

SetDatasetItemId sets DatasetItemId field to given value.

### HasDatasetItemId

`func (o *TraceView) HasDatasetItemId() bool`

HasDatasetItemId returns a boolean if a field has been set.

### GetDatasetName

`func (o *TraceView) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *TraceView) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *TraceView) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *TraceView) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### GetEndTime

`func (o *TraceView) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TraceView) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TraceView) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TraceView) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *TraceView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TraceView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TraceView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TraceView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *TraceView) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TraceView) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TraceView) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *TraceView) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetLatencyMs

`func (o *TraceView) GetLatencyMs() float32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *TraceView) GetLatencyMsOk() (*float32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *TraceView) SetLatencyMs(v float32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *TraceView) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetModel

`func (o *TraceView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TraceView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TraceView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TraceView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *TraceView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TraceView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TraceView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TraceView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *TraceView) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TraceView) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TraceView) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TraceView) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetProjectId

`func (o *TraceView) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TraceView) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TraceView) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TraceView) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRunName

`func (o *TraceView) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *TraceView) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *TraceView) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *TraceView) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetSessionId

`func (o *TraceView) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *TraceView) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *TraceView) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *TraceView) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStartTime

`func (o *TraceView) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TraceView) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TraceView) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TraceView) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTimestamp

`func (o *TraceView) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TraceView) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TraceView) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TraceView) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


