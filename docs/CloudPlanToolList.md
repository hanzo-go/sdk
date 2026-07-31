# CloudPlanToolList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | Pointer to **[]interface{}** | Tools are the metered tools, each an opaque object exactly as the catalog emits it — typically name, billing unit and price. | [optional] 

## Methods

### NewCloudPlanToolList

`func NewCloudPlanToolList() *CloudPlanToolList`

NewCloudPlanToolList instantiates a new CloudPlanToolList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPlanToolListWithDefaults

`func NewCloudPlanToolListWithDefaults() *CloudPlanToolList`

NewCloudPlanToolListWithDefaults instantiates a new CloudPlanToolList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *CloudPlanToolList) GetTools() []interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CloudPlanToolList) GetToolsOk() (*[]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CloudPlanToolList) SetTools(v []interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *CloudPlanToolList) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


