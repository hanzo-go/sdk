# SignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the document to send for signature, from the path. | [optional] 
**Signers** | Pointer to [**[]LegalSigner**](LegalSigner.md) | Signers are the people who must sign, by name and email. At least one is required. | [optional] 

## Methods

### NewSignRequest

`func NewSignRequest() *SignRequest`

NewSignRequest instantiates a new SignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignRequestWithDefaults

`func NewSignRequestWithDefaults() *SignRequest`

NewSignRequestWithDefaults instantiates a new SignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSigners

`func (o *SignRequest) GetSigners() []LegalSigner`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *SignRequest) GetSignersOk() (*[]LegalSigner, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *SignRequest) SetSigners(v []LegalSigner)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *SignRequest) HasSigners() bool`

HasSigners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


