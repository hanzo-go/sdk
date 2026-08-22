# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | CreatedAt is when the endpoint was registered, RFC3339 in UTC — stored in that spelling because it sorts as a string. | [optional] 
**Deliveries7d** | Pointer to **int32** | Deliveries7d is how many deliveries SETTLED in the trailing 7 days — the attempts that ended ok or failed, so a delivery still retrying is in neither counter yet. It is counted from the log at read time rather than stored, and it is filled only on a list or a get; a create answers 0 because there is no history, which is why it is never omitted. | [optional] 
**Description** | Pointer to **string** | Description is the operator&#39;s own label for the endpoint. Never sent anywhere. | [optional] 
**Events** | Pointer to **[]string** | Events are the subject patterns this endpoint subscribes to (\&quot;commerce.order.&gt;\&quot;). An EMPTY list means every event, not none. | [optional] 
**Failures7d** | Pointer to **int32** | Failures7d is how many of those settled as failed — the subscriber never accepted it and no further attempt will be made. It is the numerator to Deliveries7d, over the same window. | [optional] 
**Id** | Pointer to **string** | ID is the endpoint&#39;s handle, server-minted and stable for its life. It is what every other route here addresses. | [optional] 
**Org** | Pointer to **string** | Org is the tenant that owns the endpoint, taken from the validated principal rather than from any request field. | [optional] 
**Secret** | Pointer to **string** | Secret is the HMAC-SHA256 signing key a subscriber recomputes the signature with. It is returned exactly ONCE, on create: a later read of the endpoint omits it, so a lost secret is replaced rather than recovered. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;active\&quot; or \&quot;disabled\&quot; — nothing else is accepted. A disabled endpoint keeps its subscription and receives no deliveries, except a manual test send, which goes out anyway. | [optional] 
**Updated** | Pointer to **string** | UpdatedAt is when its url, events, status or description last changed. | [optional] 
**Url** | Pointer to **string** | URL is where the POST goes. Changing it is the one edit that redirects an org&#39;s events, which is why it is never bindable from a query string. | [optional] 

## Methods

### NewEndpoint

`func NewEndpoint() *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Endpoint) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Endpoint) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Endpoint) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Endpoint) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeliveries7d

`func (o *Endpoint) GetDeliveries7d() int32`

GetDeliveries7d returns the Deliveries7d field if non-nil, zero value otherwise.

### GetDeliveries7dOk

`func (o *Endpoint) GetDeliveries7dOk() (*int32, bool)`

GetDeliveries7dOk returns a tuple with the Deliveries7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries7d

`func (o *Endpoint) SetDeliveries7d(v int32)`

SetDeliveries7d sets Deliveries7d field to given value.

### HasDeliveries7d

`func (o *Endpoint) HasDeliveries7d() bool`

HasDeliveries7d returns a boolean if a field has been set.

### GetDescription

`func (o *Endpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Endpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Endpoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Endpoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvents

`func (o *Endpoint) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Endpoint) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Endpoint) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Endpoint) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFailures7d

`func (o *Endpoint) GetFailures7d() int32`

GetFailures7d returns the Failures7d field if non-nil, zero value otherwise.

### GetFailures7dOk

`func (o *Endpoint) GetFailures7dOk() (*int32, bool)`

GetFailures7dOk returns a tuple with the Failures7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures7d

`func (o *Endpoint) SetFailures7d(v int32)`

SetFailures7d sets Failures7d field to given value.

### HasFailures7d

`func (o *Endpoint) HasFailures7d() bool`

HasFailures7d returns a boolean if a field has been set.

### GetId

`func (o *Endpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Endpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Endpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Endpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *Endpoint) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Endpoint) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Endpoint) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Endpoint) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSecret

`func (o *Endpoint) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Endpoint) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Endpoint) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Endpoint) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStatus

`func (o *Endpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Endpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Endpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Endpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *Endpoint) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Endpoint) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Endpoint) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Endpoint) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUrl

`func (o *Endpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Endpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Endpoint) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Endpoint) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


