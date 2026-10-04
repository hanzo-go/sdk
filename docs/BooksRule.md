# BooksRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is the COA expense account a matching bill books to. An upsert normalizes a slug (\&quot;cloud\&quot;) to its account number. | [optional] 
**Pattern** | Pointer to **string** | Pattern is the merchant substring the rule matches on, case-insensitively. It is also the key an upsert writes by. | [optional] 
**Priority** | Pointer to **int64** | Priority breaks ties: when several patterns match, the highest wins. | [optional] 

## Methods

### NewBooksRule

`func NewBooksRule() *BooksRule`

NewBooksRule instantiates a new BooksRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksRuleWithDefaults

`func NewBooksRuleWithDefaults() *BooksRule`

NewBooksRuleWithDefaults instantiates a new BooksRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *BooksRule) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BooksRule) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BooksRule) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BooksRule) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPattern

`func (o *BooksRule) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *BooksRule) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *BooksRule) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *BooksRule) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPriority

`func (o *BooksRule) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BooksRule) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BooksRule) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BooksRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


