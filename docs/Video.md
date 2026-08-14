# Video

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
**ExcellentCount** | Pointer to **int32** |  | [optional] 
**Grade** | Pointer to **string** |  | [optional] 
**Grade2** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**LabelCount** | Pointer to **int32** |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**Lesson** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PlayAuth** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to [**[]Remark**](Remark.md) |  | [optional] 
**Remarks2** | Pointer to [**[]Remark**](Remark.md) |  | [optional] 
**ReviewState** | Pointer to **string** |  | [optional] 
**School** | Pointer to **string** |  | [optional] 
**SegmentCount** | Pointer to **int32** |  | [optional] 
**Segments** | Pointer to [**[]Label**](Label.md) |  | [optional] 
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
**WordCountMap** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewVideo

`func NewVideo() *Video`

NewVideo instantiates a new Video object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithDefaults

`func NewVideoWithDefaults() *Video`

NewVideoWithDefaults instantiates a new Video object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioUrl

`func (o *Video) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *Video) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *Video) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.

### HasAudioUrl

`func (o *Video) HasAudioUrl() bool`

HasAudioUrl returns a boolean if a field has been set.

### GetClass

`func (o *Video) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Video) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Video) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *Video) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCoverUrl

`func (o *Video) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *Video) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *Video) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *Video) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Video) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Video) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Video) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Video) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDataUrl

`func (o *Video) GetDataUrl() string`

GetDataUrl returns the DataUrl field if non-nil, zero value otherwise.

### GetDataUrlOk

`func (o *Video) GetDataUrlOk() (*string, bool)`

GetDataUrlOk returns a tuple with the DataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUrl

`func (o *Video) SetDataUrl(v string)`

SetDataUrl sets DataUrl field to given value.

### HasDataUrl

`func (o *Video) HasDataUrl() bool`

HasDataUrl returns a boolean if a field has been set.

### GetDataUrls

`func (o *Video) GetDataUrls() []string`

GetDataUrls returns the DataUrls field if non-nil, zero value otherwise.

### GetDataUrlsOk

`func (o *Video) GetDataUrlsOk() (*[]string, bool)`

GetDataUrlsOk returns a tuple with the DataUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUrls

`func (o *Video) SetDataUrls(v []string)`

SetDataUrls sets DataUrls field to given value.

### HasDataUrls

`func (o *Video) HasDataUrls() bool`

HasDataUrls returns a boolean if a field has been set.

### GetDescription

`func (o *Video) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Video) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Video) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Video) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Video) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Video) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Video) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Video) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *Video) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *Video) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *Video) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *Video) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetEditMode

`func (o *Video) GetEditMode() string`

GetEditMode returns the EditMode field if non-nil, zero value otherwise.

### GetEditModeOk

`func (o *Video) GetEditModeOk() (*string, bool)`

GetEditModeOk returns a tuple with the EditMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditMode

`func (o *Video) SetEditMode(v string)`

SetEditMode sets EditMode field to given value.

### HasEditMode

`func (o *Video) HasEditMode() bool`

HasEditMode returns a boolean if a field has been set.

### GetExcellentCount

`func (o *Video) GetExcellentCount() int32`

GetExcellentCount returns the ExcellentCount field if non-nil, zero value otherwise.

### GetExcellentCountOk

`func (o *Video) GetExcellentCountOk() (*int32, bool)`

GetExcellentCountOk returns a tuple with the ExcellentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcellentCount

`func (o *Video) SetExcellentCount(v int32)`

SetExcellentCount sets ExcellentCount field to given value.

### HasExcellentCount

`func (o *Video) HasExcellentCount() bool`

HasExcellentCount returns a boolean if a field has been set.

### GetGrade

`func (o *Video) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *Video) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *Video) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *Video) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### GetGrade2

`func (o *Video) GetGrade2() string`

GetGrade2 returns the Grade2 field if non-nil, zero value otherwise.

### GetGrade2Ok

`func (o *Video) GetGrade2Ok() (*string, bool)`

GetGrade2Ok returns a tuple with the Grade2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade2

`func (o *Video) SetGrade2(v string)`

SetGrade2 sets Grade2 field to given value.

### HasGrade2

`func (o *Video) HasGrade2() bool`

HasGrade2 returns a boolean if a field has been set.

### GetIsPublic

`func (o *Video) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Video) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Video) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *Video) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetKeywords

`func (o *Video) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *Video) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *Video) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *Video) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLabelCount

`func (o *Video) GetLabelCount() int32`

GetLabelCount returns the LabelCount field if non-nil, zero value otherwise.

### GetLabelCountOk

`func (o *Video) GetLabelCountOk() (*int32, bool)`

GetLabelCountOk returns a tuple with the LabelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCount

`func (o *Video) SetLabelCount(v int32)`

SetLabelCount sets LabelCount field to given value.

### HasLabelCount

`func (o *Video) HasLabelCount() bool`

HasLabelCount returns a boolean if a field has been set.

### GetLabels

`func (o *Video) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Video) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Video) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Video) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLesson

`func (o *Video) GetLesson() string`

GetLesson returns the Lesson field if non-nil, zero value otherwise.

### GetLessonOk

`func (o *Video) GetLessonOk() (*string, bool)`

GetLessonOk returns a tuple with the Lesson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLesson

`func (o *Video) SetLesson(v string)`

SetLesson sets Lesson field to given value.

### HasLesson

`func (o *Video) HasLesson() bool`

HasLesson returns a boolean if a field has been set.

### GetName

`func (o *Video) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Video) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Video) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Video) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *Video) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Video) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Video) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Video) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPlayAuth

`func (o *Video) GetPlayAuth() string`

GetPlayAuth returns the PlayAuth field if non-nil, zero value otherwise.

### GetPlayAuthOk

`func (o *Video) GetPlayAuthOk() (*string, bool)`

GetPlayAuthOk returns a tuple with the PlayAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayAuth

`func (o *Video) SetPlayAuth(v string)`

SetPlayAuth sets PlayAuth field to given value.

### HasPlayAuth

`func (o *Video) HasPlayAuth() bool`

HasPlayAuth returns a boolean if a field has been set.

### GetRemarks

`func (o *Video) GetRemarks() []Remark`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *Video) GetRemarksOk() (*[]Remark, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *Video) SetRemarks(v []Remark)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *Video) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetRemarks2

`func (o *Video) GetRemarks2() []Remark`

GetRemarks2 returns the Remarks2 field if non-nil, zero value otherwise.

### GetRemarks2Ok

`func (o *Video) GetRemarks2Ok() (*[]Remark, bool)`

GetRemarks2Ok returns a tuple with the Remarks2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks2

`func (o *Video) SetRemarks2(v []Remark)`

SetRemarks2 sets Remarks2 field to given value.

### HasRemarks2

`func (o *Video) HasRemarks2() bool`

HasRemarks2 returns a boolean if a field has been set.

### GetReviewState

`func (o *Video) GetReviewState() string`

GetReviewState returns the ReviewState field if non-nil, zero value otherwise.

### GetReviewStateOk

`func (o *Video) GetReviewStateOk() (*string, bool)`

GetReviewStateOk returns a tuple with the ReviewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewState

`func (o *Video) SetReviewState(v string)`

SetReviewState sets ReviewState field to given value.

### HasReviewState

`func (o *Video) HasReviewState() bool`

HasReviewState returns a boolean if a field has been set.

### GetSchool

`func (o *Video) GetSchool() string`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *Video) GetSchoolOk() (*string, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *Video) SetSchool(v string)`

SetSchool sets School field to given value.

### HasSchool

`func (o *Video) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetSegmentCount

`func (o *Video) GetSegmentCount() int32`

GetSegmentCount returns the SegmentCount field if non-nil, zero value otherwise.

### GetSegmentCountOk

`func (o *Video) GetSegmentCountOk() (*int32, bool)`

GetSegmentCountOk returns a tuple with the SegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentCount

`func (o *Video) SetSegmentCount(v int32)`

SetSegmentCount sets SegmentCount field to given value.

### HasSegmentCount

`func (o *Video) HasSegmentCount() bool`

HasSegmentCount returns a boolean if a field has been set.

### GetSegments

`func (o *Video) GetSegments() []Label`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *Video) GetSegmentsOk() (*[]Label, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *Video) SetSegments(v []Label)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *Video) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetStage

`func (o *Video) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *Video) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *Video) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *Video) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetState

`func (o *Video) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Video) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Video) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Video) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubject

`func (o *Video) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Video) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Video) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Video) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTag

`func (o *Video) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Video) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Video) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Video) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTagOnPause

`func (o *Video) GetTagOnPause() bool`

GetTagOnPause returns the TagOnPause field if non-nil, zero value otherwise.

### GetTagOnPauseOk

`func (o *Video) GetTagOnPauseOk() (*bool, bool)`

GetTagOnPauseOk returns a tuple with the TagOnPause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagOnPause

`func (o *Video) SetTagOnPause(v bool)`

SetTagOnPause sets TagOnPause field to given value.

### HasTagOnPause

`func (o *Video) HasTagOnPause() bool`

HasTagOnPause returns a boolean if a field has been set.

### GetTask1

`func (o *Video) GetTask1() string`

GetTask1 returns the Task1 field if non-nil, zero value otherwise.

### GetTask1Ok

`func (o *Video) GetTask1Ok() (*string, bool)`

GetTask1Ok returns a tuple with the Task1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask1

`func (o *Video) SetTask1(v string)`

SetTask1 sets Task1 field to given value.

### HasTask1

`func (o *Video) HasTask1() bool`

HasTask1 returns a boolean if a field has been set.

### GetTask2

`func (o *Video) GetTask2() string`

GetTask2 returns the Task2 field if non-nil, zero value otherwise.

### GetTask2Ok

`func (o *Video) GetTask2Ok() (*string, bool)`

GetTask2Ok returns a tuple with the Task2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask2

`func (o *Video) SetTask2(v string)`

SetTask2 sets Task2 field to given value.

### HasTask2

`func (o *Video) HasTask2() bool`

HasTask2 returns a boolean if a field has been set.

### GetTask3

`func (o *Video) GetTask3() string`

GetTask3 returns the Task3 field if non-nil, zero value otherwise.

### GetTask3Ok

`func (o *Video) GetTask3Ok() (*string, bool)`

GetTask3Ok returns a tuple with the Task3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask3

`func (o *Video) SetTask3(v string)`

SetTask3 sets Task3 field to given value.

### HasTask3

`func (o *Video) HasTask3() bool`

HasTask3 returns a boolean if a field has been set.

### GetTemplate

`func (o *Video) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Video) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Video) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Video) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTopic

`func (o *Video) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *Video) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *Video) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *Video) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetType

`func (o *Video) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Video) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Video) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Video) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *Video) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Video) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Video) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Video) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetVideoId

`func (o *Video) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *Video) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *Video) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *Video) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetVideoLength

`func (o *Video) GetVideoLength() string`

GetVideoLength returns the VideoLength field if non-nil, zero value otherwise.

### GetVideoLengthOk

`func (o *Video) GetVideoLengthOk() (*string, bool)`

GetVideoLengthOk returns a tuple with the VideoLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLength

`func (o *Video) SetVideoLength(v string)`

SetVideoLength sets VideoLength field to given value.

### HasVideoLength

`func (o *Video) HasVideoLength() bool`

HasVideoLength returns a boolean if a field has been set.

### GetWordCountMap

`func (o *Video) GetWordCountMap() map[string]int32`

GetWordCountMap returns the WordCountMap field if non-nil, zero value otherwise.

### GetWordCountMapOk

`func (o *Video) GetWordCountMapOk() (*map[string]int32, bool)`

GetWordCountMapOk returns a tuple with the WordCountMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordCountMap

`func (o *Video) SetWordCountMap(v map[string]int32)`

SetWordCountMap sets WordCountMap field to given value.

### HasWordCountMap

`func (o *Video) HasWordCountMap() bool`

HasWordCountMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


