# StudioRunCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Brand is a brand id; its tokens are brand.* in the graph. | [optional] 
**Flow** | Pointer to **string** | Flow is a flow id. | [optional] 
**Graph** | Pointer to [**StudioGraph**](StudioGraph.md) |  | [optional] 
**Input** | Pointer to **map[string]interface{}** | Input are the values the inputs schema asks for. | [optional] 
**Parent** | Pointer to **string** | Parent makes this run the next version of that run. | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Template** | Pointer to **string** | Template is a template id. | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewStudioRunCreate

`func NewStudioRunCreate() *StudioRunCreate`

NewStudioRunCreate instantiates a new StudioRunCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioRunCreateWithDefaults

`func NewStudioRunCreateWithDefaults() *StudioRunCreate`

NewStudioRunCreateWithDefaults instantiates a new StudioRunCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *StudioRunCreate) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *StudioRunCreate) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *StudioRunCreate) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *StudioRunCreate) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetFlow

`func (o *StudioRunCreate) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *StudioRunCreate) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *StudioRunCreate) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *StudioRunCreate) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetGraph

`func (o *StudioRunCreate) GetGraph() StudioGraph`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *StudioRunCreate) GetGraphOk() (*StudioGraph, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *StudioRunCreate) SetGraph(v StudioGraph)`

SetGraph sets Graph field to given value.

### HasGraph

`func (o *StudioRunCreate) HasGraph() bool`

HasGraph returns a boolean if a field has been set.

### GetInput

`func (o *StudioRunCreate) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *StudioRunCreate) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *StudioRunCreate) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *StudioRunCreate) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetParent

`func (o *StudioRunCreate) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StudioRunCreate) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StudioRunCreate) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StudioRunCreate) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetShared

`func (o *StudioRunCreate) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StudioRunCreate) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StudioRunCreate) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StudioRunCreate) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetTemplate

`func (o *StudioRunCreate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *StudioRunCreate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *StudioRunCreate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *StudioRunCreate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTitle

`func (o *StudioRunCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StudioRunCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StudioRunCreate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StudioRunCreate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


