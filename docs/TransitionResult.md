# TransitionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distribution** | Pointer to [**PublishResult**](PublishResult.md) |  | [optional] 
**Doctype** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Storefront** | Pointer to [**StorefrontResult**](StorefrontResult.md) |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewTransitionResult

`func NewTransitionResult() *TransitionResult`

NewTransitionResult instantiates a new TransitionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionResultWithDefaults

`func NewTransitionResultWithDefaults() *TransitionResult`

NewTransitionResultWithDefaults instantiates a new TransitionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistribution

`func (o *TransitionResult) GetDistribution() PublishResult`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *TransitionResult) GetDistributionOk() (*PublishResult, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *TransitionResult) SetDistribution(v PublishResult)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *TransitionResult) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetDoctype

`func (o *TransitionResult) GetDoctype() string`

GetDoctype returns the Doctype field if non-nil, zero value otherwise.

### GetDoctypeOk

`func (o *TransitionResult) GetDoctypeOk() (*string, bool)`

GetDoctypeOk returns a tuple with the Doctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctype

`func (o *TransitionResult) SetDoctype(v string)`

SetDoctype sets Doctype field to given value.

### HasDoctype

`func (o *TransitionResult) HasDoctype() bool`

HasDoctype returns a boolean if a field has been set.

### GetFrom

`func (o *TransitionResult) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TransitionResult) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TransitionResult) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TransitionResult) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetName

`func (o *TransitionResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransitionResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransitionResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransitionResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorefront

`func (o *TransitionResult) GetStorefront() StorefrontResult`

GetStorefront returns the Storefront field if non-nil, zero value otherwise.

### GetStorefrontOk

`func (o *TransitionResult) GetStorefrontOk() (*StorefrontResult, bool)`

GetStorefrontOk returns a tuple with the Storefront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorefront

`func (o *TransitionResult) SetStorefront(v StorefrontResult)`

SetStorefront sets Storefront field to given value.

### HasStorefront

`func (o *TransitionResult) HasStorefront() bool`

HasStorefront returns a boolean if a field has been set.

### GetTo

`func (o *TransitionResult) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TransitionResult) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TransitionResult) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TransitionResult) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


