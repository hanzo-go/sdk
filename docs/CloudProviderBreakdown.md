# CloudProviderBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the warehouse could not be read, which means \&quot;no answer\&quot; and NOT \&quot;no usage\&quot; — Items is then empty for a reason. | [optional] 
**Items** | Pointer to [**[]CloudProviderRow**](CloudProviderRow.md) | Items is one row per provider, most tokens first. | [optional] 
**Source** | Pointer to **string** | Source names the warehouse table the rows came from. | [optional] 

## Methods

### NewCloudProviderBreakdown

`func NewCloudProviderBreakdown() *CloudProviderBreakdown`

NewCloudProviderBreakdown instantiates a new CloudProviderBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderBreakdownWithDefaults

`func NewCloudProviderBreakdownWithDefaults() *CloudProviderBreakdown`

NewCloudProviderBreakdownWithDefaults instantiates a new CloudProviderBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *CloudProviderBreakdown) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CloudProviderBreakdown) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CloudProviderBreakdown) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CloudProviderBreakdown) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetItems

`func (o *CloudProviderBreakdown) GetItems() []CloudProviderRow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudProviderBreakdown) GetItemsOk() (*[]CloudProviderRow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudProviderBreakdown) SetItems(v []CloudProviderRow)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudProviderBreakdown) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSource

`func (o *CloudProviderBreakdown) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudProviderBreakdown) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudProviderBreakdown) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudProviderBreakdown) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


