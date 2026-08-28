# OpenaiAudioResponseSegmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgLogprob** | Pointer to **float32** |  | [optional] 
**CompressionRatio** | Pointer to **float32** |  | [optional] 
**End** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**NoSpeechProb** | Pointer to **float32** |  | [optional] 
**Seek** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **float32** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Tokens** | Pointer to **[]int32** |  | [optional] 
**Transient** | Pointer to **bool** |  | [optional] 

## Methods

### NewOpenaiAudioResponseSegmentsInner

`func NewOpenaiAudioResponseSegmentsInner() *OpenaiAudioResponseSegmentsInner`

NewOpenaiAudioResponseSegmentsInner instantiates a new OpenaiAudioResponseSegmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiAudioResponseSegmentsInnerWithDefaults

`func NewOpenaiAudioResponseSegmentsInnerWithDefaults() *OpenaiAudioResponseSegmentsInner`

NewOpenaiAudioResponseSegmentsInnerWithDefaults instantiates a new OpenaiAudioResponseSegmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgLogprob

`func (o *OpenaiAudioResponseSegmentsInner) GetAvgLogprob() float32`

GetAvgLogprob returns the AvgLogprob field if non-nil, zero value otherwise.

### GetAvgLogprobOk

`func (o *OpenaiAudioResponseSegmentsInner) GetAvgLogprobOk() (*float32, bool)`

GetAvgLogprobOk returns a tuple with the AvgLogprob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLogprob

`func (o *OpenaiAudioResponseSegmentsInner) SetAvgLogprob(v float32)`

SetAvgLogprob sets AvgLogprob field to given value.

### HasAvgLogprob

`func (o *OpenaiAudioResponseSegmentsInner) HasAvgLogprob() bool`

HasAvgLogprob returns a boolean if a field has been set.

### GetCompressionRatio

`func (o *OpenaiAudioResponseSegmentsInner) GetCompressionRatio() float32`

GetCompressionRatio returns the CompressionRatio field if non-nil, zero value otherwise.

### GetCompressionRatioOk

`func (o *OpenaiAudioResponseSegmentsInner) GetCompressionRatioOk() (*float32, bool)`

GetCompressionRatioOk returns a tuple with the CompressionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionRatio

`func (o *OpenaiAudioResponseSegmentsInner) SetCompressionRatio(v float32)`

SetCompressionRatio sets CompressionRatio field to given value.

### HasCompressionRatio

`func (o *OpenaiAudioResponseSegmentsInner) HasCompressionRatio() bool`

HasCompressionRatio returns a boolean if a field has been set.

### GetEnd

`func (o *OpenaiAudioResponseSegmentsInner) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *OpenaiAudioResponseSegmentsInner) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *OpenaiAudioResponseSegmentsInner) SetEnd(v float32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *OpenaiAudioResponseSegmentsInner) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetId

`func (o *OpenaiAudioResponseSegmentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenaiAudioResponseSegmentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenaiAudioResponseSegmentsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpenaiAudioResponseSegmentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNoSpeechProb

`func (o *OpenaiAudioResponseSegmentsInner) GetNoSpeechProb() float32`

GetNoSpeechProb returns the NoSpeechProb field if non-nil, zero value otherwise.

### GetNoSpeechProbOk

`func (o *OpenaiAudioResponseSegmentsInner) GetNoSpeechProbOk() (*float32, bool)`

GetNoSpeechProbOk returns a tuple with the NoSpeechProb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSpeechProb

`func (o *OpenaiAudioResponseSegmentsInner) SetNoSpeechProb(v float32)`

SetNoSpeechProb sets NoSpeechProb field to given value.

### HasNoSpeechProb

`func (o *OpenaiAudioResponseSegmentsInner) HasNoSpeechProb() bool`

HasNoSpeechProb returns a boolean if a field has been set.

### GetSeek

`func (o *OpenaiAudioResponseSegmentsInner) GetSeek() int32`

GetSeek returns the Seek field if non-nil, zero value otherwise.

### GetSeekOk

`func (o *OpenaiAudioResponseSegmentsInner) GetSeekOk() (*int32, bool)`

GetSeekOk returns a tuple with the Seek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeek

`func (o *OpenaiAudioResponseSegmentsInner) SetSeek(v int32)`

SetSeek sets Seek field to given value.

### HasSeek

`func (o *OpenaiAudioResponseSegmentsInner) HasSeek() bool`

HasSeek returns a boolean if a field has been set.

### GetStart

`func (o *OpenaiAudioResponseSegmentsInner) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *OpenaiAudioResponseSegmentsInner) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *OpenaiAudioResponseSegmentsInner) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *OpenaiAudioResponseSegmentsInner) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTemperature

`func (o *OpenaiAudioResponseSegmentsInner) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *OpenaiAudioResponseSegmentsInner) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *OpenaiAudioResponseSegmentsInner) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *OpenaiAudioResponseSegmentsInner) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetText

`func (o *OpenaiAudioResponseSegmentsInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OpenaiAudioResponseSegmentsInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OpenaiAudioResponseSegmentsInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *OpenaiAudioResponseSegmentsInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTokens

`func (o *OpenaiAudioResponseSegmentsInner) GetTokens() []int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *OpenaiAudioResponseSegmentsInner) GetTokensOk() (*[]int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *OpenaiAudioResponseSegmentsInner) SetTokens(v []int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *OpenaiAudioResponseSegmentsInner) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetTransient

`func (o *OpenaiAudioResponseSegmentsInner) GetTransient() bool`

GetTransient returns the Transient field if non-nil, zero value otherwise.

### GetTransientOk

`func (o *OpenaiAudioResponseSegmentsInner) GetTransientOk() (*bool, bool)`

GetTransientOk returns a tuple with the Transient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransient

`func (o *OpenaiAudioResponseSegmentsInner) SetTransient(v bool)`

SetTransient sets Transient field to given value.

### HasTransient

`func (o *OpenaiAudioResponseSegmentsInner) HasTransient() bool`

HasTransient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


