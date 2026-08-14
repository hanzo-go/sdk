# LicensingFingerprintResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | Pointer to **string** | Fingerprint is the OPAQUE binding value to pass to POST /v1/licensing/issue. It is one-way: the raw signals cannot be recovered from it and are never echoed back. | [optional] 
**Version** | Pointer to **string** | Version is the binding algorithm revision, so a stored fingerprint stays recognizable across a recipe rotation. | [optional] 

## Methods

### NewLicensingFingerprintResponse

`func NewLicensingFingerprintResponse() *LicensingFingerprintResponse`

NewLicensingFingerprintResponse instantiates a new LicensingFingerprintResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingFingerprintResponseWithDefaults

`func NewLicensingFingerprintResponseWithDefaults() *LicensingFingerprintResponse`

NewLicensingFingerprintResponseWithDefaults instantiates a new LicensingFingerprintResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *LicensingFingerprintResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *LicensingFingerprintResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *LicensingFingerprintResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *LicensingFingerprintResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetVersion

`func (o *LicensingFingerprintResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LicensingFingerprintResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LicensingFingerprintResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LicensingFingerprintResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


