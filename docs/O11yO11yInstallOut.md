# O11yO11yInstallOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yIntegrationsListItem**](O11yIntegrationsListItem.md) | Data holds the installed integration. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yInstallOut

`func NewO11yO11yInstallOut() *O11yO11yInstallOut`

NewO11yO11yInstallOut instantiates a new O11yO11yInstallOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yInstallOutWithDefaults

`func NewO11yO11yInstallOutWithDefaults() *O11yO11yInstallOut`

NewO11yO11yInstallOutWithDefaults instantiates a new O11yO11yInstallOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yInstallOut) GetData() O11yIntegrationsListItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yInstallOut) GetDataOk() (*O11yIntegrationsListItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yInstallOut) SetData(v O11yIntegrationsListItem)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yInstallOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yInstallOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yInstallOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yInstallOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yInstallOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


