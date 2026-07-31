# CloudLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedReason** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**Deletable** | Pointer to **bool** |  | [optional] 
**Droplets** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**MonthlyCents** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** | Service is the &#x60;namespace/name&#x60; of the live type&#x3D;LoadBalancer Service that claims this load balancer, proven from the cluster scan. Non-empty means IN USE. | [optional] 
**SizeUnit** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudLoadBalancer

`func NewCloudLoadBalancer() *CloudLoadBalancer`

NewCloudLoadBalancer instantiates a new CloudLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLoadBalancerWithDefaults

`func NewCloudLoadBalancerWithDefaults() *CloudLoadBalancer`

NewCloudLoadBalancerWithDefaults instantiates a new CloudLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedReason

`func (o *CloudLoadBalancer) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *CloudLoadBalancer) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *CloudLoadBalancer) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *CloudLoadBalancer) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### GetCluster

`func (o *CloudLoadBalancer) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CloudLoadBalancer) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CloudLoadBalancer) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *CloudLoadBalancer) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDeletable

`func (o *CloudLoadBalancer) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *CloudLoadBalancer) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *CloudLoadBalancer) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *CloudLoadBalancer) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDroplets

`func (o *CloudLoadBalancer) GetDroplets() int32`

GetDroplets returns the Droplets field if non-nil, zero value otherwise.

### GetDropletsOk

`func (o *CloudLoadBalancer) GetDropletsOk() (*int32, bool)`

GetDropletsOk returns a tuple with the Droplets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroplets

`func (o *CloudLoadBalancer) SetDroplets(v int32)`

SetDroplets sets Droplets field to given value.

### HasDroplets

`func (o *CloudLoadBalancer) HasDroplets() bool`

HasDroplets returns a boolean if a field has been set.

### GetId

`func (o *CloudLoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudLoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudLoadBalancer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudLoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *CloudLoadBalancer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CloudLoadBalancer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CloudLoadBalancer) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CloudLoadBalancer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMonthlyCents

`func (o *CloudLoadBalancer) GetMonthlyCents() int32`

GetMonthlyCents returns the MonthlyCents field if non-nil, zero value otherwise.

### GetMonthlyCentsOk

`func (o *CloudLoadBalancer) GetMonthlyCentsOk() (*int32, bool)`

GetMonthlyCentsOk returns a tuple with the MonthlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCents

`func (o *CloudLoadBalancer) SetMonthlyCents(v int32)`

SetMonthlyCents sets MonthlyCents field to given value.

### HasMonthlyCents

`func (o *CloudLoadBalancer) HasMonthlyCents() bool`

HasMonthlyCents returns a boolean if a field has been set.

### GetName

`func (o *CloudLoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudLoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudLoadBalancer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudLoadBalancer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CloudLoadBalancer) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudLoadBalancer) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudLoadBalancer) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudLoadBalancer) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *CloudLoadBalancer) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudLoadBalancer) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudLoadBalancer) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudLoadBalancer) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSizeUnit

`func (o *CloudLoadBalancer) GetSizeUnit() int32`

GetSizeUnit returns the SizeUnit field if non-nil, zero value otherwise.

### GetSizeUnitOk

`func (o *CloudLoadBalancer) GetSizeUnitOk() (*int32, bool)`

GetSizeUnitOk returns a tuple with the SizeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeUnit

`func (o *CloudLoadBalancer) SetSizeUnit(v int32)`

SetSizeUnit sets SizeUnit field to given value.

### HasSizeUnit

`func (o *CloudLoadBalancer) HasSizeUnit() bool`

HasSizeUnit returns a boolean if a field has been set.

### GetStatus

`func (o *CloudLoadBalancer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudLoadBalancer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudLoadBalancer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudLoadBalancer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


