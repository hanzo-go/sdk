# NetworkPublishedView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **string** | DNS is the name the fabric answers for this service — what a kubeconfig server, or any client on the org&#39;s overlay, dials. | [optional] 
**Id** | Pointer to **string** | ID is the fabric service&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is the service&#39;s name within the org. | [optional] 

## Methods

### NewNetworkPublishedView

`func NewNetworkPublishedView() *NetworkPublishedView`

NewNetworkPublishedView instantiates a new NetworkPublishedView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPublishedViewWithDefaults

`func NewNetworkPublishedViewWithDefaults() *NetworkPublishedView`

NewNetworkPublishedViewWithDefaults instantiates a new NetworkPublishedView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *NetworkPublishedView) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *NetworkPublishedView) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *NetworkPublishedView) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *NetworkPublishedView) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetId

`func (o *NetworkPublishedView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkPublishedView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkPublishedView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkPublishedView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkPublishedView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPublishedView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPublishedView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkPublishedView) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


