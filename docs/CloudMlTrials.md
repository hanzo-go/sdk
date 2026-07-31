# CloudMlTrials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Experiment** | Pointer to **string** | Experiment is the experiment the trials belong to, echoed from the path. | [optional] 
**Items** | Pointer to [**[]CloudMlResource**](CloudMlResource.md) | Items is one entry per Trial, in the list shape (no spec). | [optional] 

## Methods

### NewCloudMlTrials

`func NewCloudMlTrials() *CloudMlTrials`

NewCloudMlTrials instantiates a new CloudMlTrials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMlTrialsWithDefaults

`func NewCloudMlTrialsWithDefaults() *CloudMlTrials`

NewCloudMlTrialsWithDefaults instantiates a new CloudMlTrials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperiment

`func (o *CloudMlTrials) GetExperiment() string`

GetExperiment returns the Experiment field if non-nil, zero value otherwise.

### GetExperimentOk

`func (o *CloudMlTrials) GetExperimentOk() (*string, bool)`

GetExperimentOk returns a tuple with the Experiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiment

`func (o *CloudMlTrials) SetExperiment(v string)`

SetExperiment sets Experiment field to given value.

### HasExperiment

`func (o *CloudMlTrials) HasExperiment() bool`

HasExperiment returns a boolean if a field has been set.

### GetItems

`func (o *CloudMlTrials) GetItems() []CloudMlResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudMlTrials) GetItemsOk() (*[]CloudMlResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudMlTrials) SetItems(v []CloudMlResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudMlTrials) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


