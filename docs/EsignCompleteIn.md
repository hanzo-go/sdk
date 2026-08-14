# EsignCompleteIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signed** | Pointer to **bool** | Signed, when present, overrides what the provider reports — the manual path for a provider whose webhook is not wired. Omit it to take the provider&#39;s answer. | [optional] 

## Methods

### NewEsignCompleteIn

`func NewEsignCompleteIn() *EsignCompleteIn`

NewEsignCompleteIn instantiates a new EsignCompleteIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignCompleteInWithDefaults

`func NewEsignCompleteInWithDefaults() *EsignCompleteIn`

NewEsignCompleteInWithDefaults instantiates a new EsignCompleteIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigned

`func (o *EsignCompleteIn) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *EsignCompleteIn) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *EsignCompleteIn) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *EsignCompleteIn) HasSigned() bool`

HasSigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


