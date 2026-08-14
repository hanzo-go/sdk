# IamKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AccessKey (pk-*) is the publishable identifier and lookup index; AccessSecret (sk-*) is the confidential secret. | [optional] 
**AccessSecret** | Pointer to **string** |  | [optional] 
**Application** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** | CreatedTime and UpdatedTime are RFC3339 audit stamps carried as strings for byte-parity with the v1 row (orm.Model separately tracks CreatedAt / UpdatedAt as time.Time for the store&#39;s own lifecycle). | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is the human-facing label. | [optional] 
**ExpireTime** | Pointer to **string** | ExpireTime is when the key stops being honored (empty &#x3D; never). State is the lifecycle flag (\&quot;Active\&quot;, \&quot;test\&quot;, …); \&quot;test\&quot; mints test-env credentials instead of live ones. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** | Owner is the tenant that holds the key; Name is unique within Owner. | [optional] 
**Scope** | Pointer to **string** | Scope is the key&#39;s ACCESS CLASS, orthogonal to Type (which names the bound principal). Empty (the default, \&quot;secret\&quot;) is a full key: a pk- publishable half AND a confidential sk- half, the sk- authenticating a server-side reader. KeyScopePublish is a WRITE-ONLY publishable key — a pk- half only, no secret — that resolves to just an ORG (never a principal) at the ingest door and is safe to ship in client JS. A missing value on an existing row reads as the default, so every pre-Scope key is a secret key unchanged. | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type is the scope the key is bound to — \&quot;Organization\&quot;, \&quot;Application\&quot;, \&quot;User\&quot;, or \&quot;General\&quot; — and Organization / Application / User name the concrete principal for whichever scope Type selects. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewIamKey

`func NewIamKey() *IamKey`

NewIamKey instantiates a new IamKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamKeyWithDefaults

`func NewIamKeyWithDefaults() *IamKey`

NewIamKeyWithDefaults instantiates a new IamKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *IamKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *IamKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *IamKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *IamKey) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessSecret

`func (o *IamKey) GetAccessSecret() string`

GetAccessSecret returns the AccessSecret field if non-nil, zero value otherwise.

### GetAccessSecretOk

`func (o *IamKey) GetAccessSecretOk() (*string, bool)`

GetAccessSecretOk returns a tuple with the AccessSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecret

`func (o *IamKey) SetAccessSecret(v string)`

SetAccessSecret sets AccessSecret field to given value.

### HasAccessSecret

`func (o *IamKey) HasAccessSecret() bool`

HasAccessSecret returns a boolean if a field has been set.

### GetApplication

`func (o *IamKey) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamKey) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamKey) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *IamKey) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamKey) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamKey) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamKey) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamKey) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeleted

`func (o *IamKey) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamKey) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamKey) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamKey) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamKey) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamKey) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamKey) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamKey) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpireTime

`func (o *IamKey) GetExpireTime() string`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *IamKey) GetExpireTimeOk() (*string, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *IamKey) SetExpireTime(v string)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *IamKey) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetId

`func (o *IamKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IamKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *IamKey) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamKey) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamKey) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamKey) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *IamKey) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamKey) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamKey) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamKey) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetScope

`func (o *IamKey) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IamKey) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IamKey) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IamKey) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetState

`func (o *IamKey) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamKey) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamKey) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamKey) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *IamKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamKey) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamKey) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamKey) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamKey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *IamKey) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *IamKey) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *IamKey) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *IamKey) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUser

`func (o *IamKey) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamKey) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamKey) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamKey) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


