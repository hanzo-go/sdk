# FlowCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Description** | Pointer to **string** | Description says what the workflow does. | [optional] 
**Name** | Pointer to **string** | Name is the workflow&#39;s display name, unique within the org&#39;s project (the product de-duplicates by suffixing). | [optional] 

## Methods

### NewFlowCreate

`func NewFlowCreate() *FlowCreate`

NewFlowCreate instantiates a new FlowCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowCreateWithDefaults

`func NewFlowCreateWithDefaults() *FlowCreate`

NewFlowCreateWithDefaults instantiates a new FlowCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FlowCreate) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FlowCreate) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FlowCreate) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *FlowCreate) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *FlowCreate) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *FlowCreate) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDescription

`func (o *FlowCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlowCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlowCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlowCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *FlowCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowCreate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


