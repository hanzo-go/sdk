# CloudActionsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudActionRecord**](CloudActionRecord.md) | Data is the most-recent actions first, capped at listActionsLimit. | [optional] 

## Methods

### NewCloudActionsView

`func NewCloudActionsView() *CloudActionsView`

NewCloudActionsView instantiates a new CloudActionsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudActionsViewWithDefaults

`func NewCloudActionsViewWithDefaults() *CloudActionsView`

NewCloudActionsViewWithDefaults instantiates a new CloudActionsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudActionsView) GetData() []CloudActionRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudActionsView) GetDataOk() (*[]CloudActionRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudActionsView) SetData(v []CloudActionRecord)`

SetData sets Data field to given value.

### HasData

`func (o *CloudActionsView) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


