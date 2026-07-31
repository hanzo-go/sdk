# CloudCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropletsMonthly** | Pointer to **int32** |  | [optional] 
**LoadBalancersMonthly** | Pointer to **int32** |  | [optional] 
**ReclaimableMonthly** | Pointer to **int32** |  | [optional] 
**TotalMonthly** | Pointer to **int32** |  | [optional] 
**VolumesMonthly** | Pointer to **int32** |  | [optional] 
**WastedMonthly** | Pointer to **int32** | WastedMonthly is what the fleet pays every month for provisioned-but-empty space on the volumes a kubelet actually measured.  It is NOT ReclaimableMonthly and must never be added to it. Reclaimable is money a button on this board collects, by deleting volumes proven to belong to no one. Wasted is money locked inside volumes that are IN USE and holding live data: DigitalOcean can only ever grow a volume, so collecting it means copying a database onto a smaller one. See shrinkRecipe.  It is also a LOWER BOUND — unmeasured volumes contribute nothing. | [optional] 

## Methods

### NewCloudCost

`func NewCloudCost() *CloudCost`

NewCloudCost instantiates a new CloudCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCostWithDefaults

`func NewCloudCostWithDefaults() *CloudCost`

NewCloudCostWithDefaults instantiates a new CloudCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropletsMonthly

`func (o *CloudCost) GetDropletsMonthly() int32`

GetDropletsMonthly returns the DropletsMonthly field if non-nil, zero value otherwise.

### GetDropletsMonthlyOk

`func (o *CloudCost) GetDropletsMonthlyOk() (*int32, bool)`

GetDropletsMonthlyOk returns a tuple with the DropletsMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropletsMonthly

`func (o *CloudCost) SetDropletsMonthly(v int32)`

SetDropletsMonthly sets DropletsMonthly field to given value.

### HasDropletsMonthly

`func (o *CloudCost) HasDropletsMonthly() bool`

HasDropletsMonthly returns a boolean if a field has been set.

### GetLoadBalancersMonthly

`func (o *CloudCost) GetLoadBalancersMonthly() int32`

GetLoadBalancersMonthly returns the LoadBalancersMonthly field if non-nil, zero value otherwise.

### GetLoadBalancersMonthlyOk

`func (o *CloudCost) GetLoadBalancersMonthlyOk() (*int32, bool)`

GetLoadBalancersMonthlyOk returns a tuple with the LoadBalancersMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancersMonthly

`func (o *CloudCost) SetLoadBalancersMonthly(v int32)`

SetLoadBalancersMonthly sets LoadBalancersMonthly field to given value.

### HasLoadBalancersMonthly

`func (o *CloudCost) HasLoadBalancersMonthly() bool`

HasLoadBalancersMonthly returns a boolean if a field has been set.

### GetReclaimableMonthly

`func (o *CloudCost) GetReclaimableMonthly() int32`

GetReclaimableMonthly returns the ReclaimableMonthly field if non-nil, zero value otherwise.

### GetReclaimableMonthlyOk

`func (o *CloudCost) GetReclaimableMonthlyOk() (*int32, bool)`

GetReclaimableMonthlyOk returns a tuple with the ReclaimableMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimableMonthly

`func (o *CloudCost) SetReclaimableMonthly(v int32)`

SetReclaimableMonthly sets ReclaimableMonthly field to given value.

### HasReclaimableMonthly

`func (o *CloudCost) HasReclaimableMonthly() bool`

HasReclaimableMonthly returns a boolean if a field has been set.

### GetTotalMonthly

`func (o *CloudCost) GetTotalMonthly() int32`

GetTotalMonthly returns the TotalMonthly field if non-nil, zero value otherwise.

### GetTotalMonthlyOk

`func (o *CloudCost) GetTotalMonthlyOk() (*int32, bool)`

GetTotalMonthlyOk returns a tuple with the TotalMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMonthly

`func (o *CloudCost) SetTotalMonthly(v int32)`

SetTotalMonthly sets TotalMonthly field to given value.

### HasTotalMonthly

`func (o *CloudCost) HasTotalMonthly() bool`

HasTotalMonthly returns a boolean if a field has been set.

### GetVolumesMonthly

`func (o *CloudCost) GetVolumesMonthly() int32`

GetVolumesMonthly returns the VolumesMonthly field if non-nil, zero value otherwise.

### GetVolumesMonthlyOk

`func (o *CloudCost) GetVolumesMonthlyOk() (*int32, bool)`

GetVolumesMonthlyOk returns a tuple with the VolumesMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesMonthly

`func (o *CloudCost) SetVolumesMonthly(v int32)`

SetVolumesMonthly sets VolumesMonthly field to given value.

### HasVolumesMonthly

`func (o *CloudCost) HasVolumesMonthly() bool`

HasVolumesMonthly returns a boolean if a field has been set.

### GetWastedMonthly

`func (o *CloudCost) GetWastedMonthly() int32`

GetWastedMonthly returns the WastedMonthly field if non-nil, zero value otherwise.

### GetWastedMonthlyOk

`func (o *CloudCost) GetWastedMonthlyOk() (*int32, bool)`

GetWastedMonthlyOk returns a tuple with the WastedMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWastedMonthly

`func (o *CloudCost) SetWastedMonthly(v int32)`

SetWastedMonthly sets WastedMonthly field to given value.

### HasWastedMonthly

`func (o *CloudCost) HasWastedMonthly() bool`

HasWastedMonthly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


