# TrainProtect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**TrainBudget**](TrainBudget.md) | Budget bounds the regression per protected suite. | [optional] 
**Capabilities** | Pointer to **[]string** | Capabilities are capability artifacts, by sha256. | [optional] 
**Distillation** | Pointer to **bool** | Distillation adds KL to the base&#39;s answers on a preservation set drawn from the suites. | [optional] 
**Projection** | Pointer to [**TrainProjection**](TrainProjection.md) | Projection projects every update off the protected directions — the suites&#39; gradient subspace at the base and each named capability&#39;s subspace — by its strength; absent is no projection. | [optional] 
**Suites** | Pointer to **[]string** | Suites are capabilities the base already has, by suite. | [optional] 

## Methods

### NewTrainProtect

`func NewTrainProtect() *TrainProtect`

NewTrainProtect instantiates a new TrainProtect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainProtectWithDefaults

`func NewTrainProtectWithDefaults() *TrainProtect`

NewTrainProtectWithDefaults instantiates a new TrainProtect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *TrainProtect) GetBudget() TrainBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *TrainProtect) GetBudgetOk() (*TrainBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *TrainProtect) SetBudget(v TrainBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *TrainProtect) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCapabilities

`func (o *TrainProtect) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *TrainProtect) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *TrainProtect) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *TrainProtect) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetDistillation

`func (o *TrainProtect) GetDistillation() bool`

GetDistillation returns the Distillation field if non-nil, zero value otherwise.

### GetDistillationOk

`func (o *TrainProtect) GetDistillationOk() (*bool, bool)`

GetDistillationOk returns a tuple with the Distillation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistillation

`func (o *TrainProtect) SetDistillation(v bool)`

SetDistillation sets Distillation field to given value.

### HasDistillation

`func (o *TrainProtect) HasDistillation() bool`

HasDistillation returns a boolean if a field has been set.

### GetProjection

`func (o *TrainProtect) GetProjection() TrainProjection`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *TrainProtect) GetProjectionOk() (*TrainProjection, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *TrainProtect) SetProjection(v TrainProjection)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *TrainProtect) HasProjection() bool`

HasProjection returns a boolean if a field has been set.

### GetSuites

`func (o *TrainProtect) GetSuites() []string`

GetSuites returns the Suites field if non-nil, zero value otherwise.

### GetSuitesOk

`func (o *TrainProtect) GetSuitesOk() (*[]string, bool)`

GetSuitesOk returns a tuple with the Suites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuites

`func (o *TrainProtect) SetSuites(v []string)`

SetSuites sets Suites field to given value.

### HasSuites

`func (o *TrainProtect) HasSuites() bool`

HasSuites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


