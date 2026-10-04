# BooksRulesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]BooksRule**](BooksRule.md) | Rules is every rule the org has set, highest priority first — the order they are matched in. | [optional] 

## Methods

### NewBooksRulesOut

`func NewBooksRulesOut() *BooksRulesOut`

NewBooksRulesOut instantiates a new BooksRulesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksRulesOutWithDefaults

`func NewBooksRulesOutWithDefaults() *BooksRulesOut`

NewBooksRulesOutWithDefaults instantiates a new BooksRulesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *BooksRulesOut) GetRules() []BooksRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BooksRulesOut) GetRulesOk() (*[]BooksRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BooksRulesOut) SetRules(v []BooksRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *BooksRulesOut) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


