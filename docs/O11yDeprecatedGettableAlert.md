# O11yDeprecatedGettableAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to [**O11yAlert**](O11yAlert.md) |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Receivers** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to [**O11yAlertStatus**](O11yAlertStatus.md) |  | [optional] 

## Methods

### NewO11yDeprecatedGettableAlert

`func NewO11yDeprecatedGettableAlert() *O11yDeprecatedGettableAlert`

NewO11yDeprecatedGettableAlert instantiates a new O11yDeprecatedGettableAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDeprecatedGettableAlertWithDefaults

`func NewO11yDeprecatedGettableAlertWithDefaults() *O11yDeprecatedGettableAlert`

NewO11yDeprecatedGettableAlertWithDefaults instantiates a new O11yDeprecatedGettableAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *O11yDeprecatedGettableAlert) GetAlert() O11yAlert`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *O11yDeprecatedGettableAlert) GetAlertOk() (*O11yAlert, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *O11yDeprecatedGettableAlert) SetAlert(v O11yAlert)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *O11yDeprecatedGettableAlert) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetFingerprint

`func (o *O11yDeprecatedGettableAlert) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *O11yDeprecatedGettableAlert) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *O11yDeprecatedGettableAlert) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *O11yDeprecatedGettableAlert) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetReceivers

`func (o *O11yDeprecatedGettableAlert) GetReceivers() []string`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *O11yDeprecatedGettableAlert) GetReceiversOk() (*[]string, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *O11yDeprecatedGettableAlert) SetReceivers(v []string)`

SetReceivers sets Receivers field to given value.

### HasReceivers

`func (o *O11yDeprecatedGettableAlert) HasReceivers() bool`

HasReceivers returns a boolean if a field has been set.

### GetStatus

`func (o *O11yDeprecatedGettableAlert) GetStatus() O11yAlertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yDeprecatedGettableAlert) GetStatusOk() (*O11yAlertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yDeprecatedGettableAlert) SetStatus(v O11yAlertStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yDeprecatedGettableAlert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


