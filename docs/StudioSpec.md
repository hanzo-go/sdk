# StudioSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Each** | Pointer to **string** | Each fans the step out: a {{path}} naming an array, run once per item (at most 64), its outputs the items&#39; outputs in order. | [optional] 
**Id** | Pointer to **string** | ID names the step within its graph: a lowercase letter, then up to 31 lowercase letters, digits or underscores. | [optional] 
**Op** | Pointer to **string** | Op is what the step does: chat, image.generate, image.edit, voice.speak, music.generate, video.generate, model.generate, studio.render, render.html, render.document, render.slides, render.composition or video.edit. | [optional] 
**Title** | Pointer to **string** | Title is what a person sees for the step. | [optional] 
**With** | Pointer to **map[string]interface{}** | With are the op&#39;s parameters. Values may reference other steps. | [optional] 

## Methods

### NewStudioSpec

`func NewStudioSpec() *StudioSpec`

NewStudioSpec instantiates a new StudioSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioSpecWithDefaults

`func NewStudioSpecWithDefaults() *StudioSpec`

NewStudioSpecWithDefaults instantiates a new StudioSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEach

`func (o *StudioSpec) GetEach() string`

GetEach returns the Each field if non-nil, zero value otherwise.

### GetEachOk

`func (o *StudioSpec) GetEachOk() (*string, bool)`

GetEachOk returns a tuple with the Each field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEach

`func (o *StudioSpec) SetEach(v string)`

SetEach sets Each field to given value.

### HasEach

`func (o *StudioSpec) HasEach() bool`

HasEach returns a boolean if a field has been set.

### GetId

`func (o *StudioSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudioSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudioSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StudioSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOp

`func (o *StudioSpec) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *StudioSpec) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *StudioSpec) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *StudioSpec) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetTitle

`func (o *StudioSpec) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StudioSpec) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StudioSpec) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StudioSpec) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWith

`func (o *StudioSpec) GetWith() map[string]interface{}`

GetWith returns the With field if non-nil, zero value otherwise.

### GetWithOk

`func (o *StudioSpec) GetWithOk() (*map[string]interface{}, bool)`

GetWithOk returns a tuple with the With field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWith

`func (o *StudioSpec) SetWith(v map[string]interface{})`

SetWith sets With field to given value.

### HasWith

`func (o *StudioSpec) HasWith() bool`

HasWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


