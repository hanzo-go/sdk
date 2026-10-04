# NotifyNotifyCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is one of the credentials that provider reads, such as auth-token or smtp-host. A key the provider does not read is refused. | [optional] 
**Provider** | Pointer to **string** | Provider is the delivery provider the credential is for: twilio, plivo, twilio_email or mail. | [optional] 
**Value** | Pointer to **string** | Value is sealed in KMS and never answered, logged or echoed. It is never read from a query string, because a URL is logged in more places than a body is. | [optional] 

## Methods

### NewNotifyNotifyCredential

`func NewNotifyNotifyCredential() *NotifyNotifyCredential`

NewNotifyNotifyCredential instantiates a new NotifyNotifyCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyNotifyCredentialWithDefaults

`func NewNotifyNotifyCredentialWithDefaults() *NotifyNotifyCredential`

NewNotifyNotifyCredentialWithDefaults instantiates a new NotifyNotifyCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *NotifyNotifyCredential) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NotifyNotifyCredential) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NotifyNotifyCredential) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *NotifyNotifyCredential) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetProvider

`func (o *NotifyNotifyCredential) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotifyNotifyCredential) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotifyNotifyCredential) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NotifyNotifyCredential) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetValue

`func (o *NotifyNotifyCredential) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NotifyNotifyCredential) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NotifyNotifyCredential) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NotifyNotifyCredential) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


