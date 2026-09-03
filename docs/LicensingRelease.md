# LicensingRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | AppID scopes the release to an app build (\&quot;hanzo\&quot; | \&quot;lux\&quot; | \&quot;zoo\&quot;). | [optional] 
**ArtifactRef** | Pointer to **string** | ArtifactRef is where the binary lives (object-store key / OCI ref / path). | [optional] 
**CosignCert** | Pointer to **string** | CosignCert is the cosign/Fulcio cert (keyless) or public key ref used to verify CosignSignature. The download response hands this to the client. | [optional] 
**CosignSignature** | Pointer to **string** | CosignSignature is the base64 cosign signature over the artifact digest. | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** | ID is the release identifier, e.g. \&quot;engine-rocm-0.4.2-linux-amd64\&quot;. The accelerator belongs here — one product is built several ways — and never in Product below. | [optional] 
**MinFeatures** | Pointer to **[]string** | MinFeatures, when set, are features the license must include to download. | [optional] 
**Platform** | Pointer to **string** | Platform is \&quot;&lt;os&gt;/&lt;arch&gt;\&quot;, e.g. \&quot;linux/amd64\&quot;. | [optional] 
**Product** | Pointer to **string** | Product is the licensed product this artifact belongs to (commerce SKU): \&quot;engine\&quot; for every one of those builds. | [optional] 
**Sha256** | Pointer to **string** | SHA256 is the hex digest of the artifact (integrity + cosign subject). | [optional] 
**Version** | Pointer to **string** | Version is the semantic version of the binary. | [optional] 
**Yanked** | Pointer to **bool** | Yanked marks a pulled release (download refused; tokens may be revoked release-scoped too). | [optional] 

## Methods

### NewLicensingRelease

`func NewLicensingRelease() *LicensingRelease`

NewLicensingRelease instantiates a new LicensingRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingReleaseWithDefaults

`func NewLicensingReleaseWithDefaults() *LicensingRelease`

NewLicensingReleaseWithDefaults instantiates a new LicensingRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *LicensingRelease) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *LicensingRelease) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *LicensingRelease) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *LicensingRelease) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetArtifactRef

`func (o *LicensingRelease) GetArtifactRef() string`

GetArtifactRef returns the ArtifactRef field if non-nil, zero value otherwise.

### GetArtifactRefOk

`func (o *LicensingRelease) GetArtifactRefOk() (*string, bool)`

GetArtifactRefOk returns a tuple with the ArtifactRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactRef

`func (o *LicensingRelease) SetArtifactRef(v string)`

SetArtifactRef sets ArtifactRef field to given value.

### HasArtifactRef

`func (o *LicensingRelease) HasArtifactRef() bool`

HasArtifactRef returns a boolean if a field has been set.

### GetCosignCert

`func (o *LicensingRelease) GetCosignCert() string`

GetCosignCert returns the CosignCert field if non-nil, zero value otherwise.

### GetCosignCertOk

`func (o *LicensingRelease) GetCosignCertOk() (*string, bool)`

GetCosignCertOk returns a tuple with the CosignCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosignCert

`func (o *LicensingRelease) SetCosignCert(v string)`

SetCosignCert sets CosignCert field to given value.

### HasCosignCert

`func (o *LicensingRelease) HasCosignCert() bool`

HasCosignCert returns a boolean if a field has been set.

### GetCosignSignature

`func (o *LicensingRelease) GetCosignSignature() string`

GetCosignSignature returns the CosignSignature field if non-nil, zero value otherwise.

### GetCosignSignatureOk

`func (o *LicensingRelease) GetCosignSignatureOk() (*string, bool)`

GetCosignSignatureOk returns a tuple with the CosignSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosignSignature

`func (o *LicensingRelease) SetCosignSignature(v string)`

SetCosignSignature sets CosignSignature field to given value.

### HasCosignSignature

`func (o *LicensingRelease) HasCosignSignature() bool`

HasCosignSignature returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LicensingRelease) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LicensingRelease) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LicensingRelease) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LicensingRelease) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *LicensingRelease) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicensingRelease) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicensingRelease) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicensingRelease) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMinFeatures

`func (o *LicensingRelease) GetMinFeatures() []string`

GetMinFeatures returns the MinFeatures field if non-nil, zero value otherwise.

### GetMinFeaturesOk

`func (o *LicensingRelease) GetMinFeaturesOk() (*[]string, bool)`

GetMinFeaturesOk returns a tuple with the MinFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFeatures

`func (o *LicensingRelease) SetMinFeatures(v []string)`

SetMinFeatures sets MinFeatures field to given value.

### HasMinFeatures

`func (o *LicensingRelease) HasMinFeatures() bool`

HasMinFeatures returns a boolean if a field has been set.

### GetPlatform

`func (o *LicensingRelease) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *LicensingRelease) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *LicensingRelease) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *LicensingRelease) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProduct

`func (o *LicensingRelease) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LicensingRelease) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LicensingRelease) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *LicensingRelease) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSha256

`func (o *LicensingRelease) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *LicensingRelease) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *LicensingRelease) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *LicensingRelease) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetVersion

`func (o *LicensingRelease) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LicensingRelease) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LicensingRelease) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LicensingRelease) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetYanked

`func (o *LicensingRelease) GetYanked() bool`

GetYanked returns the Yanked field if non-nil, zero value otherwise.

### GetYankedOk

`func (o *LicensingRelease) GetYankedOk() (*bool, bool)`

GetYankedOk returns a tuple with the Yanked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYanked

`func (o *LicensingRelease) SetYanked(v bool)`

SetYanked sets Yanked field to given value.

### HasYanked

`func (o *LicensingRelease) HasYanked() bool`

HasYanked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


