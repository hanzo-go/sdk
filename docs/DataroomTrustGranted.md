# DataroomTrustGranted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delivery** | Pointer to **string** | Delivery is empty when the asker was mailed, and otherwise says what happened instead — so an approver is never left believing a mail went out that did not. | [optional] 
**ExpiresAt** | Pointer to **int64** | ExpiresAt is when the grant closes, in unix milliseconds. | [optional] 
**Link** | Pointer to **string** | Link is the share link&#39;s id. The link admits only the address that asked. | [optional] 
**State** | Pointer to **string** | State is \&quot;granted\&quot;. | [optional] 

## Methods

### NewDataroomTrustGranted

`func NewDataroomTrustGranted() *DataroomTrustGranted`

NewDataroomTrustGranted instantiates a new DataroomTrustGranted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomTrustGrantedWithDefaults

`func NewDataroomTrustGrantedWithDefaults() *DataroomTrustGranted`

NewDataroomTrustGrantedWithDefaults instantiates a new DataroomTrustGranted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelivery

`func (o *DataroomTrustGranted) GetDelivery() string`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *DataroomTrustGranted) GetDeliveryOk() (*string, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *DataroomTrustGranted) SetDelivery(v string)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *DataroomTrustGranted) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetExpiresAt

`func (o *DataroomTrustGranted) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DataroomTrustGranted) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DataroomTrustGranted) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DataroomTrustGranted) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLink

`func (o *DataroomTrustGranted) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *DataroomTrustGranted) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *DataroomTrustGranted) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *DataroomTrustGranted) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetState

`func (o *DataroomTrustGranted) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DataroomTrustGranted) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DataroomTrustGranted) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DataroomTrustGranted) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


