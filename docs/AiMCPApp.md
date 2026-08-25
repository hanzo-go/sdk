# AiMCPApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the subsystem, as the manifest names it. | [optional] 
**Served** | Pointer to **bool** | Served reports that THIS process mounted it, so its tools are on this process&#39;s MCP server rather than behind a sibling this process only knows the name of. | [optional] 

## Methods

### NewAiMCPApp

`func NewAiMCPApp() *AiMCPApp`

NewAiMCPApp instantiates a new AiMCPApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMCPAppWithDefaults

`func NewAiMCPAppWithDefaults() *AiMCPApp`

NewAiMCPAppWithDefaults instantiates a new AiMCPApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AiMCPApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiMCPApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiMCPApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AiMCPApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServed

`func (o *AiMCPApp) GetServed() bool`

GetServed returns the Served field if non-nil, zero value otherwise.

### GetServedOk

`func (o *AiMCPApp) GetServedOk() (*bool, bool)`

GetServedOk returns a tuple with the Served field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServed

`func (o *AiMCPApp) SetServed(v bool)`

SetServed sets Served field to given value.

### HasServed

`func (o *AiMCPApp) HasServed() bool`

HasServed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


