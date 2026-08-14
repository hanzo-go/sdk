# AutoUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Flow** | Pointer to **string** | Flow is the flow&#39;s id, taken from the path. | [optional] 
**Name** | Pointer to **string** | Name renames the flow when present. | [optional] 

## Methods

### NewAutoUpdate

`func NewAutoUpdate() *AutoUpdate`

NewAutoUpdate instantiates a new AutoUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoUpdateWithDefaults

`func NewAutoUpdateWithDefaults() *AutoUpdate`

NewAutoUpdateWithDefaults instantiates a new AutoUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AutoUpdate) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutoUpdate) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutoUpdate) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AutoUpdate) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AutoUpdate) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AutoUpdate) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetFlow

`func (o *AutoUpdate) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *AutoUpdate) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *AutoUpdate) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *AutoUpdate) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetName

`func (o *AutoUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoUpdate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


