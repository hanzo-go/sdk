# CloudSweepResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credited** | Pointer to **int32** | Credited is how many of those referrals qualified on this pass and were granted their bonuses. | [optional] 
**Swept** | Pointer to **int32** | Swept is how many pending referrals were checked. | [optional] 

## Methods

### NewCloudSweepResult

`func NewCloudSweepResult() *CloudSweepResult`

NewCloudSweepResult instantiates a new CloudSweepResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSweepResultWithDefaults

`func NewCloudSweepResultWithDefaults() *CloudSweepResult`

NewCloudSweepResultWithDefaults instantiates a new CloudSweepResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredited

`func (o *CloudSweepResult) GetCredited() int32`

GetCredited returns the Credited field if non-nil, zero value otherwise.

### GetCreditedOk

`func (o *CloudSweepResult) GetCreditedOk() (*int32, bool)`

GetCreditedOk returns a tuple with the Credited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredited

`func (o *CloudSweepResult) SetCredited(v int32)`

SetCredited sets Credited field to given value.

### HasCredited

`func (o *CloudSweepResult) HasCredited() bool`

HasCredited returns a boolean if a field has been set.

### GetSwept

`func (o *CloudSweepResult) GetSwept() int32`

GetSwept returns the Swept field if non-nil, zero value otherwise.

### GetSweptOk

`func (o *CloudSweepResult) GetSweptOk() (*int32, bool)`

GetSweptOk returns a tuple with the Swept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwept

`func (o *CloudSweepResult) SetSwept(v int32)`

SetSwept sets Swept field to given value.

### HasSwept

`func (o *CloudSweepResult) HasSwept() bool`

HasSwept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


