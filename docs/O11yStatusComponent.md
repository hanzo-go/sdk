# O11yStatusComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentStatus** | Pointer to **string** | CurrentStatus is this component&#39;s own condition: \&quot;full_outage\&quot; for a service that did not answer its health probe at all. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yStatusComponent

`func NewO11yStatusComponent() *O11yStatusComponent`

NewO11yStatusComponent instantiates a new O11yStatusComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yStatusComponentWithDefaults

`func NewO11yStatusComponentWithDefaults() *O11yStatusComponent`

NewO11yStatusComponentWithDefaults instantiates a new O11yStatusComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentStatus

`func (o *O11yStatusComponent) GetCurrentStatus() string`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *O11yStatusComponent) GetCurrentStatusOk() (*string, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *O11yStatusComponent) SetCurrentStatus(v string)`

SetCurrentStatus sets CurrentStatus field to given value.

### HasCurrentStatus

`func (o *O11yStatusComponent) HasCurrentStatus() bool`

HasCurrentStatus returns a boolean if a field has been set.

### GetId

`func (o *O11yStatusComponent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yStatusComponent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yStatusComponent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yStatusComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yStatusComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yStatusComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yStatusComponent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yStatusComponent) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


