# SyncEndpointReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **string** | Connector names the stored credential to reach this endpoint with. | [optional] 
**Locator** | Pointer to **string** | Locator addresses the resource. For a git source it is the https clone URL on the provider&#39;s own host, with no embedded credentials; for a native target it is the repository name. | [optional] 
**Provider** | Pointer to **string** | Provider is the platform: github or gitlab for a source; a target defaults to the native Hanzo Git plane. | [optional] 

## Methods

### NewSyncEndpointReq

`func NewSyncEndpointReq() *SyncEndpointReq`

NewSyncEndpointReq instantiates a new SyncEndpointReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncEndpointReqWithDefaults

`func NewSyncEndpointReqWithDefaults() *SyncEndpointReq`

NewSyncEndpointReqWithDefaults instantiates a new SyncEndpointReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *SyncEndpointReq) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *SyncEndpointReq) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *SyncEndpointReq) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *SyncEndpointReq) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetLocator

`func (o *SyncEndpointReq) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *SyncEndpointReq) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *SyncEndpointReq) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *SyncEndpointReq) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetProvider

`func (o *SyncEndpointReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SyncEndpointReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SyncEndpointReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SyncEndpointReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


