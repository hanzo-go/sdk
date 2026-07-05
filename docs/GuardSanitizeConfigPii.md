# GuardSanitizeConfigPii

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**DetectSsn** | Pointer to **bool** |  | [optional] [default to true]
**DetectCreditCard** | Pointer to **bool** |  | [optional] [default to true]
**DetectEmail** | Pointer to **bool** |  | [optional] [default to true]
**DetectPhone** | Pointer to **bool** |  | [optional] [default to true]
**DetectIp** | Pointer to **bool** |  | [optional] [default to true]
**DetectApiKeys** | Pointer to **bool** |  | [optional] [default to true]
**RedactionFormat** | Pointer to **string** |  | [optional] [default to "[REDACTED:{TYPE}]"]

## Methods

### NewGuardSanitizeConfigPii

`func NewGuardSanitizeConfigPii() *GuardSanitizeConfigPii`

NewGuardSanitizeConfigPii instantiates a new GuardSanitizeConfigPii object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardSanitizeConfigPiiWithDefaults

`func NewGuardSanitizeConfigPiiWithDefaults() *GuardSanitizeConfigPii`

NewGuardSanitizeConfigPiiWithDefaults instantiates a new GuardSanitizeConfigPii object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GuardSanitizeConfigPii) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GuardSanitizeConfigPii) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GuardSanitizeConfigPii) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GuardSanitizeConfigPii) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDetectSsn

`func (o *GuardSanitizeConfigPii) GetDetectSsn() bool`

GetDetectSsn returns the DetectSsn field if non-nil, zero value otherwise.

### GetDetectSsnOk

`func (o *GuardSanitizeConfigPii) GetDetectSsnOk() (*bool, bool)`

GetDetectSsnOk returns a tuple with the DetectSsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectSsn

`func (o *GuardSanitizeConfigPii) SetDetectSsn(v bool)`

SetDetectSsn sets DetectSsn field to given value.

### HasDetectSsn

`func (o *GuardSanitizeConfigPii) HasDetectSsn() bool`

HasDetectSsn returns a boolean if a field has been set.

### GetDetectCreditCard

`func (o *GuardSanitizeConfigPii) GetDetectCreditCard() bool`

GetDetectCreditCard returns the DetectCreditCard field if non-nil, zero value otherwise.

### GetDetectCreditCardOk

`func (o *GuardSanitizeConfigPii) GetDetectCreditCardOk() (*bool, bool)`

GetDetectCreditCardOk returns a tuple with the DetectCreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectCreditCard

`func (o *GuardSanitizeConfigPii) SetDetectCreditCard(v bool)`

SetDetectCreditCard sets DetectCreditCard field to given value.

### HasDetectCreditCard

`func (o *GuardSanitizeConfigPii) HasDetectCreditCard() bool`

HasDetectCreditCard returns a boolean if a field has been set.

### GetDetectEmail

`func (o *GuardSanitizeConfigPii) GetDetectEmail() bool`

GetDetectEmail returns the DetectEmail field if non-nil, zero value otherwise.

### GetDetectEmailOk

`func (o *GuardSanitizeConfigPii) GetDetectEmailOk() (*bool, bool)`

GetDetectEmailOk returns a tuple with the DetectEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectEmail

`func (o *GuardSanitizeConfigPii) SetDetectEmail(v bool)`

SetDetectEmail sets DetectEmail field to given value.

### HasDetectEmail

`func (o *GuardSanitizeConfigPii) HasDetectEmail() bool`

HasDetectEmail returns a boolean if a field has been set.

### GetDetectPhone

`func (o *GuardSanitizeConfigPii) GetDetectPhone() bool`

GetDetectPhone returns the DetectPhone field if non-nil, zero value otherwise.

### GetDetectPhoneOk

`func (o *GuardSanitizeConfigPii) GetDetectPhoneOk() (*bool, bool)`

GetDetectPhoneOk returns a tuple with the DetectPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectPhone

`func (o *GuardSanitizeConfigPii) SetDetectPhone(v bool)`

SetDetectPhone sets DetectPhone field to given value.

### HasDetectPhone

`func (o *GuardSanitizeConfigPii) HasDetectPhone() bool`

HasDetectPhone returns a boolean if a field has been set.

### GetDetectIp

`func (o *GuardSanitizeConfigPii) GetDetectIp() bool`

GetDetectIp returns the DetectIp field if non-nil, zero value otherwise.

### GetDetectIpOk

`func (o *GuardSanitizeConfigPii) GetDetectIpOk() (*bool, bool)`

GetDetectIpOk returns a tuple with the DetectIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectIp

`func (o *GuardSanitizeConfigPii) SetDetectIp(v bool)`

SetDetectIp sets DetectIp field to given value.

### HasDetectIp

`func (o *GuardSanitizeConfigPii) HasDetectIp() bool`

HasDetectIp returns a boolean if a field has been set.

### GetDetectApiKeys

`func (o *GuardSanitizeConfigPii) GetDetectApiKeys() bool`

GetDetectApiKeys returns the DetectApiKeys field if non-nil, zero value otherwise.

### GetDetectApiKeysOk

`func (o *GuardSanitizeConfigPii) GetDetectApiKeysOk() (*bool, bool)`

GetDetectApiKeysOk returns a tuple with the DetectApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectApiKeys

`func (o *GuardSanitizeConfigPii) SetDetectApiKeys(v bool)`

SetDetectApiKeys sets DetectApiKeys field to given value.

### HasDetectApiKeys

`func (o *GuardSanitizeConfigPii) HasDetectApiKeys() bool`

HasDetectApiKeys returns a boolean if a field has been set.

### GetRedactionFormat

`func (o *GuardSanitizeConfigPii) GetRedactionFormat() string`

GetRedactionFormat returns the RedactionFormat field if non-nil, zero value otherwise.

### GetRedactionFormatOk

`func (o *GuardSanitizeConfigPii) GetRedactionFormatOk() (*string, bool)`

GetRedactionFormatOk returns a tuple with the RedactionFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactionFormat

`func (o *GuardSanitizeConfigPii) SetRedactionFormat(v string)`

SetRedactionFormat sets RedactionFormat field to given value.

### HasRedactionFormat

`func (o *GuardSanitizeConfigPii) HasRedactionFormat() bool`

HasRedactionFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


