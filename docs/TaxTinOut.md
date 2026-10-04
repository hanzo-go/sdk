# TaxTinOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForeignTin** | Pointer to **string** | ForeignTIN is a W-8 payee&#39;s foreign tax identifying number as issued, which Form 1042-S carries for a foreign recipient. | [optional] 
**Tin** | Pointer to **string** | TIN is the payee&#39;s full U.S. taxpayer identification number, hyphenated as the IRS writes it; empty when a W-8 carries none. | [optional] 
**TinType** | Pointer to **string** | TINType is ssn or ein. | [optional] 

## Methods

### NewTaxTinOut

`func NewTaxTinOut() *TaxTinOut`

NewTaxTinOut instantiates a new TaxTinOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxTinOutWithDefaults

`func NewTaxTinOutWithDefaults() *TaxTinOut`

NewTaxTinOutWithDefaults instantiates a new TaxTinOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForeignTin

`func (o *TaxTinOut) GetForeignTin() string`

GetForeignTin returns the ForeignTin field if non-nil, zero value otherwise.

### GetForeignTinOk

`func (o *TaxTinOut) GetForeignTinOk() (*string, bool)`

GetForeignTinOk returns a tuple with the ForeignTin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTin

`func (o *TaxTinOut) SetForeignTin(v string)`

SetForeignTin sets ForeignTin field to given value.

### HasForeignTin

`func (o *TaxTinOut) HasForeignTin() bool`

HasForeignTin returns a boolean if a field has been set.

### GetTin

`func (o *TaxTinOut) GetTin() string`

GetTin returns the Tin field if non-nil, zero value otherwise.

### GetTinOk

`func (o *TaxTinOut) GetTinOk() (*string, bool)`

GetTinOk returns a tuple with the Tin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTin

`func (o *TaxTinOut) SetTin(v string)`

SetTin sets Tin field to given value.

### HasTin

`func (o *TaxTinOut) HasTin() bool`

HasTin returns a boolean if a field has been set.

### GetTinType

`func (o *TaxTinOut) GetTinType() string`

GetTinType returns the TinType field if non-nil, zero value otherwise.

### GetTinTypeOk

`func (o *TaxTinOut) GetTinTypeOk() (*string, bool)`

GetTinTypeOk returns a tuple with the TinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinType

`func (o *TaxTinOut) SetTinType(v string)`

SetTinType sets TinType field to given value.

### HasTinType

`func (o *TaxTinOut) HasTinType() bool`

HasTinType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


