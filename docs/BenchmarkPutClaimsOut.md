# BenchmarkPutClaimsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** | Org is the org the rows were recorded under: the caller&#39;s own. | [optional] 
**Recorded** | Pointer to **int64** | Recorded is how many rows were written. | [optional] 
**Rejected** | Pointer to **[]string** | Rejected names the rows that were not, and why. | [optional] 
**Visibility** | Pointer to **string** | Visibility is who may read them. | [optional] 

## Methods

### NewBenchmarkPutClaimsOut

`func NewBenchmarkPutClaimsOut() *BenchmarkPutClaimsOut`

NewBenchmarkPutClaimsOut instantiates a new BenchmarkPutClaimsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkPutClaimsOutWithDefaults

`func NewBenchmarkPutClaimsOutWithDefaults() *BenchmarkPutClaimsOut`

NewBenchmarkPutClaimsOutWithDefaults instantiates a new BenchmarkPutClaimsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *BenchmarkPutClaimsOut) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BenchmarkPutClaimsOut) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BenchmarkPutClaimsOut) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BenchmarkPutClaimsOut) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRecorded

`func (o *BenchmarkPutClaimsOut) GetRecorded() int64`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *BenchmarkPutClaimsOut) GetRecordedOk() (*int64, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *BenchmarkPutClaimsOut) SetRecorded(v int64)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *BenchmarkPutClaimsOut) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetRejected

`func (o *BenchmarkPutClaimsOut) GetRejected() []string`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *BenchmarkPutClaimsOut) GetRejectedOk() (*[]string, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *BenchmarkPutClaimsOut) SetRejected(v []string)`

SetRejected sets Rejected field to given value.

### HasRejected

`func (o *BenchmarkPutClaimsOut) HasRejected() bool`

HasRejected returns a boolean if a field has been set.

### GetVisibility

`func (o *BenchmarkPutClaimsOut) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BenchmarkPutClaimsOut) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BenchmarkPutClaimsOut) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BenchmarkPutClaimsOut) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


