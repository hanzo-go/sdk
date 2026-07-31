# CloudRoutedRunOut

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

### NewCloudRoutedRunOut

`func NewCloudRoutedRunOut() *CloudRoutedRunOut`

NewCloudRoutedRunOut instantiates a new CloudRoutedRunOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRoutedRunOutWithDefaults

`func NewCloudRoutedRunOutWithDefaults() *CloudRoutedRunOut`

NewCloudRoutedRunOutWithDefaults instantiates a new CloudRoutedRunOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *CloudRoutedRunOut) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *CloudRoutedRunOut) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *CloudRoutedRunOut) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *CloudRoutedRunOut) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBranch

`func (o *CloudRoutedRunOut) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CloudRoutedRunOut) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CloudRoutedRunOut) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CloudRoutedRunOut) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCloneUrl

`func (o *CloudRoutedRunOut) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *CloudRoutedRunOut) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *CloudRoutedRunOut) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *CloudRoutedRunOut) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetProject

`func (o *CloudRoutedRunOut) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudRoutedRunOut) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudRoutedRunOut) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudRoutedRunOut) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPrompt

`func (o *CloudRoutedRunOut) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CloudRoutedRunOut) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CloudRoutedRunOut) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *CloudRoutedRunOut) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetRepo

`func (o *CloudRoutedRunOut) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudRoutedRunOut) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudRoutedRunOut) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudRoutedRunOut) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSessionId

`func (o *CloudRoutedRunOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CloudRoutedRunOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CloudRoutedRunOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CloudRoutedRunOut) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *CloudRoutedRunOut) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *CloudRoutedRunOut) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *CloudRoutedRunOut) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *CloudRoutedRunOut) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


