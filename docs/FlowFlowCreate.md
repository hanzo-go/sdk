# FlowFlowCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Description** | Pointer to **string** | Description says what the workflow does. | [optional] 
**Name** | Pointer to **string** | Name is the workflow&#39;s display name, unique within the org&#39;s project (the product de-duplicates by suffixing). | [optional] 

## Methods

### NewFlowFlowCreate

`func NewFlowFlowCreate() *FlowFlowCreate`

NewFlowFlowCreate instantiates a new FlowFlowCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowFlowCreateWithDefaults

`func NewFlowFlowCreateWithDefaults() *FlowFlowCreate`

NewFlowFlowCreateWithDefaults instantiates a new FlowFlowCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FlowFlowCreate) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FlowFlowCreate) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FlowFlowCreate) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *FlowFlowCreate) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *FlowFlowCreate) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *FlowFlowCreate) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDescription

`func (o *FlowFlowCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlowFlowCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlowFlowCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlowFlowCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *FlowFlowCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowFlowCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowFlowCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowFlowCreate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


