# IngressIngressServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**[]IngressUpstream**](IngressUpstream.md) | Services is the org&#39;s services, ordered by id. | [optional] 

## Methods

### NewIngressIngressServices

`func NewIngressIngressServices() *IngressIngressServices`

NewIngressIngressServices instantiates a new IngressIngressServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressIngressServicesWithDefaults

`func NewIngressIngressServicesWithDefaults() *IngressIngressServices`

NewIngressIngressServicesWithDefaults instantiates a new IngressIngressServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *IngressIngressServices) GetServices() []IngressUpstream`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *IngressIngressServices) GetServicesOk() (*[]IngressUpstream, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *IngressIngressServices) SetServices(v []IngressUpstream)`

SetServices sets Services field to given value.

### HasServices

`func (o *IngressIngressServices) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


