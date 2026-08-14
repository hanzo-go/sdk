# LicensingIssueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | AppID is the brand this token runs under (\&quot;hanzo\&quot; | \&quot;lux\&quot; | \&quot;zoo\&quot;). The engine refuses a token whose app_id is not the one it was built for. | [optional] 
**Exp** | Pointer to **int32** | Exp is the token&#39;s expiry, Unix seconds. | [optional] 
**Features** | Pointer to **[]string** | Features are the capability grants copied verbatim from the plan the org bought. The engine enforces exactly these. | [optional] 
**FingerprintBound** | Pointer to **bool** | Bound reports whether a device fingerprint was folded into the token. An unbound token runs on any machine; a bound one runs only on the machine it was bound to. | [optional] 
**Holder** | Pointer to **string** | Holder is who the token was issued to. | [optional] 
**Nonce** | Pointer to **string** | Nonce uniquely identifies this token, and is what a per-token revocation names. | [optional] 
**Token** | Pointer to **string** | Token is the signed license, &#x60;base64url(payload).base64url(ed25519_sig)&#x60;. It is the credential the engine runs on — treat it as a secret. | [optional] 

## Methods

### NewLicensingIssueResponse

`func NewLicensingIssueResponse() *LicensingIssueResponse`

NewLicensingIssueResponse instantiates a new LicensingIssueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingIssueResponseWithDefaults

`func NewLicensingIssueResponseWithDefaults() *LicensingIssueResponse`

NewLicensingIssueResponseWithDefaults instantiates a new LicensingIssueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *LicensingIssueResponse) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *LicensingIssueResponse) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *LicensingIssueResponse) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *LicensingIssueResponse) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetExp

`func (o *LicensingIssueResponse) GetExp() int32`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *LicensingIssueResponse) GetExpOk() (*int32, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *LicensingIssueResponse) SetExp(v int32)`

SetExp sets Exp field to given value.

### HasExp

`func (o *LicensingIssueResponse) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetFeatures

`func (o *LicensingIssueResponse) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *LicensingIssueResponse) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *LicensingIssueResponse) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *LicensingIssueResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFingerprintBound

`func (o *LicensingIssueResponse) GetFingerprintBound() bool`

GetFingerprintBound returns the FingerprintBound field if non-nil, zero value otherwise.

### GetFingerprintBoundOk

`func (o *LicensingIssueResponse) GetFingerprintBoundOk() (*bool, bool)`

GetFingerprintBoundOk returns a tuple with the FingerprintBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintBound

`func (o *LicensingIssueResponse) SetFingerprintBound(v bool)`

SetFingerprintBound sets FingerprintBound field to given value.

### HasFingerprintBound

`func (o *LicensingIssueResponse) HasFingerprintBound() bool`

HasFingerprintBound returns a boolean if a field has been set.

### GetHolder

`func (o *LicensingIssueResponse) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *LicensingIssueResponse) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *LicensingIssueResponse) SetHolder(v string)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *LicensingIssueResponse) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetNonce

`func (o *LicensingIssueResponse) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *LicensingIssueResponse) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *LicensingIssueResponse) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *LicensingIssueResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetToken

`func (o *LicensingIssueResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LicensingIssueResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LicensingIssueResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LicensingIssueResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


