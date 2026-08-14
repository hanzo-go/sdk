# Admission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmarks** | Pointer to **[]string** | Benchmarks are the catalog ids admitted. | [optional] 
**Endpoint** | Pointer to **string** | Endpoint is the caller&#39;s own endpoint the run targets. | [optional] 
**Model** | Pointer to **string** | Model is the catalog model the run targets. | [optional] 
**Note** | Pointer to **string** | Note explains what admission does and does not promise. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;queued\&quot;: the run is admitted, not finished. | [optional] 

## Methods

### NewAdmission

`func NewAdmission() *Admission`

NewAdmission instantiates a new Admission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmissionWithDefaults

`func NewAdmissionWithDefaults() *Admission`

NewAdmissionWithDefaults instantiates a new Admission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmarks

`func (o *Admission) GetBenchmarks() []string`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *Admission) GetBenchmarksOk() (*[]string, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *Admission) SetBenchmarks(v []string)`

SetBenchmarks sets Benchmarks field to given value.

### HasBenchmarks

`func (o *Admission) HasBenchmarks() bool`

HasBenchmarks returns a boolean if a field has been set.

### GetEndpoint

`func (o *Admission) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Admission) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Admission) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Admission) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetModel

`func (o *Admission) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Admission) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Admission) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Admission) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNote

`func (o *Admission) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Admission) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Admission) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Admission) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetStatus

`func (o *Admission) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Admission) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Admission) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Admission) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


