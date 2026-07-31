# CloudEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**Deliveries7d** | Pointer to **int32** | Deliveries7d / Failures7d are cheap usage counters computed from the delivery log over usageWindow (not stored columns) and populated ONLY on list/get. They are 0 when there is no delivery history — never omitempty, so the console always sees them. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 
**Failures7d** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudEndpoint

`func NewCloudEndpoint() *CloudEndpoint`

NewCloudEndpoint instantiates a new CloudEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEndpointWithDefaults

`func NewCloudEndpointWithDefaults() *CloudEndpoint`

NewCloudEndpointWithDefaults instantiates a new CloudEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CloudEndpoint) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudEndpoint) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudEndpoint) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudEndpoint) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeliveries7d

`func (o *CloudEndpoint) GetDeliveries7d() int32`

GetDeliveries7d returns the Deliveries7d field if non-nil, zero value otherwise.

### GetDeliveries7dOk

`func (o *CloudEndpoint) GetDeliveries7dOk() (*int32, bool)`

GetDeliveries7dOk returns a tuple with the Deliveries7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries7d

`func (o *CloudEndpoint) SetDeliveries7d(v int32)`

SetDeliveries7d sets Deliveries7d field to given value.

### HasDeliveries7d

`func (o *CloudEndpoint) HasDeliveries7d() bool`

HasDeliveries7d returns a boolean if a field has been set.

### GetDescription

`func (o *CloudEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudEndpoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvents

`func (o *CloudEndpoint) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CloudEndpoint) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CloudEndpoint) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CloudEndpoint) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFailures7d

`func (o *CloudEndpoint) GetFailures7d() int32`

GetFailures7d returns the Failures7d field if non-nil, zero value otherwise.

### GetFailures7dOk

`func (o *CloudEndpoint) GetFailures7dOk() (*int32, bool)`

GetFailures7dOk returns a tuple with the Failures7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures7d

`func (o *CloudEndpoint) SetFailures7d(v int32)`

SetFailures7d sets Failures7d field to given value.

### HasFailures7d

`func (o *CloudEndpoint) HasFailures7d() bool`

HasFailures7d returns a boolean if a field has been set.

### GetId

`func (o *CloudEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *CloudEndpoint) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudEndpoint) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudEndpoint) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudEndpoint) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSecret

`func (o *CloudEndpoint) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CloudEndpoint) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CloudEndpoint) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CloudEndpoint) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStatus

`func (o *CloudEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudEndpoint) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudEndpoint) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudEndpoint) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudEndpoint) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUrl

`func (o *CloudEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudEndpoint) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


