# EndpointView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **string** | Connector names a connected account from the org&#39;s connector registry, when the endpoint reaches its provider through one. Absent means the locator stands on its own; the pair below is always sufficient either way. | [optional] 
**Locator** | Pointer to **string** | Locator addresses the thing INSIDE that provider, in the provider&#39;s own terms — an https clone URL for a hosted forge, a bare repository name for hanzo-git. It never carries a credential. | [optional] 
**Provider** | Pointer to **string** | Provider is the concrete integration: \&quot;github\&quot;, \&quot;gitlab\&quot; or \&quot;hanzo-git\&quot;. | [optional] 

## Methods

### NewEndpointView

`func NewEndpointView() *EndpointView`

NewEndpointView instantiates a new EndpointView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointViewWithDefaults

`func NewEndpointViewWithDefaults() *EndpointView`

NewEndpointViewWithDefaults instantiates a new EndpointView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *EndpointView) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *EndpointView) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *EndpointView) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *EndpointView) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetLocator

`func (o *EndpointView) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *EndpointView) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *EndpointView) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *EndpointView) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetProvider

`func (o *EndpointView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EndpointView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EndpointView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EndpointView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


