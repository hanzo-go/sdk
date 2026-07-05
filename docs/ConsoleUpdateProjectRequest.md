# ConsoleUpdateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Retention** | Pointer to **int32** |  | [optional] 

## Methods

### NewConsoleUpdateProjectRequest

`func NewConsoleUpdateProjectRequest(name string, ) *ConsoleUpdateProjectRequest`

NewConsoleUpdateProjectRequest instantiates a new ConsoleUpdateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleUpdateProjectRequestWithDefaults

`func NewConsoleUpdateProjectRequestWithDefaults() *ConsoleUpdateProjectRequest`

NewConsoleUpdateProjectRequestWithDefaults instantiates a new ConsoleUpdateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConsoleUpdateProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleUpdateProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleUpdateProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMetadata

`func (o *ConsoleUpdateProjectRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleUpdateProjectRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleUpdateProjectRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleUpdateProjectRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRetention

`func (o *ConsoleUpdateProjectRequest) GetRetention() int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *ConsoleUpdateProjectRequest) GetRetentionOk() (*int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *ConsoleUpdateProjectRequest) SetRetention(v int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *ConsoleUpdateProjectRequest) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


