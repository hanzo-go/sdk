# O11yO11yAuthDomainsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]O11yO11yAuthDomain**](O11yO11yAuthDomain.md) | Data holds the domains. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yAuthDomainsOut

`func NewO11yO11yAuthDomainsOut() *O11yO11yAuthDomainsOut`

NewO11yO11yAuthDomainsOut instantiates a new O11yO11yAuthDomainsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAuthDomainsOutWithDefaults

`func NewO11yO11yAuthDomainsOutWithDefaults() *O11yO11yAuthDomainsOut`

NewO11yO11yAuthDomainsOutWithDefaults instantiates a new O11yO11yAuthDomainsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yAuthDomainsOut) GetData() []O11yO11yAuthDomain`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yAuthDomainsOut) GetDataOk() (*[]O11yO11yAuthDomain, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yAuthDomainsOut) SetData(v []O11yO11yAuthDomain)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yAuthDomainsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yAuthDomainsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yAuthDomainsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yAuthDomainsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yAuthDomainsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


