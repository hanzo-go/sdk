# ConsoleTraceWithFullDetails

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
**Observations** | Pointer to [**[]ConsoleObservation**](ConsoleObservation.md) |  | [optional] 
**Scores** | Pointer to [**[]ConsoleScore**](ConsoleScore.md) |  | [optional] 

## Methods

### NewConsoleTraceWithFullDetails

`func NewConsoleTraceWithFullDetails() *ConsoleTraceWithFullDetails`

NewConsoleTraceWithFullDetails instantiates a new ConsoleTraceWithFullDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleTraceWithFullDetailsWithDefaults

`func NewConsoleTraceWithFullDetailsWithDefaults() *ConsoleTraceWithFullDetails`

NewConsoleTraceWithFullDetailsWithDefaults instantiates a new ConsoleTraceWithFullDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleTraceWithFullDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleTraceWithFullDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleTraceWithFullDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleTraceWithFullDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ConsoleTraceWithFullDetails) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ConsoleTraceWithFullDetails) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ConsoleTraceWithFullDetails) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ConsoleTraceWithFullDetails) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ConsoleTraceWithFullDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleTraceWithFullDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleTraceWithFullDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsoleTraceWithFullDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *ConsoleTraceWithFullDetails) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ConsoleTraceWithFullDetails) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ConsoleTraceWithFullDetails) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ConsoleTraceWithFullDetails) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSessionId

`func (o *ConsoleTraceWithFullDetails) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ConsoleTraceWithFullDetails) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ConsoleTraceWithFullDetails) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ConsoleTraceWithFullDetails) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetRelease

`func (o *ConsoleTraceWithFullDetails) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *ConsoleTraceWithFullDetails) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *ConsoleTraceWithFullDetails) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *ConsoleTraceWithFullDetails) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetVersion

`func (o *ConsoleTraceWithFullDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConsoleTraceWithFullDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConsoleTraceWithFullDetails) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConsoleTraceWithFullDetails) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnvironment

`func (o *ConsoleTraceWithFullDetails) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConsoleTraceWithFullDetails) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConsoleTraceWithFullDetails) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConsoleTraceWithFullDetails) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetMetadata

`func (o *ConsoleTraceWithFullDetails) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleTraceWithFullDetails) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleTraceWithFullDetails) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleTraceWithFullDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetInput

`func (o *ConsoleTraceWithFullDetails) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ConsoleTraceWithFullDetails) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ConsoleTraceWithFullDetails) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ConsoleTraceWithFullDetails) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ConsoleTraceWithFullDetails) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ConsoleTraceWithFullDetails) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetOutput

`func (o *ConsoleTraceWithFullDetails) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ConsoleTraceWithFullDetails) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ConsoleTraceWithFullDetails) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ConsoleTraceWithFullDetails) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *ConsoleTraceWithFullDetails) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *ConsoleTraceWithFullDetails) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetTags

`func (o *ConsoleTraceWithFullDetails) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConsoleTraceWithFullDetails) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConsoleTraceWithFullDetails) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConsoleTraceWithFullDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPublic

`func (o *ConsoleTraceWithFullDetails) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ConsoleTraceWithFullDetails) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ConsoleTraceWithFullDetails) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *ConsoleTraceWithFullDetails) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetBookmarked

`func (o *ConsoleTraceWithFullDetails) GetBookmarked() bool`

GetBookmarked returns the Bookmarked field if non-nil, zero value otherwise.

### GetBookmarkedOk

`func (o *ConsoleTraceWithFullDetails) GetBookmarkedOk() (*bool, bool)`

GetBookmarkedOk returns a tuple with the Bookmarked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarked

`func (o *ConsoleTraceWithFullDetails) SetBookmarked(v bool)`

SetBookmarked sets Bookmarked field to given value.

### HasBookmarked

`func (o *ConsoleTraceWithFullDetails) HasBookmarked() bool`

HasBookmarked returns a boolean if a field has been set.

### GetHtmlPath

`func (o *ConsoleTraceWithFullDetails) GetHtmlPath() string`

GetHtmlPath returns the HtmlPath field if non-nil, zero value otherwise.

### GetHtmlPathOk

`func (o *ConsoleTraceWithFullDetails) GetHtmlPathOk() (*string, bool)`

GetHtmlPathOk returns a tuple with the HtmlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlPath

`func (o *ConsoleTraceWithFullDetails) SetHtmlPath(v string)`

SetHtmlPath sets HtmlPath field to given value.

### HasHtmlPath

`func (o *ConsoleTraceWithFullDetails) HasHtmlPath() bool`

HasHtmlPath returns a boolean if a field has been set.

### GetLatency

`func (o *ConsoleTraceWithFullDetails) GetLatency() float32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *ConsoleTraceWithFullDetails) GetLatencyOk() (*float32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *ConsoleTraceWithFullDetails) SetLatency(v float32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *ConsoleTraceWithFullDetails) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetTotalCost

`func (o *ConsoleTraceWithFullDetails) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *ConsoleTraceWithFullDetails) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *ConsoleTraceWithFullDetails) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *ConsoleTraceWithFullDetails) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetObservations

`func (o *ConsoleTraceWithFullDetails) GetObservations() []ConsoleObservation`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *ConsoleTraceWithFullDetails) GetObservationsOk() (*[]ConsoleObservation, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *ConsoleTraceWithFullDetails) SetObservations(v []ConsoleObservation)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *ConsoleTraceWithFullDetails) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### GetScores

`func (o *ConsoleTraceWithFullDetails) GetScores() []ConsoleScore`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *ConsoleTraceWithFullDetails) GetScoresOk() (*[]ConsoleScore, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *ConsoleTraceWithFullDetails) SetScores(v []ConsoleScore)`

SetScores sets Scores field to given value.

### HasScores

`func (o *ConsoleTraceWithFullDetails) HasScores() bool`

HasScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


