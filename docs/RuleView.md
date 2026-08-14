# RuleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | what kind of secret this rule recognises | [optional] 
**Id** | Pointer to **string** | the rule identifier a finding cites | [optional] 
**Name** | Pointer to **string** | the rule&#39;s human name | [optional] 
**Severity** | Pointer to **string** | how serious a match is: critical, high, medium or low | [optional] 

## Methods

### NewRuleView

`func NewRuleView() *RuleView`

NewRuleView instantiates a new RuleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleViewWithDefaults

`func NewRuleViewWithDefaults() *RuleView`

NewRuleViewWithDefaults instantiates a new RuleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RuleView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *RuleView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RuleView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *RuleView) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RuleView) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RuleView) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RuleView) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


