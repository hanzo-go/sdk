# NetworkRouterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routers** | Pointer to [**[]NetworkRouterView**](NetworkRouterView.md) | Routers is one row per ZT edge-router tagged with the caller&#39;s org role. | [optional] 

## Methods

### NewNetworkRouterList

`func NewNetworkRouterList() *NetworkRouterList`

NewNetworkRouterList instantiates a new NetworkRouterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouterListWithDefaults

`func NewNetworkRouterListWithDefaults() *NetworkRouterList`

NewNetworkRouterListWithDefaults instantiates a new NetworkRouterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouters

`func (o *NetworkRouterList) GetRouters() []NetworkRouterView`

GetRouters returns the Routers field if non-nil, zero value otherwise.

### GetRoutersOk

`func (o *NetworkRouterList) GetRoutersOk() (*[]NetworkRouterView, bool)`

GetRoutersOk returns a tuple with the Routers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouters

`func (o *NetworkRouterList) SetRouters(v []NetworkRouterView)`

SetRouters sets Routers field to given value.

### HasRouters

`func (o *NetworkRouterList) HasRouters() bool`

HasRouters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


