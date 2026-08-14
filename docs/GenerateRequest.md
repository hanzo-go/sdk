# GenerateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]string** | Data supplies every merge field the template declares, keyed by field key. Every declared field is REQUIRED: a missing one is refused with 400 rather than rendered as a blank into a contract. | [optional] 
**TemplateId** | Pointer to **string** | TemplateID is the template to render. Required; resolved for the caller&#39;s org, so an override wins over the built-in. | [optional] 

## Methods

### NewGenerateRequest

`func NewGenerateRequest() *GenerateRequest`

NewGenerateRequest instantiates a new GenerateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateRequestWithDefaults

`func NewGenerateRequestWithDefaults() *GenerateRequest`

NewGenerateRequestWithDefaults instantiates a new GenerateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GenerateRequest) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GenerateRequest) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GenerateRequest) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *GenerateRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTemplateId

`func (o *GenerateRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *GenerateRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *GenerateRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *GenerateRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


