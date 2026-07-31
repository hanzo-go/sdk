# CloudPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevenueShareBps** | Pointer to **int32** | RevenueShareBps is the share of net platform revenue a sweep accrues into the reserve fund, in basis points. 0–10000; 2000 (20%) is the platform default. | [optional] 

## Methods

### NewCloudPolicyRequest

`func NewCloudPolicyRequest() *CloudPolicyRequest`

NewCloudPolicyRequest instantiates a new CloudPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPolicyRequestWithDefaults

`func NewCloudPolicyRequestWithDefaults() *CloudPolicyRequest`

NewCloudPolicyRequestWithDefaults instantiates a new CloudPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevenueShareBps

`func (o *CloudPolicyRequest) GetRevenueShareBps() int32`

GetRevenueShareBps returns the RevenueShareBps field if non-nil, zero value otherwise.

### GetRevenueShareBpsOk

`func (o *CloudPolicyRequest) GetRevenueShareBpsOk() (*int32, bool)`

GetRevenueShareBpsOk returns a tuple with the RevenueShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueShareBps

`func (o *CloudPolicyRequest) SetRevenueShareBps(v int32)`

SetRevenueShareBps sets RevenueShareBps field to given value.

### HasRevenueShareBps

`func (o *CloudPolicyRequest) HasRevenueShareBps() bool`

HasRevenueShareBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


