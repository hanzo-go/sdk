# ConsoleTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 
**Output** | Pointer to **interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**Bookmarked** | Pointer to **bool** |  | [optional] 
**HtmlPath** | Pointer to **string** |  | [optional] 
**Latency** | Pointer to **float32** | Latency in seconds | [optional] 
**TotalCost** | Pointer to **float32** | Total cost in USD | [optional] 
**Observations** | Pointer to **[]string** |  | [optional] 
**Scores** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConsoleTrace

`func NewConsoleTrace() *ConsoleTrace`

NewConsoleTrace instantiates a new ConsoleTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleTraceWithDefaults

`func NewConsoleTraceWithDefaults() *ConsoleTrace`

NewConsoleTraceWithDefaults instantiates a new ConsoleTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleTrace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleTrace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleTrace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleTrace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ConsoleTrace) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ConsoleTrace) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ConsoleTrace) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ConsoleTrace) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ConsoleTrace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleTrace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleTrace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsoleTrace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *ConsoleTrace) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ConsoleTrace) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ConsoleTrace) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ConsoleTrace) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSessionId

`func (o *ConsoleTrace) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ConsoleTrace) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ConsoleTrace) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ConsoleTrace) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetRelease

`func (o *ConsoleTrace) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *ConsoleTrace) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *ConsoleTrace) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *ConsoleTrace) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetVersion

`func (o *ConsoleTrace) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConsoleTrace) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConsoleTrace) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConsoleTrace) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnvironment

`func (o *ConsoleTrace) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConsoleTrace) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConsoleTrace) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConsoleTrace) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetMetadata

`func (o *ConsoleTrace) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleTrace) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleTrace) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleTrace) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetInput

`func (o *ConsoleTrace) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ConsoleTrace) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ConsoleTrace) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ConsoleTrace) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ConsoleTrace) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ConsoleTrace) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetOutput

`func (o *ConsoleTrace) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ConsoleTrace) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ConsoleTrace) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ConsoleTrace) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *ConsoleTrace) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *ConsoleTrace) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetTags

`func (o *ConsoleTrace) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConsoleTrace) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConsoleTrace) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConsoleTrace) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPublic

`func (o *ConsoleTrace) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ConsoleTrace) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ConsoleTrace) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *ConsoleTrace) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetBookmarked

`func (o *ConsoleTrace) GetBookmarked() bool`

GetBookmarked returns the Bookmarked field if non-nil, zero value otherwise.

### GetBookmarkedOk

`func (o *ConsoleTrace) GetBookmarkedOk() (*bool, bool)`

GetBookmarkedOk returns a tuple with the Bookmarked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarked

`func (o *ConsoleTrace) SetBookmarked(v bool)`

SetBookmarked sets Bookmarked field to given value.

### HasBookmarked

`func (o *ConsoleTrace) HasBookmarked() bool`

HasBookmarked returns a boolean if a field has been set.

### GetHtmlPath

`func (o *ConsoleTrace) GetHtmlPath() string`

GetHtmlPath returns the HtmlPath field if non-nil, zero value otherwise.

### GetHtmlPathOk

`func (o *ConsoleTrace) GetHtmlPathOk() (*string, bool)`

GetHtmlPathOk returns a tuple with the HtmlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlPath

`func (o *ConsoleTrace) SetHtmlPath(v string)`

SetHtmlPath sets HtmlPath field to given value.

### HasHtmlPath

`func (o *ConsoleTrace) HasHtmlPath() bool`

HasHtmlPath returns a boolean if a field has been set.

### GetLatency

`func (o *ConsoleTrace) GetLatency() float32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *ConsoleTrace) GetLatencyOk() (*float32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *ConsoleTrace) SetLatency(v float32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *ConsoleTrace) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetTotalCost

`func (o *ConsoleTrace) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *ConsoleTrace) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *ConsoleTrace) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *ConsoleTrace) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetObservations

`func (o *ConsoleTrace) GetObservations() []string`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *ConsoleTrace) GetObservationsOk() (*[]string, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *ConsoleTrace) SetObservations(v []string)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *ConsoleTrace) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### GetScores

`func (o *ConsoleTrace) GetScores() []string`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *ConsoleTrace) GetScoresOk() (*[]string, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *ConsoleTrace) SetScores(v []string)`

SetScores sets Scores field to given value.

### HasScores

`func (o *ConsoleTrace) HasScores() bool`

HasScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


