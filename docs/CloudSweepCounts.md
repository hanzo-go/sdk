# CloudSweepCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accrued** | Pointer to **int32** | Accrued is how many new royalty accruals it latched. | [optional] 
**Swept** | Pointer to **int32** | Swept is how many (author, deploying org) pairs the sweep examined. | [optional] 

## Methods

### NewCloudSweepCounts

`func NewCloudSweepCounts() *CloudSweepCounts`

NewCloudSweepCounts instantiates a new CloudSweepCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSweepCountsWithDefaults

`func NewCloudSweepCountsWithDefaults() *CloudSweepCounts`

NewCloudSweepCountsWithDefaults instantiates a new CloudSweepCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrued

`func (o *CloudSweepCounts) GetAccrued() int32`

GetAccrued returns the Accrued field if non-nil, zero value otherwise.

### GetAccruedOk

`func (o *CloudSweepCounts) GetAccruedOk() (*int32, bool)`

GetAccruedOk returns a tuple with the Accrued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrued

`func (o *CloudSweepCounts) SetAccrued(v int32)`

SetAccrued sets Accrued field to given value.

### HasAccrued

`func (o *CloudSweepCounts) HasAccrued() bool`

HasAccrued returns a boolean if a field has been set.

### GetSwept

`func (o *CloudSweepCounts) GetSwept() int32`

GetSwept returns the Swept field if non-nil, zero value otherwise.

### GetSweptOk

`func (o *CloudSweepCounts) GetSweptOk() (*int32, bool)`

GetSweptOk returns a tuple with the Swept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwept

`func (o *CloudSweepCounts) SetSwept(v int32)`

SetSwept sets Swept field to given value.

### HasSwept

`func (o *CloudSweepCounts) HasSwept() bool`

HasSwept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


