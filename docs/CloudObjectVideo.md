# CloudObjectVideo

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
**Labels** | Pointer to [**[]CloudObjectLabel**](CloudObjectLabel.md) |  | [optional] 
**Lesson** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PlayAuth** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to [**[]CloudObjectRemark**](CloudObjectRemark.md) |  | [optional] 
**Remarks2** | Pointer to [**[]CloudObjectRemark**](CloudObjectRemark.md) |  | [optional] 
**ReviewState** | Pointer to **string** |  | [optional] 
**School** | Pointer to **string** |  | [optional] 
**SegmentCount** | Pointer to **int64** |  | [optional] 
**Segments** | Pointer to [**[]CloudObjectLabel**](CloudObjectLabel.md) |  | [optional] 
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
**WordCountMap** | Pointer to  |  | [optional] 

## Methods

### NewCloudObjectVideo

`func NewCloudObjectVideo() *CloudObjectVideo`

NewCloudObjectVideo instantiates a new CloudObjectVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectVideoWithDefaults

`func NewCloudObjectVideoWithDefaults() *CloudObjectVideo`

NewCloudObjectVideoWithDefaults instantiates a new CloudObjectVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioUrl

`func (o *CloudObjectVideo) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *CloudObjectVideo) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *CloudObjectVideo) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.

### HasAudioUrl

`func (o *CloudObjectVideo) HasAudioUrl() bool`

HasAudioUrl returns a boolean if a field has been set.

### GetClass

`func (o *CloudObjectVideo) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *CloudObjectVideo) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *CloudObjectVideo) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *CloudObjectVideo) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCoverUrl

`func (o *CloudObjectVideo) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *CloudObjectVideo) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *CloudObjectVideo) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *CloudObjectVideo) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CloudObjectVideo) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectVideo) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectVideo) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectVideo) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDataUrl

`func (o *CloudObjectVideo) GetDataUrl() string`

GetDataUrl returns the DataUrl field if non-nil, zero value otherwise.

### GetDataUrlOk

`func (o *CloudObjectVideo) GetDataUrlOk() (*string, bool)`

GetDataUrlOk returns a tuple with the DataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUrl

`func (o *CloudObjectVideo) SetDataUrl(v string)`

SetDataUrl sets DataUrl field to given value.

### HasDataUrl

`func (o *CloudObjectVideo) HasDataUrl() bool`

HasDataUrl returns a boolean if a field has been set.

### GetDataUrls

`func (o *CloudObjectVideo) GetDataUrls() []string`

GetDataUrls returns the DataUrls field if non-nil, zero value otherwise.

### GetDataUrlsOk

`func (o *CloudObjectVideo) GetDataUrlsOk() (*[]string, bool)`

GetDataUrlsOk returns a tuple with the DataUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUrls

`func (o *CloudObjectVideo) SetDataUrls(v []string)`

SetDataUrls sets DataUrls field to given value.

### HasDataUrls

`func (o *CloudObjectVideo) HasDataUrls() bool`

HasDataUrls returns a boolean if a field has been set.

### GetDescription

`func (o *CloudObjectVideo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudObjectVideo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudObjectVideo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudObjectVideo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudObjectVideo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudObjectVideo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudObjectVideo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudObjectVideo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *CloudObjectVideo) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *CloudObjectVideo) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *CloudObjectVideo) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *CloudObjectVideo) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetEditMode

`func (o *CloudObjectVideo) GetEditMode() string`

GetEditMode returns the EditMode field if non-nil, zero value otherwise.

### GetEditModeOk

`func (o *CloudObjectVideo) GetEditModeOk() (*string, bool)`

GetEditModeOk returns a tuple with the EditMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditMode

`func (o *CloudObjectVideo) SetEditMode(v string)`

SetEditMode sets EditMode field to given value.

### HasEditMode

`func (o *CloudObjectVideo) HasEditMode() bool`

HasEditMode returns a boolean if a field has been set.

### GetExcellentCount

`func (o *CloudObjectVideo) GetExcellentCount() int64`

GetExcellentCount returns the ExcellentCount field if non-nil, zero value otherwise.

### GetExcellentCountOk

`func (o *CloudObjectVideo) GetExcellentCountOk() (*int64, bool)`

GetExcellentCountOk returns a tuple with the ExcellentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcellentCount

`func (o *CloudObjectVideo) SetExcellentCount(v int64)`

SetExcellentCount sets ExcellentCount field to given value.

### HasExcellentCount

`func (o *CloudObjectVideo) HasExcellentCount() bool`

HasExcellentCount returns a boolean if a field has been set.

### GetGrade

`func (o *CloudObjectVideo) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *CloudObjectVideo) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *CloudObjectVideo) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *CloudObjectVideo) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### GetGrade2

`func (o *CloudObjectVideo) GetGrade2() string`

GetGrade2 returns the Grade2 field if non-nil, zero value otherwise.

### GetGrade2Ok

`func (o *CloudObjectVideo) GetGrade2Ok() (*string, bool)`

GetGrade2Ok returns a tuple with the Grade2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade2

`func (o *CloudObjectVideo) SetGrade2(v string)`

SetGrade2 sets Grade2 field to given value.

### HasGrade2

`func (o *CloudObjectVideo) HasGrade2() bool`

HasGrade2 returns a boolean if a field has been set.

### GetIsPublic

`func (o *CloudObjectVideo) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CloudObjectVideo) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CloudObjectVideo) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *CloudObjectVideo) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetKeywords

`func (o *CloudObjectVideo) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *CloudObjectVideo) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *CloudObjectVideo) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *CloudObjectVideo) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLabelCount

`func (o *CloudObjectVideo) GetLabelCount() int64`

GetLabelCount returns the LabelCount field if non-nil, zero value otherwise.

### GetLabelCountOk

`func (o *CloudObjectVideo) GetLabelCountOk() (*int64, bool)`

GetLabelCountOk returns a tuple with the LabelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCount

`func (o *CloudObjectVideo) SetLabelCount(v int64)`

SetLabelCount sets LabelCount field to given value.

### HasLabelCount

`func (o *CloudObjectVideo) HasLabelCount() bool`

HasLabelCount returns a boolean if a field has been set.

### GetLabels

`func (o *CloudObjectVideo) GetLabels() []CloudObjectLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CloudObjectVideo) GetLabelsOk() (*[]CloudObjectLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CloudObjectVideo) SetLabels(v []CloudObjectLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CloudObjectVideo) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLesson

`func (o *CloudObjectVideo) GetLesson() string`

GetLesson returns the Lesson field if non-nil, zero value otherwise.

### GetLessonOk

`func (o *CloudObjectVideo) GetLessonOk() (*string, bool)`

GetLessonOk returns a tuple with the Lesson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLesson

`func (o *CloudObjectVideo) SetLesson(v string)`

SetLesson sets Lesson field to given value.

### HasLesson

`func (o *CloudObjectVideo) HasLesson() bool`

HasLesson returns a boolean if a field has been set.

### GetName

`func (o *CloudObjectVideo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectVideo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectVideo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudObjectVideo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *CloudObjectVideo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectVideo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectVideo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudObjectVideo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPlayAuth

`func (o *CloudObjectVideo) GetPlayAuth() string`

GetPlayAuth returns the PlayAuth field if non-nil, zero value otherwise.

### GetPlayAuthOk

`func (o *CloudObjectVideo) GetPlayAuthOk() (*string, bool)`

GetPlayAuthOk returns a tuple with the PlayAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayAuth

`func (o *CloudObjectVideo) SetPlayAuth(v string)`

SetPlayAuth sets PlayAuth field to given value.

### HasPlayAuth

`func (o *CloudObjectVideo) HasPlayAuth() bool`

HasPlayAuth returns a boolean if a field has been set.

### GetRemarks

`func (o *CloudObjectVideo) GetRemarks() []CloudObjectRemark`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *CloudObjectVideo) GetRemarksOk() (*[]CloudObjectRemark, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *CloudObjectVideo) SetRemarks(v []CloudObjectRemark)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *CloudObjectVideo) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetRemarks2

`func (o *CloudObjectVideo) GetRemarks2() []CloudObjectRemark`

GetRemarks2 returns the Remarks2 field if non-nil, zero value otherwise.

### GetRemarks2Ok

`func (o *CloudObjectVideo) GetRemarks2Ok() (*[]CloudObjectRemark, bool)`

GetRemarks2Ok returns a tuple with the Remarks2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks2

`func (o *CloudObjectVideo) SetRemarks2(v []CloudObjectRemark)`

SetRemarks2 sets Remarks2 field to given value.

### HasRemarks2

`func (o *CloudObjectVideo) HasRemarks2() bool`

HasRemarks2 returns a boolean if a field has been set.

### GetReviewState

`func (o *CloudObjectVideo) GetReviewState() string`

GetReviewState returns the ReviewState field if non-nil, zero value otherwise.

### GetReviewStateOk

`func (o *CloudObjectVideo) GetReviewStateOk() (*string, bool)`

GetReviewStateOk returns a tuple with the ReviewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewState

`func (o *CloudObjectVideo) SetReviewState(v string)`

SetReviewState sets ReviewState field to given value.

### HasReviewState

`func (o *CloudObjectVideo) HasReviewState() bool`

HasReviewState returns a boolean if a field has been set.

### GetSchool

`func (o *CloudObjectVideo) GetSchool() string`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *CloudObjectVideo) GetSchoolOk() (*string, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *CloudObjectVideo) SetSchool(v string)`

SetSchool sets School field to given value.

### HasSchool

`func (o *CloudObjectVideo) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetSegmentCount

`func (o *CloudObjectVideo) GetSegmentCount() int64`

GetSegmentCount returns the SegmentCount field if non-nil, zero value otherwise.

### GetSegmentCountOk

`func (o *CloudObjectVideo) GetSegmentCountOk() (*int64, bool)`

GetSegmentCountOk returns a tuple with the SegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentCount

`func (o *CloudObjectVideo) SetSegmentCount(v int64)`

SetSegmentCount sets SegmentCount field to given value.

### HasSegmentCount

`func (o *CloudObjectVideo) HasSegmentCount() bool`

HasSegmentCount returns a boolean if a field has been set.

### GetSegments

`func (o *CloudObjectVideo) GetSegments() []CloudObjectLabel`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *CloudObjectVideo) GetSegmentsOk() (*[]CloudObjectLabel, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *CloudObjectVideo) SetSegments(v []CloudObjectLabel)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *CloudObjectVideo) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetStage

`func (o *CloudObjectVideo) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CloudObjectVideo) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CloudObjectVideo) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CloudObjectVideo) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetState

`func (o *CloudObjectVideo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudObjectVideo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudObjectVideo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudObjectVideo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubject

`func (o *CloudObjectVideo) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudObjectVideo) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudObjectVideo) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudObjectVideo) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTag

`func (o *CloudObjectVideo) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CloudObjectVideo) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CloudObjectVideo) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CloudObjectVideo) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTagOnPause

`func (o *CloudObjectVideo) GetTagOnPause() bool`

GetTagOnPause returns the TagOnPause field if non-nil, zero value otherwise.

### GetTagOnPauseOk

`func (o *CloudObjectVideo) GetTagOnPauseOk() (*bool, bool)`

GetTagOnPauseOk returns a tuple with the TagOnPause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagOnPause

`func (o *CloudObjectVideo) SetTagOnPause(v bool)`

SetTagOnPause sets TagOnPause field to given value.

### HasTagOnPause

`func (o *CloudObjectVideo) HasTagOnPause() bool`

HasTagOnPause returns a boolean if a field has been set.

### GetTask1

`func (o *CloudObjectVideo) GetTask1() string`

GetTask1 returns the Task1 field if non-nil, zero value otherwise.

### GetTask1Ok

`func (o *CloudObjectVideo) GetTask1Ok() (*string, bool)`

GetTask1Ok returns a tuple with the Task1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask1

`func (o *CloudObjectVideo) SetTask1(v string)`

SetTask1 sets Task1 field to given value.

### HasTask1

`func (o *CloudObjectVideo) HasTask1() bool`

HasTask1 returns a boolean if a field has been set.

### GetTask2

`func (o *CloudObjectVideo) GetTask2() string`

GetTask2 returns the Task2 field if non-nil, zero value otherwise.

### GetTask2Ok

`func (o *CloudObjectVideo) GetTask2Ok() (*string, bool)`

GetTask2Ok returns a tuple with the Task2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask2

`func (o *CloudObjectVideo) SetTask2(v string)`

SetTask2 sets Task2 field to given value.

### HasTask2

`func (o *CloudObjectVideo) HasTask2() bool`

HasTask2 returns a boolean if a field has been set.

### GetTask3

`func (o *CloudObjectVideo) GetTask3() string`

GetTask3 returns the Task3 field if non-nil, zero value otherwise.

### GetTask3Ok

`func (o *CloudObjectVideo) GetTask3Ok() (*string, bool)`

GetTask3Ok returns a tuple with the Task3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask3

`func (o *CloudObjectVideo) SetTask3(v string)`

SetTask3 sets Task3 field to given value.

### HasTask3

`func (o *CloudObjectVideo) HasTask3() bool`

HasTask3 returns a boolean if a field has been set.

### GetTemplate

`func (o *CloudObjectVideo) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CloudObjectVideo) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CloudObjectVideo) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CloudObjectVideo) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTopic

`func (o *CloudObjectVideo) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *CloudObjectVideo) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *CloudObjectVideo) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *CloudObjectVideo) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetType

`func (o *CloudObjectVideo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudObjectVideo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudObjectVideo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudObjectVideo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *CloudObjectVideo) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CloudObjectVideo) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CloudObjectVideo) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CloudObjectVideo) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetVideoId

`func (o *CloudObjectVideo) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *CloudObjectVideo) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *CloudObjectVideo) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *CloudObjectVideo) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetVideoLength

`func (o *CloudObjectVideo) GetVideoLength() string`

GetVideoLength returns the VideoLength field if non-nil, zero value otherwise.

### GetVideoLengthOk

`func (o *CloudObjectVideo) GetVideoLengthOk() (*string, bool)`

GetVideoLengthOk returns a tuple with the VideoLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLength

`func (o *CloudObjectVideo) SetVideoLength(v string)`

SetVideoLength sets VideoLength field to given value.

### HasVideoLength

`func (o *CloudObjectVideo) HasVideoLength() bool`

HasVideoLength returns a boolean if a field has been set.

### GetWordCountMap

`func (o *CloudObjectVideo) GetWordCountMap() map[string]int64`

GetWordCountMap returns the WordCountMap field if non-nil, zero value otherwise.

### GetWordCountMapOk

`func (o *CloudObjectVideo) GetWordCountMapOk() (*map[string]int64, bool)`

GetWordCountMapOk returns a tuple with the WordCountMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordCountMap

`func (o *CloudObjectVideo) SetWordCountMap(v map[string]int64)`

SetWordCountMap sets WordCountMap field to given value.

### HasWordCountMap

`func (o *CloudObjectVideo) HasWordCountMap() bool`

HasWordCountMap returns a boolean if a field has been set.

### SetWordCountMapNil

`func (o *CloudObjectVideo) SetWordCountMapNil(b bool)`

 SetWordCountMapNil sets the value for WordCountMap to be an explicit nil

### UnsetWordCountMap
`func (o *CloudObjectVideo) UnsetWordCountMap()`

UnsetWordCountMap ensures that no value is present for WordCountMap, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


