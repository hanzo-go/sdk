# CategoryIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brands** | Pointer to **[]string** | Brands are the brands whose console shows it. Omit for every brand. | [optional] 
**Id** | Pointer to **string** | ID is the category slug to write, from the path. | [optional] 
**Label** | Pointer to **string** | Label is the display name. Required. | [optional] 
**Order** | Pointer to **int32** | Order is where the category sits among its siblings, ascending. | [optional] 
**Summary** | Pointer to **string** | Summary is the one line describing what the category groups. | [optional] 

## Methods

### NewCategoryIn

`func NewCategoryIn() *CategoryIn`

NewCategoryIn instantiates a new CategoryIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryInWithDefaults

`func NewCategoryInWithDefaults() *CategoryIn`

NewCategoryInWithDefaults instantiates a new CategoryIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrands

`func (o *CategoryIn) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *CategoryIn) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *CategoryIn) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *CategoryIn) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetId

`func (o *CategoryIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CategoryIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *CategoryIn) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CategoryIn) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CategoryIn) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CategoryIn) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOrder

`func (o *CategoryIn) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CategoryIn) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CategoryIn) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CategoryIn) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSummary

`func (o *CategoryIn) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CategoryIn) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CategoryIn) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CategoryIn) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


