# DoLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to **int32** | Count of attached backend droplets | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewDoLoadBalancer

`func NewDoLoadBalancer() *DoLoadBalancer`

NewDoLoadBalancer instantiates a new DoLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoLoadBalancerWithDefaults

`func NewDoLoadBalancerWithDefaults() *DoLoadBalancer`

NewDoLoadBalancerWithDefaults instantiates a new DoLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DoLoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DoLoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DoLoadBalancer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DoLoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DoLoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DoLoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DoLoadBalancer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DoLoadBalancer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DoLoadBalancer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DoLoadBalancer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DoLoadBalancer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DoLoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTargets

`func (o *DoLoadBalancer) GetTargets() int32`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *DoLoadBalancer) GetTargetsOk() (*int32, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *DoLoadBalancer) SetTargets(v int32)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *DoLoadBalancer) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetIp

`func (o *DoLoadBalancer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DoLoadBalancer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DoLoadBalancer) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DoLoadBalancer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetStatus

`func (o *DoLoadBalancer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DoLoadBalancer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DoLoadBalancer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DoLoadBalancer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


