# LicensingReleaseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Releases** | Pointer to [**[]LicensingRelease**](LicensingRelease.md) | Releases is the published releases, always an array and never null. | [optional] 

## Methods

### NewLicensingReleaseList

`func NewLicensingReleaseList() *LicensingReleaseList`

NewLicensingReleaseList instantiates a new LicensingReleaseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingReleaseListWithDefaults

`func NewLicensingReleaseListWithDefaults() *LicensingReleaseList`

NewLicensingReleaseListWithDefaults instantiates a new LicensingReleaseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleases

`func (o *LicensingReleaseList) GetReleases() []LicensingRelease`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *LicensingReleaseList) GetReleasesOk() (*[]LicensingRelease, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *LicensingReleaseList) SetReleases(v []LicensingRelease)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *LicensingReleaseList) HasReleases() bool`

HasReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


