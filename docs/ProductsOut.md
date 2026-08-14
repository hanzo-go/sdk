# ProductsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ProductRow**](ProductRow.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewProductsOut

`func NewProductsOut() *ProductsOut`

NewProductsOut instantiates a new ProductsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsOutWithDefaults

`func NewProductsOutWithDefaults() *ProductsOut`

NewProductsOutWithDefaults instantiates a new ProductsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProductsOut) GetData() []ProductRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProductsOut) GetDataOk() (*[]ProductRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProductsOut) SetData(v []ProductRow)`

SetData sets Data field to given value.

### HasData

`func (o *ProductsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *ProductsOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ProductsOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ProductsOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ProductsOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *ProductsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *ProductsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProductsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProductsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProductsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


