# ProviderBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the warehouse could not be read, which means \&quot;no answer\&quot; and NOT \&quot;no usage\&quot; — Items is then empty for a reason. | [optional] 
**Items** | Pointer to [**[]ProviderRow**](ProviderRow.md) | Items is one row per provider, most tokens first. | [optional] 
**Source** | Pointer to **string** | Source names the warehouse table the rows came from. | [optional] 

## Methods

### NewProviderBreakdown

`func NewProviderBreakdown() *ProviderBreakdown`

NewProviderBreakdown instantiates a new ProviderBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderBreakdownWithDefaults

`func NewProviderBreakdownWithDefaults() *ProviderBreakdown`

NewProviderBreakdownWithDefaults instantiates a new ProviderBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ProviderBreakdown) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ProviderBreakdown) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ProviderBreakdown) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ProviderBreakdown) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetItems

`func (o *ProviderBreakdown) GetItems() []ProviderRow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProviderBreakdown) GetItemsOk() (*[]ProviderRow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProviderBreakdown) SetItems(v []ProviderRow)`

SetItems sets Items field to given value.

### HasItems

`func (o *ProviderBreakdown) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSource

`func (o *ProviderBreakdown) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProviderBreakdown) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProviderBreakdown) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProviderBreakdown) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


