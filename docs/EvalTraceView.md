# EvalTraceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyHash** | Pointer to **string** | APIKeyHash is the non-reversible credential ref (never a plaintext key), so a trace correlates to the key that drove it without the store holding a secret. | [optional] 
**DatasetItemId** | Pointer to **string** | ItemID is the example the call answered. | [optional] 
**DatasetName** | Pointer to **string** | Dataset is the set the graded example came from. | [optional] 
**EndTime** | Pointer to **string** | EndTime is when it returned. | [optional] 
**Id** | Pointer to **string** | ID is the trace&#39;s handle, the value a score points at. | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 
**LatencyMs** | Pointer to **float64** | LatencyMs is EndTime-StartTime in milliseconds, nil when the trace carries no timing (so the console renders \&quot;—\&quot;, never a fabricated 0). | [optional] 
**Model** | Pointer to **string** | Model is the model that answered. | [optional] 
**Name** | Pointer to **string** | Name is the trace&#39;s label, \&quot;eval:&lt;run&gt;\&quot; for a call a run made. | [optional] 
**Output** | Pointer to **string** | Output is what it answered. | [optional] 
**ProjectId** | Pointer to **string** | ProjectID is the sub-scope within the org the call was made under. | [optional] 
**RunName** | Pointer to **string** | RunName is the run the call belongs to. | [optional] 
**SessionId** | Pointer to **string** | SessionID groups the calls of one run. | [optional] 
**StartTime** | Pointer to **string** | StartTime is when the call began. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp is the trace&#39;s own clock, equal to StartTime for a timed call. | [optional] 

## Methods

### NewEvalTraceView

`func NewEvalTraceView() *EvalTraceView`

NewEvalTraceView instantiates a new EvalTraceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalTraceViewWithDefaults

`func NewEvalTraceViewWithDefaults() *EvalTraceView`

NewEvalTraceViewWithDefaults instantiates a new EvalTraceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyHash

`func (o *EvalTraceView) GetApiKeyHash() string`

GetApiKeyHash returns the ApiKeyHash field if non-nil, zero value otherwise.

### GetApiKeyHashOk

`func (o *EvalTraceView) GetApiKeyHashOk() (*string, bool)`

GetApiKeyHashOk returns a tuple with the ApiKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyHash

`func (o *EvalTraceView) SetApiKeyHash(v string)`

SetApiKeyHash sets ApiKeyHash field to given value.

### HasApiKeyHash

`func (o *EvalTraceView) HasApiKeyHash() bool`

HasApiKeyHash returns a boolean if a field has been set.

### GetDatasetItemId

`func (o *EvalTraceView) GetDatasetItemId() string`

GetDatasetItemId returns the DatasetItemId field if non-nil, zero value otherwise.

### GetDatasetItemIdOk

`func (o *EvalTraceView) GetDatasetItemIdOk() (*string, bool)`

GetDatasetItemIdOk returns a tuple with the DatasetItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetItemId

`func (o *EvalTraceView) SetDatasetItemId(v string)`

SetDatasetItemId sets DatasetItemId field to given value.

### HasDatasetItemId

`func (o *EvalTraceView) HasDatasetItemId() bool`

HasDatasetItemId returns a boolean if a field has been set.

### GetDatasetName

`func (o *EvalTraceView) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *EvalTraceView) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *EvalTraceView) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *EvalTraceView) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### GetEndTime

`func (o *EvalTraceView) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EvalTraceView) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EvalTraceView) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *EvalTraceView) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *EvalTraceView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvalTraceView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvalTraceView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EvalTraceView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *EvalTraceView) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *EvalTraceView) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *EvalTraceView) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *EvalTraceView) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *EvalTraceView) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *EvalTraceView) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetLatencyMs

`func (o *EvalTraceView) GetLatencyMs() float64`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *EvalTraceView) GetLatencyMsOk() (*float64, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *EvalTraceView) SetLatencyMs(v float64)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *EvalTraceView) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetModel

`func (o *EvalTraceView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvalTraceView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvalTraceView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EvalTraceView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *EvalTraceView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvalTraceView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvalTraceView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EvalTraceView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *EvalTraceView) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *EvalTraceView) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *EvalTraceView) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *EvalTraceView) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetProjectId

`func (o *EvalTraceView) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EvalTraceView) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EvalTraceView) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *EvalTraceView) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRunName

`func (o *EvalTraceView) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *EvalTraceView) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *EvalTraceView) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *EvalTraceView) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetSessionId

`func (o *EvalTraceView) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *EvalTraceView) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *EvalTraceView) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *EvalTraceView) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStartTime

`func (o *EvalTraceView) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EvalTraceView) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EvalTraceView) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *EvalTraceView) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTimestamp

`func (o *EvalTraceView) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EvalTraceView) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EvalTraceView) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EvalTraceView) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


