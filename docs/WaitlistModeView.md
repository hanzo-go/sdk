# WaitlistModeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the queried host, normalized (lowercased, port stripped). | [optional] 
**Known** | Pointer to **bool** | Known is false when no registered service claims this host, or when the registry is unavailable — the guard then lets the request through, which is why the two cases answer alike. | [optional] 
**Service** | Pointer to **string** | Service is the registered service that governs this host, empty when none does. | [optional] 
**WaitlistMode** | Pointer to **bool** | WaitlistMode is true when the service is GATED to approved users, false when it is open. Always false for an ungoverned host. | [optional] 

## Methods

### NewWaitlistModeView

`func NewWaitlistModeView() *WaitlistModeView`

NewWaitlistModeView instantiates a new WaitlistModeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitlistModeViewWithDefaults

`func NewWaitlistModeViewWithDefaults() *WaitlistModeView`

NewWaitlistModeViewWithDefaults instantiates a new WaitlistModeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *WaitlistModeView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *WaitlistModeView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *WaitlistModeView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *WaitlistModeView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKnown

`func (o *WaitlistModeView) GetKnown() bool`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *WaitlistModeView) GetKnownOk() (*bool, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnown

`func (o *WaitlistModeView) SetKnown(v bool)`

SetKnown sets Known field to given value.

### HasKnown

`func (o *WaitlistModeView) HasKnown() bool`

HasKnown returns a boolean if a field has been set.

### GetService

`func (o *WaitlistModeView) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *WaitlistModeView) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *WaitlistModeView) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *WaitlistModeView) HasService() bool`

HasService returns a boolean if a field has been set.

### GetWaitlistMode

`func (o *WaitlistModeView) GetWaitlistMode() bool`

GetWaitlistMode returns the WaitlistMode field if non-nil, zero value otherwise.

### GetWaitlistModeOk

`func (o *WaitlistModeView) GetWaitlistModeOk() (*bool, bool)`

GetWaitlistModeOk returns a tuple with the WaitlistMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistMode

`func (o *WaitlistModeView) SetWaitlistMode(v bool)`

SetWaitlistMode sets WaitlistMode field to given value.

### HasWaitlistMode

`func (o *WaitlistModeView) HasWaitlistMode() bool`

HasWaitlistMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


