# PlatformProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]PlatformDeclRef**](PlatformDeclRef.md) | Apps are the declarations it starts with, each by values directory (&#x60;org&#x60;, defaulting to the caller&#39;s own) and name. At least one: a project is the apps that name it. | [optional] 
**Mode** | Pointer to **string** | Mode is &#x60;branch&#x60; (the default: a review branch, nothing deploys) or &#x60;commit&#x60; (main). | [optional] 
**Name** | Pointer to **string** | Name is the project, a DNS-1123 label; it becomes each app&#39;s &#x60;partOf&#x60;. | [optional] 
**Org** | Pointer to **string** | Org names the project&#39;s owner, defaulting to the caller&#39;s own scope; naming another is an act-as only a SuperAdmin has. | [optional] 

## Methods

### NewPlatformProjectCreate

`func NewPlatformProjectCreate() *PlatformProjectCreate`

NewPlatformProjectCreate instantiates a new PlatformProjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformProjectCreateWithDefaults

`func NewPlatformProjectCreateWithDefaults() *PlatformProjectCreate`

NewPlatformProjectCreateWithDefaults instantiates a new PlatformProjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *PlatformProjectCreate) GetApps() []PlatformDeclRef`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *PlatformProjectCreate) GetAppsOk() (*[]PlatformDeclRef, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *PlatformProjectCreate) SetApps(v []PlatformDeclRef)`

SetApps sets Apps field to given value.

### HasApps

`func (o *PlatformProjectCreate) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetMode

`func (o *PlatformProjectCreate) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PlatformProjectCreate) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PlatformProjectCreate) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PlatformProjectCreate) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *PlatformProjectCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformProjectCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformProjectCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformProjectCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformProjectCreate) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformProjectCreate) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformProjectCreate) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformProjectCreate) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


