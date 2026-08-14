# MeshView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the ZT edge service&#39;s id. | [optional] 
**Mtls** | Pointer to **string** | Mtls is \&quot;required\&quot; when the service mandates end-to-end encryption, else \&quot;enabled\&quot; — the fabric mutually authenticates every link, so it is never truly off. | [optional] 
**Service** | Pointer to **string** | Service is the edge service&#39;s name. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;active\&quot;: a listed service is a configured, dialable mesh entry. | [optional] 

## Methods

### NewMeshView

`func NewMeshView() *MeshView`

NewMeshView instantiates a new MeshView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeshViewWithDefaults

`func NewMeshViewWithDefaults() *MeshView`

NewMeshViewWithDefaults instantiates a new MeshView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeshView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeshView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeshView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MeshView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtls

`func (o *MeshView) GetMtls() string`

GetMtls returns the Mtls field if non-nil, zero value otherwise.

### GetMtlsOk

`func (o *MeshView) GetMtlsOk() (*string, bool)`

GetMtlsOk returns a tuple with the Mtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtls

`func (o *MeshView) SetMtls(v string)`

SetMtls sets Mtls field to given value.

### HasMtls

`func (o *MeshView) HasMtls() bool`

HasMtls returns a boolean if a field has been set.

### GetService

`func (o *MeshView) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MeshView) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MeshView) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MeshView) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *MeshView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeshView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeshView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MeshView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


