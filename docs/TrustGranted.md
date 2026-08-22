# TrustGranted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delivery** | Pointer to **string** | Delivery is empty when the asker was mailed, and otherwise says what happened instead — so an approver is never left believing a mail went out that did not. | [optional] 
**ExpiresAt** | Pointer to **int32** | ExpiresAt is when the grant closes, in unix milliseconds. | [optional] 
**Link** | Pointer to **string** | Link is the share link&#39;s id. The link admits only the address that asked. | [optional] 
**State** | Pointer to **string** | State is \&quot;granted\&quot;. | [optional] 

## Methods

### NewTrustGranted

`func NewTrustGranted() *TrustGranted`

NewTrustGranted instantiates a new TrustGranted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustGrantedWithDefaults

`func NewTrustGrantedWithDefaults() *TrustGranted`

NewTrustGrantedWithDefaults instantiates a new TrustGranted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelivery

`func (o *TrustGranted) GetDelivery() string`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *TrustGranted) GetDeliveryOk() (*string, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *TrustGranted) SetDelivery(v string)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *TrustGranted) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TrustGranted) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TrustGranted) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TrustGranted) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TrustGranted) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLink

`func (o *TrustGranted) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *TrustGranted) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *TrustGranted) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *TrustGranted) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetState

`func (o *TrustGranted) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TrustGranted) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TrustGranted) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TrustGranted) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


