# LicensingRevokeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | Pointer to [**LicensingRevocationEntry**](LicensingRevocationEntry.md) | Entry is the stored revocation, including who recorded it and when. | [optional] 
**Revoked** | Pointer to **bool** | Revoked is always true — a failure is an error status, not a false here. | [optional] 

## Methods

### NewLicensingRevokeResponse

`func NewLicensingRevokeResponse() *LicensingRevokeResponse`

NewLicensingRevokeResponse instantiates a new LicensingRevokeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingRevokeResponseWithDefaults

`func NewLicensingRevokeResponseWithDefaults() *LicensingRevokeResponse`

NewLicensingRevokeResponseWithDefaults instantiates a new LicensingRevokeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *LicensingRevokeResponse) GetEntry() LicensingRevocationEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *LicensingRevokeResponse) GetEntryOk() (*LicensingRevocationEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *LicensingRevokeResponse) SetEntry(v LicensingRevocationEntry)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *LicensingRevokeResponse) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetRevoked

`func (o *LicensingRevokeResponse) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *LicensingRevokeResponse) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *LicensingRevokeResponse) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *LicensingRevokeResponse) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


