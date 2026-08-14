# CordonIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cordon** | Pointer to **bool** | Cordon true marks the node unschedulable; false restores it. | [optional] 
**Drain** | Pointer to **bool** | Drain additionally evicts the pods already running there. | [optional] 
**Id** | Pointer to **string** | ID is the node&#39;s droplet id, from the path. | [optional] 

## Methods

### NewCordonIn

`func NewCordonIn() *CordonIn`

NewCordonIn instantiates a new CordonIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCordonInWithDefaults

`func NewCordonInWithDefaults() *CordonIn`

NewCordonInWithDefaults instantiates a new CordonIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCordon

`func (o *CordonIn) GetCordon() bool`

GetCordon returns the Cordon field if non-nil, zero value otherwise.

### GetCordonOk

`func (o *CordonIn) GetCordonOk() (*bool, bool)`

GetCordonOk returns a tuple with the Cordon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCordon

`func (o *CordonIn) SetCordon(v bool)`

SetCordon sets Cordon field to given value.

### HasCordon

`func (o *CordonIn) HasCordon() bool`

HasCordon returns a boolean if a field has been set.

### GetDrain

`func (o *CordonIn) GetDrain() bool`

GetDrain returns the Drain field if non-nil, zero value otherwise.

### GetDrainOk

`func (o *CordonIn) GetDrainOk() (*bool, bool)`

GetDrainOk returns a tuple with the Drain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrain

`func (o *CordonIn) SetDrain(v bool)`

SetDrain sets Drain field to given value.

### HasDrain

`func (o *CordonIn) HasDrain() bool`

HasDrain returns a boolean if a field has been set.

### GetId

`func (o *CordonIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CordonIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CordonIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CordonIn) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


