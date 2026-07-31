# CloudWaitlistModeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the queried host, normalized (lowercased, port stripped). | [optional] 
**Known** | Pointer to **bool** | Known is false when no registered service claims this host, or when the registry is unavailable — the guard then lets the request through, which is why the two cases answer alike. | [optional] 
**Service** | Pointer to **string** | Service is the registered service that governs this host, empty when none does. | [optional] 
**WaitlistMode** | Pointer to **bool** | WaitlistMode is true when the service is GATED to approved users, false when it is open. Always false for an ungoverned host. | [optional] 

## Methods

### NewCloudWaitlistModeView

`func NewCloudWaitlistModeView() *CloudWaitlistModeView`

NewCloudWaitlistModeView instantiates a new CloudWaitlistModeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWaitlistModeViewWithDefaults

`func NewCloudWaitlistModeViewWithDefaults() *CloudWaitlistModeView`

NewCloudWaitlistModeViewWithDefaults instantiates a new CloudWaitlistModeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CloudWaitlistModeView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudWaitlistModeView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudWaitlistModeView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudWaitlistModeView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKnown

`func (o *CloudWaitlistModeView) GetKnown() bool`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *CloudWaitlistModeView) GetKnownOk() (*bool, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnown

`func (o *CloudWaitlistModeView) SetKnown(v bool)`

SetKnown sets Known field to given value.

### HasKnown

`func (o *CloudWaitlistModeView) HasKnown() bool`

HasKnown returns a boolean if a field has been set.

### GetService

`func (o *CloudWaitlistModeView) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudWaitlistModeView) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudWaitlistModeView) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudWaitlistModeView) HasService() bool`

HasService returns a boolean if a field has been set.

### GetWaitlistMode

`func (o *CloudWaitlistModeView) GetWaitlistMode() bool`

GetWaitlistMode returns the WaitlistMode field if non-nil, zero value otherwise.

### GetWaitlistModeOk

`func (o *CloudWaitlistModeView) GetWaitlistModeOk() (*bool, bool)`

GetWaitlistModeOk returns a tuple with the WaitlistMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistMode

`func (o *CloudWaitlistModeView) SetWaitlistMode(v bool)`

SetWaitlistMode sets WaitlistMode field to given value.

### HasWaitlistMode

`func (o *CloudWaitlistModeView) HasWaitlistMode() bool`

HasWaitlistMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


