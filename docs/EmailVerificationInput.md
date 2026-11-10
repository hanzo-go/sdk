# EmailVerificationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to **[]string** | Checks explicitly selects which checks to run. Defaults to [\&quot;syntax\&quot;, \&quot;domain\&quot;, \&quot;mx\&quot;, \&quot;classification\&quot;]. | [optional] 
**Email** | Pointer to **string** | Email is the address to evaluate. | [optional] 
**FromEmail** | Pointer to **string** | FromEmail sets sender used in MAIL FROM for SMTP probe when explicitly enabled. | [optional] 
**HelloName** | Pointer to **string** | HelloName sets domain used in EHLO for SMTP probe when explicitly enabled. | [optional] 
**ProxyUri** | Pointer to **string** | ProxyURI sets optional proxy for network lookups. | [optional] 
**Suggestions** | Pointer to **bool** | Suggestions requests domain typo suggestions without altering the original email. | [optional] 

## Methods

### NewEmailVerificationInput

`func NewEmailVerificationInput() *EmailVerificationInput`

NewEmailVerificationInput instantiates a new EmailVerificationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailVerificationInputWithDefaults

`func NewEmailVerificationInputWithDefaults() *EmailVerificationInput`

NewEmailVerificationInputWithDefaults instantiates a new EmailVerificationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *EmailVerificationInput) GetChecks() []string`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *EmailVerificationInput) GetChecksOk() (*[]string, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *EmailVerificationInput) SetChecks(v []string)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *EmailVerificationInput) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetEmail

`func (o *EmailVerificationInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailVerificationInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailVerificationInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailVerificationInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFromEmail

`func (o *EmailVerificationInput) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *EmailVerificationInput) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *EmailVerificationInput) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *EmailVerificationInput) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetHelloName

`func (o *EmailVerificationInput) GetHelloName() string`

GetHelloName returns the HelloName field if non-nil, zero value otherwise.

### GetHelloNameOk

`func (o *EmailVerificationInput) GetHelloNameOk() (*string, bool)`

GetHelloNameOk returns a tuple with the HelloName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloName

`func (o *EmailVerificationInput) SetHelloName(v string)`

SetHelloName sets HelloName field to given value.

### HasHelloName

`func (o *EmailVerificationInput) HasHelloName() bool`

HasHelloName returns a boolean if a field has been set.

### GetProxyUri

`func (o *EmailVerificationInput) GetProxyUri() string`

GetProxyUri returns the ProxyUri field if non-nil, zero value otherwise.

### GetProxyUriOk

`func (o *EmailVerificationInput) GetProxyUriOk() (*string, bool)`

GetProxyUriOk returns a tuple with the ProxyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUri

`func (o *EmailVerificationInput) SetProxyUri(v string)`

SetProxyUri sets ProxyUri field to given value.

### HasProxyUri

`func (o *EmailVerificationInput) HasProxyUri() bool`

HasProxyUri returns a boolean if a field has been set.

### GetSuggestions

`func (o *EmailVerificationInput) GetSuggestions() bool`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *EmailVerificationInput) GetSuggestionsOk() (*bool, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *EmailVerificationInput) SetSuggestions(v bool)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *EmailVerificationInput) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


