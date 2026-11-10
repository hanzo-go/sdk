# TrainCreateClientIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adaptation** | Pointer to [**TrainAdaptation**](TrainAdaptation.md) | Adaptation is lora, the one mode clients run; its shape is lora_config. | [optional] 
**BaseModel** | Pointer to **string** | BaseModel is a model the engine loads: a Hugging Face repo id or a local path. | [optional] 
**LoraConfig** | Pointer to **interface{}** |  | [optional] 
**Protect** | Pointer to [**TrainProtect**](TrainProtect.md) | Protect is refused: clients protect nothing yet. | [optional] 

## Methods

### NewTrainCreateClientIn

`func NewTrainCreateClientIn() *TrainCreateClientIn`

NewTrainCreateClientIn instantiates a new TrainCreateClientIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainCreateClientInWithDefaults

`func NewTrainCreateClientInWithDefaults() *TrainCreateClientIn`

NewTrainCreateClientInWithDefaults instantiates a new TrainCreateClientIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptation

`func (o *TrainCreateClientIn) GetAdaptation() TrainAdaptation`

GetAdaptation returns the Adaptation field if non-nil, zero value otherwise.

### GetAdaptationOk

`func (o *TrainCreateClientIn) GetAdaptationOk() (*TrainAdaptation, bool)`

GetAdaptationOk returns a tuple with the Adaptation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptation

`func (o *TrainCreateClientIn) SetAdaptation(v TrainAdaptation)`

SetAdaptation sets Adaptation field to given value.

### HasAdaptation

`func (o *TrainCreateClientIn) HasAdaptation() bool`

HasAdaptation returns a boolean if a field has been set.

### GetBaseModel

`func (o *TrainCreateClientIn) GetBaseModel() string`

GetBaseModel returns the BaseModel field if non-nil, zero value otherwise.

### GetBaseModelOk

`func (o *TrainCreateClientIn) GetBaseModelOk() (*string, bool)`

GetBaseModelOk returns a tuple with the BaseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseModel

`func (o *TrainCreateClientIn) SetBaseModel(v string)`

SetBaseModel sets BaseModel field to given value.

### HasBaseModel

`func (o *TrainCreateClientIn) HasBaseModel() bool`

HasBaseModel returns a boolean if a field has been set.

### GetLoraConfig

`func (o *TrainCreateClientIn) GetLoraConfig() interface{}`

GetLoraConfig returns the LoraConfig field if non-nil, zero value otherwise.

### GetLoraConfigOk

`func (o *TrainCreateClientIn) GetLoraConfigOk() (*interface{}, bool)`

GetLoraConfigOk returns a tuple with the LoraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoraConfig

`func (o *TrainCreateClientIn) SetLoraConfig(v interface{})`

SetLoraConfig sets LoraConfig field to given value.

### HasLoraConfig

`func (o *TrainCreateClientIn) HasLoraConfig() bool`

HasLoraConfig returns a boolean if a field has been set.

### SetLoraConfigNil

`func (o *TrainCreateClientIn) SetLoraConfigNil(b bool)`

 SetLoraConfigNil sets the value for LoraConfig to be an explicit nil

### UnsetLoraConfig
`func (o *TrainCreateClientIn) UnsetLoraConfig()`

UnsetLoraConfig ensures that no value is present for LoraConfig, not even an explicit nil
### GetProtect

`func (o *TrainCreateClientIn) GetProtect() TrainProtect`

GetProtect returns the Protect field if non-nil, zero value otherwise.

### GetProtectOk

`func (o *TrainCreateClientIn) GetProtectOk() (*TrainProtect, bool)`

GetProtectOk returns a tuple with the Protect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtect

`func (o *TrainCreateClientIn) SetProtect(v TrainProtect)`

SetProtect sets Protect field to given value.

### HasProtect

`func (o *TrainCreateClientIn) HasProtect() bool`

HasProtect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


