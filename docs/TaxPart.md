# TaxPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csv** | Pointer to **string** | CSV is the file itself. It carries full TINs, and is answered only to an org admin of the payer, on the audit trail. | [optional] 
**Digest** | Pointer to **string** | Digest is the SHA-256 of its bytes, hex, so a re-read proves it is the same file. | [optional] 
**Name** | Pointer to **string** | Name is the file name, e.g. \&quot;1099-NEC-2026-original-1.csv\&quot;. | [optional] 
**Records** | Pointer to **int64** | Records is how many forms it holds, at most 250. | [optional] 
**Type** | Pointer to **string** | Type is original or correction: the portal takes them apart. | [optional] 

## Methods

### NewTaxPart

`func NewTaxPart() *TaxPart`

NewTaxPart instantiates a new TaxPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxPartWithDefaults

`func NewTaxPartWithDefaults() *TaxPart`

NewTaxPartWithDefaults instantiates a new TaxPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsv

`func (o *TaxPart) GetCsv() string`

GetCsv returns the Csv field if non-nil, zero value otherwise.

### GetCsvOk

`func (o *TaxPart) GetCsvOk() (*string, bool)`

GetCsvOk returns a tuple with the Csv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsv

`func (o *TaxPart) SetCsv(v string)`

SetCsv sets Csv field to given value.

### HasCsv

`func (o *TaxPart) HasCsv() bool`

HasCsv returns a boolean if a field has been set.

### GetDigest

`func (o *TaxPart) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *TaxPart) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *TaxPart) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *TaxPart) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetName

`func (o *TaxPart) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxPart) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxPart) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxPart) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecords

`func (o *TaxPart) GetRecords() int64`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *TaxPart) GetRecordsOk() (*int64, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *TaxPart) SetRecords(v int64)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *TaxPart) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetType

`func (o *TaxPart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxPart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxPart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaxPart) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


