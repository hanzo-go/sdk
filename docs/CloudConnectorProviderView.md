# CloudConnectorProviderView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category groups the card. | [optional] 
**Description** | Pointer to **string** | Description is the one-line pitch the console card shows. | [optional] 
**Id** | Pointer to **string** | ID is the provider&#39;s registry id and the :provider path segment. | [optional] 
**Methods** | Pointer to **[]string** | Methods are the intake paths this provider supports, derived from its capabilities: \&quot;device\&quot;, \&quot;oauth\&quot; (adopt an externally obtained bundle) and \&quot;token\&quot; (a customer-held credential). At least one, always. | [optional] 
**Name** | Pointer to **string** | Name is the provider&#39;s display name. | [optional] 
**Scopes** | Pointer to **[]string** | Scopes are the permissions a connection will ask for. Never null. | [optional] 

## Methods

### NewCloudConnectorProviderView

`func NewCloudConnectorProviderView() *CloudConnectorProviderView`

NewCloudConnectorProviderView instantiates a new CloudConnectorProviderView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConnectorProviderViewWithDefaults

`func NewCloudConnectorProviderViewWithDefaults() *CloudConnectorProviderView`

NewCloudConnectorProviderViewWithDefaults instantiates a new CloudConnectorProviderView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CloudConnectorProviderView) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudConnectorProviderView) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudConnectorProviderView) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudConnectorProviderView) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *CloudConnectorProviderView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudConnectorProviderView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudConnectorProviderView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudConnectorProviderView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CloudConnectorProviderView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudConnectorProviderView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudConnectorProviderView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudConnectorProviderView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethods

`func (o *CloudConnectorProviderView) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *CloudConnectorProviderView) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *CloudConnectorProviderView) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *CloudConnectorProviderView) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetName

`func (o *CloudConnectorProviderView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudConnectorProviderView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudConnectorProviderView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudConnectorProviderView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *CloudConnectorProviderView) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CloudConnectorProviderView) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CloudConnectorProviderView) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CloudConnectorProviderView) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


