# O11yAzureLogsCollectionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryGroups** | Pointer to **[]string** | List of categories to enable for diagnostic settings, to start with it will have &#39;allLogs&#39; and no filtering. | [optional] 

## Methods

### NewO11yAzureLogsCollectionStrategy

`func NewO11yAzureLogsCollectionStrategy() *O11yAzureLogsCollectionStrategy`

NewO11yAzureLogsCollectionStrategy instantiates a new O11yAzureLogsCollectionStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAzureLogsCollectionStrategyWithDefaults

`func NewO11yAzureLogsCollectionStrategyWithDefaults() *O11yAzureLogsCollectionStrategy`

NewO11yAzureLogsCollectionStrategyWithDefaults instantiates a new O11yAzureLogsCollectionStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryGroups

`func (o *O11yAzureLogsCollectionStrategy) GetCategoryGroups() []string`

GetCategoryGroups returns the CategoryGroups field if non-nil, zero value otherwise.

### GetCategoryGroupsOk

`func (o *O11yAzureLogsCollectionStrategy) GetCategoryGroupsOk() (*[]string, bool)`

GetCategoryGroupsOk returns a tuple with the CategoryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGroups

`func (o *O11yAzureLogsCollectionStrategy) SetCategoryGroups(v []string)`

SetCategoryGroups sets CategoryGroups field to given value.

### HasCategoryGroups

`func (o *O11yAzureLogsCollectionStrategy) HasCategoryGroups() bool`

HasCategoryGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


