# ProviderView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is whether THIS DEPLOYMENT has the provider&#39;s app credentials, so connect can succeed. False renders the card without a working Connect button. | [optional] 
**Category** | Pointer to **string** | Category groups the card (\&quot;Communication\&quot;, \&quot;Developer\&quot;, \&quot;Marketing\&quot;). | [optional] 
**Connected** | Pointer to **bool** | Connected is whether this org has a live connection to the provider. | [optional] 
**Connection** | Pointer to [**ConnectionView**](ConnectionView.md) | Connection is the connected account&#39;s non-secret detail. Absent when the org has no connection; tokens NEVER appear here (they live only in KMS). | [optional] 
**Description** | Pointer to **string** | Description is the one-line pitch the console card shows. | [optional] 
**Id** | Pointer to **string** | ID is the provider&#39;s registry id and the :provider path segment (\&quot;slack\&quot;). | [optional] 
**Name** | Pointer to **string** | Name is the provider&#39;s display name (\&quot;Slack\&quot;). | [optional] 

## Methods

### NewProviderView

`func NewProviderView() *ProviderView`

NewProviderView instantiates a new ProviderView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderViewWithDefaults

`func NewProviderViewWithDefaults() *ProviderView`

NewProviderViewWithDefaults instantiates a new ProviderView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ProviderView) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ProviderView) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ProviderView) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ProviderView) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCategory

`func (o *ProviderView) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ProviderView) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ProviderView) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ProviderView) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConnected

`func (o *ProviderView) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ProviderView) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ProviderView) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ProviderView) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetConnection

`func (o *ProviderView) GetConnection() ConnectionView`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ProviderView) GetConnectionOk() (*ConnectionView, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ProviderView) SetConnection(v ConnectionView)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ProviderView) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetDescription

`func (o *ProviderView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProviderView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ProviderView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProviderView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderView) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


