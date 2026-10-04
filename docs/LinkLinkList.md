# LinkLinkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]LinkDeviceView**](LinkDeviceView.md) | Devices is the same rows folded per machine — the cross-machine \&quot;AI Providers / Accounts\&quot; view. | [optional] 
**Links** | Pointer to [**[]LinkLinkView**](LinkLinkView.md) | Links is every link the caller registered, newest first. Revoked links are INCLUDED rather than dropped, because a logged-out account keeps its usage history and audit trail. | [optional] 

## Methods

### NewLinkLinkList

`func NewLinkLinkList() *LinkLinkList`

NewLinkLinkList instantiates a new LinkLinkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkLinkListWithDefaults

`func NewLinkLinkListWithDefaults() *LinkLinkList`

NewLinkLinkListWithDefaults instantiates a new LinkLinkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *LinkLinkList) GetDevices() []LinkDeviceView`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *LinkLinkList) GetDevicesOk() (*[]LinkDeviceView, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *LinkLinkList) SetDevices(v []LinkDeviceView)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *LinkLinkList) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetLinks

`func (o *LinkLinkList) GetLinks() []LinkLinkView`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LinkLinkList) GetLinksOk() (*[]LinkLinkView, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LinkLinkList) SetLinks(v []LinkLinkView)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LinkLinkList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


