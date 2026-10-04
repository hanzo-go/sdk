# TaxStatementList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TaxStatement**](TaxStatement.md) | Data are the statements, oldest first. | [optional] 

## Methods

### NewTaxStatementList

`func NewTaxStatementList() *TaxStatementList`

NewTaxStatementList instantiates a new TaxStatementList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxStatementListWithDefaults

`func NewTaxStatementListWithDefaults() *TaxStatementList`

NewTaxStatementListWithDefaults instantiates a new TaxStatementList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TaxStatementList) GetData() []TaxStatement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TaxStatementList) GetDataOk() (*[]TaxStatement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TaxStatementList) SetData(v []TaxStatement)`

SetData sets Data field to given value.

### HasData

`func (o *TaxStatementList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


