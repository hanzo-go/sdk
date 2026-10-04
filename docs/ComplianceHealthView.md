# ComplianceHealthView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** | Provider is the wired verification provider&#39;s name (\&quot;manual\&quot; by default). | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when the subsystem is live. | [optional] 

## Methods

### NewComplianceHealthView

`func NewComplianceHealthView() *ComplianceHealthView`

NewComplianceHealthView instantiates a new ComplianceHealthView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceHealthViewWithDefaults

`func NewComplianceHealthViewWithDefaults() *ComplianceHealthView`

NewComplianceHealthViewWithDefaults instantiates a new ComplianceHealthView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ComplianceHealthView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ComplianceHealthView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ComplianceHealthView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ComplianceHealthView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *ComplianceHealthView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComplianceHealthView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComplianceHealthView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComplianceHealthView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


