# BenchmarkAdmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmarks** | Pointer to **[]string** | Benchmarks are the catalog ids admitted. | [optional] 
**By** | Pointer to **string** | By is the verified user who asked for it. | [optional] 
**Endpoint** | Pointer to **string** | Endpoint is the caller&#39;s own endpoint the run targets. | [optional] 
**Model** | Pointer to **string** | Model is the catalog model the run targets. | [optional] 
**Note** | Pointer to **string** | Note explains what admission does and does not promise. | [optional] 
**Org** | Pointer to **string** | Org is the org the run is admitted for: the caller&#39;s verified org. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;queued\&quot;: the run is admitted, not finished. | [optional] 

## Methods

### NewBenchmarkAdmission

`func NewBenchmarkAdmission() *BenchmarkAdmission`

NewBenchmarkAdmission instantiates a new BenchmarkAdmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkAdmissionWithDefaults

`func NewBenchmarkAdmissionWithDefaults() *BenchmarkAdmission`

NewBenchmarkAdmissionWithDefaults instantiates a new BenchmarkAdmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmarks

`func (o *BenchmarkAdmission) GetBenchmarks() []string`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *BenchmarkAdmission) GetBenchmarksOk() (*[]string, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *BenchmarkAdmission) SetBenchmarks(v []string)`

SetBenchmarks sets Benchmarks field to given value.

### HasBenchmarks

`func (o *BenchmarkAdmission) HasBenchmarks() bool`

HasBenchmarks returns a boolean if a field has been set.

### GetBy

`func (o *BenchmarkAdmission) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *BenchmarkAdmission) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *BenchmarkAdmission) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *BenchmarkAdmission) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetEndpoint

`func (o *BenchmarkAdmission) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *BenchmarkAdmission) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *BenchmarkAdmission) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *BenchmarkAdmission) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetModel

`func (o *BenchmarkAdmission) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BenchmarkAdmission) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BenchmarkAdmission) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BenchmarkAdmission) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNote

`func (o *BenchmarkAdmission) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BenchmarkAdmission) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BenchmarkAdmission) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BenchmarkAdmission) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOrg

`func (o *BenchmarkAdmission) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BenchmarkAdmission) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BenchmarkAdmission) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BenchmarkAdmission) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetStatus

`func (o *BenchmarkAdmission) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BenchmarkAdmission) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BenchmarkAdmission) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BenchmarkAdmission) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


