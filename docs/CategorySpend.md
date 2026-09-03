# CategorySpend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is the bucket the ledger&#39;s own tag mapped to. An untagged or unrecognised line gets its own honest bucket rather than being folded away. | [optional] 
**Cents** | Pointer to **int64** | Cents is what the org spent in that bucket over the window, in US cents. | [optional] 
**Count** | Pointer to **int64** | Count is how many ledger lines rolled up into it. | [optional] 

## Methods

### NewCategorySpend

`func NewCategorySpend() *CategorySpend`

NewCategorySpend instantiates a new CategorySpend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorySpendWithDefaults

`func NewCategorySpendWithDefaults() *CategorySpend`

NewCategorySpendWithDefaults instantiates a new CategorySpend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CategorySpend) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CategorySpend) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CategorySpend) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CategorySpend) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCents

`func (o *CategorySpend) GetCents() int64`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *CategorySpend) GetCentsOk() (*int64, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *CategorySpend) SetCents(v int64)`

SetCents sets Cents field to given value.

### HasCents

`func (o *CategorySpend) HasCents() bool`

HasCents returns a boolean if a field has been set.

### GetCount

`func (o *CategorySpend) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CategorySpend) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CategorySpend) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *CategorySpend) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


