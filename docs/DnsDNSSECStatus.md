# DnsDNSSECStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Algorithm** | Pointer to **string** |  | [optional] [default to "ECDSAP256SHA256"]
**DsRecord** | Pointer to **string** | DS record to add at registrar | [optional] 
**KeyTag** | Pointer to **int32** |  | [optional] 
**DigestType** | Pointer to **string** |  | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 

## Methods

### NewDnsDNSSECStatus

`func NewDnsDNSSECStatus() *DnsDNSSECStatus`

NewDnsDNSSECStatus instantiates a new DnsDNSSECStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsDNSSECStatusWithDefaults

`func NewDnsDNSSECStatusWithDefaults() *DnsDNSSECStatus`

NewDnsDNSSECStatusWithDefaults instantiates a new DnsDNSSECStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DnsDNSSECStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DnsDNSSECStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DnsDNSSECStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DnsDNSSECStatus) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *DnsDNSSECStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DnsDNSSECStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DnsDNSSECStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DnsDNSSECStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAlgorithm

`func (o *DnsDNSSECStatus) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *DnsDNSSECStatus) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *DnsDNSSECStatus) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *DnsDNSSECStatus) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetDsRecord

`func (o *DnsDNSSECStatus) GetDsRecord() string`

GetDsRecord returns the DsRecord field if non-nil, zero value otherwise.

### GetDsRecordOk

`func (o *DnsDNSSECStatus) GetDsRecordOk() (*string, bool)`

GetDsRecordOk returns a tuple with the DsRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsRecord

`func (o *DnsDNSSECStatus) SetDsRecord(v string)`

SetDsRecord sets DsRecord field to given value.

### HasDsRecord

`func (o *DnsDNSSECStatus) HasDsRecord() bool`

HasDsRecord returns a boolean if a field has been set.

### GetKeyTag

`func (o *DnsDNSSECStatus) GetKeyTag() int32`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *DnsDNSSECStatus) GetKeyTagOk() (*int32, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *DnsDNSSECStatus) SetKeyTag(v int32)`

SetKeyTag sets KeyTag field to given value.

### HasKeyTag

`func (o *DnsDNSSECStatus) HasKeyTag() bool`

HasKeyTag returns a boolean if a field has been set.

### GetDigestType

`func (o *DnsDNSSECStatus) GetDigestType() string`

GetDigestType returns the DigestType field if non-nil, zero value otherwise.

### GetDigestTypeOk

`func (o *DnsDNSSECStatus) GetDigestTypeOk() (*string, bool)`

GetDigestTypeOk returns a tuple with the DigestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestType

`func (o *DnsDNSSECStatus) SetDigestType(v string)`

SetDigestType sets DigestType field to given value.

### HasDigestType

`func (o *DnsDNSSECStatus) HasDigestType() bool`

HasDigestType returns a boolean if a field has been set.

### GetDigest

`func (o *DnsDNSSECStatus) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DnsDNSSECStatus) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DnsDNSSECStatus) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *DnsDNSSECStatus) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetPublicKey

`func (o *DnsDNSSECStatus) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DnsDNSSECStatus) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DnsDNSSECStatus) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *DnsDNSSECStatus) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


