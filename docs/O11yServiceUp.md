# O11yServiceUp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the service as the fleet prober knows it (probes.go&#39;s target name, which is the &#x60;service&#x60; label on hanzo_service_up). | [optional] 
**Up** | Pointer to **bool** | Up is true when the service answered its own health URL on the last cycle. | [optional] 

## Methods

### NewO11yServiceUp

`func NewO11yServiceUp() *O11yServiceUp`

NewO11yServiceUp instantiates a new O11yServiceUp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yServiceUpWithDefaults

`func NewO11yServiceUpWithDefaults() *O11yServiceUp`

NewO11yServiceUpWithDefaults instantiates a new O11yServiceUp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *O11yServiceUp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yServiceUp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yServiceUp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yServiceUp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUp

`func (o *O11yServiceUp) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *O11yServiceUp) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *O11yServiceUp) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *O11yServiceUp) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


