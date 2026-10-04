# ProviderGithubInstallationsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallUrl** | Pointer to **string** | InstallURL is where to grant a new account, so a UI with an empty list has somewhere to send the reader instead of a dead end. | [optional] 
**Installations** | Pointer to [**[]ProviderGithubInstallationView**](ProviderGithubInstallationView.md) | Installations is every account the caller may see: the ones its own org has bound, or — for a super admin — every account the App is installed on. Never null; [] when none. | [optional] 

## Methods

### NewProviderGithubInstallationsOut

`func NewProviderGithubInstallationsOut() *ProviderGithubInstallationsOut`

NewProviderGithubInstallationsOut instantiates a new ProviderGithubInstallationsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubInstallationsOutWithDefaults

`func NewProviderGithubInstallationsOutWithDefaults() *ProviderGithubInstallationsOut`

NewProviderGithubInstallationsOutWithDefaults instantiates a new ProviderGithubInstallationsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallUrl

`func (o *ProviderGithubInstallationsOut) GetInstallUrl() string`

GetInstallUrl returns the InstallUrl field if non-nil, zero value otherwise.

### GetInstallUrlOk

`func (o *ProviderGithubInstallationsOut) GetInstallUrlOk() (*string, bool)`

GetInstallUrlOk returns a tuple with the InstallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallUrl

`func (o *ProviderGithubInstallationsOut) SetInstallUrl(v string)`

SetInstallUrl sets InstallUrl field to given value.

### HasInstallUrl

`func (o *ProviderGithubInstallationsOut) HasInstallUrl() bool`

HasInstallUrl returns a boolean if a field has been set.

### GetInstallations

`func (o *ProviderGithubInstallationsOut) GetInstallations() []ProviderGithubInstallationView`

GetInstallations returns the Installations field if non-nil, zero value otherwise.

### GetInstallationsOk

`func (o *ProviderGithubInstallationsOut) GetInstallationsOk() (*[]ProviderGithubInstallationView, bool)`

GetInstallationsOk returns a tuple with the Installations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallations

`func (o *ProviderGithubInstallationsOut) SetInstallations(v []ProviderGithubInstallationView)`

SetInstallations sets Installations field to given value.

### HasInstallations

`func (o *ProviderGithubInstallationsOut) HasInstallations() bool`

HasInstallations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


