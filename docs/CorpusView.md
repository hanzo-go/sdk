# CorpusView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count is how many tactics survived every filter. | [optional] 
**Stage** | Pointer to **string** | Stage is the growth stage the tag join ran at — the org&#39;s observed stage, or the one ?stage&#x3D; previewed. | [optional] 
**Strategies** | Pointer to [**[]StrategyView**](StrategyView.md) | Strategies are the surviving tactics, in corpus authoring order. | [optional] 

## Methods

### NewCorpusView

`func NewCorpusView() *CorpusView`

NewCorpusView instantiates a new CorpusView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorpusViewWithDefaults

`func NewCorpusViewWithDefaults() *CorpusView`

NewCorpusViewWithDefaults instantiates a new CorpusView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CorpusView) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CorpusView) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CorpusView) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CorpusView) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetStage

`func (o *CorpusView) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CorpusView) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CorpusView) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CorpusView) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStrategies

`func (o *CorpusView) GetStrategies() []StrategyView`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *CorpusView) GetStrategiesOk() (*[]StrategyView, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *CorpusView) SetStrategies(v []StrategyView)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *CorpusView) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


