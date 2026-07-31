# CloudArgoDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | ArgoCD allows a destination by cluster name; omitted for the in-cluster projection. | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudArgoDestination

`func NewCloudArgoDestination() *CloudArgoDestination`

NewCloudArgoDestination instantiates a new CloudArgoDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoDestinationWithDefaults

`func NewCloudArgoDestinationWithDefaults() *CloudArgoDestination`

NewCloudArgoDestinationWithDefaults instantiates a new CloudArgoDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudArgoDestination) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudArgoDestination) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudArgoDestination) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudArgoDestination) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CloudArgoDestination) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CloudArgoDestination) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CloudArgoDestination) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CloudArgoDestination) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetServer

`func (o *CloudArgoDestination) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CloudArgoDestination) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CloudArgoDestination) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *CloudArgoDestination) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


