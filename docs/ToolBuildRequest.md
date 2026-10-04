# ToolBuildRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the plugin&#39;s name: one lowercase path segment (a-z0-9, _ or -), and the id the runtime loads it by. | [optional] 
**Provider** | Pointer to **string** | Provider is the connectors provider whose credential the plugin reads at run time. Empty for a plugin that needs none. | [optional] 
**Source** | Pointer to **string** | Source is TypeScript to build as-is. Exactly one of Source or Spec. | [optional] 
**Spec** | Pointer to **string** | Spec is API documentation — an OpenAPI document, or prose describing the endpoints — that the generator turns into Source. The generated source is returned in the response, so a caller can read what will run before it runs. | [optional] 

## Methods

### NewToolBuildRequest

`func NewToolBuildRequest() *ToolBuildRequest`

NewToolBuildRequest instantiates a new ToolBuildRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolBuildRequestWithDefaults

`func NewToolBuildRequestWithDefaults() *ToolBuildRequest`

NewToolBuildRequestWithDefaults instantiates a new ToolBuildRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ToolBuildRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolBuildRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolBuildRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ToolBuildRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *ToolBuildRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ToolBuildRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ToolBuildRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ToolBuildRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSource

`func (o *ToolBuildRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ToolBuildRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ToolBuildRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ToolBuildRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpec

`func (o *ToolBuildRequest) GetSpec() string`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ToolBuildRequest) GetSpecOk() (*string, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ToolBuildRequest) SetSpec(v string)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ToolBuildRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


