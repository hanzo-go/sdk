# NexusVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioUrl** | Pointer to **string** |  | [optional] 
**Class** | Pointer to **string** |  | [optional] 
**CoverUrl** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DataUrl** | Pointer to **string** |  | [optional] 
**DataUrls** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**EditMode** | Pointer to **string** |  | [optional] 
**ExcellentCount** | Pointer to **int64** |  | [optional] 
**Grade** | Pointer to **string** |  | [optional] 
**Grade2** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**LabelCount** | Pointer to **int64** |  | [optional] 
**Labels** | Pointer to [**[]NexusLabel**](NexusLabel.md) |  | [optional] 
**Lesson** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PlayAuth** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to [**[]NexusRemark**](NexusRemark.md) |  | [optional] 
**Remarks2** | Pointer to [**[]NexusRemark**](NexusRemark.md) |  | [optional] 
**ReviewState** | Pointer to **string** |  | [optional] 
**School** | Pointer to **string** |  | [optional] 
**SegmentCount** | Pointer to **int64** |  | [optional] 
**Segments** | Pointer to [**[]NexusLabel**](NexusLabel.md) |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**TagOnPause** | Pointer to **bool** |  | [optional] 
**Task1** | Pointer to **string** |  | [optional] 
**Task2** | Pointer to **string** |  | [optional] 
**Task3** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**VideoId** | Pointer to **string** |  | [optional] 
**VideoLength** | Pointer to **string** |  | [optional] 
**WordCountMap** | Pointer to **map[string]int64** |  | [optional] 

## Methods

### NewNexusVideo

`func NewNexusVideo() *NexusVideo`

NewNexusVideo instantiates a new NexusVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusVideoWithDefaults

`func NewNexusVideoWithDefaults() *NexusVideo`

NewNexusVideoWithDefaults instantiates a new NexusVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioUrl

`func (o *NexusVideo) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *NexusVideo) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *NexusVideo) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.

### HasAudioUrl

`func (o *NexusVideo) HasAudioUrl() bool`

HasAudioUrl returns a boolean if a field has been set.

### GetClass

`func (o *NexusVideo) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *NexusVideo) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *NexusVideo) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *NexusVideo) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCoverUrl

`func (o *NexusVideo) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *NexusVideo) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *NexusVideo) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *NexusVideo) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NexusVideo) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NexusVideo) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NexusVideo) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NexusVideo) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDataUrl

`func (o *NexusVideo) GetDataUrl() string`

GetDataUrl returns the DataUrl field if non-nil, zero value otherwise.

### GetDataUrlOk

`func (o *NexusVideo) GetDataUrlOk() (*string, bool)`

GetDataUrlOk returns a tuple with the DataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUrl

`func (o *NexusVideo) SetDataUrl(v string)`

SetDataUrl sets DataUrl field to given value.

### HasDataUrl

`func (o *NexusVideo) HasDataUrl() bool`

HasDataUrl returns a boolean if a field has been set.

### GetDataUrls

`func (o *NexusVideo) GetDataUrls() []string`

GetDataUrls returns the DataUrls field if non-nil, zero value otherwise.

### GetDataUrlsOk

`func (o *NexusVideo) GetDataUrlsOk() (*[]string, bool)`

GetDataUrlsOk returns a tuple with the DataUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUrls

`func (o *NexusVideo) SetDataUrls(v []string)`

SetDataUrls sets DataUrls field to given value.

### HasDataUrls

`func (o *NexusVideo) HasDataUrls() bool`

HasDataUrls returns a boolean if a field has been set.

### GetDescription

`func (o *NexusVideo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NexusVideo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NexusVideo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NexusVideo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *NexusVideo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NexusVideo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NexusVideo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NexusVideo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *NexusVideo) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *NexusVideo) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *NexusVideo) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *NexusVideo) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetEditMode

`func (o *NexusVideo) GetEditMode() string`

GetEditMode returns the EditMode field if non-nil, zero value otherwise.

### GetEditModeOk

`func (o *NexusVideo) GetEditModeOk() (*string, bool)`

GetEditModeOk returns a tuple with the EditMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditMode

`func (o *NexusVideo) SetEditMode(v string)`

SetEditMode sets EditMode field to given value.

### HasEditMode

`func (o *NexusVideo) HasEditMode() bool`

HasEditMode returns a boolean if a field has been set.

### GetExcellentCount

`func (o *NexusVideo) GetExcellentCount() int64`

GetExcellentCount returns the ExcellentCount field if non-nil, zero value otherwise.

### GetExcellentCountOk

`func (o *NexusVideo) GetExcellentCountOk() (*int64, bool)`

GetExcellentCountOk returns a tuple with the ExcellentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcellentCount

`func (o *NexusVideo) SetExcellentCount(v int64)`

SetExcellentCount sets ExcellentCount field to given value.

### HasExcellentCount

`func (o *NexusVideo) HasExcellentCount() bool`

HasExcellentCount returns a boolean if a field has been set.

### GetGrade

`func (o *NexusVideo) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *NexusVideo) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *NexusVideo) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *NexusVideo) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### GetGrade2

`func (o *NexusVideo) GetGrade2() string`

GetGrade2 returns the Grade2 field if non-nil, zero value otherwise.

### GetGrade2Ok

`func (o *NexusVideo) GetGrade2Ok() (*string, bool)`

GetGrade2Ok returns a tuple with the Grade2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade2

`func (o *NexusVideo) SetGrade2(v string)`

SetGrade2 sets Grade2 field to given value.

### HasGrade2

`func (o *NexusVideo) HasGrade2() bool`

HasGrade2 returns a boolean if a field has been set.

### GetIsPublic

`func (o *NexusVideo) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *NexusVideo) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *NexusVideo) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *NexusVideo) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetKeywords

`func (o *NexusVideo) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *NexusVideo) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *NexusVideo) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *NexusVideo) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLabelCount

`func (o *NexusVideo) GetLabelCount() int64`

GetLabelCount returns the LabelCount field if non-nil, zero value otherwise.

### GetLabelCountOk

`func (o *NexusVideo) GetLabelCountOk() (*int64, bool)`

GetLabelCountOk returns a tuple with the LabelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCount

`func (o *NexusVideo) SetLabelCount(v int64)`

SetLabelCount sets LabelCount field to given value.

### HasLabelCount

`func (o *NexusVideo) HasLabelCount() bool`

HasLabelCount returns a boolean if a field has been set.

### GetLabels

`func (o *NexusVideo) GetLabels() []NexusLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NexusVideo) GetLabelsOk() (*[]NexusLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NexusVideo) SetLabels(v []NexusLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NexusVideo) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLesson

`func (o *NexusVideo) GetLesson() string`

GetLesson returns the Lesson field if non-nil, zero value otherwise.

### GetLessonOk

`func (o *NexusVideo) GetLessonOk() (*string, bool)`

GetLessonOk returns a tuple with the Lesson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLesson

`func (o *NexusVideo) SetLesson(v string)`

SetLesson sets Lesson field to given value.

### HasLesson

`func (o *NexusVideo) HasLesson() bool`

HasLesson returns a boolean if a field has been set.

### GetName

`func (o *NexusVideo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NexusVideo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NexusVideo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NexusVideo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *NexusVideo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NexusVideo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NexusVideo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NexusVideo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPlayAuth

`func (o *NexusVideo) GetPlayAuth() string`

GetPlayAuth returns the PlayAuth field if non-nil, zero value otherwise.

### GetPlayAuthOk

`func (o *NexusVideo) GetPlayAuthOk() (*string, bool)`

GetPlayAuthOk returns a tuple with the PlayAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayAuth

`func (o *NexusVideo) SetPlayAuth(v string)`

SetPlayAuth sets PlayAuth field to given value.

### HasPlayAuth

`func (o *NexusVideo) HasPlayAuth() bool`

HasPlayAuth returns a boolean if a field has been set.

### GetRemarks

`func (o *NexusVideo) GetRemarks() []NexusRemark`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *NexusVideo) GetRemarksOk() (*[]NexusRemark, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *NexusVideo) SetRemarks(v []NexusRemark)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *NexusVideo) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetRemarks2

`func (o *NexusVideo) GetRemarks2() []NexusRemark`

GetRemarks2 returns the Remarks2 field if non-nil, zero value otherwise.

### GetRemarks2Ok

`func (o *NexusVideo) GetRemarks2Ok() (*[]NexusRemark, bool)`

GetRemarks2Ok returns a tuple with the Remarks2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks2

`func (o *NexusVideo) SetRemarks2(v []NexusRemark)`

SetRemarks2 sets Remarks2 field to given value.

### HasRemarks2

`func (o *NexusVideo) HasRemarks2() bool`

HasRemarks2 returns a boolean if a field has been set.

### GetReviewState

`func (o *NexusVideo) GetReviewState() string`

GetReviewState returns the ReviewState field if non-nil, zero value otherwise.

### GetReviewStateOk

`func (o *NexusVideo) GetReviewStateOk() (*string, bool)`

GetReviewStateOk returns a tuple with the ReviewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewState

`func (o *NexusVideo) SetReviewState(v string)`

SetReviewState sets ReviewState field to given value.

### HasReviewState

`func (o *NexusVideo) HasReviewState() bool`

HasReviewState returns a boolean if a field has been set.

### GetSchool

`func (o *NexusVideo) GetSchool() string`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *NexusVideo) GetSchoolOk() (*string, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *NexusVideo) SetSchool(v string)`

SetSchool sets School field to given value.

### HasSchool

`func (o *NexusVideo) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetSegmentCount

`func (o *NexusVideo) GetSegmentCount() int64`

GetSegmentCount returns the SegmentCount field if non-nil, zero value otherwise.

### GetSegmentCountOk

`func (o *NexusVideo) GetSegmentCountOk() (*int64, bool)`

GetSegmentCountOk returns a tuple with the SegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentCount

`func (o *NexusVideo) SetSegmentCount(v int64)`

SetSegmentCount sets SegmentCount field to given value.

### HasSegmentCount

`func (o *NexusVideo) HasSegmentCount() bool`

HasSegmentCount returns a boolean if a field has been set.

### GetSegments

`func (o *NexusVideo) GetSegments() []NexusLabel`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *NexusVideo) GetSegmentsOk() (*[]NexusLabel, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *NexusVideo) SetSegments(v []NexusLabel)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *NexusVideo) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetStage

`func (o *NexusVideo) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *NexusVideo) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *NexusVideo) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *NexusVideo) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetState

`func (o *NexusVideo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NexusVideo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NexusVideo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NexusVideo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubject

`func (o *NexusVideo) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NexusVideo) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NexusVideo) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *NexusVideo) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTag

`func (o *NexusVideo) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *NexusVideo) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *NexusVideo) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *NexusVideo) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTagOnPause

`func (o *NexusVideo) GetTagOnPause() bool`

GetTagOnPause returns the TagOnPause field if non-nil, zero value otherwise.

### GetTagOnPauseOk

`func (o *NexusVideo) GetTagOnPauseOk() (*bool, bool)`

GetTagOnPauseOk returns a tuple with the TagOnPause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagOnPause

`func (o *NexusVideo) SetTagOnPause(v bool)`

SetTagOnPause sets TagOnPause field to given value.

### HasTagOnPause

`func (o *NexusVideo) HasTagOnPause() bool`

HasTagOnPause returns a boolean if a field has been set.

### GetTask1

`func (o *NexusVideo) GetTask1() string`

GetTask1 returns the Task1 field if non-nil, zero value otherwise.

### GetTask1Ok

`func (o *NexusVideo) GetTask1Ok() (*string, bool)`

GetTask1Ok returns a tuple with the Task1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask1

`func (o *NexusVideo) SetTask1(v string)`

SetTask1 sets Task1 field to given value.

### HasTask1

`func (o *NexusVideo) HasTask1() bool`

HasTask1 returns a boolean if a field has been set.

### GetTask2

`func (o *NexusVideo) GetTask2() string`

GetTask2 returns the Task2 field if non-nil, zero value otherwise.

### GetTask2Ok

`func (o *NexusVideo) GetTask2Ok() (*string, bool)`

GetTask2Ok returns a tuple with the Task2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask2

`func (o *NexusVideo) SetTask2(v string)`

SetTask2 sets Task2 field to given value.

### HasTask2

`func (o *NexusVideo) HasTask2() bool`

HasTask2 returns a boolean if a field has been set.

### GetTask3

`func (o *NexusVideo) GetTask3() string`

GetTask3 returns the Task3 field if non-nil, zero value otherwise.

### GetTask3Ok

`func (o *NexusVideo) GetTask3Ok() (*string, bool)`

GetTask3Ok returns a tuple with the Task3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask3

`func (o *NexusVideo) SetTask3(v string)`

SetTask3 sets Task3 field to given value.

### HasTask3

`func (o *NexusVideo) HasTask3() bool`

HasTask3 returns a boolean if a field has been set.

### GetTemplate

`func (o *NexusVideo) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *NexusVideo) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *NexusVideo) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *NexusVideo) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTopic

`func (o *NexusVideo) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *NexusVideo) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *NexusVideo) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *NexusVideo) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetType

`func (o *NexusVideo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NexusVideo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NexusVideo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NexusVideo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *NexusVideo) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *NexusVideo) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *NexusVideo) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *NexusVideo) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetVideoId

`func (o *NexusVideo) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *NexusVideo) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *NexusVideo) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *NexusVideo) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetVideoLength

`func (o *NexusVideo) GetVideoLength() string`

GetVideoLength returns the VideoLength field if non-nil, zero value otherwise.

### GetVideoLengthOk

`func (o *NexusVideo) GetVideoLengthOk() (*string, bool)`

GetVideoLengthOk returns a tuple with the VideoLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLength

`func (o *NexusVideo) SetVideoLength(v string)`

SetVideoLength sets VideoLength field to given value.

### HasVideoLength

`func (o *NexusVideo) HasVideoLength() bool`

HasVideoLength returns a boolean if a field has been set.

### GetWordCountMap

`func (o *NexusVideo) GetWordCountMap() map[string]int64`

GetWordCountMap returns the WordCountMap field if non-nil, zero value otherwise.

### GetWordCountMapOk

`func (o *NexusVideo) GetWordCountMapOk() (*map[string]int64, bool)`

GetWordCountMapOk returns a tuple with the WordCountMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordCountMap

`func (o *NexusVideo) SetWordCountMap(v map[string]int64)`

SetWordCountMap sets WordCountMap field to given value.

### HasWordCountMap

`func (o *NexusVideo) HasWordCountMap() bool`

HasWordCountMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


