# ConsoleObservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**CompletionStartTime** | Pointer to **time.Time** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 
**Output** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**ParentObservationId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**PromptId** | Pointer to **string** |  | [optional] 
**PromptName** | Pointer to **string** |  | [optional] 
**PromptVersion** | Pointer to **int32** |  | [optional] 
**Usage** | Pointer to [**ConsoleUsage**](ConsoleUsage.md) |  | [optional] 
**CalculatedInputCost** | Pointer to **float32** |  | [optional] 
**CalculatedOutputCost** | Pointer to **float32** |  | [optional] 
**CalculatedTotalCost** | Pointer to **float32** |  | [optional] 
**Latency** | Pointer to **float32** | Latency in seconds | [optional] 
**TimeToFirstToken** | Pointer to **float32** | Time to first token in seconds | [optional] 

## Methods

### NewConsoleObservation

`func NewConsoleObservation() *ConsoleObservation`

NewConsoleObservation instantiates a new ConsoleObservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleObservationWithDefaults

`func NewConsoleObservationWithDefaults() *ConsoleObservation`

NewConsoleObservationWithDefaults instantiates a new ConsoleObservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleObservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleObservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleObservation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleObservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTraceId

`func (o *ConsoleObservation) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ConsoleObservation) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ConsoleObservation) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ConsoleObservation) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *ConsoleObservation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsoleObservation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsoleObservation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConsoleObservation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ConsoleObservation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleObservation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleObservation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsoleObservation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *ConsoleObservation) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ConsoleObservation) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ConsoleObservation) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ConsoleObservation) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ConsoleObservation) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ConsoleObservation) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ConsoleObservation) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ConsoleObservation) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetCompletionStartTime

`func (o *ConsoleObservation) GetCompletionStartTime() time.Time`

GetCompletionStartTime returns the CompletionStartTime field if non-nil, zero value otherwise.

### GetCompletionStartTimeOk

`func (o *ConsoleObservation) GetCompletionStartTimeOk() (*time.Time, bool)`

GetCompletionStartTimeOk returns a tuple with the CompletionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionStartTime

`func (o *ConsoleObservation) SetCompletionStartTime(v time.Time)`

SetCompletionStartTime sets CompletionStartTime field to given value.

### HasCompletionStartTime

`func (o *ConsoleObservation) HasCompletionStartTime() bool`

HasCompletionStartTime returns a boolean if a field has been set.

### GetModel

`func (o *ConsoleObservation) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConsoleObservation) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConsoleObservation) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConsoleObservation) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelParameters

`func (o *ConsoleObservation) GetModelParameters() map[string]interface{}`

GetModelParameters returns the ModelParameters field if non-nil, zero value otherwise.

### GetModelParametersOk

`func (o *ConsoleObservation) GetModelParametersOk() (*map[string]interface{}, bool)`

GetModelParametersOk returns a tuple with the ModelParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelParameters

`func (o *ConsoleObservation) SetModelParameters(v map[string]interface{})`

SetModelParameters sets ModelParameters field to given value.

### HasModelParameters

`func (o *ConsoleObservation) HasModelParameters() bool`

HasModelParameters returns a boolean if a field has been set.

### GetInput

`func (o *ConsoleObservation) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ConsoleObservation) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ConsoleObservation) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ConsoleObservation) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ConsoleObservation) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ConsoleObservation) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetOutput

`func (o *ConsoleObservation) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ConsoleObservation) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ConsoleObservation) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ConsoleObservation) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *ConsoleObservation) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *ConsoleObservation) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetMetadata

`func (o *ConsoleObservation) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleObservation) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleObservation) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleObservation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLevel

`func (o *ConsoleObservation) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ConsoleObservation) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ConsoleObservation) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ConsoleObservation) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ConsoleObservation) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ConsoleObservation) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ConsoleObservation) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ConsoleObservation) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetParentObservationId

`func (o *ConsoleObservation) GetParentObservationId() string`

GetParentObservationId returns the ParentObservationId field if non-nil, zero value otherwise.

### GetParentObservationIdOk

`func (o *ConsoleObservation) GetParentObservationIdOk() (*string, bool)`

GetParentObservationIdOk returns a tuple with the ParentObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObservationId

`func (o *ConsoleObservation) SetParentObservationId(v string)`

SetParentObservationId sets ParentObservationId field to given value.

### HasParentObservationId

`func (o *ConsoleObservation) HasParentObservationId() bool`

HasParentObservationId returns a boolean if a field has been set.

### GetVersion

`func (o *ConsoleObservation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConsoleObservation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConsoleObservation) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConsoleObservation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnvironment

`func (o *ConsoleObservation) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConsoleObservation) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConsoleObservation) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConsoleObservation) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetPromptId

`func (o *ConsoleObservation) GetPromptId() string`

GetPromptId returns the PromptId field if non-nil, zero value otherwise.

### GetPromptIdOk

`func (o *ConsoleObservation) GetPromptIdOk() (*string, bool)`

GetPromptIdOk returns a tuple with the PromptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptId

`func (o *ConsoleObservation) SetPromptId(v string)`

SetPromptId sets PromptId field to given value.

### HasPromptId

`func (o *ConsoleObservation) HasPromptId() bool`

HasPromptId returns a boolean if a field has been set.

### GetPromptName

`func (o *ConsoleObservation) GetPromptName() string`

GetPromptName returns the PromptName field if non-nil, zero value otherwise.

### GetPromptNameOk

`func (o *ConsoleObservation) GetPromptNameOk() (*string, bool)`

GetPromptNameOk returns a tuple with the PromptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptName

`func (o *ConsoleObservation) SetPromptName(v string)`

SetPromptName sets PromptName field to given value.

### HasPromptName

`func (o *ConsoleObservation) HasPromptName() bool`

HasPromptName returns a boolean if a field has been set.

### GetPromptVersion

`func (o *ConsoleObservation) GetPromptVersion() int32`

GetPromptVersion returns the PromptVersion field if non-nil, zero value otherwise.

### GetPromptVersionOk

`func (o *ConsoleObservation) GetPromptVersionOk() (*int32, bool)`

GetPromptVersionOk returns a tuple with the PromptVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptVersion

`func (o *ConsoleObservation) SetPromptVersion(v int32)`

SetPromptVersion sets PromptVersion field to given value.

### HasPromptVersion

`func (o *ConsoleObservation) HasPromptVersion() bool`

HasPromptVersion returns a boolean if a field has been set.

### GetUsage

`func (o *ConsoleObservation) GetUsage() ConsoleUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ConsoleObservation) GetUsageOk() (*ConsoleUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ConsoleObservation) SetUsage(v ConsoleUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ConsoleObservation) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetCalculatedInputCost

`func (o *ConsoleObservation) GetCalculatedInputCost() float32`

GetCalculatedInputCost returns the CalculatedInputCost field if non-nil, zero value otherwise.

### GetCalculatedInputCostOk

`func (o *ConsoleObservation) GetCalculatedInputCostOk() (*float32, bool)`

GetCalculatedInputCostOk returns a tuple with the CalculatedInputCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedInputCost

`func (o *ConsoleObservation) SetCalculatedInputCost(v float32)`

SetCalculatedInputCost sets CalculatedInputCost field to given value.

### HasCalculatedInputCost

`func (o *ConsoleObservation) HasCalculatedInputCost() bool`

HasCalculatedInputCost returns a boolean if a field has been set.

### GetCalculatedOutputCost

`func (o *ConsoleObservation) GetCalculatedOutputCost() float32`

GetCalculatedOutputCost returns the CalculatedOutputCost field if non-nil, zero value otherwise.

### GetCalculatedOutputCostOk

`func (o *ConsoleObservation) GetCalculatedOutputCostOk() (*float32, bool)`

GetCalculatedOutputCostOk returns a tuple with the CalculatedOutputCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedOutputCost

`func (o *ConsoleObservation) SetCalculatedOutputCost(v float32)`

SetCalculatedOutputCost sets CalculatedOutputCost field to given value.

### HasCalculatedOutputCost

`func (o *ConsoleObservation) HasCalculatedOutputCost() bool`

HasCalculatedOutputCost returns a boolean if a field has been set.

### GetCalculatedTotalCost

`func (o *ConsoleObservation) GetCalculatedTotalCost() float32`

GetCalculatedTotalCost returns the CalculatedTotalCost field if non-nil, zero value otherwise.

### GetCalculatedTotalCostOk

`func (o *ConsoleObservation) GetCalculatedTotalCostOk() (*float32, bool)`

GetCalculatedTotalCostOk returns a tuple with the CalculatedTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedTotalCost

`func (o *ConsoleObservation) SetCalculatedTotalCost(v float32)`

SetCalculatedTotalCost sets CalculatedTotalCost field to given value.

### HasCalculatedTotalCost

`func (o *ConsoleObservation) HasCalculatedTotalCost() bool`

HasCalculatedTotalCost returns a boolean if a field has been set.

### GetLatency

`func (o *ConsoleObservation) GetLatency() float32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *ConsoleObservation) GetLatencyOk() (*float32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *ConsoleObservation) SetLatency(v float32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *ConsoleObservation) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetTimeToFirstToken

`func (o *ConsoleObservation) GetTimeToFirstToken() float32`

GetTimeToFirstToken returns the TimeToFirstToken field if non-nil, zero value otherwise.

### GetTimeToFirstTokenOk

`func (o *ConsoleObservation) GetTimeToFirstTokenOk() (*float32, bool)`

GetTimeToFirstTokenOk returns a tuple with the TimeToFirstToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToFirstToken

`func (o *ConsoleObservation) SetTimeToFirstToken(v float32)`

SetTimeToFirstToken sets TimeToFirstToken field to given value.

### HasTimeToFirstToken

`func (o *ConsoleObservation) HasTimeToFirstToken() bool`

HasTimeToFirstToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


