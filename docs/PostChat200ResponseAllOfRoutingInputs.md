# PostChat200ResponseAllOfRoutingInputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depth** | **string** | The reasoning depth the request states; default when it states none. | 
**Key** | **string** | The table row that depth was read at. | 
**Policy** | **string** | Whether the caller&#39;s plan or credit pays for the paid model at that depth. | 

## Methods

### NewPostChat200ResponseAllOfRoutingInputs

`func NewPostChat200ResponseAllOfRoutingInputs(depth string, key string, policy string, ) *PostChat200ResponseAllOfRoutingInputs`

NewPostChat200ResponseAllOfRoutingInputs instantiates a new PostChat200ResponseAllOfRoutingInputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostChat200ResponseAllOfRoutingInputsWithDefaults

`func NewPostChat200ResponseAllOfRoutingInputsWithDefaults() *PostChat200ResponseAllOfRoutingInputs`

NewPostChat200ResponseAllOfRoutingInputsWithDefaults instantiates a new PostChat200ResponseAllOfRoutingInputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepth

`func (o *PostChat200ResponseAllOfRoutingInputs) GetDepth() string`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *PostChat200ResponseAllOfRoutingInputs) GetDepthOk() (*string, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *PostChat200ResponseAllOfRoutingInputs) SetDepth(v string)`

SetDepth sets Depth field to given value.


### GetKey

`func (o *PostChat200ResponseAllOfRoutingInputs) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostChat200ResponseAllOfRoutingInputs) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostChat200ResponseAllOfRoutingInputs) SetKey(v string)`

SetKey sets Key field to given value.


### GetPolicy

`func (o *PostChat200ResponseAllOfRoutingInputs) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PostChat200ResponseAllOfRoutingInputs) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PostChat200ResponseAllOfRoutingInputs) SetPolicy(v string)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


