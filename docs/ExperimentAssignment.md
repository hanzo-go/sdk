# ExperimentAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Experiment** | Pointer to **string** | Trial is the experiment that was evaluated. | [optional] 
**On** | Pointer to **bool** | On is false when the flag returned nothing for this subject, which means the subject is not enrolled — not an error. | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**Subject** | Pointer to **string** | Subject is the unit that was bucketed. | [optional] 
**Variant** | Pointer to **string** | Arm is the arm the subject falls in, empty when the flag enrolled it in none. | [optional] 

## Methods

### NewExperimentAssignment

`func NewExperimentAssignment() *ExperimentAssignment`

NewExperimentAssignment instantiates a new ExperimentAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentAssignmentWithDefaults

`func NewExperimentAssignmentWithDefaults() *ExperimentAssignment`

NewExperimentAssignmentWithDefaults instantiates a new ExperimentAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperiment

`func (o *ExperimentAssignment) GetExperiment() string`

GetExperiment returns the Experiment field if non-nil, zero value otherwise.

### GetExperimentOk

`func (o *ExperimentAssignment) GetExperimentOk() (*string, bool)`

GetExperimentOk returns a tuple with the Experiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiment

`func (o *ExperimentAssignment) SetExperiment(v string)`

SetExperiment sets Experiment field to given value.

### HasExperiment

`func (o *ExperimentAssignment) HasExperiment() bool`

HasExperiment returns a boolean if a field has been set.

### GetOn

`func (o *ExperimentAssignment) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *ExperimentAssignment) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *ExperimentAssignment) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *ExperimentAssignment) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetPayload

`func (o *ExperimentAssignment) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ExperimentAssignment) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ExperimentAssignment) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ExperimentAssignment) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *ExperimentAssignment) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *ExperimentAssignment) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetSubject

`func (o *ExperimentAssignment) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ExperimentAssignment) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ExperimentAssignment) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ExperimentAssignment) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVariant

`func (o *ExperimentAssignment) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ExperimentAssignment) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ExperimentAssignment) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *ExperimentAssignment) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


