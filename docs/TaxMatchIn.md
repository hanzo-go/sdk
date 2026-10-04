# TaxMatchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the relationship, from the path. | [optional] 
**Result** | Pointer to **string** | Result is what IRS TIN Matching answered for this payee&#39;s name and TIN: matched or mismatched. | [optional] 

## Methods

### NewTaxMatchIn

`func NewTaxMatchIn() *TaxMatchIn`

NewTaxMatchIn instantiates a new TaxMatchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxMatchInWithDefaults

`func NewTaxMatchInWithDefaults() *TaxMatchIn`

NewTaxMatchInWithDefaults instantiates a new TaxMatchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxMatchIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxMatchIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxMatchIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxMatchIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResult

`func (o *TaxMatchIn) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TaxMatchIn) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TaxMatchIn) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *TaxMatchIn) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


