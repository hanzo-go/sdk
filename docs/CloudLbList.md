# CloudLbList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancers** | Pointer to [**[]CloudLbView**](CloudLbView.md) | LoadBalancers are the caller org&#39;s load balancers under their friendly names. | [optional] 

## Methods

### NewCloudLbList

`func NewCloudLbList() *CloudLbList`

NewCloudLbList instantiates a new CloudLbList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLbListWithDefaults

`func NewCloudLbListWithDefaults() *CloudLbList`

NewCloudLbListWithDefaults instantiates a new CloudLbList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancers

`func (o *CloudLbList) GetLoadBalancers() []CloudLbView`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *CloudLbList) GetLoadBalancersOk() (*[]CloudLbView, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *CloudLbList) SetLoadBalancers(v []CloudLbView)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *CloudLbList) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


