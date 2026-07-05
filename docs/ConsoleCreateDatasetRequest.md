# ConsoleCreateDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConsoleCreateDatasetRequest

`func NewConsoleCreateDatasetRequest(name string, ) *ConsoleCreateDatasetRequest`

NewConsoleCreateDatasetRequest instantiates a new ConsoleCreateDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleCreateDatasetRequestWithDefaults

`func NewConsoleCreateDatasetRequestWithDefaults() *ConsoleCreateDatasetRequest`

NewConsoleCreateDatasetRequestWithDefaults instantiates a new ConsoleCreateDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConsoleCreateDatasetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleCreateDatasetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleCreateDatasetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConsoleCreateDatasetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleCreateDatasetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleCreateDatasetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleCreateDatasetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *ConsoleCreateDatasetRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleCreateDatasetRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleCreateDatasetRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleCreateDatasetRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


