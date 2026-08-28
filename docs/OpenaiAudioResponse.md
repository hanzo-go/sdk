# OpenaiAudioResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **float32** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Segments** | Pointer to [**[]OpenaiAudioResponseSegmentsInner**](OpenaiAudioResponseSegmentsInner.md) |  | [optional] 
**Task** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Words** | Pointer to [**[]OpenaiAudioResponseWordsInner**](OpenaiAudioResponseWordsInner.md) |  | [optional] 

## Methods

### NewOpenaiAudioResponse

`func NewOpenaiAudioResponse() *OpenaiAudioResponse`

NewOpenaiAudioResponse instantiates a new OpenaiAudioResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiAudioResponseWithDefaults

`func NewOpenaiAudioResponseWithDefaults() *OpenaiAudioResponse`

NewOpenaiAudioResponseWithDefaults instantiates a new OpenaiAudioResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *OpenaiAudioResponse) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *OpenaiAudioResponse) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *OpenaiAudioResponse) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *OpenaiAudioResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLanguage

`func (o *OpenaiAudioResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *OpenaiAudioResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *OpenaiAudioResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *OpenaiAudioResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetSegments

`func (o *OpenaiAudioResponse) GetSegments() []OpenaiAudioResponseSegmentsInner`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *OpenaiAudioResponse) GetSegmentsOk() (*[]OpenaiAudioResponseSegmentsInner, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *OpenaiAudioResponse) SetSegments(v []OpenaiAudioResponseSegmentsInner)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *OpenaiAudioResponse) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetTask

`func (o *OpenaiAudioResponse) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *OpenaiAudioResponse) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *OpenaiAudioResponse) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *OpenaiAudioResponse) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetText

`func (o *OpenaiAudioResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OpenaiAudioResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OpenaiAudioResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *OpenaiAudioResponse) HasText() bool`

HasText returns a boolean if a field has been set.

### GetWords

`func (o *OpenaiAudioResponse) GetWords() []OpenaiAudioResponseWordsInner`

GetWords returns the Words field if non-nil, zero value otherwise.

### GetWordsOk

`func (o *OpenaiAudioResponse) GetWordsOk() (*[]OpenaiAudioResponseWordsInner, bool)`

GetWordsOk returns a tuple with the Words field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWords

`func (o *OpenaiAudioResponse) SetWords(v []OpenaiAudioResponseWordsInner)`

SetWords sets Words field to given value.

### HasWords

`func (o *OpenaiAudioResponse) HasWords() bool`

HasWords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


