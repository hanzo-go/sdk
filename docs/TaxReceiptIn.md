# TaxReceiptIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the filing, from the path. | [optional] 
**ReceiptId** | Pointer to **string** | ReceiptID is the Receipt ID IRIS returned for the submission. | [optional] 
**Status** | Pointer to **string** | Status is submitted (IRIS returned a Receipt ID), accepted, accepted_with_errors or rejected (IRIS&#39;s acknowledgment). | [optional] 
**Tcc** | Pointer to **string** | TCC is the IRIS Transmitter Control Code the payer submitted under: five uppercase letters and digits. | [optional] 

## Methods

### NewTaxReceiptIn

`func NewTaxReceiptIn() *TaxReceiptIn`

NewTaxReceiptIn instantiates a new TaxReceiptIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxReceiptInWithDefaults

`func NewTaxReceiptInWithDefaults() *TaxReceiptIn`

NewTaxReceiptInWithDefaults instantiates a new TaxReceiptIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxReceiptIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxReceiptIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxReceiptIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxReceiptIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReceiptId

`func (o *TaxReceiptIn) GetReceiptId() string`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *TaxReceiptIn) GetReceiptIdOk() (*string, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *TaxReceiptIn) SetReceiptId(v string)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *TaxReceiptIn) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### GetStatus

`func (o *TaxReceiptIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaxReceiptIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaxReceiptIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaxReceiptIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTcc

`func (o *TaxReceiptIn) GetTcc() string`

GetTcc returns the Tcc field if non-nil, zero value otherwise.

### GetTccOk

`func (o *TaxReceiptIn) GetTccOk() (*string, bool)`

GetTccOk returns a tuple with the Tcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcc

`func (o *TaxReceiptIn) SetTcc(v string)`

SetTcc sets Tcc field to given value.

### HasTcc

`func (o *TaxReceiptIn) HasTcc() bool`

HasTcc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


