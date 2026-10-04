# TaxCorrectIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boxes** | Pointer to [**[]TaxAmount**](TaxAmount.md) | Boxes are the corrected amounts, by box. Omitted, the boxes are derived again from the payments and the payee&#39;s W-9 is read again — which is also how a wrong name or TIN is corrected. Zero in every box withdraws the return. | [optional] 
**Id** | Pointer to **string** | ID is the furnished form to correct, from the path. | [optional] 
**Reason** | Pointer to **string** | Reason says what was wrong. Required; it goes on the record. | [optional] 

## Methods

### NewTaxCorrectIn

`func NewTaxCorrectIn() *TaxCorrectIn`

NewTaxCorrectIn instantiates a new TaxCorrectIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCorrectInWithDefaults

`func NewTaxCorrectInWithDefaults() *TaxCorrectIn`

NewTaxCorrectInWithDefaults instantiates a new TaxCorrectIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxes

`func (o *TaxCorrectIn) GetBoxes() []TaxAmount`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *TaxCorrectIn) GetBoxesOk() (*[]TaxAmount, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *TaxCorrectIn) SetBoxes(v []TaxAmount)`

SetBoxes sets Boxes field to given value.

### HasBoxes

`func (o *TaxCorrectIn) HasBoxes() bool`

HasBoxes returns a boolean if a field has been set.

### GetId

`func (o *TaxCorrectIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxCorrectIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxCorrectIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxCorrectIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *TaxCorrectIn) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TaxCorrectIn) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TaxCorrectIn) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TaxCorrectIn) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


