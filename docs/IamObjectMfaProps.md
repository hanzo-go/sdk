# IamObjectMfaProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**IsPreferred** | Pointer to **bool** |  | [optional] 
**MfaRememberInHours** | Pointer to **int64** |  | [optional] 
**MfaType** | Pointer to **string** |  | [optional] 
**RecoveryCodes** | Pointer to **[]string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectMfaProps

`func NewIamObjectMfaProps() *IamObjectMfaProps`

NewIamObjectMfaProps instantiates a new IamObjectMfaProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectMfaPropsWithDefaults

`func NewIamObjectMfaPropsWithDefaults() *IamObjectMfaProps`

NewIamObjectMfaPropsWithDefaults instantiates a new IamObjectMfaProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *IamObjectMfaProps) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *IamObjectMfaProps) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *IamObjectMfaProps) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *IamObjectMfaProps) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEnabled

`func (o *IamObjectMfaProps) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IamObjectMfaProps) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IamObjectMfaProps) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IamObjectMfaProps) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsPreferred

`func (o *IamObjectMfaProps) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *IamObjectMfaProps) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *IamObjectMfaProps) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *IamObjectMfaProps) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetMfaRememberInHours

`func (o *IamObjectMfaProps) GetMfaRememberInHours() int64`

GetMfaRememberInHours returns the MfaRememberInHours field if non-nil, zero value otherwise.

### GetMfaRememberInHoursOk

`func (o *IamObjectMfaProps) GetMfaRememberInHoursOk() (*int64, bool)`

GetMfaRememberInHoursOk returns a tuple with the MfaRememberInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRememberInHours

`func (o *IamObjectMfaProps) SetMfaRememberInHours(v int64)`

SetMfaRememberInHours sets MfaRememberInHours field to given value.

### HasMfaRememberInHours

`func (o *IamObjectMfaProps) HasMfaRememberInHours() bool`

HasMfaRememberInHours returns a boolean if a field has been set.

### GetMfaType

`func (o *IamObjectMfaProps) GetMfaType() string`

GetMfaType returns the MfaType field if non-nil, zero value otherwise.

### GetMfaTypeOk

`func (o *IamObjectMfaProps) GetMfaTypeOk() (*string, bool)`

GetMfaTypeOk returns a tuple with the MfaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaType

`func (o *IamObjectMfaProps) SetMfaType(v string)`

SetMfaType sets MfaType field to given value.

### HasMfaType

`func (o *IamObjectMfaProps) HasMfaType() bool`

HasMfaType returns a boolean if a field has been set.

### GetRecoveryCodes

`func (o *IamObjectMfaProps) GetRecoveryCodes() []string`

GetRecoveryCodes returns the RecoveryCodes field if non-nil, zero value otherwise.

### GetRecoveryCodesOk

`func (o *IamObjectMfaProps) GetRecoveryCodesOk() (*[]string, bool)`

GetRecoveryCodesOk returns a tuple with the RecoveryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCodes

`func (o *IamObjectMfaProps) SetRecoveryCodes(v []string)`

SetRecoveryCodes sets RecoveryCodes field to given value.

### HasRecoveryCodes

`func (o *IamObjectMfaProps) HasRecoveryCodes() bool`

HasRecoveryCodes returns a boolean if a field has been set.

### GetSecret

`func (o *IamObjectMfaProps) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IamObjectMfaProps) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IamObjectMfaProps) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IamObjectMfaProps) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUrl

`func (o *IamObjectMfaProps) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IamObjectMfaProps) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IamObjectMfaProps) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IamObjectMfaProps) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


