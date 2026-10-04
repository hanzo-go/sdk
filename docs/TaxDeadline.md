# TaxDeadline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to **string** | File is the date the return is due to the IRS, electronically (IRIS). | [optional] 
**Furnish** | Pointer to **string** | Furnish is the date Copy B is due to the recipient. | [optional] 

## Methods

### NewTaxDeadline

`func NewTaxDeadline() *TaxDeadline`

NewTaxDeadline instantiates a new TaxDeadline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxDeadlineWithDefaults

`func NewTaxDeadlineWithDefaults() *TaxDeadline`

NewTaxDeadlineWithDefaults instantiates a new TaxDeadline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *TaxDeadline) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TaxDeadline) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TaxDeadline) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *TaxDeadline) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFurnish

`func (o *TaxDeadline) GetFurnish() string`

GetFurnish returns the Furnish field if non-nil, zero value otherwise.

### GetFurnishOk

`func (o *TaxDeadline) GetFurnishOk() (*string, bool)`

GetFurnishOk returns a tuple with the Furnish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFurnish

`func (o *TaxDeadline) SetFurnish(v string)`

SetFurnish sets Furnish field to given value.

### HasFurnish

`func (o *TaxDeadline) HasFurnish() bool`

HasFurnish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


