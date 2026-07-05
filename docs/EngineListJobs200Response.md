# EngineListJobs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]EngineJob**](EngineJob.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineListJobs200Response

`func NewEngineListJobs200Response() *EngineListJobs200Response`

NewEngineListJobs200Response instantiates a new EngineListJobs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineListJobs200ResponseWithDefaults

`func NewEngineListJobs200ResponseWithDefaults() *EngineListJobs200Response`

NewEngineListJobs200ResponseWithDefaults instantiates a new EngineListJobs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *EngineListJobs200Response) GetJobs() []EngineJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *EngineListJobs200Response) GetJobsOk() (*[]EngineJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *EngineListJobs200Response) SetJobs(v []EngineJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *EngineListJobs200Response) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetTotal

`func (o *EngineListJobs200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EngineListJobs200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EngineListJobs200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EngineListJobs200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


