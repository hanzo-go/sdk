# TaxFilingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TaxFiling**](TaxFiling.md) | Data are the filings, oldest first, without their files. | [optional] 

## Methods

### NewTaxFilingList

`func NewTaxFilingList() *TaxFilingList`

NewTaxFilingList instantiates a new TaxFilingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxFilingListWithDefaults

`func NewTaxFilingListWithDefaults() *TaxFilingList`

NewTaxFilingListWithDefaults instantiates a new TaxFilingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TaxFilingList) GetData() []TaxFiling`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TaxFilingList) GetDataOk() (*[]TaxFiling, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TaxFilingList) SetData(v []TaxFiling)`

SetData sets Data field to given value.

### HasData

`func (o *TaxFilingList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


