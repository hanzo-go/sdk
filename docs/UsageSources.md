# UsageSources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commerce** | Pointer to **bool** | Commerce is whether the billing ledger answered the spend block. | [optional] 
**Warehouse** | Pointer to **bool** | Warehouse is whether the usage warehouse answered the LLM block. | [optional] 

## Methods

### NewUsageSources

`func NewUsageSources() *UsageSources`

NewUsageSources instantiates a new UsageSources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSourcesWithDefaults

`func NewUsageSourcesWithDefaults() *UsageSources`

NewUsageSourcesWithDefaults instantiates a new UsageSources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommerce

`func (o *UsageSources) GetCommerce() bool`

GetCommerce returns the Commerce field if non-nil, zero value otherwise.

### GetCommerceOk

`func (o *UsageSources) GetCommerceOk() (*bool, bool)`

GetCommerceOk returns a tuple with the Commerce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommerce

`func (o *UsageSources) SetCommerce(v bool)`

SetCommerce sets Commerce field to given value.

### HasCommerce

`func (o *UsageSources) HasCommerce() bool`

HasCommerce returns a boolean if a field has been set.

### GetWarehouse

`func (o *UsageSources) GetWarehouse() bool`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *UsageSources) GetWarehouseOk() (*bool, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *UsageSources) SetWarehouse(v bool)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *UsageSources) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


