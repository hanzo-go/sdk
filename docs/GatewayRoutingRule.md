# GatewayRoutingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Match** | [**GatewayRoutingRuleMatch**](GatewayRoutingRuleMatch.md) |  | 
**Action** | [**GatewayRoutingRuleAction**](GatewayRoutingRuleAction.md) |  | 
**Priority** | Pointer to **int32** |  | [optional] [default to 100]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGatewayRoutingRule

`func NewGatewayRoutingRule(id string, match GatewayRoutingRuleMatch, action GatewayRoutingRuleAction, ) *GatewayRoutingRule`

NewGatewayRoutingRule instantiates a new GatewayRoutingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRoutingRuleWithDefaults

`func NewGatewayRoutingRuleWithDefaults() *GatewayRoutingRule`

NewGatewayRoutingRuleWithDefaults instantiates a new GatewayRoutingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayRoutingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayRoutingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayRoutingRule) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GatewayRoutingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayRoutingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayRoutingRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayRoutingRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMatch

`func (o *GatewayRoutingRule) GetMatch() GatewayRoutingRuleMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *GatewayRoutingRule) GetMatchOk() (*GatewayRoutingRuleMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *GatewayRoutingRule) SetMatch(v GatewayRoutingRuleMatch)`

SetMatch sets Match field to given value.


### GetAction

`func (o *GatewayRoutingRule) GetAction() GatewayRoutingRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GatewayRoutingRule) GetActionOk() (*GatewayRoutingRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GatewayRoutingRule) SetAction(v GatewayRoutingRuleAction)`

SetAction sets Action field to given value.


### GetPriority

`func (o *GatewayRoutingRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GatewayRoutingRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GatewayRoutingRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GatewayRoutingRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEnabled

`func (o *GatewayRoutingRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayRoutingRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayRoutingRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GatewayRoutingRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GatewayRoutingRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GatewayRoutingRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GatewayRoutingRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GatewayRoutingRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


