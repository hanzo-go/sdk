# NexusImage

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

### NewNexusImage

`func NewNexusImage() *NexusImage`

NewNexusImage instantiates a new NexusImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusImageWithDefaults

`func NewNexusImageWithDefaults() *NexusImage`

NewNexusImageWithDefaults instantiates a new NexusImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageFamily

`func (o *NexusImage) GetImageFamily() string`

GetImageFamily returns the ImageFamily field if non-nil, zero value otherwise.

### GetImageFamilyOk

`func (o *NexusImage) GetImageFamilyOk() (*string, bool)`

GetImageFamilyOk returns a tuple with the ImageFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFamily

`func (o *NexusImage) SetImageFamily(v string)`

SetImageFamily sets ImageFamily field to given value.

### HasImageFamily

`func (o *NexusImage) HasImageFamily() bool`

HasImageFamily returns a boolean if a field has been set.

### GetImageName

`func (o *NexusImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *NexusImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *NexusImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *NexusImage) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageOwnerAlias

`func (o *NexusImage) GetImageOwnerAlias() string`

GetImageOwnerAlias returns the ImageOwnerAlias field if non-nil, zero value otherwise.

### GetImageOwnerAliasOk

`func (o *NexusImage) GetImageOwnerAliasOk() (*string, bool)`

GetImageOwnerAliasOk returns a tuple with the ImageOwnerAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOwnerAlias

`func (o *NexusImage) SetImageOwnerAlias(v string)`

SetImageOwnerAlias sets ImageOwnerAlias field to given value.

### HasImageOwnerAlias

`func (o *NexusImage) HasImageOwnerAlias() bool`

HasImageOwnerAlias returns a boolean if a field has been set.

### GetImageOwnerId

`func (o *NexusImage) GetImageOwnerId() int64`

GetImageOwnerId returns the ImageOwnerId field if non-nil, zero value otherwise.

### GetImageOwnerIdOk

`func (o *NexusImage) GetImageOwnerIdOk() (*int64, bool)`

GetImageOwnerIdOk returns a tuple with the ImageOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOwnerId

`func (o *NexusImage) SetImageOwnerId(v int64)`

SetImageOwnerId sets ImageOwnerId field to given value.

### HasImageOwnerId

`func (o *NexusImage) HasImageOwnerId() bool`

HasImageOwnerId returns a boolean if a field has been set.

### GetImageVersion

`func (o *NexusImage) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *NexusImage) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *NexusImage) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *NexusImage) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.

### GetIsCopied

`func (o *NexusImage) GetIsCopied() bool`

GetIsCopied returns the IsCopied field if non-nil, zero value otherwise.

### GetIsCopiedOk

`func (o *NexusImage) GetIsCopiedOk() (*bool, bool)`

GetIsCopiedOk returns a tuple with the IsCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopied

`func (o *NexusImage) SetIsCopied(v bool)`

SetIsCopied sets IsCopied field to given value.

### HasIsCopied

`func (o *NexusImage) HasIsCopied() bool`

HasIsCopied returns a boolean if a field has been set.

### GetIsPublic

`func (o *NexusImage) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *NexusImage) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *NexusImage) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *NexusImage) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsSelfShared

`func (o *NexusImage) GetIsSelfShared() string`

GetIsSelfShared returns the IsSelfShared field if non-nil, zero value otherwise.

### GetIsSelfSharedOk

`func (o *NexusImage) GetIsSelfSharedOk() (*string, bool)`

GetIsSelfSharedOk returns a tuple with the IsSelfShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfShared

`func (o *NexusImage) SetIsSelfShared(v string)`

SetIsSelfShared sets IsSelfShared field to given value.

### HasIsSelfShared

`func (o *NexusImage) HasIsSelfShared() bool`

HasIsSelfShared returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *NexusImage) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *NexusImage) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *NexusImage) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *NexusImage) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetIsSupportCloudinit

`func (o *NexusImage) GetIsSupportCloudinit() bool`

GetIsSupportCloudinit returns the IsSupportCloudinit field if non-nil, zero value otherwise.

### GetIsSupportCloudinitOk

`func (o *NexusImage) GetIsSupportCloudinitOk() (*bool, bool)`

GetIsSupportCloudinitOk returns a tuple with the IsSupportCloudinit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupportCloudinit

`func (o *NexusImage) SetIsSupportCloudinit(v bool)`

SetIsSupportCloudinit sets IsSupportCloudinit field to given value.

### HasIsSupportCloudinit

`func (o *NexusImage) HasIsSupportCloudinit() bool`

HasIsSupportCloudinit returns a boolean if a field has been set.

### GetIsSupportIoOptimized

`func (o *NexusImage) GetIsSupportIoOptimized() bool`

GetIsSupportIoOptimized returns the IsSupportIoOptimized field if non-nil, zero value otherwise.

### GetIsSupportIoOptimizedOk

`func (o *NexusImage) GetIsSupportIoOptimizedOk() (*bool, bool)`

GetIsSupportIoOptimizedOk returns a tuple with the IsSupportIoOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupportIoOptimized

`func (o *NexusImage) SetIsSupportIoOptimized(v bool)`

SetIsSupportIoOptimized sets IsSupportIoOptimized field to given value.

### HasIsSupportIoOptimized

`func (o *NexusImage) HasIsSupportIoOptimized() bool`

HasIsSupportIoOptimized returns a boolean if a field has been set.

### GetLoginAsNonRootSupported

`func (o *NexusImage) GetLoginAsNonRootSupported() bool`

GetLoginAsNonRootSupported returns the LoginAsNonRootSupported field if non-nil, zero value otherwise.

### GetLoginAsNonRootSupportedOk

`func (o *NexusImage) GetLoginAsNonRootSupportedOk() (*bool, bool)`

GetLoginAsNonRootSupportedOk returns a tuple with the LoginAsNonRootSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAsNonRootSupported

`func (o *NexusImage) SetLoginAsNonRootSupported(v bool)`

SetLoginAsNonRootSupported sets LoginAsNonRootSupported field to given value.

### HasLoginAsNonRootSupported

`func (o *NexusImage) HasLoginAsNonRootSupported() bool`

HasLoginAsNonRootSupported returns a boolean if a field has been set.

### GetOSNameEn

`func (o *NexusImage) GetOSNameEn() string`

GetOSNameEn returns the OSNameEn field if non-nil, zero value otherwise.

### GetOSNameEnOk

`func (o *NexusImage) GetOSNameEnOk() (*string, bool)`

GetOSNameEnOk returns a tuple with the OSNameEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSNameEn

`func (o *NexusImage) SetOSNameEn(v string)`

SetOSNameEn sets OSNameEn field to given value.

### HasOSNameEn

`func (o *NexusImage) HasOSNameEn() bool`

HasOSNameEn returns a boolean if a field has been set.

### GetOSType

`func (o *NexusImage) GetOSType() string`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *NexusImage) GetOSTypeOk() (*string, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *NexusImage) SetOSType(v string)`

SetOSType sets OSType field to given value.

### HasOSType

`func (o *NexusImage) HasOSType() bool`

HasOSType returns a boolean if a field has been set.

### GetProductCode

`func (o *NexusImage) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *NexusImage) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *NexusImage) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *NexusImage) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetResourceGroupId

`func (o *NexusImage) GetResourceGroupId() string`

GetResourceGroupId returns the ResourceGroupId field if non-nil, zero value otherwise.

### GetResourceGroupIdOk

`func (o *NexusImage) GetResourceGroupIdOk() (*string, bool)`

GetResourceGroupIdOk returns a tuple with the ResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupId

`func (o *NexusImage) SetResourceGroupId(v string)`

SetResourceGroupId sets ResourceGroupId field to given value.

### HasResourceGroupId

`func (o *NexusImage) HasResourceGroupId() bool`

HasResourceGroupId returns a boolean if a field has been set.

### GetSupplierName

`func (o *NexusImage) GetSupplierName() string`

GetSupplierName returns the SupplierName field if non-nil, zero value otherwise.

### GetSupplierNameOk

`func (o *NexusImage) GetSupplierNameOk() (*string, bool)`

GetSupplierNameOk returns a tuple with the SupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierName

`func (o *NexusImage) SetSupplierName(v string)`

SetSupplierName sets SupplierName field to given value.

### HasSupplierName

`func (o *NexusImage) HasSupplierName() bool`

HasSupplierName returns a boolean if a field has been set.

### GetUsage

`func (o *NexusImage) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *NexusImage) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *NexusImage) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *NexusImage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetBootMode

`func (o *NexusImage) GetBootMode() string`

GetBootMode returns the BootMode field if non-nil, zero value otherwise.

### GetBootModeOk

`func (o *NexusImage) GetBootModeOk() (*string, bool)`

GetBootModeOk returns a tuple with the BootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootMode

`func (o *NexusImage) SetBootMode(v string)`

SetBootMode sets BootMode field to given value.

### HasBootMode

`func (o *NexusImage) HasBootMode() bool`

HasBootMode returns a boolean if a field has been set.

### GetCategory

`func (o *NexusImage) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NexusImage) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NexusImage) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NexusImage) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NexusImage) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NexusImage) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NexusImage) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NexusImage) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *NexusImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NexusImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NexusImage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NexusImage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *NexusImage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NexusImage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NexusImage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NexusImage) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetImageId

`func (o *NexusImage) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *NexusImage) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *NexusImage) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *NexusImage) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetName

`func (o *NexusImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NexusImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NexusImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NexusImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *NexusImage) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *NexusImage) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *NexusImage) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *NexusImage) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOwner

`func (o *NexusImage) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NexusImage) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NexusImage) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NexusImage) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPlatform

`func (o *NexusImage) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *NexusImage) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *NexusImage) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *NexusImage) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProgress

`func (o *NexusImage) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *NexusImage) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *NexusImage) SetProgress(v string)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *NexusImage) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProvider

`func (o *NexusImage) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NexusImage) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NexusImage) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NexusImage) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRemotePassword

`func (o *NexusImage) GetRemotePassword() string`

GetRemotePassword returns the RemotePassword field if non-nil, zero value otherwise.

### GetRemotePasswordOk

`func (o *NexusImage) GetRemotePasswordOk() (*string, bool)`

GetRemotePasswordOk returns a tuple with the RemotePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePassword

`func (o *NexusImage) SetRemotePassword(v string)`

SetRemotePassword sets RemotePassword field to given value.

### HasRemotePassword

`func (o *NexusImage) HasRemotePassword() bool`

HasRemotePassword returns a boolean if a field has been set.

### GetRemotePort

`func (o *NexusImage) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *NexusImage) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *NexusImage) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *NexusImage) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetRemoteProtocol

`func (o *NexusImage) GetRemoteProtocol() string`

GetRemoteProtocol returns the RemoteProtocol field if non-nil, zero value otherwise.

### GetRemoteProtocolOk

`func (o *NexusImage) GetRemoteProtocolOk() (*string, bool)`

GetRemoteProtocolOk returns a tuple with the RemoteProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProtocol

`func (o *NexusImage) SetRemoteProtocol(v string)`

SetRemoteProtocol sets RemoteProtocol field to given value.

### HasRemoteProtocol

`func (o *NexusImage) HasRemoteProtocol() bool`

HasRemoteProtocol returns a boolean if a field has been set.

### GetRemoteUsername

`func (o *NexusImage) GetRemoteUsername() string`

GetRemoteUsername returns the RemoteUsername field if non-nil, zero value otherwise.

### GetRemoteUsernameOk

`func (o *NexusImage) GetRemoteUsernameOk() (*string, bool)`

GetRemoteUsernameOk returns a tuple with the RemoteUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUsername

`func (o *NexusImage) SetRemoteUsername(v string)`

SetRemoteUsername sets RemoteUsername field to given value.

### HasRemoteUsername

`func (o *NexusImage) HasRemoteUsername() bool`

HasRemoteUsername returns a boolean if a field has been set.

### GetSize

`func (o *NexusImage) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NexusImage) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NexusImage) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *NexusImage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *NexusImage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NexusImage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NexusImage) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NexusImage) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSystemArchitecture

`func (o *NexusImage) GetSystemArchitecture() string`

GetSystemArchitecture returns the SystemArchitecture field if non-nil, zero value otherwise.

### GetSystemArchitectureOk

`func (o *NexusImage) GetSystemArchitectureOk() (*string, bool)`

GetSystemArchitectureOk returns a tuple with the SystemArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemArchitecture

`func (o *NexusImage) SetSystemArchitecture(v string)`

SetSystemArchitecture sets SystemArchitecture field to given value.

### HasSystemArchitecture

`func (o *NexusImage) HasSystemArchitecture() bool`

HasSystemArchitecture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


