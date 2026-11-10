# ProviderWhatsappGroupIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is the group&#39;s description. Optional, at most 2048 characters. | [optional] 
**Subject** | Pointer to **string** | Subject is the group&#39;s name, shown to its members. Required, at most 128 characters. | [optional] 

## Methods

### NewProviderWhatsappGroupIn

`func NewProviderWhatsappGroupIn() *ProviderWhatsappGroupIn`

NewProviderWhatsappGroupIn instantiates a new ProviderWhatsappGroupIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWhatsappGroupInWithDefaults

`func NewProviderWhatsappGroupInWithDefaults() *ProviderWhatsappGroupIn`

NewProviderWhatsappGroupInWithDefaults instantiates a new ProviderWhatsappGroupIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProviderWhatsappGroupIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderWhatsappGroupIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderWhatsappGroupIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProviderWhatsappGroupIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubject

`func (o *ProviderWhatsappGroupIn) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ProviderWhatsappGroupIn) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ProviderWhatsappGroupIn) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ProviderWhatsappGroupIn) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


