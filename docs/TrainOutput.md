# TrainOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind is checkpoint, lora, capability, basis or merged. | [optional] 
**Name** | Pointer to **string** | Name labels it. | [optional] 

## Methods

### NewTrainOutput

`func NewTrainOutput() *TrainOutput`

NewTrainOutput instantiates a new TrainOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainOutputWithDefaults

`func NewTrainOutputWithDefaults() *TrainOutput`

NewTrainOutputWithDefaults instantiates a new TrainOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TrainOutput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrainOutput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrainOutput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrainOutput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *TrainOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrainOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrainOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrainOutput) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


