# CloudSignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the document to send for signature, from the path. | [optional] 
**Signers** | Pointer to [**[]CloudLegalSigner**](CloudLegalSigner.md) | Signers are the people who must sign, by name and email. At least one is required. | [optional] 

## Methods

### NewCloudSignRequest

`func NewCloudSignRequest() *CloudSignRequest`

NewCloudSignRequest instantiates a new CloudSignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSignRequestWithDefaults

`func NewCloudSignRequestWithDefaults() *CloudSignRequest`

NewCloudSignRequestWithDefaults instantiates a new CloudSignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudSignRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSignRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSignRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSignRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSigners

`func (o *CloudSignRequest) GetSigners() []CloudLegalSigner`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *CloudSignRequest) GetSignersOk() (*[]CloudLegalSigner, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *CloudSignRequest) SetSigners(v []CloudLegalSigner)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *CloudSignRequest) HasSigners() bool`

HasSigners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


