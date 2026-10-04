# EsignEsignRejection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientId** | Pointer to **string** | RecipientID is the recipient who declined. | [optional] 
**Status** | Pointer to **string** | Status is REJECTED — one declining signer ends the document for everyone, and there is no route back. | [optional] 

## Methods

### NewEsignEsignRejection

`func NewEsignEsignRejection() *EsignEsignRejection`

NewEsignEsignRejection instantiates a new EsignEsignRejection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEsignRejectionWithDefaults

`func NewEsignEsignRejectionWithDefaults() *EsignEsignRejection`

NewEsignEsignRejectionWithDefaults instantiates a new EsignEsignRejection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientId

`func (o *EsignEsignRejection) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EsignEsignRejection) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EsignEsignRejection) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EsignEsignRejection) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetStatus

`func (o *EsignEsignRejection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsignEsignRejection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsignEsignRejection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EsignEsignRejection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


