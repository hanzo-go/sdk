# LicensingIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | Pointer to **string** | Fingerprint is a previously-registered device binding value, as returned by POST /v1/licensing/fingerprint. Leave it empty and pass Signals to bind the device at issue time instead. | [optional] 
**Holder** | Pointer to **string** | Holder overrides who the token is issued to; defaults to the caller&#39;s own validated subject. It NAMES the token&#39;s bearer and grants nothing on its own — the entitlement checked is always the caller&#39;s org&#39;s. | [optional] 
**Product** | **string** | Product is the licensed commerce product the caller wants a token for. | 
**Release** | Pointer to **string** | Release scopes the token to one signed binary release, recorded as a \&quot;release:&lt;id&gt;\&quot; feature so a single bad release can be revoked on its own. | [optional] 
**Signals** | Pointer to [**LicensingDeviceSignals**](LicensingDeviceSignals.md) | Signals binds the device at issue time, as an alternative to a pre-registered fingerprint. The raw signals are never stored or echoed — they are folded immediately into the one-way binding value. | [optional] 
**TtlSeconds** | Pointer to **int32** | TTLSeconds requests a token lifetime in seconds. It is clamped to the deployment maximum AND to the entitlement&#39;s own expiry — a token never outlives the subscription that paid for it. | [optional] 

## Methods

### NewLicensingIssueRequest

`func NewLicensingIssueRequest(product string, ) *LicensingIssueRequest`

NewLicensingIssueRequest instantiates a new LicensingIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingIssueRequestWithDefaults

`func NewLicensingIssueRequestWithDefaults() *LicensingIssueRequest`

NewLicensingIssueRequestWithDefaults instantiates a new LicensingIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *LicensingIssueRequest) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *LicensingIssueRequest) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *LicensingIssueRequest) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *LicensingIssueRequest) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHolder

`func (o *LicensingIssueRequest) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *LicensingIssueRequest) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *LicensingIssueRequest) SetHolder(v string)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *LicensingIssueRequest) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetProduct

`func (o *LicensingIssueRequest) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LicensingIssueRequest) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LicensingIssueRequest) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetRelease

`func (o *LicensingIssueRequest) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *LicensingIssueRequest) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *LicensingIssueRequest) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *LicensingIssueRequest) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetSignals

`func (o *LicensingIssueRequest) GetSignals() LicensingDeviceSignals`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *LicensingIssueRequest) GetSignalsOk() (*LicensingDeviceSignals, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *LicensingIssueRequest) SetSignals(v LicensingDeviceSignals)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *LicensingIssueRequest) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### GetTtlSeconds

`func (o *LicensingIssueRequest) GetTtlSeconds() int32`

GetTtlSeconds returns the TtlSeconds field if non-nil, zero value otherwise.

### GetTtlSecondsOk

`func (o *LicensingIssueRequest) GetTtlSecondsOk() (*int32, bool)`

GetTtlSecondsOk returns a tuple with the TtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSeconds

`func (o *LicensingIssueRequest) SetTtlSeconds(v int32)`

SetTtlSeconds sets TtlSeconds field to given value.

### HasTtlSeconds

`func (o *LicensingIssueRequest) HasTtlSeconds() bool`

HasTtlSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


