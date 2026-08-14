# LbList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancers** | Pointer to [**[]LbView**](LbView.md) | LoadBalancers are the caller org&#39;s load balancers under their friendly names. | [optional] 

## Methods

### NewLbList

`func NewLbList() *LbList`

NewLbList instantiates a new LbList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbListWithDefaults

`func NewLbListWithDefaults() *LbList`

NewLbListWithDefaults instantiates a new LbList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancers

`func (o *LbList) GetLoadBalancers() []LbView`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *LbList) GetLoadBalancersOk() (*[]LbView, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *LbList) SetLoadBalancers(v []LbView)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *LbList) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


