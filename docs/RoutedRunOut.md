# RoutedRunOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**CloneUrl** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** | Repo is the repository to work in and CloneURL is how to fetch it. | [optional] 
**SessionId** | Pointer to **string** | SessionID is the live session opened at dispatch; the machine streams its turns into it. | [optional] 
**TimeoutSeconds** | Pointer to **int32** | TimeoutSeconds bounds the run on the machine; 0 means the machine&#39;s own default. | [optional] 

## Methods

### NewRoutedRunOut

`func NewRoutedRunOut() *RoutedRunOut`

NewRoutedRunOut instantiates a new RoutedRunOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutedRunOutWithDefaults

`func NewRoutedRunOutWithDefaults() *RoutedRunOut`

NewRoutedRunOutWithDefaults instantiates a new RoutedRunOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *RoutedRunOut) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *RoutedRunOut) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *RoutedRunOut) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *RoutedRunOut) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBranch

`func (o *RoutedRunOut) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *RoutedRunOut) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *RoutedRunOut) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *RoutedRunOut) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCloneUrl

`func (o *RoutedRunOut) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *RoutedRunOut) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *RoutedRunOut) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *RoutedRunOut) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetProject

`func (o *RoutedRunOut) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RoutedRunOut) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RoutedRunOut) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *RoutedRunOut) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPrompt

`func (o *RoutedRunOut) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *RoutedRunOut) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *RoutedRunOut) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *RoutedRunOut) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetRepo

`func (o *RoutedRunOut) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *RoutedRunOut) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *RoutedRunOut) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *RoutedRunOut) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSessionId

`func (o *RoutedRunOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *RoutedRunOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *RoutedRunOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *RoutedRunOut) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *RoutedRunOut) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *RoutedRunOut) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *RoutedRunOut) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *RoutedRunOut) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


