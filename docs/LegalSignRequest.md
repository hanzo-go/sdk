# LegalSignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the document to send for signature, from the path. | [optional] 
**Signers** | Pointer to [**[]LegalLegalSigner**](LegalLegalSigner.md) | Signers are the people who must sign, by name and email. At least one is required. | [optional] 

## Methods

### NewLegalSignRequest

`func NewLegalSignRequest() *LegalSignRequest`

NewLegalSignRequest instantiates a new LegalSignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalSignRequestWithDefaults

`func NewLegalSignRequestWithDefaults() *LegalSignRequest`

NewLegalSignRequestWithDefaults instantiates a new LegalSignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LegalSignRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegalSignRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegalSignRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegalSignRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSigners

`func (o *LegalSignRequest) GetSigners() []LegalLegalSigner`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *LegalSignRequest) GetSignersOk() (*[]LegalLegalSigner, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *LegalSignRequest) SetSigners(v []LegalLegalSigner)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *LegalSignRequest) HasSigners() bool`

HasSigners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


