# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DocumentText** | Pointer to **string** |  | [optional] 
**DocumentUrl** | Pointer to **string** |  | [optional] 
**Example** | Pointer to **string** |  | [optional] 
**Grade** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Log** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Scale** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *Task) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *Task) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *Task) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *Task) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Task) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Task) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Task) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Task) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisplayName

`func (o *Task) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Task) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Task) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Task) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDocumentText

`func (o *Task) GetDocumentText() string`

GetDocumentText returns the DocumentText field if non-nil, zero value otherwise.

### GetDocumentTextOk

`func (o *Task) GetDocumentTextOk() (*string, bool)`

GetDocumentTextOk returns a tuple with the DocumentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentText

`func (o *Task) SetDocumentText(v string)`

SetDocumentText sets DocumentText field to given value.

### HasDocumentText

`func (o *Task) HasDocumentText() bool`

HasDocumentText returns a boolean if a field has been set.

### GetDocumentUrl

`func (o *Task) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *Task) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *Task) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.

### HasDocumentUrl

`func (o *Task) HasDocumentUrl() bool`

HasDocumentUrl returns a boolean if a field has been set.

### GetExample

`func (o *Task) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *Task) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *Task) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *Task) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetGrade

`func (o *Task) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *Task) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *Task) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *Task) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### GetLabels

`func (o *Task) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Task) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Task) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Task) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLog

`func (o *Task) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *Task) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *Task) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *Task) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Task) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *Task) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Task) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Task) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Task) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPath

`func (o *Task) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Task) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Task) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Task) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProvider

`func (o *Task) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Task) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Task) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Task) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetResult

`func (o *Task) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Task) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Task) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Task) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetScale

`func (o *Task) GetScale() string`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *Task) GetScaleOk() (*string, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *Task) SetScale(v string)`

SetScale sets Scale field to given value.

### HasScale

`func (o *Task) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetScore

`func (o *Task) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Task) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Task) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Task) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSubject

`func (o *Task) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Task) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Task) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Task) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTopic

`func (o *Task) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *Task) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *Task) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *Task) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetType

`func (o *Task) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Task) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Task) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Task) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


