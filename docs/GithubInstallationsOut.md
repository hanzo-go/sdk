# GithubInstallationsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallUrl** | Pointer to **string** | InstallURL is where to grant a new account, so a UI with an empty list has somewhere to send the reader instead of a dead end. | [optional] 
**Installations** | Pointer to [**[]GithubInstallationView**](GithubInstallationView.md) | Installations is every account the caller may see: the ones its own org has bound, or — for a super admin — every account the App is installed on. Never null; [] when none. | [optional] 

## Methods

### NewGithubInstallationsOut

`func NewGithubInstallationsOut() *GithubInstallationsOut`

NewGithubInstallationsOut instantiates a new GithubInstallationsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubInstallationsOutWithDefaults

`func NewGithubInstallationsOutWithDefaults() *GithubInstallationsOut`

NewGithubInstallationsOutWithDefaults instantiates a new GithubInstallationsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallUrl

`func (o *GithubInstallationsOut) GetInstallUrl() string`

GetInstallUrl returns the InstallUrl field if non-nil, zero value otherwise.

### GetInstallUrlOk

`func (o *GithubInstallationsOut) GetInstallUrlOk() (*string, bool)`

GetInstallUrlOk returns a tuple with the InstallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallUrl

`func (o *GithubInstallationsOut) SetInstallUrl(v string)`

SetInstallUrl sets InstallUrl field to given value.

### HasInstallUrl

`func (o *GithubInstallationsOut) HasInstallUrl() bool`

HasInstallUrl returns a boolean if a field has been set.

### GetInstallations

`func (o *GithubInstallationsOut) GetInstallations() []GithubInstallationView`

GetInstallations returns the Installations field if non-nil, zero value otherwise.

### GetInstallationsOk

`func (o *GithubInstallationsOut) GetInstallationsOk() (*[]GithubInstallationView, bool)`

GetInstallationsOk returns a tuple with the Installations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallations

`func (o *GithubInstallationsOut) SetInstallations(v []GithubInstallationView)`

SetInstallations sets Installations field to given value.

### HasInstallations

`func (o *GithubInstallationsOut) HasInstallations() bool`

HasInstallations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


