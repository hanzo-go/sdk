# RegistryWebhookCreateTargetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Address** | **string** |  | 
**AuthHeader** | Pointer to **string** |  | [optional] 
**SkipCertVerify** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegistryWebhookCreateTargetsInner

`func NewRegistryWebhookCreateTargetsInner(type_ string, address string, ) *RegistryWebhookCreateTargetsInner`

NewRegistryWebhookCreateTargetsInner instantiates a new RegistryWebhookCreateTargetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryWebhookCreateTargetsInnerWithDefaults

`func NewRegistryWebhookCreateTargetsInnerWithDefaults() *RegistryWebhookCreateTargetsInner`

NewRegistryWebhookCreateTargetsInnerWithDefaults instantiates a new RegistryWebhookCreateTargetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RegistryWebhookCreateTargetsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistryWebhookCreateTargetsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistryWebhookCreateTargetsInner) SetType(v string)`

SetType sets Type field to given value.


### GetAddress

`func (o *RegistryWebhookCreateTargetsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RegistryWebhookCreateTargetsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RegistryWebhookCreateTargetsInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAuthHeader

`func (o *RegistryWebhookCreateTargetsInner) GetAuthHeader() string`

GetAuthHeader returns the AuthHeader field if non-nil, zero value otherwise.

### GetAuthHeaderOk

`func (o *RegistryWebhookCreateTargetsInner) GetAuthHeaderOk() (*string, bool)`

GetAuthHeaderOk returns a tuple with the AuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeader

`func (o *RegistryWebhookCreateTargetsInner) SetAuthHeader(v string)`

SetAuthHeader sets AuthHeader field to given value.

### HasAuthHeader

`func (o *RegistryWebhookCreateTargetsInner) HasAuthHeader() bool`

HasAuthHeader returns a boolean if a field has been set.

### GetSkipCertVerify

`func (o *RegistryWebhookCreateTargetsInner) GetSkipCertVerify() bool`

GetSkipCertVerify returns the SkipCertVerify field if non-nil, zero value otherwise.

### GetSkipCertVerifyOk

`func (o *RegistryWebhookCreateTargetsInner) GetSkipCertVerifyOk() (*bool, bool)`

GetSkipCertVerifyOk returns a tuple with the SkipCertVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCertVerify

`func (o *RegistryWebhookCreateTargetsInner) SetSkipCertVerify(v bool)`

SetSkipCertVerify sets SkipCertVerify field to given value.

### HasSkipCertVerify

`func (o *RegistryWebhookCreateTargetsInner) HasSkipCertVerify() bool`

HasSkipCertVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


