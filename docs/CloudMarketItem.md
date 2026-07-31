# CloudMarketItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Dispatchable** | Pointer to **bool** |  | [optional] 
**InputSchema** | Pointer to **interface{}** |  | [optional] 
**Installed** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to [**CloudPrice**](CloudPrice.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudMarketItem

`func NewCloudMarketItem() *CloudMarketItem`

NewCloudMarketItem instantiates a new CloudMarketItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMarketItemWithDefaults

`func NewCloudMarketItemWithDefaults() *CloudMarketItem`

NewCloudMarketItemWithDefaults instantiates a new CloudMarketItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *CloudMarketItem) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *CloudMarketItem) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *CloudMarketItem) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *CloudMarketItem) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetCategory

`func (o *CloudMarketItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudMarketItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudMarketItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudMarketItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *CloudMarketItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudMarketItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudMarketItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudMarketItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDispatchable

`func (o *CloudMarketItem) GetDispatchable() bool`

GetDispatchable returns the Dispatchable field if non-nil, zero value otherwise.

### GetDispatchableOk

`func (o *CloudMarketItem) GetDispatchableOk() (*bool, bool)`

GetDispatchableOk returns a tuple with the Dispatchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchable

`func (o *CloudMarketItem) SetDispatchable(v bool)`

SetDispatchable sets Dispatchable field to given value.

### HasDispatchable

`func (o *CloudMarketItem) HasDispatchable() bool`

HasDispatchable returns a boolean if a field has been set.

### GetInputSchema

`func (o *CloudMarketItem) GetInputSchema() interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *CloudMarketItem) GetInputSchemaOk() (*interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *CloudMarketItem) SetInputSchema(v interface{})`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *CloudMarketItem) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### SetInputSchemaNil

`func (o *CloudMarketItem) SetInputSchemaNil(b bool)`

 SetInputSchemaNil sets the value for InputSchema to be an explicit nil

### UnsetInputSchema
`func (o *CloudMarketItem) UnsetInputSchema()`

UnsetInputSchema ensures that no value is present for InputSchema, not even an explicit nil
### GetInstalled

`func (o *CloudMarketItem) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *CloudMarketItem) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *CloudMarketItem) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *CloudMarketItem) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetName

`func (o *CloudMarketItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudMarketItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudMarketItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudMarketItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *CloudMarketItem) GetPrice() CloudPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CloudMarketItem) GetPriceOk() (*CloudPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CloudMarketItem) SetPrice(v CloudPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CloudMarketItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSource

`func (o *CloudMarketItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudMarketItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudMarketItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudMarketItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *CloudMarketItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudMarketItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudMarketItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudMarketItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


