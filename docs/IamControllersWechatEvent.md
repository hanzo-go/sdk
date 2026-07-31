# IamControllersWechatEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgType** | Pointer to **string** |  | [optional] 
**Event** | **string** | SCAN or SUBSCRIBE | 
**EventKey** | Pointer to **string** | Used as providerId | [optional] 
**FromUserName** | Pointer to **string** |  | [optional] 
**Ticket** | **string** |  | 

## Methods

### NewIamControllersWechatEvent

`func NewIamControllersWechatEvent(event string, ticket string, ) *IamControllersWechatEvent`

NewIamControllersWechatEvent instantiates a new IamControllersWechatEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamControllersWechatEventWithDefaults

`func NewIamControllersWechatEventWithDefaults() *IamControllersWechatEvent`

NewIamControllersWechatEventWithDefaults instantiates a new IamControllersWechatEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgType

`func (o *IamControllersWechatEvent) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *IamControllersWechatEvent) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *IamControllersWechatEvent) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *IamControllersWechatEvent) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetEvent

`func (o *IamControllersWechatEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *IamControllersWechatEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *IamControllersWechatEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetEventKey

`func (o *IamControllersWechatEvent) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *IamControllersWechatEvent) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *IamControllersWechatEvent) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *IamControllersWechatEvent) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetFromUserName

`func (o *IamControllersWechatEvent) GetFromUserName() string`

GetFromUserName returns the FromUserName field if non-nil, zero value otherwise.

### GetFromUserNameOk

`func (o *IamControllersWechatEvent) GetFromUserNameOk() (*string, bool)`

GetFromUserNameOk returns a tuple with the FromUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserName

`func (o *IamControllersWechatEvent) SetFromUserName(v string)`

SetFromUserName sets FromUserName field to given value.

### HasFromUserName

`func (o *IamControllersWechatEvent) HasFromUserName() bool`

HasFromUserName returns a boolean if a field has been set.

### GetTicket

`func (o *IamControllersWechatEvent) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *IamControllersWechatEvent) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *IamControllersWechatEvent) SetTicket(v string)`

SetTicket sets Ticket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


