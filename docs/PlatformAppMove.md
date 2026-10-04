# PlatformAppMove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the declaration&#39;s name, from the path. | [optional] 
**Mode** | Pointer to **string** | Mode is &#x60;branch&#x60; (the default) or &#x60;commit&#x60;. | [optional] 
**Org** | Pointer to **string** | Org is the values directory the app lives in, defaulting to the caller&#39;s own org. | [optional] 
**PartOf** | Pointer to **string** | PartOf is the project to move it to. A name nothing names yet creates it. | [optional] 

## Methods

### NewPlatformAppMove

`func NewPlatformAppMove() *PlatformAppMove`

NewPlatformAppMove instantiates a new PlatformAppMove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformAppMoveWithDefaults

`func NewPlatformAppMoveWithDefaults() *PlatformAppMove`

NewPlatformAppMoveWithDefaults instantiates a new PlatformAppMove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PlatformAppMove) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PlatformAppMove) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PlatformAppMove) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PlatformAppMove) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetMode

`func (o *PlatformAppMove) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PlatformAppMove) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PlatformAppMove) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PlatformAppMove) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformAppMove) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformAppMove) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformAppMove) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformAppMove) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPartOf

`func (o *PlatformAppMove) GetPartOf() string`

GetPartOf returns the PartOf field if non-nil, zero value otherwise.

### GetPartOfOk

`func (o *PlatformAppMove) GetPartOfOk() (*string, bool)`

GetPartOfOk returns a tuple with the PartOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOf

`func (o *PlatformAppMove) SetPartOf(v string)`

SetPartOf sets PartOf field to given value.

### HasPartOf

`func (o *PlatformAppMove) HasPartOf() bool`

HasPartOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


