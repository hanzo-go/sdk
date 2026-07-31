# IamObjectWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 
**Headers** | Pointer to [**[]IamObjectHeader**](IamObjectHeader.md) |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsUserExtended** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectFields** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**SingleOrgOnly** | Pointer to **bool** |  | [optional] 
**TokenFields** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectWebhook

`func NewIamObjectWebhook() *IamObjectWebhook`

NewIamObjectWebhook instantiates a new IamObjectWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectWebhookWithDefaults

`func NewIamObjectWebhookWithDefaults() *IamObjectWebhook`

NewIamObjectWebhookWithDefaults instantiates a new IamObjectWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *IamObjectWebhook) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *IamObjectWebhook) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *IamObjectWebhook) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *IamObjectWebhook) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectWebhook) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectWebhook) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectWebhook) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectWebhook) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEvents

`func (o *IamObjectWebhook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *IamObjectWebhook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *IamObjectWebhook) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *IamObjectWebhook) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHeaders

`func (o *IamObjectWebhook) GetHeaders() []IamObjectHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *IamObjectWebhook) GetHeadersOk() (*[]IamObjectHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *IamObjectWebhook) SetHeaders(v []IamObjectHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *IamObjectWebhook) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetIsEnabled

`func (o *IamObjectWebhook) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IamObjectWebhook) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IamObjectWebhook) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IamObjectWebhook) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsUserExtended

`func (o *IamObjectWebhook) GetIsUserExtended() bool`

GetIsUserExtended returns the IsUserExtended field if non-nil, zero value otherwise.

### GetIsUserExtendedOk

`func (o *IamObjectWebhook) GetIsUserExtendedOk() (*bool, bool)`

GetIsUserExtendedOk returns a tuple with the IsUserExtended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserExtended

`func (o *IamObjectWebhook) SetIsUserExtended(v bool)`

SetIsUserExtended sets IsUserExtended field to given value.

### HasIsUserExtended

`func (o *IamObjectWebhook) HasIsUserExtended() bool`

HasIsUserExtended returns a boolean if a field has been set.

### GetMethod

`func (o *IamObjectWebhook) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IamObjectWebhook) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IamObjectWebhook) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *IamObjectWebhook) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *IamObjectWebhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectWebhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectWebhook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectWebhook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectFields

`func (o *IamObjectWebhook) GetObjectFields() []string`

GetObjectFields returns the ObjectFields field if non-nil, zero value otherwise.

### GetObjectFieldsOk

`func (o *IamObjectWebhook) GetObjectFieldsOk() (*[]string, bool)`

GetObjectFieldsOk returns a tuple with the ObjectFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectFields

`func (o *IamObjectWebhook) SetObjectFields(v []string)`

SetObjectFields sets ObjectFields field to given value.

### HasObjectFields

`func (o *IamObjectWebhook) HasObjectFields() bool`

HasObjectFields returns a boolean if a field has been set.

### GetOrganization

`func (o *IamObjectWebhook) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamObjectWebhook) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamObjectWebhook) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamObjectWebhook) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectWebhook) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectWebhook) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectWebhook) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectWebhook) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSingleOrgOnly

`func (o *IamObjectWebhook) GetSingleOrgOnly() bool`

GetSingleOrgOnly returns the SingleOrgOnly field if non-nil, zero value otherwise.

### GetSingleOrgOnlyOk

`func (o *IamObjectWebhook) GetSingleOrgOnlyOk() (*bool, bool)`

GetSingleOrgOnlyOk returns a tuple with the SingleOrgOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleOrgOnly

`func (o *IamObjectWebhook) SetSingleOrgOnly(v bool)`

SetSingleOrgOnly sets SingleOrgOnly field to given value.

### HasSingleOrgOnly

`func (o *IamObjectWebhook) HasSingleOrgOnly() bool`

HasSingleOrgOnly returns a boolean if a field has been set.

### GetTokenFields

`func (o *IamObjectWebhook) GetTokenFields() []string`

GetTokenFields returns the TokenFields field if non-nil, zero value otherwise.

### GetTokenFieldsOk

`func (o *IamObjectWebhook) GetTokenFieldsOk() (*[]string, bool)`

GetTokenFieldsOk returns a tuple with the TokenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFields

`func (o *IamObjectWebhook) SetTokenFields(v []string)`

SetTokenFields sets TokenFields field to given value.

### HasTokenFields

`func (o *IamObjectWebhook) HasTokenFields() bool`

HasTokenFields returns a boolean if a field has been set.

### GetUrl

`func (o *IamObjectWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IamObjectWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IamObjectWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IamObjectWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


