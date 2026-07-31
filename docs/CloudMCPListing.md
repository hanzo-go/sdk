# CloudMCPListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is the publisher&#39;s one-line summary. | [optional] 
**Featured** | Pointer to **bool** | Featured puts the listing on the front of the shelf. Curation. | [optional] 
**Hidden** | Pointer to **bool** | Hidden keeps the listing out of the org-visible catalog. Curation: a sync never changes it. Only a SuperAdmin sets it, and only a SuperAdmin sees a hidden entry listed. | [optional] 
**Id** | Pointer to **string** | ID addresses the listing in a URL. It is the reverse-DNS NAME with its one slash written as an underscore — reversible, because a namespace never contains an underscore — so the id is readable and stable rather than a hash that means nothing to whoever reads a link. | [optional] 
**Logo** | Pointer to **string** | Logo is the brand mark to render for the listing — the publisher&#39;s icon when the entry carries one, or the one an admin set. Curation. | [optional] 
**Name** | Pointer to **string** | Name is the publisher&#39;s reverse-DNS name, e.g. \&quot;com.stripe/mcp\&quot;. | [optional] 
**Official** | Pointer to **bool** | Official is whether this is the vendor&#39;s OWN server rather than someone else&#39;s copy of it. Derived on every sync (see isOfficial) until a SuperAdmin sets it explicitly, after which the admin&#39;s answer stands. | [optional] 
**Packages** | Pointer to [**[]CloudMCPPackage**](CloudMCPPackage.md) | Packages are the runnable package forms — npm, pypi, oci — each with the runtime that launches it and the transport it then speaks. | [optional] 
**Registry** | Pointer to **string** | Registry is the upstream this row was synced from. | [optional] 
**Remotes** | Pointer to [**[]CloudMCPRemote**](CloudMCPRemote.md) | Remotes are the hosted endpoints the publisher serves the server at. | [optional] 
**Repo** | Pointer to **string** | Repo is the source repository URL, when the entry names one. | [optional] 
**Site** | Pointer to **string** | Site is the project&#39;s homepage, when the entry names one. | [optional] 
**Synced** | Pointer to **int32** | Synced is when this row was last confirmed against upstream, Unix seconds. | [optional] 
**Title** | Pointer to **string** | Title is the human-readable display name, when the entry carries one. | [optional] 
**Transports** | Pointer to **[]string** | Transports are the distinct transports this server can be reached over, sorted: some of \&quot;stdio\&quot;, \&quot;streamable-http\&quot;, \&quot;sse\&quot;. A listing with \&quot;streamable-http\&quot; is one an org can enable here and now; a listing that is only \&quot;stdio\&quot; needs a process to run it. | [optional] 
**Vendor** | Pointer to **string** | Vendor is the namespace half of Name — the publisher, e.g. \&quot;com.stripe\&quot;. | [optional] 
**Version** | Pointer to **string** | Version is the published version of this listing. | [optional] 

## Methods

### NewCloudMCPListing

`func NewCloudMCPListing() *CloudMCPListing`

NewCloudMCPListing instantiates a new CloudMCPListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMCPListingWithDefaults

`func NewCloudMCPListingWithDefaults() *CloudMCPListing`

NewCloudMCPListingWithDefaults instantiates a new CloudMCPListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudMCPListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudMCPListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudMCPListing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudMCPListing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatured

`func (o *CloudMCPListing) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *CloudMCPListing) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *CloudMCPListing) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *CloudMCPListing) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetHidden

`func (o *CloudMCPListing) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *CloudMCPListing) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *CloudMCPListing) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *CloudMCPListing) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *CloudMCPListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudMCPListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudMCPListing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudMCPListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogo

`func (o *CloudMCPListing) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CloudMCPListing) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CloudMCPListing) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CloudMCPListing) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *CloudMCPListing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudMCPListing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudMCPListing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudMCPListing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfficial

`func (o *CloudMCPListing) GetOfficial() bool`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *CloudMCPListing) GetOfficialOk() (*bool, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *CloudMCPListing) SetOfficial(v bool)`

SetOfficial sets Official field to given value.

### HasOfficial

`func (o *CloudMCPListing) HasOfficial() bool`

HasOfficial returns a boolean if a field has been set.

### GetPackages

`func (o *CloudMCPListing) GetPackages() []CloudMCPPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *CloudMCPListing) GetPackagesOk() (*[]CloudMCPPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *CloudMCPListing) SetPackages(v []CloudMCPPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *CloudMCPListing) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetRegistry

`func (o *CloudMCPListing) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *CloudMCPListing) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *CloudMCPListing) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *CloudMCPListing) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRemotes

`func (o *CloudMCPListing) GetRemotes() []CloudMCPRemote`

GetRemotes returns the Remotes field if non-nil, zero value otherwise.

### GetRemotesOk

`func (o *CloudMCPListing) GetRemotesOk() (*[]CloudMCPRemote, bool)`

GetRemotesOk returns a tuple with the Remotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotes

`func (o *CloudMCPListing) SetRemotes(v []CloudMCPRemote)`

SetRemotes sets Remotes field to given value.

### HasRemotes

`func (o *CloudMCPListing) HasRemotes() bool`

HasRemotes returns a boolean if a field has been set.

### GetRepo

`func (o *CloudMCPListing) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudMCPListing) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudMCPListing) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudMCPListing) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSite

`func (o *CloudMCPListing) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CloudMCPListing) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CloudMCPListing) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *CloudMCPListing) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSynced

`func (o *CloudMCPListing) GetSynced() int32`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *CloudMCPListing) GetSyncedOk() (*int32, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *CloudMCPListing) SetSynced(v int32)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *CloudMCPListing) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### GetTitle

`func (o *CloudMCPListing) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudMCPListing) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudMCPListing) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudMCPListing) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTransports

`func (o *CloudMCPListing) GetTransports() []string`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *CloudMCPListing) GetTransportsOk() (*[]string, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *CloudMCPListing) SetTransports(v []string)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *CloudMCPListing) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetVendor

`func (o *CloudMCPListing) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CloudMCPListing) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CloudMCPListing) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CloudMCPListing) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *CloudMCPListing) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudMCPListing) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudMCPListing) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudMCPListing) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


