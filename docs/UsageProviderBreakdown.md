# UsageProviderBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the warehouse could not be read, which means \&quot;no answer\&quot; and NOT \&quot;no usage\&quot; — Items is then empty for a reason. | [optional] 
**Items** | Pointer to [**[]UsageProviderRow**](UsageProviderRow.md) | Items is one row per provider, most tokens first. | [optional] 
**Source** | Pointer to **string** | Source names the warehouse table the rows came from. | [optional] 

## Methods

### NewUsageProviderBreakdown

`func NewUsageProviderBreakdown() *UsageProviderBreakdown`

NewUsageProviderBreakdown instantiates a new UsageProviderBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageProviderBreakdownWithDefaults

`func NewUsageProviderBreakdownWithDefaults() *UsageProviderBreakdown`

NewUsageProviderBreakdownWithDefaults instantiates a new UsageProviderBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *UsageProviderBreakdown) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *UsageProviderBreakdown) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *UsageProviderBreakdown) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *UsageProviderBreakdown) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetItems

`func (o *UsageProviderBreakdown) GetItems() []UsageProviderRow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UsageProviderBreakdown) GetItemsOk() (*[]UsageProviderRow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UsageProviderBreakdown) SetItems(v []UsageProviderRow)`

SetItems sets Items field to given value.

### HasItems

`func (o *UsageProviderBreakdown) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSource

`func (o *UsageProviderBreakdown) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UsageProviderBreakdown) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UsageProviderBreakdown) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UsageProviderBreakdown) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


