# TrainSupports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adaptations** | Pointer to **[]string** | Adaptations are the modes it runs. | [optional] 
**Outputs** | Pointer to **[]string** | Outputs are the output kinds it produces. | [optional] 
**Protect** | Pointer to **[]string** | Protect are the protections it runs: projection, distillation, capabilities. | [optional] 
**Terms** | Pointer to **[]string** | Terms are the objective terms it runs. | [optional] 

## Methods

### NewTrainSupports

`func NewTrainSupports() *TrainSupports`

NewTrainSupports instantiates a new TrainSupports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainSupportsWithDefaults

`func NewTrainSupportsWithDefaults() *TrainSupports`

NewTrainSupportsWithDefaults instantiates a new TrainSupports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptations

`func (o *TrainSupports) GetAdaptations() []string`

GetAdaptations returns the Adaptations field if non-nil, zero value otherwise.

### GetAdaptationsOk

`func (o *TrainSupports) GetAdaptationsOk() (*[]string, bool)`

GetAdaptationsOk returns a tuple with the Adaptations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptations

`func (o *TrainSupports) SetAdaptations(v []string)`

SetAdaptations sets Adaptations field to given value.

### HasAdaptations

`func (o *TrainSupports) HasAdaptations() bool`

HasAdaptations returns a boolean if a field has been set.

### GetOutputs

`func (o *TrainSupports) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *TrainSupports) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *TrainSupports) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *TrainSupports) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetProtect

`func (o *TrainSupports) GetProtect() []string`

GetProtect returns the Protect field if non-nil, zero value otherwise.

### GetProtectOk

`func (o *TrainSupports) GetProtectOk() (*[]string, bool)`

GetProtectOk returns a tuple with the Protect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtect

`func (o *TrainSupports) SetProtect(v []string)`

SetProtect sets Protect field to given value.

### HasProtect

`func (o *TrainSupports) HasProtect() bool`

HasProtect returns a boolean if a field has been set.

### GetTerms

`func (o *TrainSupports) GetTerms() []string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *TrainSupports) GetTermsOk() (*[]string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *TrainSupports) SetTerms(v []string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *TrainSupports) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


