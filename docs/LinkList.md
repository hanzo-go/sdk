# LinkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]DeviceView**](DeviceView.md) | Devices is the same rows folded per machine — the cross-machine \&quot;AI Providers / Accounts\&quot; view. | [optional] 
**Links** | Pointer to [**[]LinkView**](LinkView.md) | Links is every link the caller registered, newest first. Revoked links are INCLUDED rather than dropped, because a logged-out account keeps its usage history and audit trail. | [optional] 

## Methods

### NewLinkList

`func NewLinkList() *LinkList`

NewLinkList instantiates a new LinkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkListWithDefaults

`func NewLinkListWithDefaults() *LinkList`

NewLinkListWithDefaults instantiates a new LinkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *LinkList) GetDevices() []DeviceView`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *LinkList) GetDevicesOk() (*[]DeviceView, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *LinkList) SetDevices(v []DeviceView)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *LinkList) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetLinks

`func (o *LinkList) GetLinks() []LinkView`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LinkList) GetLinksOk() (*[]LinkView, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LinkList) SetLinks(v []LinkView)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LinkList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


