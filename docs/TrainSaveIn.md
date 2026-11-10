# TrainSaveIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID names the client. | [optional] 
**Name** | Pointer to **string** | Name is the adapter&#39;s name. | [optional] 
**Output** | Pointer to [**TrainOutput**](TrainOutput.md) | Output is lora, the one kind a client writes. | [optional] 

## Methods

### NewTrainSaveIn

`func NewTrainSaveIn() *TrainSaveIn`

NewTrainSaveIn instantiates a new TrainSaveIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainSaveInWithDefaults

`func NewTrainSaveInWithDefaults() *TrainSaveIn`

NewTrainSaveInWithDefaults instantiates a new TrainSaveIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrainSaveIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainSaveIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainSaveIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainSaveIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TrainSaveIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrainSaveIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrainSaveIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrainSaveIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *TrainSaveIn) GetOutput() TrainOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TrainSaveIn) GetOutputOk() (*TrainOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TrainSaveIn) SetOutput(v TrainOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TrainSaveIn) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


