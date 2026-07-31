# AiEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Msg** | Pointer to **string** | Human-readable failure reason; empty when &#x60;status&#x60; is &#x60;ok&#x60;. | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
**Data2** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAiEnvelope

`func NewAiEnvelope(status string, ) *AiEnvelope`

NewAiEnvelope instantiates a new AiEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEnvelopeWithDefaults

`func NewAiEnvelopeWithDefaults() *AiEnvelope`

NewAiEnvelopeWithDefaults instantiates a new AiEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AiEnvelope) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiEnvelope) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiEnvelope) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMsg

`func (o *AiEnvelope) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AiEnvelope) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AiEnvelope) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AiEnvelope) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *AiEnvelope) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AiEnvelope) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AiEnvelope) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AiEnvelope) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AiEnvelope) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AiEnvelope) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetData2

`func (o *AiEnvelope) GetData2() interface{}`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *AiEnvelope) GetData2Ok() (*interface{}, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *AiEnvelope) SetData2(v interface{})`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *AiEnvelope) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *AiEnvelope) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *AiEnvelope) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


