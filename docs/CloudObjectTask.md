# CloudObjectTask

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
**ModelUsageMap** | Pointer to [**CloudObjectUsageInfo**](CloudObjectUsageInfo.md) |  | [optional] 
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

### NewCloudObjectTask

`func NewCloudObjectTask() *CloudObjectTask`

NewCloudObjectTask instantiates a new CloudObjectTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectTaskWithDefaults

`func NewCloudObjectTaskWithDefaults() *CloudObjectTask`

NewCloudObjectTaskWithDefaults instantiates a new CloudObjectTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *CloudObjectTask) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *CloudObjectTask) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *CloudObjectTask) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *CloudObjectTask) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetApplication

`func (o *CloudObjectTask) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *CloudObjectTask) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *CloudObjectTask) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *CloudObjectTask) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CloudObjectTask) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectTask) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectTask) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectTask) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudObjectTask) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudObjectTask) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudObjectTask) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudObjectTask) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExample

`func (o *CloudObjectTask) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *CloudObjectTask) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *CloudObjectTask) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *CloudObjectTask) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetGrade

`func (o *CloudObjectTask) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *CloudObjectTask) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *CloudObjectTask) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *CloudObjectTask) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### GetLabels

`func (o *CloudObjectTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CloudObjectTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CloudObjectTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CloudObjectTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLog

`func (o *CloudObjectTask) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *CloudObjectTask) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *CloudObjectTask) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *CloudObjectTask) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetModelUsageMap

`func (o *CloudObjectTask) GetModelUsageMap() CloudObjectUsageInfo`

GetModelUsageMap returns the ModelUsageMap field if non-nil, zero value otherwise.

### GetModelUsageMapOk

`func (o *CloudObjectTask) GetModelUsageMapOk() (*CloudObjectUsageInfo, bool)`

GetModelUsageMapOk returns a tuple with the ModelUsageMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelUsageMap

`func (o *CloudObjectTask) SetModelUsageMap(v CloudObjectUsageInfo)`

SetModelUsageMap sets ModelUsageMap field to given value.

### HasModelUsageMap

`func (o *CloudObjectTask) HasModelUsageMap() bool`

HasModelUsageMap returns a boolean if a field has been set.

### GetName

`func (o *CloudObjectTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudObjectTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *CloudObjectTask) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectTask) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectTask) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudObjectTask) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPath

`func (o *CloudObjectTask) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudObjectTask) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudObjectTask) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudObjectTask) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProvider

`func (o *CloudObjectTask) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudObjectTask) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudObjectTask) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudObjectTask) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviders

`func (o *CloudObjectTask) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CloudObjectTask) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CloudObjectTask) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CloudObjectTask) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetResult

`func (o *CloudObjectTask) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CloudObjectTask) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CloudObjectTask) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CloudObjectTask) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubject

`func (o *CloudObjectTask) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudObjectTask) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudObjectTask) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudObjectTask) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *CloudObjectTask) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CloudObjectTask) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CloudObjectTask) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CloudObjectTask) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTopic

`func (o *CloudObjectTask) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *CloudObjectTask) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *CloudObjectTask) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *CloudObjectTask) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetType

`func (o *CloudObjectTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudObjectTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudObjectTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudObjectTask) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


