# CloudflarePagesBuildConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildCommand** | Pointer to **string** | BuildCommand is what Cloudflare runs to build the site (\&quot;npm run build\&quot;). Omitted means no build step: the repository is published as it stands. | [optional] 
**DestinationDir** | Pointer to **string** | DestinationDir is the directory the build leaves the site in (\&quot;dist\&quot;), relative to RootDir. It is what gets served. | [optional] 
**RootDir** | Pointer to **string** | RootDir is where in the repository the build runs, for a project that is not at the repository root. Omitted means the root. | [optional] 

## Methods

### NewCloudflarePagesBuildConfig

`func NewCloudflarePagesBuildConfig() *CloudflarePagesBuildConfig`

NewCloudflarePagesBuildConfig instantiates a new CloudflarePagesBuildConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudflarePagesBuildConfigWithDefaults

`func NewCloudflarePagesBuildConfigWithDefaults() *CloudflarePagesBuildConfig`

NewCloudflarePagesBuildConfigWithDefaults instantiates a new CloudflarePagesBuildConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildCommand

`func (o *CloudflarePagesBuildConfig) GetBuildCommand() string`

GetBuildCommand returns the BuildCommand field if non-nil, zero value otherwise.

### GetBuildCommandOk

`func (o *CloudflarePagesBuildConfig) GetBuildCommandOk() (*string, bool)`

GetBuildCommandOk returns a tuple with the BuildCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCommand

`func (o *CloudflarePagesBuildConfig) SetBuildCommand(v string)`

SetBuildCommand sets BuildCommand field to given value.

### HasBuildCommand

`func (o *CloudflarePagesBuildConfig) HasBuildCommand() bool`

HasBuildCommand returns a boolean if a field has been set.

### GetDestinationDir

`func (o *CloudflarePagesBuildConfig) GetDestinationDir() string`

GetDestinationDir returns the DestinationDir field if non-nil, zero value otherwise.

### GetDestinationDirOk

`func (o *CloudflarePagesBuildConfig) GetDestinationDirOk() (*string, bool)`

GetDestinationDirOk returns a tuple with the DestinationDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDir

`func (o *CloudflarePagesBuildConfig) SetDestinationDir(v string)`

SetDestinationDir sets DestinationDir field to given value.

### HasDestinationDir

`func (o *CloudflarePagesBuildConfig) HasDestinationDir() bool`

HasDestinationDir returns a boolean if a field has been set.

### GetRootDir

`func (o *CloudflarePagesBuildConfig) GetRootDir() string`

GetRootDir returns the RootDir field if non-nil, zero value otherwise.

### GetRootDirOk

`func (o *CloudflarePagesBuildConfig) GetRootDirOk() (*string, bool)`

GetRootDirOk returns a tuple with the RootDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDir

`func (o *CloudflarePagesBuildConfig) SetRootDir(v string)`

SetRootDir sets RootDir field to given value.

### HasRootDir

`func (o *CloudflarePagesBuildConfig) HasRootDir() bool`

HasRootDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


