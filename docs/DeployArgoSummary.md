# DeployArgoSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to **[]string** | Images are the container images the application runs. One entry for an App CR, built from its spec.image as \&quot;repository:tag\&quot; — the bare repository when it declares no tag, and absent when it declares neither. Absent on a CD row, which tracks commits rather than images. | [optional] 

## Methods

### NewDeployArgoSummary

`func NewDeployArgoSummary() *DeployArgoSummary`

NewDeployArgoSummary instantiates a new DeployArgoSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoSummaryWithDefaults

`func NewDeployArgoSummaryWithDefaults() *DeployArgoSummary`

NewDeployArgoSummaryWithDefaults instantiates a new DeployArgoSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *DeployArgoSummary) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *DeployArgoSummary) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *DeployArgoSummary) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *DeployArgoSummary) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


