# SbomSbomComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | Pointer to **string** | License is the FIRST license fact found for the component: its SPDX id, else its name, else the expression. Empty when the document declares none. | [optional] 
**Name** | Pointer to **string** | Name is the dependency&#39;s package name as CycloneDX records it. | [optional] 
**Purl** | Pointer to **string** | Purl is the package URL — the ecosystem-qualified identity a vulnerability feed can be joined on. | [optional] 
**Type** | Pointer to **string** | Type is the CycloneDX component type: library, application, framework, … | [optional] 
**Version** | Pointer to **string** | Version is the resolved version of that package. | [optional] 

## Methods

### NewSbomSbomComponent

`func NewSbomSbomComponent() *SbomSbomComponent`

NewSbomSbomComponent instantiates a new SbomSbomComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSbomSbomComponentWithDefaults

`func NewSbomSbomComponentWithDefaults() *SbomSbomComponent`

NewSbomSbomComponentWithDefaults instantiates a new SbomSbomComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicense

`func (o *SbomSbomComponent) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SbomSbomComponent) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SbomSbomComponent) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *SbomSbomComponent) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *SbomSbomComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SbomSbomComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SbomSbomComponent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SbomSbomComponent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPurl

`func (o *SbomSbomComponent) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *SbomSbomComponent) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *SbomSbomComponent) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *SbomSbomComponent) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetType

`func (o *SbomSbomComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SbomSbomComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SbomSbomComponent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SbomSbomComponent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *SbomSbomComponent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SbomSbomComponent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SbomSbomComponent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SbomSbomComponent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


