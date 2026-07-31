# CloudCategorySpend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is the bucket the ledger&#39;s own tag mapped to. An untagged or unrecognised line gets its own honest bucket rather than being folded away. | [optional] 
**Cents** | Pointer to **int32** | Cents is what the org spent in that bucket over the window, in US cents. | [optional] 
**Count** | Pointer to **int32** | Count is how many ledger lines rolled up into it. | [optional] 

## Methods

### NewCloudCategorySpend

`func NewCloudCategorySpend() *CloudCategorySpend`

NewCloudCategorySpend instantiates a new CloudCategorySpend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCategorySpendWithDefaults

`func NewCloudCategorySpendWithDefaults() *CloudCategorySpend`

NewCloudCategorySpendWithDefaults instantiates a new CloudCategorySpend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CloudCategorySpend) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudCategorySpend) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudCategorySpend) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudCategorySpend) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCents

`func (o *CloudCategorySpend) GetCents() int32`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *CloudCategorySpend) GetCentsOk() (*int32, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *CloudCategorySpend) SetCents(v int32)`

SetCents sets Cents field to given value.

### HasCents

`func (o *CloudCategorySpend) HasCents() bool`

HasCents returns a boolean if a field has been set.

### GetCount

`func (o *CloudCategorySpend) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudCategorySpend) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudCategorySpend) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudCategorySpend) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


