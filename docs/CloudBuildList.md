# CloudBuildList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Builds** | Pointer to [**[]CloudBuildSummary**](CloudBuildSummary.md) | Builds is every published build, most recently updated first. | [optional] 

## Methods

### NewCloudBuildList

`func NewCloudBuildList() *CloudBuildList`

NewCloudBuildList instantiates a new CloudBuildList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBuildListWithDefaults

`func NewCloudBuildListWithDefaults() *CloudBuildList`

NewCloudBuildListWithDefaults instantiates a new CloudBuildList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilds

`func (o *CloudBuildList) GetBuilds() []CloudBuildSummary`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *CloudBuildList) GetBuildsOk() (*[]CloudBuildSummary, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *CloudBuildList) SetBuilds(v []CloudBuildSummary)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *CloudBuildList) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


