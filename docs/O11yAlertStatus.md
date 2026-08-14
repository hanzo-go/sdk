# O11yAlertStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InhibitedBy** | Pointer to **[]string** |  | [optional] 
**SilencedBy** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yAlertStatus

`func NewO11yAlertStatus() *O11yAlertStatus`

NewO11yAlertStatus instantiates a new O11yAlertStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAlertStatusWithDefaults

`func NewO11yAlertStatusWithDefaults() *O11yAlertStatus`

NewO11yAlertStatusWithDefaults instantiates a new O11yAlertStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInhibitedBy

`func (o *O11yAlertStatus) GetInhibitedBy() []string`

GetInhibitedBy returns the InhibitedBy field if non-nil, zero value otherwise.

### GetInhibitedByOk

`func (o *O11yAlertStatus) GetInhibitedByOk() (*[]string, bool)`

GetInhibitedByOk returns a tuple with the InhibitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInhibitedBy

`func (o *O11yAlertStatus) SetInhibitedBy(v []string)`

SetInhibitedBy sets InhibitedBy field to given value.

### HasInhibitedBy

`func (o *O11yAlertStatus) HasInhibitedBy() bool`

HasInhibitedBy returns a boolean if a field has been set.

### GetSilencedBy

`func (o *O11yAlertStatus) GetSilencedBy() []string`

GetSilencedBy returns the SilencedBy field if non-nil, zero value otherwise.

### GetSilencedByOk

`func (o *O11yAlertStatus) GetSilencedByOk() (*[]string, bool)`

GetSilencedByOk returns a tuple with the SilencedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilencedBy

`func (o *O11yAlertStatus) SetSilencedBy(v []string)`

SetSilencedBy sets SilencedBy field to given value.

### HasSilencedBy

`func (o *O11yAlertStatus) HasSilencedBy() bool`

HasSilencedBy returns a boolean if a field has been set.

### GetState

`func (o *O11yAlertStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *O11yAlertStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *O11yAlertStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *O11yAlertStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


