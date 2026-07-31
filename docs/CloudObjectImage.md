# CloudObjectImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageFamily** | Pointer to **string** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**ImageOwnerAlias** | Pointer to **string** |  | [optional] 
**ImageOwnerId** | Pointer to **int64** |  | [optional] 
**ImageVersion** | Pointer to **string** |  | [optional] 
**IsCopied** | Pointer to **bool** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**IsSelfShared** | Pointer to **string** |  | [optional] 
**IsSubscribed** | Pointer to **bool** |  | [optional] 
**IsSupportCloudinit** | Pointer to **bool** |  | [optional] 
**IsSupportIoOptimized** | Pointer to **bool** |  | [optional] 
**LoginAsNonRootSupported** | Pointer to **bool** |  | [optional] 
**OSNameEn** | Pointer to **string** |  | [optional] 
**OSType** | Pointer to **string** |  | [optional] 
**ProductCode** | Pointer to **string** |  | [optional] 
**ResourceGroupId** | Pointer to **string** |  | [optional] 
**SupplierName** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to **string** |  | [optional] 
**BootMode** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**RemotePassword** | Pointer to **string** |  | [optional] 
**RemotePort** | Pointer to **int64** |  | [optional] 
**RemoteProtocol** | Pointer to **string** |  | [optional] 
**RemoteUsername** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SystemArchitecture** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudObjectImage

`func NewCloudObjectImage() *CloudObjectImage`

NewCloudObjectImage instantiates a new CloudObjectImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectImageWithDefaults

`func NewCloudObjectImageWithDefaults() *CloudObjectImage`

NewCloudObjectImageWithDefaults instantiates a new CloudObjectImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageFamily

`func (o *CloudObjectImage) GetImageFamily() string`

GetImageFamily returns the ImageFamily field if non-nil, zero value otherwise.

### GetImageFamilyOk

`func (o *CloudObjectImage) GetImageFamilyOk() (*string, bool)`

GetImageFamilyOk returns a tuple with the ImageFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFamily

`func (o *CloudObjectImage) SetImageFamily(v string)`

SetImageFamily sets ImageFamily field to given value.

### HasImageFamily

`func (o *CloudObjectImage) HasImageFamily() bool`

HasImageFamily returns a boolean if a field has been set.

### GetImageName

`func (o *CloudObjectImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *CloudObjectImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *CloudObjectImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *CloudObjectImage) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageOwnerAlias

`func (o *CloudObjectImage) GetImageOwnerAlias() string`

GetImageOwnerAlias returns the ImageOwnerAlias field if non-nil, zero value otherwise.

### GetImageOwnerAliasOk

`func (o *CloudObjectImage) GetImageOwnerAliasOk() (*string, bool)`

GetImageOwnerAliasOk returns a tuple with the ImageOwnerAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOwnerAlias

`func (o *CloudObjectImage) SetImageOwnerAlias(v string)`

SetImageOwnerAlias sets ImageOwnerAlias field to given value.

### HasImageOwnerAlias

`func (o *CloudObjectImage) HasImageOwnerAlias() bool`

HasImageOwnerAlias returns a boolean if a field has been set.

### GetImageOwnerId

`func (o *CloudObjectImage) GetImageOwnerId() int64`

GetImageOwnerId returns the ImageOwnerId field if non-nil, zero value otherwise.

### GetImageOwnerIdOk

`func (o *CloudObjectImage) GetImageOwnerIdOk() (*int64, bool)`

GetImageOwnerIdOk returns a tuple with the ImageOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOwnerId

`func (o *CloudObjectImage) SetImageOwnerId(v int64)`

SetImageOwnerId sets ImageOwnerId field to given value.

### HasImageOwnerId

`func (o *CloudObjectImage) HasImageOwnerId() bool`

HasImageOwnerId returns a boolean if a field has been set.

### GetImageVersion

`func (o *CloudObjectImage) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *CloudObjectImage) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *CloudObjectImage) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *CloudObjectImage) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.

### GetIsCopied

`func (o *CloudObjectImage) GetIsCopied() bool`

GetIsCopied returns the IsCopied field if non-nil, zero value otherwise.

### GetIsCopiedOk

`func (o *CloudObjectImage) GetIsCopiedOk() (*bool, bool)`

GetIsCopiedOk returns a tuple with the IsCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopied

`func (o *CloudObjectImage) SetIsCopied(v bool)`

SetIsCopied sets IsCopied field to given value.

### HasIsCopied

`func (o *CloudObjectImage) HasIsCopied() bool`

HasIsCopied returns a boolean if a field has been set.

### GetIsPublic

`func (o *CloudObjectImage) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CloudObjectImage) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CloudObjectImage) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *CloudObjectImage) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsSelfShared

`func (o *CloudObjectImage) GetIsSelfShared() string`

GetIsSelfShared returns the IsSelfShared field if non-nil, zero value otherwise.

### GetIsSelfSharedOk

`func (o *CloudObjectImage) GetIsSelfSharedOk() (*string, bool)`

GetIsSelfSharedOk returns a tuple with the IsSelfShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfShared

`func (o *CloudObjectImage) SetIsSelfShared(v string)`

SetIsSelfShared sets IsSelfShared field to given value.

### HasIsSelfShared

`func (o *CloudObjectImage) HasIsSelfShared() bool`

HasIsSelfShared returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *CloudObjectImage) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *CloudObjectImage) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *CloudObjectImage) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *CloudObjectImage) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetIsSupportCloudinit

`func (o *CloudObjectImage) GetIsSupportCloudinit() bool`

GetIsSupportCloudinit returns the IsSupportCloudinit field if non-nil, zero value otherwise.

### GetIsSupportCloudinitOk

`func (o *CloudObjectImage) GetIsSupportCloudinitOk() (*bool, bool)`

GetIsSupportCloudinitOk returns a tuple with the IsSupportCloudinit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupportCloudinit

`func (o *CloudObjectImage) SetIsSupportCloudinit(v bool)`

SetIsSupportCloudinit sets IsSupportCloudinit field to given value.

### HasIsSupportCloudinit

`func (o *CloudObjectImage) HasIsSupportCloudinit() bool`

HasIsSupportCloudinit returns a boolean if a field has been set.

### GetIsSupportIoOptimized

`func (o *CloudObjectImage) GetIsSupportIoOptimized() bool`

GetIsSupportIoOptimized returns the IsSupportIoOptimized field if non-nil, zero value otherwise.

### GetIsSupportIoOptimizedOk

`func (o *CloudObjectImage) GetIsSupportIoOptimizedOk() (*bool, bool)`

GetIsSupportIoOptimizedOk returns a tuple with the IsSupportIoOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupportIoOptimized

`func (o *CloudObjectImage) SetIsSupportIoOptimized(v bool)`

SetIsSupportIoOptimized sets IsSupportIoOptimized field to given value.

### HasIsSupportIoOptimized

`func (o *CloudObjectImage) HasIsSupportIoOptimized() bool`

HasIsSupportIoOptimized returns a boolean if a field has been set.

### GetLoginAsNonRootSupported

`func (o *CloudObjectImage) GetLoginAsNonRootSupported() bool`

GetLoginAsNonRootSupported returns the LoginAsNonRootSupported field if non-nil, zero value otherwise.

### GetLoginAsNonRootSupportedOk

`func (o *CloudObjectImage) GetLoginAsNonRootSupportedOk() (*bool, bool)`

GetLoginAsNonRootSupportedOk returns a tuple with the LoginAsNonRootSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAsNonRootSupported

`func (o *CloudObjectImage) SetLoginAsNonRootSupported(v bool)`

SetLoginAsNonRootSupported sets LoginAsNonRootSupported field to given value.

### HasLoginAsNonRootSupported

`func (o *CloudObjectImage) HasLoginAsNonRootSupported() bool`

HasLoginAsNonRootSupported returns a boolean if a field has been set.

### GetOSNameEn

`func (o *CloudObjectImage) GetOSNameEn() string`

GetOSNameEn returns the OSNameEn field if non-nil, zero value otherwise.

### GetOSNameEnOk

`func (o *CloudObjectImage) GetOSNameEnOk() (*string, bool)`

GetOSNameEnOk returns a tuple with the OSNameEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSNameEn

`func (o *CloudObjectImage) SetOSNameEn(v string)`

SetOSNameEn sets OSNameEn field to given value.

### HasOSNameEn

`func (o *CloudObjectImage) HasOSNameEn() bool`

HasOSNameEn returns a boolean if a field has been set.

### GetOSType

`func (o *CloudObjectImage) GetOSType() string`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *CloudObjectImage) GetOSTypeOk() (*string, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *CloudObjectImage) SetOSType(v string)`

SetOSType sets OSType field to given value.

### HasOSType

`func (o *CloudObjectImage) HasOSType() bool`

HasOSType returns a boolean if a field has been set.

### GetProductCode

`func (o *CloudObjectImage) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *CloudObjectImage) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *CloudObjectImage) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *CloudObjectImage) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetResourceGroupId

`func (o *CloudObjectImage) GetResourceGroupId() string`

GetResourceGroupId returns the ResourceGroupId field if non-nil, zero value otherwise.

### GetResourceGroupIdOk

`func (o *CloudObjectImage) GetResourceGroupIdOk() (*string, bool)`

GetResourceGroupIdOk returns a tuple with the ResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupId

`func (o *CloudObjectImage) SetResourceGroupId(v string)`

SetResourceGroupId sets ResourceGroupId field to given value.

### HasResourceGroupId

`func (o *CloudObjectImage) HasResourceGroupId() bool`

HasResourceGroupId returns a boolean if a field has been set.

### GetSupplierName

`func (o *CloudObjectImage) GetSupplierName() string`

GetSupplierName returns the SupplierName field if non-nil, zero value otherwise.

### GetSupplierNameOk

`func (o *CloudObjectImage) GetSupplierNameOk() (*string, bool)`

GetSupplierNameOk returns a tuple with the SupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierName

`func (o *CloudObjectImage) SetSupplierName(v string)`

SetSupplierName sets SupplierName field to given value.

### HasSupplierName

`func (o *CloudObjectImage) HasSupplierName() bool`

HasSupplierName returns a boolean if a field has been set.

### GetUsage

`func (o *CloudObjectImage) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CloudObjectImage) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CloudObjectImage) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CloudObjectImage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetBootMode

`func (o *CloudObjectImage) GetBootMode() string`

GetBootMode returns the BootMode field if non-nil, zero value otherwise.

### GetBootModeOk

`func (o *CloudObjectImage) GetBootModeOk() (*string, bool)`

GetBootModeOk returns a tuple with the BootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootMode

`func (o *CloudObjectImage) SetBootMode(v string)`

SetBootMode sets BootMode field to given value.

### HasBootMode

`func (o *CloudObjectImage) HasBootMode() bool`

HasBootMode returns a boolean if a field has been set.

### GetCategory

`func (o *CloudObjectImage) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudObjectImage) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudObjectImage) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudObjectImage) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CloudObjectImage) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectImage) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectImage) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectImage) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *CloudObjectImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudObjectImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudObjectImage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudObjectImage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudObjectImage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudObjectImage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudObjectImage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudObjectImage) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetImageId

`func (o *CloudObjectImage) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CloudObjectImage) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CloudObjectImage) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CloudObjectImage) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetName

`func (o *CloudObjectImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudObjectImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *CloudObjectImage) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CloudObjectImage) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CloudObjectImage) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CloudObjectImage) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOwner

`func (o *CloudObjectImage) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectImage) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectImage) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudObjectImage) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPlatform

`func (o *CloudObjectImage) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CloudObjectImage) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CloudObjectImage) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CloudObjectImage) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProgress

`func (o *CloudObjectImage) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *CloudObjectImage) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *CloudObjectImage) SetProgress(v string)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *CloudObjectImage) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProvider

`func (o *CloudObjectImage) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudObjectImage) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudObjectImage) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudObjectImage) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRemotePassword

`func (o *CloudObjectImage) GetRemotePassword() string`

GetRemotePassword returns the RemotePassword field if non-nil, zero value otherwise.

### GetRemotePasswordOk

`func (o *CloudObjectImage) GetRemotePasswordOk() (*string, bool)`

GetRemotePasswordOk returns a tuple with the RemotePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePassword

`func (o *CloudObjectImage) SetRemotePassword(v string)`

SetRemotePassword sets RemotePassword field to given value.

### HasRemotePassword

`func (o *CloudObjectImage) HasRemotePassword() bool`

HasRemotePassword returns a boolean if a field has been set.

### GetRemotePort

`func (o *CloudObjectImage) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *CloudObjectImage) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *CloudObjectImage) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *CloudObjectImage) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetRemoteProtocol

`func (o *CloudObjectImage) GetRemoteProtocol() string`

GetRemoteProtocol returns the RemoteProtocol field if non-nil, zero value otherwise.

### GetRemoteProtocolOk

`func (o *CloudObjectImage) GetRemoteProtocolOk() (*string, bool)`

GetRemoteProtocolOk returns a tuple with the RemoteProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProtocol

`func (o *CloudObjectImage) SetRemoteProtocol(v string)`

SetRemoteProtocol sets RemoteProtocol field to given value.

### HasRemoteProtocol

`func (o *CloudObjectImage) HasRemoteProtocol() bool`

HasRemoteProtocol returns a boolean if a field has been set.

### GetRemoteUsername

`func (o *CloudObjectImage) GetRemoteUsername() string`

GetRemoteUsername returns the RemoteUsername field if non-nil, zero value otherwise.

### GetRemoteUsernameOk

`func (o *CloudObjectImage) GetRemoteUsernameOk() (*string, bool)`

GetRemoteUsernameOk returns a tuple with the RemoteUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUsername

`func (o *CloudObjectImage) SetRemoteUsername(v string)`

SetRemoteUsername sets RemoteUsername field to given value.

### HasRemoteUsername

`func (o *CloudObjectImage) HasRemoteUsername() bool`

HasRemoteUsername returns a boolean if a field has been set.

### GetSize

`func (o *CloudObjectImage) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CloudObjectImage) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CloudObjectImage) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CloudObjectImage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *CloudObjectImage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudObjectImage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudObjectImage) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudObjectImage) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSystemArchitecture

`func (o *CloudObjectImage) GetSystemArchitecture() string`

GetSystemArchitecture returns the SystemArchitecture field if non-nil, zero value otherwise.

### GetSystemArchitectureOk

`func (o *CloudObjectImage) GetSystemArchitectureOk() (*string, bool)`

GetSystemArchitectureOk returns a tuple with the SystemArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemArchitecture

`func (o *CloudObjectImage) SetSystemArchitecture(v string)`

SetSystemArchitecture sets SystemArchitecture field to given value.

### HasSystemArchitecture

`func (o *CloudObjectImage) HasSystemArchitecture() bool`

HasSystemArchitecture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


