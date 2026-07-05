# AutomationsCreateVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to [**AutomationsFlowTrigger**](AutomationsFlowTrigger.md) |  | [optional] 

## Methods

### NewAutomationsCreateVersionRequest

`func NewAutomationsCreateVersionRequest() *AutomationsCreateVersionRequest`

NewAutomationsCreateVersionRequest instantiates a new AutomationsCreateVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationsCreateVersionRequestWithDefaults

`func NewAutomationsCreateVersionRequestWithDefaults() *AutomationsCreateVersionRequest`

NewAutomationsCreateVersionRequestWithDefaults instantiates a new AutomationsCreateVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AutomationsCreateVersionRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutomationsCreateVersionRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutomationsCreateVersionRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutomationsCreateVersionRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTrigger

`func (o *AutomationsCreateVersionRequest) GetTrigger() AutomationsFlowTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *AutomationsCreateVersionRequest) GetTriggerOk() (*AutomationsFlowTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *AutomationsCreateVersionRequest) SetTrigger(v AutomationsFlowTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *AutomationsCreateVersionRequest) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


