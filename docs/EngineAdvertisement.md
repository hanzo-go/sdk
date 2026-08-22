# EngineAdvertisement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apis** | Pointer to **[]string** | APIs are the wire formats the engine serves on that one port: \&quot;openai\&quot;, \&quot;anthropic\&quot;, or both. | [optional] 
**Models** | Pointer to **[]string** | Models are the model ids the node&#39;s own GET /v1/models answered with — what this GPU can actually be asked for. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ready\&quot; when the node&#39;s engine answered, \&quot;unreachable\&quot; when it did not. Advertised is not the same as serving, and this is the difference. | [optional] 
**Url** | Pointer to **string** | URL is the base address the node advertised its engine on — where a model call to this GPU is sent. The node chose it, so reaching it is a question about the node&#39;s network, not about this surface. | [optional] 

## Methods

### NewEngineAdvertisement

`func NewEngineAdvertisement() *EngineAdvertisement`

NewEngineAdvertisement instantiates a new EngineAdvertisement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineAdvertisementWithDefaults

`func NewEngineAdvertisementWithDefaults() *EngineAdvertisement`

NewEngineAdvertisementWithDefaults instantiates a new EngineAdvertisement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApis

`func (o *EngineAdvertisement) GetApis() []string`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *EngineAdvertisement) GetApisOk() (*[]string, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *EngineAdvertisement) SetApis(v []string)`

SetApis sets Apis field to given value.

### HasApis

`func (o *EngineAdvertisement) HasApis() bool`

HasApis returns a boolean if a field has been set.

### GetModels

`func (o *EngineAdvertisement) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *EngineAdvertisement) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *EngineAdvertisement) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *EngineAdvertisement) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetStatus

`func (o *EngineAdvertisement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngineAdvertisement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngineAdvertisement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EngineAdvertisement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *EngineAdvertisement) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EngineAdvertisement) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EngineAdvertisement) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EngineAdvertisement) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


