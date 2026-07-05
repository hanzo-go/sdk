# PaasLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **string** |  | [optional] 

## Methods

### NewPaasLoginRequest

`func NewPaasLoginRequest() *PaasLoginRequest`

NewPaasLoginRequest instantiates a new PaasLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasLoginRequestWithDefaults

`func NewPaasLoginRequestWithDefaults() *PaasLoginRequest`

NewPaasLoginRequestWithDefaults instantiates a new PaasLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *PaasLoginRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PaasLoginRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PaasLoginRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PaasLoginRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCode

`func (o *PaasLoginRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaasLoginRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaasLoginRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PaasLoginRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRedirectUri

`func (o *PaasLoginRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *PaasLoginRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *PaasLoginRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *PaasLoginRequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


