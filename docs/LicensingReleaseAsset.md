# LicensingReleaseAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CosignCert** | Pointer to **string** |  | [optional] 
**CosignSignature** | Pointer to **string** |  | [optional] 
**DownloadUrl** | Pointer to **string** | DownloadURL is a short-lived signed URL to the artifact bytes. The scaffold returns the ArtifactRef as-is; production issues a signed URL. | [optional] 
**Release** | Pointer to [**LicensingRelease**](LicensingRelease.md) |  | [optional] 

## Methods

### NewLicensingReleaseAsset

`func NewLicensingReleaseAsset() *LicensingReleaseAsset`

NewLicensingReleaseAsset instantiates a new LicensingReleaseAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingReleaseAssetWithDefaults

`func NewLicensingReleaseAssetWithDefaults() *LicensingReleaseAsset`

NewLicensingReleaseAssetWithDefaults instantiates a new LicensingReleaseAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCosignCert

`func (o *LicensingReleaseAsset) GetCosignCert() string`

GetCosignCert returns the CosignCert field if non-nil, zero value otherwise.

### GetCosignCertOk

`func (o *LicensingReleaseAsset) GetCosignCertOk() (*string, bool)`

GetCosignCertOk returns a tuple with the CosignCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosignCert

`func (o *LicensingReleaseAsset) SetCosignCert(v string)`

SetCosignCert sets CosignCert field to given value.

### HasCosignCert

`func (o *LicensingReleaseAsset) HasCosignCert() bool`

HasCosignCert returns a boolean if a field has been set.

### GetCosignSignature

`func (o *LicensingReleaseAsset) GetCosignSignature() string`

GetCosignSignature returns the CosignSignature field if non-nil, zero value otherwise.

### GetCosignSignatureOk

`func (o *LicensingReleaseAsset) GetCosignSignatureOk() (*string, bool)`

GetCosignSignatureOk returns a tuple with the CosignSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosignSignature

`func (o *LicensingReleaseAsset) SetCosignSignature(v string)`

SetCosignSignature sets CosignSignature field to given value.

### HasCosignSignature

`func (o *LicensingReleaseAsset) HasCosignSignature() bool`

HasCosignSignature returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *LicensingReleaseAsset) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *LicensingReleaseAsset) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *LicensingReleaseAsset) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *LicensingReleaseAsset) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetRelease

`func (o *LicensingReleaseAsset) GetRelease() LicensingRelease`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *LicensingReleaseAsset) GetReleaseOk() (*LicensingRelease, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *LicensingReleaseAsset) SetRelease(v LicensingRelease)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *LicensingReleaseAsset) HasRelease() bool`

HasRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


