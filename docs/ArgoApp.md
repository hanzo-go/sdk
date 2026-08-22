# ArgoApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion is the constant \&quot;argoproj.io/v1alpha1\&quot; — the shape, not the source. These are projections of operator App CRs and Hanzo CD Applications; no argoproj.io object is stored anywhere behind this plane. | [optional] 
**Kind** | Pointer to **string** | Kind is the constant \&quot;Application\&quot;. | [optional] 
**Metadata** | Pointer to [**ArgoMeta**](ArgoMeta.md) | Metadata is the projected object&#39;s identity. | [optional] 
**Spec** | Pointer to [**ArgoSpec**](ArgoSpec.md) | Spec is the desired state: where it comes from, where it lands, what project it belongs to. | [optional] 
**Status** | Pointer to [**ArgoStatus**](ArgoStatus.md) | Status is what was observed: the sync verdict, the health, and the owned objects when this is a detail read. | [optional] 

## Methods

### NewArgoApp

`func NewArgoApp() *ArgoApp`

NewArgoApp instantiates a new ArgoApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoAppWithDefaults

`func NewArgoAppWithDefaults() *ArgoApp`

NewArgoAppWithDefaults instantiates a new ArgoApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArgoApp) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArgoApp) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArgoApp) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArgoApp) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ArgoApp) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArgoApp) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArgoApp) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArgoApp) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ArgoApp) GetMetadata() ArgoMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArgoApp) GetMetadataOk() (*ArgoMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArgoApp) SetMetadata(v ArgoMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArgoApp) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *ArgoApp) GetSpec() ArgoSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ArgoApp) GetSpecOk() (*ArgoSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ArgoApp) SetSpec(v ArgoSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ArgoApp) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ArgoApp) GetStatus() ArgoStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArgoApp) GetStatusOk() (*ArgoStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArgoApp) SetStatus(v ArgoStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArgoApp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


