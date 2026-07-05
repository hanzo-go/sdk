# DoLoadBalancerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Friendly name; must match ^[a-z0-9]([a-z0-9-]{0,38}[a-z0-9])?$ | 
**Region** | **string** | DO region slug | 
**Type** | Pointer to **string** | empty → DO default (REGIONAL) | [optional] 
**Size** | Pointer to **string** | DO size slug | [optional] 
**ForwardingRules** | Pointer to [**[]DoForwardingRule**](DoForwardingRule.md) | empty → default http 80→80 | [optional] 

## Methods

### NewDoLoadBalancerCreate

`func NewDoLoadBalancerCreate(name string, region string, ) *DoLoadBalancerCreate`

NewDoLoadBalancerCreate instantiates a new DoLoadBalancerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoLoadBalancerCreateWithDefaults

`func NewDoLoadBalancerCreateWithDefaults() *DoLoadBalancerCreate`

NewDoLoadBalancerCreateWithDefaults instantiates a new DoLoadBalancerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DoLoadBalancerCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DoLoadBalancerCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DoLoadBalancerCreate) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *DoLoadBalancerCreate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DoLoadBalancerCreate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DoLoadBalancerCreate) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetType

`func (o *DoLoadBalancerCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DoLoadBalancerCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DoLoadBalancerCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DoLoadBalancerCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *DoLoadBalancerCreate) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DoLoadBalancerCreate) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DoLoadBalancerCreate) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *DoLoadBalancerCreate) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetForwardingRules

`func (o *DoLoadBalancerCreate) GetForwardingRules() []DoForwardingRule`

GetForwardingRules returns the ForwardingRules field if non-nil, zero value otherwise.

### GetForwardingRulesOk

`func (o *DoLoadBalancerCreate) GetForwardingRulesOk() (*[]DoForwardingRule, bool)`

GetForwardingRulesOk returns a tuple with the ForwardingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingRules

`func (o *DoLoadBalancerCreate) SetForwardingRules(v []DoForwardingRule)`

SetForwardingRules sets ForwardingRules field to given value.

### HasForwardingRules

`func (o *DoLoadBalancerCreate) HasForwardingRules() bool`

HasForwardingRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


