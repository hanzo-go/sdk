# LicensingVerifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | AppID is the brand the token runs under. | [optional] 
**Exp** | Pointer to **int32** | Exp is the token&#39;s expiry, Unix seconds. | [optional] 
**Features** | Pointer to **[]string** | Features are the capability grants the token carries. | [optional] 
**FingerprintBound** | Pointer to **bool** | Bound reports that the token carries a device binding. | [optional] 
**Holder** | Pointer to **string** | Holder is who the token was issued to. | [optional] 
**Nonce** | Pointer to **string** | Nonce uniquely identifies the token. | [optional] 
**Reason** | Pointer to **string** | Reason says why an invalid token was rejected. Empty when Valid. | [optional] 
**Revoked** | Pointer to **bool** | Revoked reports that the signature was good but the token has been revoked. | [optional] 
**Valid** | Pointer to **bool** | Valid is the single answer: signature, schema, expiry, app and revocation all passed. | [optional] 

## Methods

### NewLicensingVerifyResponse

`func NewLicensingVerifyResponse() *LicensingVerifyResponse`

NewLicensingVerifyResponse instantiates a new LicensingVerifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingVerifyResponseWithDefaults

`func NewLicensingVerifyResponseWithDefaults() *LicensingVerifyResponse`

NewLicensingVerifyResponseWithDefaults instantiates a new LicensingVerifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *LicensingVerifyResponse) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *LicensingVerifyResponse) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *LicensingVerifyResponse) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *LicensingVerifyResponse) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetExp

`func (o *LicensingVerifyResponse) GetExp() int32`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *LicensingVerifyResponse) GetExpOk() (*int32, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *LicensingVerifyResponse) SetExp(v int32)`

SetExp sets Exp field to given value.

### HasExp

`func (o *LicensingVerifyResponse) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetFeatures

`func (o *LicensingVerifyResponse) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *LicensingVerifyResponse) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *LicensingVerifyResponse) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *LicensingVerifyResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFingerprintBound

`func (o *LicensingVerifyResponse) GetFingerprintBound() bool`

GetFingerprintBound returns the FingerprintBound field if non-nil, zero value otherwise.

### GetFingerprintBoundOk

`func (o *LicensingVerifyResponse) GetFingerprintBoundOk() (*bool, bool)`

GetFingerprintBoundOk returns a tuple with the FingerprintBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintBound

`func (o *LicensingVerifyResponse) SetFingerprintBound(v bool)`

SetFingerprintBound sets FingerprintBound field to given value.

### HasFingerprintBound

`func (o *LicensingVerifyResponse) HasFingerprintBound() bool`

HasFingerprintBound returns a boolean if a field has been set.

### GetHolder

`func (o *LicensingVerifyResponse) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *LicensingVerifyResponse) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *LicensingVerifyResponse) SetHolder(v string)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *LicensingVerifyResponse) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetNonce

`func (o *LicensingVerifyResponse) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *LicensingVerifyResponse) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *LicensingVerifyResponse) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *LicensingVerifyResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetReason

`func (o *LicensingVerifyResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LicensingVerifyResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LicensingVerifyResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *LicensingVerifyResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRevoked

`func (o *LicensingVerifyResponse) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *LicensingVerifyResponse) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *LicensingVerifyResponse) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *LicensingVerifyResponse) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetValid

`func (o *LicensingVerifyResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *LicensingVerifyResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *LicensingVerifyResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *LicensingVerifyResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


