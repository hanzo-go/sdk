# AutoConnectorAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Props** | Pointer to [**[]AutoPropSpec**](AutoPropSpec.md) |  | [optional] 

## Methods

### NewAutoConnectorAction

`func NewAutoConnectorAction() *AutoConnectorAction`

NewAutoConnectorAction instantiates a new AutoConnectorAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoConnectorActionWithDefaults

`func NewAutoConnectorActionWithDefaults() *AutoConnectorAction`

NewAutoConnectorActionWithDefaults instantiates a new AutoConnectorAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AutoConnectorAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoConnectorAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoConnectorAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoConnectorAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AutoConnectorAction) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutoConnectorAction) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutoConnectorAction) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutoConnectorAction) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *AutoConnectorAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoConnectorAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoConnectorAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoConnectorAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProps

`func (o *AutoConnectorAction) GetProps() []AutoPropSpec`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *AutoConnectorAction) GetPropsOk() (*[]AutoPropSpec, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *AutoConnectorAction) SetProps(v []AutoPropSpec)`

SetProps sets Props field to given value.

### HasProps

`func (o *AutoConnectorAction) HasProps() bool`

HasProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


