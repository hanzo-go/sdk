# CloudJobCancel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the job (activity) id, from the URL path. | [optional] 
**Reason** | Pointer to **string** | Reason is recorded on the cancellation; empty records \&quot;canceled from console\&quot;. | [optional] 
**Run** | Pointer to **string** | Run is the run id; empty defaults to the job id, which is what the dispatcher sets (runId &#x3D;&#x3D; activityId &#x3D;&#x3D; prompt_id), so the common case sends no body. | [optional] 

## Methods

### NewCloudJobCancel

`func NewCloudJobCancel() *CloudJobCancel`

NewCloudJobCancel instantiates a new CloudJobCancel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudJobCancelWithDefaults

`func NewCloudJobCancelWithDefaults() *CloudJobCancel`

NewCloudJobCancelWithDefaults instantiates a new CloudJobCancel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudJobCancel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudJobCancel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudJobCancel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudJobCancel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *CloudJobCancel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudJobCancel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudJobCancel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudJobCancel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRun

`func (o *CloudJobCancel) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *CloudJobCancel) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *CloudJobCancel) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *CloudJobCancel) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


