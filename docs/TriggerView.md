# TriggerView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | whether it currently fires | [optional] 
**FunctionName** | Pointer to **string** | the function it calls | [optional] 
**Id** | Pointer to **string** | the trigger&#39;s handle | [optional] 
**Name** | Pointer to **string** | a human label for it | [optional] 
**Target** | Pointer to **string** | the path it calls | [optional] 
**Type** | Pointer to **string** | what kind of trigger it is; HTTP is the only one today | [optional] 

## Methods

### NewTriggerView

`func NewTriggerView() *TriggerView`

NewTriggerView instantiates a new TriggerView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerViewWithDefaults

`func NewTriggerViewWithDefaults() *TriggerView`

NewTriggerViewWithDefaults instantiates a new TriggerView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TriggerView) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TriggerView) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TriggerView) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TriggerView) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFunctionName

`func (o *TriggerView) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *TriggerView) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *TriggerView) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *TriggerView) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetId

`func (o *TriggerView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TriggerView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TriggerView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TriggerView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *TriggerView) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TriggerView) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TriggerView) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TriggerView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetType

`func (o *TriggerView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TriggerView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


