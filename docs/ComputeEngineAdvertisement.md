# ComputeEngineAdvertisement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apis** | Pointer to **[]string** | APIs are the wire formats the engine serves on that one port: \&quot;openai\&quot;, \&quot;anthropic\&quot;, or both. | [optional] 
**Models** | Pointer to **[]string** | Models are the model ids the node&#39;s own GET /v1/models answered with — what this GPU can actually be asked for. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ready\&quot; when the node&#39;s engine answered, \&quot;unreachable\&quot; when it did not. Advertised is not the same as serving, and this is the difference. | [optional] 
**Url** | Pointer to **string** | URL is the base address the node advertised its engine on — where a model call to this GPU is sent. The node chose it, so reaching it is a question about the node&#39;s network, not about this surface. | [optional] 

## Methods

### NewComputeEngineAdvertisement

`func NewComputeEngineAdvertisement() *ComputeEngineAdvertisement`

NewComputeEngineAdvertisement instantiates a new ComputeEngineAdvertisement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeEngineAdvertisementWithDefaults

`func NewComputeEngineAdvertisementWithDefaults() *ComputeEngineAdvertisement`

NewComputeEngineAdvertisementWithDefaults instantiates a new ComputeEngineAdvertisement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApis

`func (o *ComputeEngineAdvertisement) GetApis() []string`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *ComputeEngineAdvertisement) GetApisOk() (*[]string, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *ComputeEngineAdvertisement) SetApis(v []string)`

SetApis sets Apis field to given value.

### HasApis

`func (o *ComputeEngineAdvertisement) HasApis() bool`

HasApis returns a boolean if a field has been set.

### GetModels

`func (o *ComputeEngineAdvertisement) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ComputeEngineAdvertisement) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ComputeEngineAdvertisement) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *ComputeEngineAdvertisement) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetStatus

`func (o *ComputeEngineAdvertisement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComputeEngineAdvertisement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComputeEngineAdvertisement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComputeEngineAdvertisement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *ComputeEngineAdvertisement) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ComputeEngineAdvertisement) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ComputeEngineAdvertisement) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ComputeEngineAdvertisement) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


