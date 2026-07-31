# CloudAutoUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Flow** | Pointer to **string** | Flow is the flow&#39;s id, taken from the path. | [optional] 
**Name** | Pointer to **string** | Name renames the flow when present. | [optional] 

## Methods

### NewCloudAutoUpdate

`func NewCloudAutoUpdate() *CloudAutoUpdate`

NewCloudAutoUpdate instantiates a new CloudAutoUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAutoUpdateWithDefaults

`func NewCloudAutoUpdateWithDefaults() *CloudAutoUpdate`

NewCloudAutoUpdateWithDefaults instantiates a new CloudAutoUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudAutoUpdate) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudAutoUpdate) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudAutoUpdate) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CloudAutoUpdate) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *CloudAutoUpdate) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CloudAutoUpdate) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetFlow

`func (o *CloudAutoUpdate) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *CloudAutoUpdate) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *CloudAutoUpdate) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *CloudAutoUpdate) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetName

`func (o *CloudAutoUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAutoUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAutoUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAutoUpdate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


