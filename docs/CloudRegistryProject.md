# CloudRegistryProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to **int32** | Images is how many of the org&#39;s repositories the OCI catalog holds. | [optional] 
**Packages** | Pointer to **int32** | Packages is how many of the org&#39;s packages the npm registry reports. | [optional] 
**Project** | Pointer to **string** | Project is the namespace: the org&#39;s slug, which prefixes its image names and scopes its npm packages. | [optional] 

## Methods

### NewCloudRegistryProject

`func NewCloudRegistryProject() *CloudRegistryProject`

NewCloudRegistryProject instantiates a new CloudRegistryProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegistryProjectWithDefaults

`func NewCloudRegistryProjectWithDefaults() *CloudRegistryProject`

NewCloudRegistryProjectWithDefaults instantiates a new CloudRegistryProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *CloudRegistryProject) GetImages() int32`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CloudRegistryProject) GetImagesOk() (*int32, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CloudRegistryProject) SetImages(v int32)`

SetImages sets Images field to given value.

### HasImages

`func (o *CloudRegistryProject) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetPackages

`func (o *CloudRegistryProject) GetPackages() int32`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *CloudRegistryProject) GetPackagesOk() (*int32, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *CloudRegistryProject) SetPackages(v int32)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *CloudRegistryProject) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetProject

`func (o *CloudRegistryProject) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudRegistryProject) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudRegistryProject) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudRegistryProject) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


