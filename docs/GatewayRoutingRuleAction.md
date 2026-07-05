# GatewayRoutingRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** | Target URL or service | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**StripPrefix** | Pointer to **bool** |  | [optional] 

## Methods

### NewGatewayRoutingRuleAction

`func NewGatewayRoutingRuleAction() *GatewayRoutingRuleAction`

NewGatewayRoutingRuleAction instantiates a new GatewayRoutingRuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRoutingRuleActionWithDefaults

`func NewGatewayRoutingRuleActionWithDefaults() *GatewayRoutingRuleAction`

NewGatewayRoutingRuleActionWithDefaults instantiates a new GatewayRoutingRuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GatewayRoutingRuleAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayRoutingRuleAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayRoutingRuleAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayRoutingRuleAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTarget

`func (o *GatewayRoutingRuleAction) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GatewayRoutingRuleAction) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GatewayRoutingRuleAction) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GatewayRoutingRuleAction) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetHeaders

`func (o *GatewayRoutingRuleAction) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *GatewayRoutingRuleAction) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *GatewayRoutingRuleAction) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *GatewayRoutingRuleAction) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetStripPrefix

`func (o *GatewayRoutingRuleAction) GetStripPrefix() bool`

GetStripPrefix returns the StripPrefix field if non-nil, zero value otherwise.

### GetStripPrefixOk

`func (o *GatewayRoutingRuleAction) GetStripPrefixOk() (*bool, bool)`

GetStripPrefixOk returns a tuple with the StripPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripPrefix

`func (o *GatewayRoutingRuleAction) SetStripPrefix(v bool)`

SetStripPrefix sets StripPrefix field to given value.

### HasStripPrefix

`func (o *GatewayRoutingRuleAction) HasStripPrefix() bool`

HasStripPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


