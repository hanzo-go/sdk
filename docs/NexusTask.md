# NexusTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **string** |  | [optional] 
**Application** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Example** | Pointer to **string** |  | [optional] 
**Grade** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Log** | Pointer to **string** |  | [optional] 
**ModelUsageMap** | Pointer to [**NexusUsageInfo**](NexusUsageInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to **[]string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewNexusTask

`func NewNexusTask() *NexusTask`

NewNexusTask instantiates a new NexusTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusTaskWithDefaults

`func NewNexusTaskWithDefaults() *NexusTask`

NewNexusTaskWithDefaults instantiates a new NexusTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *NexusTask) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *NexusTask) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *NexusTask) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *NexusTask) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetApplication

`func (o *NexusTask) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *NexusTask) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *NexusTask) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *NexusTask) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NexusTask) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NexusTask) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NexusTask) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NexusTask) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisplayName

`func (o *NexusTask) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NexusTask) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NexusTask) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NexusTask) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExample

`func (o *NexusTask) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *NexusTask) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *NexusTask) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *NexusTask) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetGrade

`func (o *NexusTask) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *NexusTask) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *NexusTask) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *NexusTask) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### GetLabels

`func (o *NexusTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NexusTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NexusTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NexusTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLog

`func (o *NexusTask) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *NexusTask) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *NexusTask) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *NexusTask) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetModelUsageMap

`func (o *NexusTask) GetModelUsageMap() NexusUsageInfo`

GetModelUsageMap returns the ModelUsageMap field if non-nil, zero value otherwise.

### GetModelUsageMapOk

`func (o *NexusTask) GetModelUsageMapOk() (*NexusUsageInfo, bool)`

GetModelUsageMapOk returns a tuple with the ModelUsageMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelUsageMap

`func (o *NexusTask) SetModelUsageMap(v NexusUsageInfo)`

SetModelUsageMap sets ModelUsageMap field to given value.

### HasModelUsageMap

`func (o *NexusTask) HasModelUsageMap() bool`

HasModelUsageMap returns a boolean if a field has been set.

### GetName

`func (o *NexusTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NexusTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NexusTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NexusTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *NexusTask) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NexusTask) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NexusTask) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NexusTask) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPath

`func (o *NexusTask) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *NexusTask) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *NexusTask) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *NexusTask) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProvider

`func (o *NexusTask) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NexusTask) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NexusTask) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NexusTask) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviders

`func (o *NexusTask) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *NexusTask) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *NexusTask) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *NexusTask) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetResult

`func (o *NexusTask) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *NexusTask) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *NexusTask) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *NexusTask) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubject

`func (o *NexusTask) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NexusTask) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NexusTask) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *NexusTask) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *NexusTask) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NexusTask) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NexusTask) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *NexusTask) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTopic

`func (o *NexusTask) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *NexusTask) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *NexusTask) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *NexusTask) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetType

`func (o *NexusTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NexusTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NexusTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NexusTask) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


