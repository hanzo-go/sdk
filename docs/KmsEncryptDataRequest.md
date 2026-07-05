# KmsEncryptDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plaintext** | **string** | Base64-encoded plaintext to encrypt | 

## Methods

### NewKmsEncryptDataRequest

`func NewKmsEncryptDataRequest(plaintext string, ) *KmsEncryptDataRequest`

NewKmsEncryptDataRequest instantiates a new KmsEncryptDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsEncryptDataRequestWithDefaults

`func NewKmsEncryptDataRequestWithDefaults() *KmsEncryptDataRequest`

NewKmsEncryptDataRequestWithDefaults instantiates a new KmsEncryptDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaintext

`func (o *KmsEncryptDataRequest) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *KmsEncryptDataRequest) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *KmsEncryptDataRequest) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


