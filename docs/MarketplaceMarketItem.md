# MarketplaceMarketItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** | Category is that same listing&#39;s grouping. Free text chosen by the publisher, absent when there is no public listing or the publisher left it blank. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Dispatchable** | Pointer to **bool** |  | [optional] 
**InputSchema** | Pointer to **interface{}** |  | [optional] 
**Installed** | Pointer to **bool** | Installed is whether the tool is activated for THIS caller&#39;s (org, project): the same bit as Activated, under the shop&#39;s name for it, which install and uninstall are the writes for. It is per caller, so one listing reads installed for one org and not for another. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to [**MarketplacePrice**](MarketplacePrice.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** | Title is the shop-window name, painted over the registry Name from the CHEAPEST public listing of this very tool — the platform&#39;s, for a tool every org reaches — the row a call would pay. Absent when the tool is not listed: that row is a plain capability, not an offer. | [optional] 

## Methods

### NewMarketplaceMarketItem

`func NewMarketplaceMarketItem() *MarketplaceMarketItem`

NewMarketplaceMarketItem instantiates a new MarketplaceMarketItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceMarketItemWithDefaults

`func NewMarketplaceMarketItemWithDefaults() *MarketplaceMarketItem`

NewMarketplaceMarketItemWithDefaults instantiates a new MarketplaceMarketItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *MarketplaceMarketItem) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *MarketplaceMarketItem) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *MarketplaceMarketItem) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *MarketplaceMarketItem) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetCategory

`func (o *MarketplaceMarketItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketplaceMarketItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketplaceMarketItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketplaceMarketItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *MarketplaceMarketItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceMarketItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceMarketItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceMarketItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDispatchable

`func (o *MarketplaceMarketItem) GetDispatchable() bool`

GetDispatchable returns the Dispatchable field if non-nil, zero value otherwise.

### GetDispatchableOk

`func (o *MarketplaceMarketItem) GetDispatchableOk() (*bool, bool)`

GetDispatchableOk returns a tuple with the Dispatchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchable

`func (o *MarketplaceMarketItem) SetDispatchable(v bool)`

SetDispatchable sets Dispatchable field to given value.

### HasDispatchable

`func (o *MarketplaceMarketItem) HasDispatchable() bool`

HasDispatchable returns a boolean if a field has been set.

### GetInputSchema

`func (o *MarketplaceMarketItem) GetInputSchema() interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *MarketplaceMarketItem) GetInputSchemaOk() (*interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *MarketplaceMarketItem) SetInputSchema(v interface{})`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *MarketplaceMarketItem) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### SetInputSchemaNil

`func (o *MarketplaceMarketItem) SetInputSchemaNil(b bool)`

 SetInputSchemaNil sets the value for InputSchema to be an explicit nil

### UnsetInputSchema
`func (o *MarketplaceMarketItem) UnsetInputSchema()`

UnsetInputSchema ensures that no value is present for InputSchema, not even an explicit nil
### GetInstalled

`func (o *MarketplaceMarketItem) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *MarketplaceMarketItem) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *MarketplaceMarketItem) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *MarketplaceMarketItem) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetName

`func (o *MarketplaceMarketItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceMarketItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceMarketItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketplaceMarketItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *MarketplaceMarketItem) GetPrice() MarketplacePrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarketplaceMarketItem) GetPriceOk() (*MarketplacePrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarketplaceMarketItem) SetPrice(v MarketplacePrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarketplaceMarketItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSource

`func (o *MarketplaceMarketItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketplaceMarketItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketplaceMarketItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketplaceMarketItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *MarketplaceMarketItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarketplaceMarketItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarketplaceMarketItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MarketplaceMarketItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


