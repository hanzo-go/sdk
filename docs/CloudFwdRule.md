# CloudFwdRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryPort** | Pointer to **int32** | EntryPort is the port the load balancer listens on. | [optional] 
**EntryProtocol** | Pointer to **string** | EntryProtocol is the protocol the load balancer listens with (http, https, tcp). | [optional] 
**TargetPort** | Pointer to **int32** | TargetPort is the backend port traffic is forwarded to. | [optional] 
**TargetProtocol** | Pointer to **string** | TargetProtocol is the protocol used to reach the backend droplets. | [optional] 

## Methods

### NewCloudFwdRule

`func NewCloudFwdRule() *CloudFwdRule`

NewCloudFwdRule instantiates a new CloudFwdRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFwdRuleWithDefaults

`func NewCloudFwdRuleWithDefaults() *CloudFwdRule`

NewCloudFwdRuleWithDefaults instantiates a new CloudFwdRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryPort

`func (o *CloudFwdRule) GetEntryPort() int32`

GetEntryPort returns the EntryPort field if non-nil, zero value otherwise.

### GetEntryPortOk

`func (o *CloudFwdRule) GetEntryPortOk() (*int32, bool)`

GetEntryPortOk returns a tuple with the EntryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPort

`func (o *CloudFwdRule) SetEntryPort(v int32)`

SetEntryPort sets EntryPort field to given value.

### HasEntryPort

`func (o *CloudFwdRule) HasEntryPort() bool`

HasEntryPort returns a boolean if a field has been set.

### GetEntryProtocol

`func (o *CloudFwdRule) GetEntryProtocol() string`

GetEntryProtocol returns the EntryProtocol field if non-nil, zero value otherwise.

### GetEntryProtocolOk

`func (o *CloudFwdRule) GetEntryProtocolOk() (*string, bool)`

GetEntryProtocolOk returns a tuple with the EntryProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryProtocol

`func (o *CloudFwdRule) SetEntryProtocol(v string)`

SetEntryProtocol sets EntryProtocol field to given value.

### HasEntryProtocol

`func (o *CloudFwdRule) HasEntryProtocol() bool`

HasEntryProtocol returns a boolean if a field has been set.

### GetTargetPort

`func (o *CloudFwdRule) GetTargetPort() int32`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *CloudFwdRule) GetTargetPortOk() (*int32, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *CloudFwdRule) SetTargetPort(v int32)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *CloudFwdRule) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.

### GetTargetProtocol

`func (o *CloudFwdRule) GetTargetProtocol() string`

GetTargetProtocol returns the TargetProtocol field if non-nil, zero value otherwise.

### GetTargetProtocolOk

`func (o *CloudFwdRule) GetTargetProtocolOk() (*string, bool)`

GetTargetProtocolOk returns a tuple with the TargetProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProtocol

`func (o *CloudFwdRule) SetTargetProtocol(v string)`

SetTargetProtocol sets TargetProtocol field to given value.

### HasTargetProtocol

`func (o *CloudFwdRule) HasTargetProtocol() bool`

HasTargetProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


