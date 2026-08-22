# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the reusable prompt or snippet itself. It may carry {placeholder} tokens for the client-specific bits — {client_name}, {domain}, {product} — which are substituted where the template is used, not here. | [optional] 
**Enabled** | Pointer to **bool** | Enabled is the admin lever. Absent reads as ON; an explicit false withdraws the template from org-facing reads. | [optional] 
**Id** | Pointer to **string** | ID is the slug a step references to pull this template in. | [optional] 
**Title** | Pointer to **string** | Title names the template in the authoring plane and in a picker. | [optional] 

## Methods

### NewPage

`func NewPage() *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *Page) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Page) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Page) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Page) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEnabled

`func (o *Page) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Page) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Page) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Page) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *Page) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Page) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Page) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Page) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Page) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Page) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Page) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Page) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


