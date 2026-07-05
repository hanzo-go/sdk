# GatewayRoutingRuleMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path pattern (supports wildcards) | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGatewayRoutingRuleMatch

`func NewGatewayRoutingRuleMatch() *GatewayRoutingRuleMatch`

NewGatewayRoutingRuleMatch instantiates a new GatewayRoutingRuleMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRoutingRuleMatchWithDefaults

`func NewGatewayRoutingRuleMatchWithDefaults() *GatewayRoutingRuleMatch`

NewGatewayRoutingRuleMatchWithDefaults instantiates a new GatewayRoutingRuleMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *GatewayRoutingRuleMatch) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GatewayRoutingRuleMatch) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GatewayRoutingRuleMatch) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GatewayRoutingRuleMatch) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMethod

`func (o *GatewayRoutingRuleMatch) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GatewayRoutingRuleMatch) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GatewayRoutingRuleMatch) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *GatewayRoutingRuleMatch) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetHeaders

`func (o *GatewayRoutingRuleMatch) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *GatewayRoutingRuleMatch) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *GatewayRoutingRuleMatch) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *GatewayRoutingRuleMatch) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


