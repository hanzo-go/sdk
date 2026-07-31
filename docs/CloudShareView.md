# CloudShareView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | Pointer to **string** | Backend is the local endpoint the share proxies to. | [optional] 
**BackendMode** | Pointer to **string** | BackendMode is how the tunnel serves the backend, e.g. proxy or web. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the share was opened, unix seconds. | [optional] 
**Token** | Pointer to **string** | Token is the share&#39;s own identifier, the leaf of its public URL. | [optional] 
**Url** | Pointer to **string** | URL is the share&#39;s public address, rendered from the deployment&#39;s URL template. | [optional] 

## Methods

### NewCloudShareView

`func NewCloudShareView() *CloudShareView`

NewCloudShareView instantiates a new CloudShareView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudShareViewWithDefaults

`func NewCloudShareViewWithDefaults() *CloudShareView`

NewCloudShareViewWithDefaults instantiates a new CloudShareView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackend

`func (o *CloudShareView) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *CloudShareView) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *CloudShareView) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *CloudShareView) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetBackendMode

`func (o *CloudShareView) GetBackendMode() string`

GetBackendMode returns the BackendMode field if non-nil, zero value otherwise.

### GetBackendModeOk

`func (o *CloudShareView) GetBackendModeOk() (*string, bool)`

GetBackendModeOk returns a tuple with the BackendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendMode

`func (o *CloudShareView) SetBackendMode(v string)`

SetBackendMode sets BackendMode field to given value.

### HasBackendMode

`func (o *CloudShareView) HasBackendMode() bool`

HasBackendMode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudShareView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudShareView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudShareView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudShareView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetToken

`func (o *CloudShareView) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CloudShareView) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CloudShareView) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CloudShareView) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUrl

`func (o *CloudShareView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudShareView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudShareView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudShareView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


