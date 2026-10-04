# CiPipelines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchedAt** | Pointer to **time.Time** |  | [optional] 
**Orgs** | Pointer to **[]string** |  | [optional] 
**Services** | Pointer to [**[]CiPipeline**](CiPipeline.md) |  | [optional] 
**SourceErr** | Pointer to **string** |  | [optional] 
**Stale** | Pointer to **bool** |  | [optional] 

## Methods

### NewCiPipelines

`func NewCiPipelines() *CiPipelines`

NewCiPipelines instantiates a new CiPipelines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiPipelinesWithDefaults

`func NewCiPipelinesWithDefaults() *CiPipelines`

NewCiPipelinesWithDefaults instantiates a new CiPipelines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFetchedAt

`func (o *CiPipelines) GetFetchedAt() time.Time`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *CiPipelines) GetFetchedAtOk() (*time.Time, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *CiPipelines) SetFetchedAt(v time.Time)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *CiPipelines) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.

### GetOrgs

`func (o *CiPipelines) GetOrgs() []string`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *CiPipelines) GetOrgsOk() (*[]string, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *CiPipelines) SetOrgs(v []string)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *CiPipelines) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetServices

`func (o *CiPipelines) GetServices() []CiPipeline`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CiPipelines) GetServicesOk() (*[]CiPipeline, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CiPipelines) SetServices(v []CiPipeline)`

SetServices sets Services field to given value.

### HasServices

`func (o *CiPipelines) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSourceErr

`func (o *CiPipelines) GetSourceErr() string`

GetSourceErr returns the SourceErr field if non-nil, zero value otherwise.

### GetSourceErrOk

`func (o *CiPipelines) GetSourceErrOk() (*string, bool)`

GetSourceErrOk returns a tuple with the SourceErr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceErr

`func (o *CiPipelines) SetSourceErr(v string)`

SetSourceErr sets SourceErr field to given value.

### HasSourceErr

`func (o *CiPipelines) HasSourceErr() bool`

HasSourceErr returns a boolean if a field has been set.

### GetStale

`func (o *CiPipelines) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *CiPipelines) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *CiPipelines) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *CiPipelines) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


