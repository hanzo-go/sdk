# CloudEngineAdvertisement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apis** | Pointer to **[]string** | [\&quot;openai\&quot;,\&quot;anthropic\&quot;] | [optional] 
**Models** | Pointer to **[]string** | ids from the node&#39;s GET /v1/models | [optional] 
**Status** | Pointer to **string** | \&quot;ready\&quot; | \&quot;unreachable\&quot; | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudEngineAdvertisement

`func NewCloudEngineAdvertisement() *CloudEngineAdvertisement`

NewCloudEngineAdvertisement instantiates a new CloudEngineAdvertisement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEngineAdvertisementWithDefaults

`func NewCloudEngineAdvertisementWithDefaults() *CloudEngineAdvertisement`

NewCloudEngineAdvertisementWithDefaults instantiates a new CloudEngineAdvertisement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApis

`func (o *CloudEngineAdvertisement) GetApis() []string`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *CloudEngineAdvertisement) GetApisOk() (*[]string, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *CloudEngineAdvertisement) SetApis(v []string)`

SetApis sets Apis field to given value.

### HasApis

`func (o *CloudEngineAdvertisement) HasApis() bool`

HasApis returns a boolean if a field has been set.

### GetModels

`func (o *CloudEngineAdvertisement) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CloudEngineAdvertisement) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CloudEngineAdvertisement) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *CloudEngineAdvertisement) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetStatus

`func (o *CloudEngineAdvertisement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudEngineAdvertisement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudEngineAdvertisement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudEngineAdvertisement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *CloudEngineAdvertisement) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudEngineAdvertisement) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudEngineAdvertisement) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudEngineAdvertisement) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


