# EsignRejection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientId** | Pointer to **string** | RecipientID is the recipient who declined. | [optional] 
**Status** | Pointer to **string** | Status is REJECTED — one declining signer ends the document for everyone, and there is no route back. | [optional] 

## Methods

### NewEsignRejection

`func NewEsignRejection() *EsignRejection`

NewEsignRejection instantiates a new EsignRejection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignRejectionWithDefaults

`func NewEsignRejectionWithDefaults() *EsignRejection`

NewEsignRejectionWithDefaults instantiates a new EsignRejection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientId

`func (o *EsignRejection) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EsignRejection) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EsignRejection) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EsignRejection) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetStatus

`func (o *EsignRejection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsignRejection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsignRejection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EsignRejection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


