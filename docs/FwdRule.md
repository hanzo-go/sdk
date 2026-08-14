# FwdRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryPort** | Pointer to **int32** | EntryPort is the port the load balancer listens on. | [optional] 
**EntryProtocol** | Pointer to **string** | EntryProtocol is the protocol the load balancer listens with (http, https, tcp). | [optional] 
**TargetPort** | Pointer to **int32** | TargetPort is the backend port traffic is forwarded to. | [optional] 
**TargetProtocol** | Pointer to **string** | TargetProtocol is the protocol used to reach the backend droplets. | [optional] 

## Methods

### NewFwdRule

`func NewFwdRule() *FwdRule`

NewFwdRule instantiates a new FwdRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwdRuleWithDefaults

`func NewFwdRuleWithDefaults() *FwdRule`

NewFwdRuleWithDefaults instantiates a new FwdRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryPort

`func (o *FwdRule) GetEntryPort() int32`

GetEntryPort returns the EntryPort field if non-nil, zero value otherwise.

### GetEntryPortOk

`func (o *FwdRule) GetEntryPortOk() (*int32, bool)`

GetEntryPortOk returns a tuple with the EntryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPort

`func (o *FwdRule) SetEntryPort(v int32)`

SetEntryPort sets EntryPort field to given value.

### HasEntryPort

`func (o *FwdRule) HasEntryPort() bool`

HasEntryPort returns a boolean if a field has been set.

### GetEntryProtocol

`func (o *FwdRule) GetEntryProtocol() string`

GetEntryProtocol returns the EntryProtocol field if non-nil, zero value otherwise.

### GetEntryProtocolOk

`func (o *FwdRule) GetEntryProtocolOk() (*string, bool)`

GetEntryProtocolOk returns a tuple with the EntryProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryProtocol

`func (o *FwdRule) SetEntryProtocol(v string)`

SetEntryProtocol sets EntryProtocol field to given value.

### HasEntryProtocol

`func (o *FwdRule) HasEntryProtocol() bool`

HasEntryProtocol returns a boolean if a field has been set.

### GetTargetPort

`func (o *FwdRule) GetTargetPort() int32`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *FwdRule) GetTargetPortOk() (*int32, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *FwdRule) SetTargetPort(v int32)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *FwdRule) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.

### GetTargetProtocol

`func (o *FwdRule) GetTargetProtocol() string`

GetTargetProtocol returns the TargetProtocol field if non-nil, zero value otherwise.

### GetTargetProtocolOk

`func (o *FwdRule) GetTargetProtocolOk() (*string, bool)`

GetTargetProtocolOk returns a tuple with the TargetProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProtocol

`func (o *FwdRule) SetTargetProtocol(v string)`

SetTargetProtocol sets TargetProtocol field to given value.

### HasTargetProtocol

`func (o *FwdRule) HasTargetProtocol() bool`

HasTargetProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


