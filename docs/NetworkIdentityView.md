# NetworkIdentityView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | ExternalID is the IAM subject the identity logs in as: the controller admits a bearer whose &#x60;sub&#x60; equals it. | [optional] 
**Id** | Pointer to **string** | ID is the identity&#39;s fabric id — the key DELETE addresses. | [optional] 
**Name** | Pointer to **string** | Name is the identity&#39;s name within the org. | [optional] 
**Roles** | Pointer to **[]string** | Roles are the identity&#39;s role attributes as the fabric holds them: the \&quot;org-&lt;org&gt;\&quot; of every org it was ensured in, plus any org-scoped roles. | [optional] 

## Methods

### NewNetworkIdentityView

`func NewNetworkIdentityView() *NetworkIdentityView`

NewNetworkIdentityView instantiates a new NetworkIdentityView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIdentityViewWithDefaults

`func NewNetworkIdentityViewWithDefaults() *NetworkIdentityView`

NewNetworkIdentityViewWithDefaults instantiates a new NetworkIdentityView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *NetworkIdentityView) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetworkIdentityView) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetworkIdentityView) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NetworkIdentityView) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *NetworkIdentityView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkIdentityView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkIdentityView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkIdentityView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkIdentityView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkIdentityView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkIdentityView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkIdentityView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *NetworkIdentityView) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *NetworkIdentityView) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *NetworkIdentityView) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *NetworkIdentityView) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


