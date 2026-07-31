# CloudPromptList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudPromptMeta**](CloudPromptMeta.md) | Data is one row per prompt the org owns, each with its version numbers and taxonomy — never the template bodies. | [optional] 

## Methods

### NewCloudPromptList

`func NewCloudPromptList() *CloudPromptList`

NewCloudPromptList instantiates a new CloudPromptList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPromptListWithDefaults

`func NewCloudPromptListWithDefaults() *CloudPromptList`

NewCloudPromptListWithDefaults instantiates a new CloudPromptList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudPromptList) GetData() []CloudPromptMeta`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudPromptList) GetDataOk() (*[]CloudPromptMeta, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudPromptList) SetData(v []CloudPromptMeta)`

SetData sets Data field to given value.

### HasData

`func (o *CloudPromptList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


