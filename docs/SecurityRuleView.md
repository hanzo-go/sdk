# SecurityRuleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | what kind of secret this rule recognises | [optional] 
**Id** | Pointer to **string** | the rule identifier a finding cites | [optional] 
**Name** | Pointer to **string** | the rule&#39;s human name | [optional] 
**Severity** | Pointer to **string** | how serious a match is: critical, high, medium or low | [optional] 

## Methods

### NewSecurityRuleView

`func NewSecurityRuleView() *SecurityRuleView`

NewSecurityRuleView instantiates a new SecurityRuleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRuleViewWithDefaults

`func NewSecurityRuleViewWithDefaults() *SecurityRuleView`

NewSecurityRuleViewWithDefaults instantiates a new SecurityRuleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecurityRuleView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityRuleView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityRuleView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityRuleView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SecurityRuleView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityRuleView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityRuleView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityRuleView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityRuleView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityRuleView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityRuleView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityRuleView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *SecurityRuleView) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SecurityRuleView) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SecurityRuleView) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SecurityRuleView) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


