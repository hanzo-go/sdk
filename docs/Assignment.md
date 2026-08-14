# Assignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Experiment** | Pointer to **string** | Trial is the experiment that was evaluated. | [optional] 
**On** | Pointer to **bool** | On is false when the flag returned nothing for this subject, which means the subject is not enrolled — not an error. | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**Subject** | Pointer to **string** | Subject is the unit that was bucketed. | [optional] 
**Variant** | Pointer to **string** | Arm is the arm the subject falls in, empty when the flag enrolled it in none. | [optional] 

## Methods

### NewAssignment

`func NewAssignment() *Assignment`

NewAssignment instantiates a new Assignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentWithDefaults

`func NewAssignmentWithDefaults() *Assignment`

NewAssignmentWithDefaults instantiates a new Assignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperiment

`func (o *Assignment) GetExperiment() string`

GetExperiment returns the Experiment field if non-nil, zero value otherwise.

### GetExperimentOk

`func (o *Assignment) GetExperimentOk() (*string, bool)`

GetExperimentOk returns a tuple with the Experiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiment

`func (o *Assignment) SetExperiment(v string)`

SetExperiment sets Experiment field to given value.

### HasExperiment

`func (o *Assignment) HasExperiment() bool`

HasExperiment returns a boolean if a field has been set.

### GetOn

`func (o *Assignment) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *Assignment) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *Assignment) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *Assignment) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetPayload

`func (o *Assignment) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Assignment) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Assignment) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Assignment) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *Assignment) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *Assignment) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetSubject

`func (o *Assignment) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Assignment) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Assignment) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Assignment) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVariant

`func (o *Assignment) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *Assignment) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *Assignment) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *Assignment) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


