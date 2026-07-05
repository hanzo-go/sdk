# IntegrationsProviderView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Available** | Pointer to **bool** | App credentials configured on this deployment | [optional] 
**Connected** | Pointer to **bool** | This org has a connection | [optional] 
**Connection** | Pointer to [**IntegrationsConnectionView**](IntegrationsConnectionView.md) |  | [optional] 

## Methods

### NewIntegrationsProviderView

`func NewIntegrationsProviderView() *IntegrationsProviderView`

NewIntegrationsProviderView instantiates a new IntegrationsProviderView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsProviderViewWithDefaults

`func NewIntegrationsProviderViewWithDefaults() *IntegrationsProviderView`

NewIntegrationsProviderViewWithDefaults instantiates a new IntegrationsProviderView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationsProviderView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationsProviderView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationsProviderView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationsProviderView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IntegrationsProviderView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationsProviderView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationsProviderView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationsProviderView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IntegrationsProviderView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationsProviderView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationsProviderView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationsProviderView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *IntegrationsProviderView) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IntegrationsProviderView) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IntegrationsProviderView) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IntegrationsProviderView) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetAvailable

`func (o *IntegrationsProviderView) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *IntegrationsProviderView) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *IntegrationsProviderView) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *IntegrationsProviderView) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetConnected

`func (o *IntegrationsProviderView) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *IntegrationsProviderView) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *IntegrationsProviderView) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *IntegrationsProviderView) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetConnection

`func (o *IntegrationsProviderView) GetConnection() IntegrationsConnectionView`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *IntegrationsProviderView) GetConnectionOk() (*IntegrationsConnectionView, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *IntegrationsProviderView) SetConnection(v IntegrationsConnectionView)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *IntegrationsProviderView) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


