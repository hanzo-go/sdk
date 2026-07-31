# CloudConnectorAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Props** | Pointer to [**[]CloudPropSpec**](CloudPropSpec.md) |  | [optional] 

## Methods

### NewCloudConnectorAction

`func NewCloudConnectorAction() *CloudConnectorAction`

NewCloudConnectorAction instantiates a new CloudConnectorAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConnectorActionWithDefaults

`func NewCloudConnectorActionWithDefaults() *CloudConnectorAction`

NewCloudConnectorActionWithDefaults instantiates a new CloudConnectorAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudConnectorAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudConnectorAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudConnectorAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudConnectorAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudConnectorAction) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudConnectorAction) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudConnectorAction) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudConnectorAction) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *CloudConnectorAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudConnectorAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudConnectorAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudConnectorAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProps

`func (o *CloudConnectorAction) GetProps() []CloudPropSpec`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *CloudConnectorAction) GetPropsOk() (*[]CloudPropSpec, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *CloudConnectorAction) SetProps(v []CloudPropSpec)`

SetProps sets Props field to given value.

### HasProps

`func (o *CloudConnectorAction) HasProps() bool`

HasProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


