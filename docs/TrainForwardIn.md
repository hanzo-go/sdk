# TrainForwardIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]interface{}** | Data are the examples: {prompt, completion} text, or tokenized datums. | [optional] 
**Id** | Pointer to **string** | ID names the client. | [optional] 

## Methods

### NewTrainForwardIn

`func NewTrainForwardIn() *TrainForwardIn`

NewTrainForwardIn instantiates a new TrainForwardIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainForwardInWithDefaults

`func NewTrainForwardInWithDefaults() *TrainForwardIn`

NewTrainForwardInWithDefaults instantiates a new TrainForwardIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TrainForwardIn) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrainForwardIn) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrainForwardIn) SetData(v []interface{})`

SetData sets Data field to given value.

### HasData

`func (o *TrainForwardIn) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *TrainForwardIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainForwardIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainForwardIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainForwardIn) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


