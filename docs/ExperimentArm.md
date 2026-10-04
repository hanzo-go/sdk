# ExperimentArm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Control** | Pointer to **bool** | true on the baseline arm every other arm is compared to | [optional] 
**Key** | Pointer to **string** | the arm&#39;s slug, unique within the experiment | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**Weight** | Pointer to **float64** | its share of the rollout; the arms sum to 100 | [optional] 

## Methods

### NewExperimentArm

`func NewExperimentArm() *ExperimentArm`

NewExperimentArm instantiates a new ExperimentArm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentArmWithDefaults

`func NewExperimentArmWithDefaults() *ExperimentArm`

NewExperimentArmWithDefaults instantiates a new ExperimentArm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControl

`func (o *ExperimentArm) GetControl() bool`

GetControl returns the Control field if non-nil, zero value otherwise.

### GetControlOk

`func (o *ExperimentArm) GetControlOk() (*bool, bool)`

GetControlOk returns a tuple with the Control field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControl

`func (o *ExperimentArm) SetControl(v bool)`

SetControl sets Control field to given value.

### HasControl

`func (o *ExperimentArm) HasControl() bool`

HasControl returns a boolean if a field has been set.

### GetKey

`func (o *ExperimentArm) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExperimentArm) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExperimentArm) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ExperimentArm) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPayload

`func (o *ExperimentArm) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ExperimentArm) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ExperimentArm) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ExperimentArm) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *ExperimentArm) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *ExperimentArm) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetWeight

`func (o *ExperimentArm) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ExperimentArm) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ExperimentArm) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ExperimentArm) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


