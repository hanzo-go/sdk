# EmailTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bounced** | Pointer to **int64** | Bounced is messages that bounced, hard or soft (SendGrid&#39;s dropped included). | [optional] 
**Clicked** | Pointer to **int64** | Clicked is messages a link of which was clicked. | [optional] 
**Complained** | Pointer to **int64** | Complained is messages reported as spam. | [optional] 
**Delivered** | Pointer to **int64** | Delivered is messages the receiving server took. | [optional] 
**Opened** | Pointer to **int64** | Opened is messages opened. | [optional] 
**Sent** | Pointer to **int64** | Sent is messages the provider accepted. | [optional] 
**Unsubscribed** | Pointer to **int64** | Unsubscribed is messages whose recipient unsubscribed at the provider. | [optional] 

## Methods

### NewEmailTally

`func NewEmailTally() *EmailTally`

NewEmailTally instantiates a new EmailTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTallyWithDefaults

`func NewEmailTallyWithDefaults() *EmailTally`

NewEmailTallyWithDefaults instantiates a new EmailTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBounced

`func (o *EmailTally) GetBounced() int64`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *EmailTally) GetBouncedOk() (*int64, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *EmailTally) SetBounced(v int64)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *EmailTally) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetClicked

`func (o *EmailTally) GetClicked() int64`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *EmailTally) GetClickedOk() (*int64, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *EmailTally) SetClicked(v int64)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *EmailTally) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetComplained

`func (o *EmailTally) GetComplained() int64`

GetComplained returns the Complained field if non-nil, zero value otherwise.

### GetComplainedOk

`func (o *EmailTally) GetComplainedOk() (*int64, bool)`

GetComplainedOk returns a tuple with the Complained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplained

`func (o *EmailTally) SetComplained(v int64)`

SetComplained sets Complained field to given value.

### HasComplained

`func (o *EmailTally) HasComplained() bool`

HasComplained returns a boolean if a field has been set.

### GetDelivered

`func (o *EmailTally) GetDelivered() int64`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *EmailTally) GetDeliveredOk() (*int64, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *EmailTally) SetDelivered(v int64)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *EmailTally) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetOpened

`func (o *EmailTally) GetOpened() int64`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *EmailTally) GetOpenedOk() (*int64, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *EmailTally) SetOpened(v int64)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *EmailTally) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetSent

`func (o *EmailTally) GetSent() int64`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *EmailTally) GetSentOk() (*int64, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *EmailTally) SetSent(v int64)`

SetSent sets Sent field to given value.

### HasSent

`func (o *EmailTally) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *EmailTally) GetUnsubscribed() int64`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *EmailTally) GetUnsubscribedOk() (*int64, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *EmailTally) SetUnsubscribed(v int64)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *EmailTally) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


