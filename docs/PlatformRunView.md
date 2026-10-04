# PlatformRunView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the application id the run created or converged. | [optional] 
**Name** | Pointer to **string** | Name is the run&#39;s name, as stored. | [optional] 
**Shape** | Pointer to **string** | Shape is the compute size label the request asked for, or \&quot;auto\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is the application&#39;s state — &#x60;deploying&#x60; on a fresh accept. | [optional] 
**Url** | Pointer to **string** | URL is the run&#39;s live HTTPS address. | [optional] 

## Methods

### NewPlatformRunView

`func NewPlatformRunView() *PlatformRunView`

NewPlatformRunView instantiates a new PlatformRunView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformRunViewWithDefaults

`func NewPlatformRunViewWithDefaults() *PlatformRunView`

NewPlatformRunViewWithDefaults instantiates a new PlatformRunView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlatformRunView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformRunView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformRunView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformRunView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PlatformRunView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformRunView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformRunView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformRunView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShape

`func (o *PlatformRunView) GetShape() string`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *PlatformRunView) GetShapeOk() (*string, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *PlatformRunView) SetShape(v string)`

SetShape sets Shape field to given value.

### HasShape

`func (o *PlatformRunView) HasShape() bool`

HasShape returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformRunView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformRunView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformRunView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformRunView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *PlatformRunView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PlatformRunView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PlatformRunView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PlatformRunView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


