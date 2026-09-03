# TrustGrantView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the one address the link admits. | [optional] 
**ExpiresAt** | Pointer to **int64** | ExpiresAt is when the grant closes, in unix milliseconds. | [optional] 
**Item** | Pointer to **string** | Item is the item granted, empty when the whole released tier was granted. | [optional] 
**Link** | Pointer to **string** | Link is the share link&#39;s id — the token the party opens. Reading it here does not widen it: the link admits only Email whoever holds the id. | [optional] 
**Live** | Pointer to **bool** | Live is whether the grant is still open at the time of reading. | [optional] 

## Methods

### NewTrustGrantView

`func NewTrustGrantView() *TrustGrantView`

NewTrustGrantView instantiates a new TrustGrantView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustGrantViewWithDefaults

`func NewTrustGrantViewWithDefaults() *TrustGrantView`

NewTrustGrantViewWithDefaults instantiates a new TrustGrantView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *TrustGrantView) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TrustGrantView) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TrustGrantView) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TrustGrantView) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TrustGrantView) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TrustGrantView) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TrustGrantView) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TrustGrantView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetItem

`func (o *TrustGrantView) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *TrustGrantView) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *TrustGrantView) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *TrustGrantView) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetLink

`func (o *TrustGrantView) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *TrustGrantView) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *TrustGrantView) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *TrustGrantView) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetLive

`func (o *TrustGrantView) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *TrustGrantView) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *TrustGrantView) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *TrustGrantView) HasLive() bool`

HasLive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


