# CloudGenerateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]string** | Data supplies every merge field the template declares, keyed by field key. Every declared field is REQUIRED: a missing one is refused with 400 rather than rendered as a blank into a contract. | [optional] 
**TemplateId** | Pointer to **string** | TemplateID is the template to render. Required; resolved for the caller&#39;s org, so an override wins over the built-in. | [optional] 

## Methods

### NewCloudGenerateRequest

`func NewCloudGenerateRequest() *CloudGenerateRequest`

NewCloudGenerateRequest instantiates a new CloudGenerateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGenerateRequestWithDefaults

`func NewCloudGenerateRequestWithDefaults() *CloudGenerateRequest`

NewCloudGenerateRequestWithDefaults instantiates a new CloudGenerateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudGenerateRequest) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudGenerateRequest) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudGenerateRequest) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *CloudGenerateRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTemplateId

`func (o *CloudGenerateRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CloudGenerateRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CloudGenerateRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CloudGenerateRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


