# Sources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commerce** | Pointer to **bool** | Commerce is whether the billing ledger answered the spend block. | [optional] 
**Warehouse** | Pointer to **bool** | Warehouse is whether the usage warehouse answered the LLM block. | [optional] 

## Methods

### NewSources

`func NewSources() *Sources`

NewSources instantiates a new Sources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesWithDefaults

`func NewSourcesWithDefaults() *Sources`

NewSourcesWithDefaults instantiates a new Sources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommerce

`func (o *Sources) GetCommerce() bool`

GetCommerce returns the Commerce field if non-nil, zero value otherwise.

### GetCommerceOk

`func (o *Sources) GetCommerceOk() (*bool, bool)`

GetCommerceOk returns a tuple with the Commerce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommerce

`func (o *Sources) SetCommerce(v bool)`

SetCommerce sets Commerce field to given value.

### HasCommerce

`func (o *Sources) HasCommerce() bool`

HasCommerce returns a boolean if a field has been set.

### GetWarehouse

`func (o *Sources) GetWarehouse() bool`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *Sources) GetWarehouseOk() (*bool, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *Sources) SetWarehouse(v bool)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *Sources) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


