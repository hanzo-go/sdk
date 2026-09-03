# PublishedView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **string** | DNS is the name the fabric answers for this service — what a kubeconfig server, or any client on the org&#39;s overlay, dials. | [optional] 
**Id** | Pointer to **string** | ID is the fabric service&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is the service&#39;s name within the org. | [optional] 

## Methods

### NewPublishedView

`func NewPublishedView() *PublishedView`

NewPublishedView instantiates a new PublishedView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedViewWithDefaults

`func NewPublishedViewWithDefaults() *PublishedView`

NewPublishedViewWithDefaults instantiates a new PublishedView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *PublishedView) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *PublishedView) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *PublishedView) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *PublishedView) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetId

`func (o *PublishedView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublishedView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublishedView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublishedView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PublishedView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublishedView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublishedView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublishedView) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


