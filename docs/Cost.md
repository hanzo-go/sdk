# Cost

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

### NewCost

`func NewCost() *Cost`

NewCost instantiates a new Cost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostWithDefaults

`func NewCostWithDefaults() *Cost`

NewCostWithDefaults instantiates a new Cost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropletsMonthly

`func (o *Cost) GetDropletsMonthly() int32`

GetDropletsMonthly returns the DropletsMonthly field if non-nil, zero value otherwise.

### GetDropletsMonthlyOk

`func (o *Cost) GetDropletsMonthlyOk() (*int32, bool)`

GetDropletsMonthlyOk returns a tuple with the DropletsMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropletsMonthly

`func (o *Cost) SetDropletsMonthly(v int32)`

SetDropletsMonthly sets DropletsMonthly field to given value.

### HasDropletsMonthly

`func (o *Cost) HasDropletsMonthly() bool`

HasDropletsMonthly returns a boolean if a field has been set.

### GetLoadBalancersMonthly

`func (o *Cost) GetLoadBalancersMonthly() int32`

GetLoadBalancersMonthly returns the LoadBalancersMonthly field if non-nil, zero value otherwise.

### GetLoadBalancersMonthlyOk

`func (o *Cost) GetLoadBalancersMonthlyOk() (*int32, bool)`

GetLoadBalancersMonthlyOk returns a tuple with the LoadBalancersMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancersMonthly

`func (o *Cost) SetLoadBalancersMonthly(v int32)`

SetLoadBalancersMonthly sets LoadBalancersMonthly field to given value.

### HasLoadBalancersMonthly

`func (o *Cost) HasLoadBalancersMonthly() bool`

HasLoadBalancersMonthly returns a boolean if a field has been set.

### GetReclaimableMonthly

`func (o *Cost) GetReclaimableMonthly() int32`

GetReclaimableMonthly returns the ReclaimableMonthly field if non-nil, zero value otherwise.

### GetReclaimableMonthlyOk

`func (o *Cost) GetReclaimableMonthlyOk() (*int32, bool)`

GetReclaimableMonthlyOk returns a tuple with the ReclaimableMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimableMonthly

`func (o *Cost) SetReclaimableMonthly(v int32)`

SetReclaimableMonthly sets ReclaimableMonthly field to given value.

### HasReclaimableMonthly

`func (o *Cost) HasReclaimableMonthly() bool`

HasReclaimableMonthly returns a boolean if a field has been set.

### GetTotalMonthly

`func (o *Cost) GetTotalMonthly() int32`

GetTotalMonthly returns the TotalMonthly field if non-nil, zero value otherwise.

### GetTotalMonthlyOk

`func (o *Cost) GetTotalMonthlyOk() (*int32, bool)`

GetTotalMonthlyOk returns a tuple with the TotalMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMonthly

`func (o *Cost) SetTotalMonthly(v int32)`

SetTotalMonthly sets TotalMonthly field to given value.

### HasTotalMonthly

`func (o *Cost) HasTotalMonthly() bool`

HasTotalMonthly returns a boolean if a field has been set.

### GetVolumesMonthly

`func (o *Cost) GetVolumesMonthly() int32`

GetVolumesMonthly returns the VolumesMonthly field if non-nil, zero value otherwise.

### GetVolumesMonthlyOk

`func (o *Cost) GetVolumesMonthlyOk() (*int32, bool)`

GetVolumesMonthlyOk returns a tuple with the VolumesMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesMonthly

`func (o *Cost) SetVolumesMonthly(v int32)`

SetVolumesMonthly sets VolumesMonthly field to given value.

### HasVolumesMonthly

`func (o *Cost) HasVolumesMonthly() bool`

HasVolumesMonthly returns a boolean if a field has been set.

### GetWastedMonthly

`func (o *Cost) GetWastedMonthly() int32`

GetWastedMonthly returns the WastedMonthly field if non-nil, zero value otherwise.

### GetWastedMonthlyOk

`func (o *Cost) GetWastedMonthlyOk() (*int32, bool)`

GetWastedMonthlyOk returns a tuple with the WastedMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWastedMonthly

`func (o *Cost) SetWastedMonthly(v int32)`

SetWastedMonthly sets WastedMonthly field to given value.

### HasWastedMonthly

`func (o *Cost) HasWastedMonthly() bool`

HasWastedMonthly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


