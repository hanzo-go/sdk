# \IamAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProvider**](IamAPI.md#AddProvider) | **Post** /v1/iam/providers | Adds an identity provider your people can sign in with, or a service your applications send through — a social or enterprise login, an email or SMS sender, a storage or payment connector.
[**AddToken**](IamAPI.md#AddToken) | **Post** /v1/iam/tokens | Records an access token — the credential an application or integration presents on a caller&#39;s behalf.
[**AddWebauthnCredential**](IamAPI.md#AddWebauthnCredential) | **Post** /v1/iam/webauthn-credentials | Registers a passkey or security key for a person, so they can sign in with their device instead of a password.
[**CreateOrganization**](IamAPI.md#CreateOrganization) | **Post** /v1/iam/organizations | Makes a new organization — the account your users, applications, roles, projects and workspaces are all named inside.
[**CreateSession**](IamAPI.md#CreateSession) | **Post** /v1/iam/sessions/create | Records a sign-in.
[**DeleteIamApplication**](IamAPI.md#DeleteIamApplication) | **Delete** /v1/iam/application | Removes an application.
[**DeleteIamScimV2UsersByOwnerByName**](IamAPI.md#DeleteIamScimV2UsersByOwnerByName) | **Delete** /v1/iam/scim/v2/Users/{owner}/{name} | Deprovisions a person — how removing someone in your identity provider removes their access here.
[**DeleteIamServiceAccountsByName**](IamAPI.md#DeleteIamServiceAccountsByName) | **Delete** /v1/iam/service-accounts/{name} | Serves DELETE /v1/iam/service-accounts/:name.
[**DeleteOrganization**](IamAPI.md#DeleteOrganization) | **Post** /v1/iam/organizations/delete | Removes an organization and everything named inside it.
[**DeleteProvider**](IamAPI.md#DeleteProvider) | **Post** /v1/iam/providers/delete | Removes a provider.
[**DeleteSession**](IamAPI.md#DeleteSession) | **Post** /v1/iam/sessions/delete | Signs a person out of one application — the session ends and every browser carrying it stops being authenticated.
[**DeleteToken**](IamAPI.md#DeleteToken) | **Post** /v1/iam/tokens/delete | Revokes an access token.
[**DeleteWebauthnCredential**](IamAPI.md#DeleteWebauthnCredential) | **Post** /v1/iam/webauthn-credentials/delete | Removes a passkey or security key — what you call when a device is lost.
[**GetIamAccount**](IamAPI.md#GetIamAccount) | **Get** /v1/iam/account | Returns the signed-in person&#39;s own account and the organization they belong to — what a console reads to draw the account menu.
[**GetIamApplication**](IamAPI.md#GetIamApplication) | **Get** /v1/iam/application | Returns one application: its sign-in methods, its allowed redirect URIs and the client credentials your integration authenticates with.
[**GetIamApplications**](IamAPI.md#GetIamApplications) | **Get** /v1/iam/applications | Returns the applications in one organization, newest first — each product or site your people sign in to, with the sign-in methods and redirect URIs it allows.
[**GetIamApplicationsGet**](IamAPI.md#GetIamApplicationsGet) | **Get** /v1/iam/applications/get | Returns one application: its sign-in methods, its allowed redirect URIs and the client credentials your integration authenticates with.
[**GetIamAuditLogs**](IamAPI.md#GetIamAuditLogs) | **Get** /v1/iam/audit-logs | Returns your organization&#39;s audit trail, newest first — who did what, when, and from where.
[**GetIamAuthApplication**](IamAPI.md#GetIamAuthApplication) | **Get** /v1/iam/auth/application | Returns everything a login screen needs to draw itself for one application: its branding, and each sign-in method it offers with the provider details that method needs.
[**GetIamAuthMethods**](IamAPI.md#GetIamAuthMethods) | **Get** /v1/iam/auth/methods | Returns the sign-in methods one application actually has switched on, so a login screen can render the right buttons for it without you hard-coding a list that drifts the moment you add a provider.
[**GetIamCerts**](IamAPI.md#GetIamCerts) | **Get** /v1/iam/certs | Returns your organization&#39;s signing certificates, newest first — the keys the tokens your applications verify are signed with.
[**GetIamConsent**](IamAPI.md#GetIamConsent) | **Get** /v1/iam/consent | Returns the calling person&#39;s own privacy and communication choices.
[**GetIamGetAccount**](IamAPI.md#GetIamGetAccount) | **Get** /v1/iam/get-account | Returns the signed-in person&#39;s own account and the organization they belong to — what a console reads to draw the account menu.
[**GetIamGetAppLogin**](IamAPI.md#GetIamGetAppLogin) | **Get** /v1/iam/get-app-login | Returns everything a login screen needs to draw itself for one application: its branding, and each sign-in method it offers with the provider details that method needs.
[**GetIamGetApplication**](IamAPI.md#GetIamGetApplication) | **Get** /v1/iam/get-application | Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.
[**GetIamGetApplications**](IamAPI.md#GetIamGetApplications) | **Get** /v1/iam/get-applications | Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.
[**GetIamGetCert**](IamAPI.md#GetIamGetCert) | **Get** /v1/iam/get-cert | Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.
[**GetIamGetCerts**](IamAPI.md#GetIamGetCerts) | **Get** /v1/iam/get-certs | Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.
[**GetIamGetGlobalUsers**](IamAPI.md#GetIamGetGlobalUsers) | **Get** /v1/iam/get-global-users | Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.
[**GetIamGetInvitations**](IamAPI.md#GetIamGetInvitations) | **Get** /v1/iam/get-invitations | Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.
[**GetIamGetMemberships**](IamAPI.md#GetIamGetMemberships) | **Get** /v1/iam/get-memberships | Answers either question about who belongs where: which organizations one person can act in, or who can act in one organization.
[**GetIamGetOrganization**](IamAPI.md#GetIamGetOrganization) | **Get** /v1/iam/get-organization | Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.
[**GetIamGetOrganizationProjects**](IamAPI.md#GetIamGetOrganizationProjects) | **Get** /v1/iam/get-organization-projects | Returns one organization&#39;s projects — what a scope switcher lists so somebody can move between them.
[**GetIamGetOrganizationWorkspaces**](IamAPI.md#GetIamGetOrganizationWorkspaces) | **Get** /v1/iam/get-organization-workspaces | Returns one organization&#39;s workspaces — what a scope switcher lists so somebody can move between them.
[**GetIamGetOrganizations**](IamAPI.md#GetIamGetOrganizations) | **Get** /v1/iam/get-organizations | Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.
[**GetIamGetPermission**](IamAPI.md#GetIamGetPermission) | **Get** /v1/iam/get-permission | Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.
[**GetIamGetPermissions**](IamAPI.md#GetIamGetPermissions) | **Get** /v1/iam/get-permissions | Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.
[**GetIamGetProvider**](IamAPI.md#GetIamGetProvider) | **Get** /v1/iam/get-provider | Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.
[**GetIamGetProviders**](IamAPI.md#GetIamGetProviders) | **Get** /v1/iam/get-providers | Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.
[**GetIamGetRecords**](IamAPI.md#GetIamGetRecords) | **Get** /v1/iam/get-records | Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.
[**GetIamGetRole**](IamAPI.md#GetIamGetRole) | **Get** /v1/iam/get-role | Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.
[**GetIamGetRoles**](IamAPI.md#GetIamGetRoles) | **Get** /v1/iam/get-roles | Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.
[**GetIamGetUser**](IamAPI.md#GetIamGetUser) | **Get** /v1/iam/get-user | Reads one person, two ways.
[**GetIamGetUsers**](IamAPI.md#GetIamGetUsers) | **Get** /v1/iam/get-users | Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.
[**GetIamInvitations**](IamAPI.md#GetIamInvitations) | **Get** /v1/iam/invitations | Returns your organization&#39;s invitations, newest first — who has been asked to join, on what terms, and how many seats each invitation still has left.
[**GetIamKeys**](IamAPI.md#GetIamKeys) | **Get** /v1/iam/keys | Returns your organization&#39;s API keys, newest first — what each is called, what it may reach, and its publishable half.
[**GetIamKeysGet**](IamAPI.md#GetIamKeysGet) | **Get** /v1/iam/keys/get | Returns one API key: what it is called, what it may reach, and when it was issued.
[**GetIamLinkedAccounts**](IamAPI.md#GetIamLinkedAccounts) | **Get** /v1/iam/linked-accounts | Returns the sign-in identities linked to the calling person&#39;s account — every provider they can currently sign in with.
[**GetIamMemberships**](IamAPI.md#GetIamMemberships) | **Get** /v1/iam/memberships | Answers either question about who belongs where: which organizations one person can act in, or who can act in one organization.
[**GetIamOauthAuthorize**](IamAPI.md#GetIamOauthAuthorize) | **Get** /v1/iam/oauth/authorize | Starts a sign-in — the address you send a browser to, and the beginning of every OAuth and OpenID Connect flow.
[**GetIamOauthCallback**](IamAPI.md#GetIamOauthCallback) | **Get** /v1/iam/oauth/callback | Completes the round-trip: it resolves and burns the single-use transaction (checking expiry + browser binding), exchanges and verifies the IdP response, links or provisions the local user, and mints the iam authorization code the relying party expects — then redirects to the original redirect_uri with code + state.
[**GetIamOauthLogout**](IamAPI.md#GetIamOauthLogout) | **Get** /v1/iam/oauth/logout | Ends a sign-in and sends the browser somewhere sensible.
[**GetIamOauthUserinfo**](IamAPI.md#GetIamOauthUserinfo) | **Get** /v1/iam/oauth/userinfo | Returns the profile claims for whoever the access token belongs to — the standard OpenID Connect way to find out who is calling you without your application storing anything itself.
[**GetIamPermissions**](IamAPI.md#GetIamPermissions) | **Get** /v1/iam/permissions | Returns the permissions in one organization, newest first — each one a grant saying which people or roles may do what, and to which resources.
[**GetIamPermissionsGet**](IamAPI.md#GetIamPermissionsGet) | **Get** /v1/iam/permissions/get | Returns one permission: who it grants to, what it allows, and the resources it covers.
[**GetIamProjects**](IamAPI.md#GetIamProjects) | **Get** /v1/iam/projects | Returns your organization&#39;s projects, newest first — the scope people pick between when their work is separated by product or client rather than by team.
[**GetIamRegistryJwks**](IamAPI.md#GetIamRegistryJwks) | **Get** /v1/iam/registry/jwks | Publishes the public key your registry uses to verify the tokens issued above — the one URL to configure so the registry trusts logins without holding any secret of its own.
[**GetIamRegistryToken**](IamAPI.md#GetIamRegistryToken) | **Get** /v1/iam/registry/token | Signs a container client in to your registry.
[**GetIamResolveKey**](IamAPI.md#GetIamResolveKey) | **Get** /v1/iam/resolve-key | Answers which organization a PUBLISHABLE key belongs to — what a service of yours calls to attribute a request that arrived carrying a key shipped in a browser.
[**GetIamRoles**](IamAPI.md#GetIamRoles) | **Get** /v1/iam/roles | Returns your organization&#39;s roles, newest first — each a named group of people that permissions are granted to.
[**GetIamScimV2Resourcetypes**](IamAPI.md#GetIamScimV2Resourcetypes) | **Get** /v1/iam/scim/v2/ResourceTypes | Returns the kinds of record this directory provisions and the address of each, so your identity provider discovers them rather than having them configured by hand.
[**GetIamScimV2ResourcetypesByName**](IamAPI.md#GetIamScimV2ResourcetypesByName) | **Get** /v1/iam/scim/v2/ResourceTypes/{name} | Returns one provisionable record kind in full.
[**GetIamScimV2Schemas**](IamAPI.md#GetIamScimV2Schemas) | **Get** /v1/iam/scim/v2/Schemas | Returns the attribute definitions this directory understands, so your identity provider knows which fields it may send and what they mean before it sends any.
[**GetIamScimV2SchemasById**](IamAPI.md#GetIamScimV2SchemasById) | **Get** /v1/iam/scim/v2/Schemas/{id} | Returns one attribute definition in full.
[**GetIamScimV2Serviceproviderconfig**](IamAPI.md#GetIamScimV2Serviceproviderconfig) | **Get** /v1/iam/scim/v2/ServiceProviderConfig | Tells your identity provider which parts of SCIM this directory supports, so it configures itself instead of you filling in a form.
[**GetIamScimV2Users**](IamAPI.md#GetIamScimV2Users) | **Get** /v1/iam/scim/v2/Users | Returns the people in your organization to your identity provider, in the standard SCIM shape, so an IdP can reconcile its directory against ours.
[**GetIamScimV2UsersByOwnerByName**](IamAPI.md#GetIamScimV2UsersByOwnerByName) | **Get** /v1/iam/scim/v2/Users/{owner}/{name} | Returns one person in the standard SCIM shape.
[**GetIamServiceAccounts**](IamAPI.md#GetIamServiceAccounts) | **Get** /v1/iam/service-accounts | Returns your organization&#39;s service accounts — what each is called and when it was created.
[**GetIamUsers**](IamAPI.md#GetIamUsers) | **Get** /v1/iam/users | Returns a page of the people in your organization, with the total so you can page through the rest.
[**GetIamUsersGet**](IamAPI.md#GetIamUsersGet) | **Get** /v1/iam/users/get | Returns one person in your organization, addressed by their username or by their email address.
[**GetIamWeb3Nonce**](IamAPI.md#GetIamWeb3Nonce) | **Get** /v1/iam/web3/nonce | Starts a wallet sign-in: it returns a one-time challenge for the wallet to sign.
[**GetIamWebauthnSigninBegin**](IamAPI.md#GetIamWebauthnSigninBegin) | **Get** /v1/iam/webauthn/signin/begin | Starts a passkey sign-in: it returns the challenge the person&#39;s authenticator signs.
[**GetIamWebauthnSignupBegin**](IamAPI.md#GetIamWebauthnSignupBegin) | **Get** /v1/iam/webauthn/signup/begin | Starts enrolling a passkey for the signed-in person: it returns the options their browser hands to the authenticator.
[**GetIamWellKnownJwks**](IamAPI.md#GetIamWellKnownJwks) | **Get** /v1/iam/.well-known/jwks | Publishes the public keys that verify the tokens issued here — the one URL you point a service at so it can check a token itself, offline, without calling back and without holding any secret of ours.
[**GetIamWellKnownOauthAuthorizationServer**](IamAPI.md#GetIamWellKnownOauthAuthorizationServer) | **Get** /v1/iam/.well-known/oauth-authorization-server | Returns the OpenID Connect discovery document — the one URL you point a standards-compliant client at so it can find every other endpoint on its own, instead of you configuring them by hand.
[**GetIamWellKnownOpenidConfiguration**](IamAPI.md#GetIamWellKnownOpenidConfiguration) | **Get** /v1/iam/.well-known/openid-configuration | Returns the OpenID Connect discovery document — the one URL you point a standards-compliant client at so it can find every other endpoint on its own, instead of you configuring them by hand.
[**GetIamWhoami**](IamAPI.md#GetIamWhoami) | **Get** /v1/iam/whoami | Tells you who the current caller is — the lightweight check a page makes on load to decide whether to render signed-in or signed-out.
[**GetIamWorkspaces**](IamAPI.md#GetIamWorkspaces) | **Get** /v1/iam/workspaces | Returns your organization&#39;s workspaces, newest first — the scope a team works in, alongside projects rather than instead of them.
[**GetOrganization**](IamAPI.md#GetOrganization) | **Get** /v1/iam/organizations/get | Returns one organization: its display, its defaults and the sign-in rules everyone in it inherits.
[**GetProvider**](IamAPI.md#GetProvider) | **Post** /v1/iam/providers/get | Returns one provider: what it connects to and how it is configured.
[**GetSession**](IamAPI.md#GetSession) | **Post** /v1/iam/sessions/get | Returns one person&#39;s session in one application — when it began and which browsers or devices are still carrying it.
[**GetToken**](IamAPI.md#GetToken) | **Post** /v1/iam/tokens/get | Returns one access token: who and what it was issued to, and when it expires.
[**GetWebauthnCredential**](IamAPI.md#GetWebauthnCredential) | **Post** /v1/iam/webauthn-credentials/get | Returns one passkey or security key: whose it is, what device it lives on, and when it was registered.
[**ListOrganizations**](IamAPI.md#ListOrganizations) | **Get** /v1/iam/organizations | Returns the organizations you can see, newest first.
[**ListProviders**](IamAPI.md#ListProviders) | **Get** /v1/iam/providers | Returns your organization&#39;s providers, newest first — the identity providers your people sign in with, and the senders and connectors your applications go through.
[**ListSessions**](IamAPI.md#ListSessions) | **Post** /v1/iam/sessions/list | Returns who is currently signed in to your organization, newest first, and can be narrowed to one person or one application.
[**ListTokens**](IamAPI.md#ListTokens) | **Get** /v1/iam/tokens | Returns the access tokens issued in your organization, newest first, and can be narrowed to one organization.
[**ListWebauthnCredentials**](IamAPI.md#ListWebauthnCredentials) | **Get** /v1/iam/webauthn-credentials | Returns the passkeys and security keys registered to one person, newest first — which device each lives on and when it was registered.
[**PatchIamScimV2UsersByOwnerByName**](IamAPI.md#PatchIamScimV2UsersByOwnerByName) | **Patch** /v1/iam/scim/v2/Users/{owner}/{name} | Applies a partial change from your identity provider — one attribute moved, not the whole record resent.
[**PostIamAddApplication**](IamAPI.md#PostIamAddApplication) | **Post** /v1/iam/add-application | Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.
[**PostIamAddMembership**](IamAPI.md#PostIamAddMembership) | **Post** /v1/iam/add-membership | Lets a person or an application act in an organization.
[**PostIamAddOrganization**](IamAPI.md#PostIamAddOrganization) | **Post** /v1/iam/add-organization | Creates an organization — the account everything else in your directory hangs from.
[**PostIamAddProject**](IamAPI.md#PostIamAddProject) | **Post** /v1/iam/add-project | Creates a project inside your organization — the scope people pick between when their work is separated by product or client rather than by team.
[**PostIamAddProvider**](IamAPI.md#PostIamAddProvider) | **Post** /v1/iam/add-provider | Adds an identity provider your people can sign in with, or a service your applications send through — a social or enterprise login, an email or SMS sender, a storage or payment connector.
[**PostIamAddRole**](IamAPI.md#PostIamAddRole) | **Post** /v1/iam/add-role | Creates a role — a named group of people that permissions are granted to.
[**PostIamAddUser**](IamAPI.md#PostIamAddUser) | **Post** /v1/iam/add-user | Adds a person to your organization and, if you send a password, sets the one they will sign in with.
[**PostIamAddWorkspace**](IamAPI.md#PostIamAddWorkspace) | **Post** /v1/iam/add-workspace | Creates a workspace inside your organization — the scope a team works in, alongside projects rather than instead of them.
[**PostIamAdminProvision**](IamAPI.md#PostIamAdminProvision) | **Post** /v1/iam/admin/provision | Sets up an account on someone&#39;s behalf — the same onboarding a person gets themselves, driven by one of your own services instead of by them.
[**PostIamApplication**](IamAPI.md#PostIamApplication) | **Post** /v1/iam/application | Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.
[**PostIamApplications**](IamAPI.md#PostIamApplications) | **Post** /v1/iam/applications | Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.
[**PostIamApplicationsDelete**](IamAPI.md#PostIamApplicationsDelete) | **Post** /v1/iam/applications/delete | Removes an application.
[**PostIamApplicationsUpdate**](IamAPI.md#PostIamApplicationsUpdate) | **Post** /v1/iam/applications/update | Changes an application&#39;s display, its sign-in methods and the redirect URIs it may return to — the call that makes login work from a new host.
[**PostIamAssume**](IamAPI.md#PostIamAssume) | **Post** /v1/iam/assume | Steps a platform operator into an organization: it returns their own access token re-scoped to that tenant, so they see what the tenant sees.
[**PostIamAuditLogs**](IamAPI.md#PostIamAuditLogs) | **Post** /v1/iam/audit-logs | Records an audit entry, so activity from your own systems lands in the same trail as everything the Hanzo Cloud records for you.
[**PostIamAuditLogsDelete**](IamAPI.md#PostIamAuditLogsDelete) | **Post** /v1/iam/audit-logs/delete | Removes an audit entry.
[**PostIamAuditLogsGet**](IamAPI.md#PostIamAuditLogsGet) | **Post** /v1/iam/audit-logs/get | Returns one audit entry in full: the action, the person or key behind it, and the request it came in on.
[**PostIamAuditLogsUpdate**](IamAPI.md#PostIamAuditLogsUpdate) | **Post** /v1/iam/audit-logs/update | Corrects an audit entry.
[**PostIamCerts**](IamAPI.md#PostIamCerts) | **Post** /v1/iam/certs | Adds a signing certificate your applications can verify tokens against — the call you make to bring your own key, or to stage the next one before a rotation.
[**PostIamCertsDelete**](IamAPI.md#PostIamCertsDelete) | **Post** /v1/iam/certs/delete | Removes a signing certificate.
[**PostIamCertsGet**](IamAPI.md#PostIamCertsGet) | **Post** /v1/iam/certs/get | Returns one signing certificate — its algorithm, its validity window and its public half.
[**PostIamCertsUpdate**](IamAPI.md#PostIamCertsUpdate) | **Post** /v1/iam/certs/update | Changes a signing certificate&#39;s settings.
[**PostIamDeleteApplication**](IamAPI.md#PostIamDeleteApplication) | **Post** /v1/iam/delete-application | Deletes an application.
[**PostIamDeleteMembership**](IamAPI.md#PostIamDeleteMembership) | **Post** /v1/iam/delete-membership | Takes away a person&#39;s or an application&#39;s right to act in an organization.
[**PostIamDeleteMfa**](IamAPI.md#PostIamDeleteMfa) | **Post** /v1/iam/delete-mfa | Turns a factor off, so sign-in stops asking for it.
[**PostIamDeleteOrganization**](IamAPI.md#PostIamDeleteOrganization) | **Post** /v1/iam/delete-organization | Deletes an organization and everything named inside it — its users, applications, roles, projects and workspaces.
[**PostIamDeleteProject**](IamAPI.md#PostIamDeleteProject) | **Post** /v1/iam/delete-project | Deletes a project.
[**PostIamDeleteProvider**](IamAPI.md#PostIamDeleteProvider) | **Post** /v1/iam/delete-provider | Removes a provider.
[**PostIamDeleteRole**](IamAPI.md#PostIamDeleteRole) | **Post** /v1/iam/delete-role | Deletes a role.
[**PostIamDeleteUser**](IamAPI.md#PostIamDeleteUser) | **Post** /v1/iam/delete-user | Removes a person from your organization.
[**PostIamDeleteWorkspace**](IamAPI.md#PostIamDeleteWorkspace) | **Post** /v1/iam/delete-workspace | Deletes a workspace.
[**PostIamInvitations**](IamAPI.md#PostIamInvitations) | **Post** /v1/iam/invitations | Issues an invitation to join your organization — the code or link a new member redeems, with the role they arrive holding and the date it stops working.
[**PostIamInvitationsDelete**](IamAPI.md#PostIamInvitationsDelete) | **Post** /v1/iam/invitations/delete | Withdraws an invitation.
[**PostIamInvitationsGet**](IamAPI.md#PostIamInvitationsGet) | **Post** /v1/iam/invitations/get | Returns one invitation: who it is for, what it grants on acceptance, and when it expires.
[**PostIamInvitationsUpdate**](IamAPI.md#PostIamInvitationsUpdate) | **Post** /v1/iam/invitations/update | Changes an invitation&#39;s terms — the role it grants, how many may redeem it, or when it expires.
[**PostIamIssueUserToken**](IamAPI.md#PostIamIssueUserToken) | **Post** /v1/iam/issue-user-token | Mints an access token for the &#x60;?id&#x3D;&lt;owner&gt;/&lt;name&gt;&#x60; target user (optional &#x60;?aud&#x3D;&#x60; resource, RFC 8707), issued by the authenticated + allow-listed confidential client.
[**PostIamKeys**](IamAPI.md#PostIamKeys) | **Post** /v1/iam/keys | Issues an API key.
[**PostIamKeysDelete**](IamAPI.md#PostIamKeysDelete) | **Post** /v1/iam/keys/delete | Revokes an API key.
[**PostIamKeysMint**](IamAPI.md#PostIamKeysMint) | **Post** /v1/iam/keys/mint | (re)generates the target user&#39;s key of the requested TYPE and returns it once, over the shared authorizeMinter + mintTarget seam.
[**PostIamKeysRevoke**](IamAPI.md#PostIamKeysRevoke) | **Post** /v1/iam/keys/revoke | Clears the target user&#39;s key of the requested TYPE (immediate revoke).
[**PostIamKeysUpdate**](IamAPI.md#PostIamKeysUpdate) | **Post** /v1/iam/keys/update | Changes what a key is called or what it may reach.
[**PostIamLink**](IamAPI.md#PostIamLink) | **Post** /v1/iam/link | Starts connecting another sign-in identity to the account you are already signed in as.
[**PostIamLogin**](IamAPI.md#PostIamLogin) | **Post** /v1/iam/login | Signs a person in with the credential they typed, and — when the request is part of an OAuth flow — hands back the one-time code that finishes it.
[**PostIamMemberships**](IamAPI.md#PostIamMemberships) | **Post** /v1/iam/memberships | Lets a person or an application act in an organization.
[**PostIamMfaDisable**](IamAPI.md#PostIamMfaDisable) | **Post** /v1/iam/mfa/disable | Turns a factor off, so sign-in stops asking for it.
[**PostIamMfaPreferred**](IamAPI.md#PostIamMfaPreferred) | **Post** /v1/iam/mfa/preferred | Picks which second factor an account is asked for first when it has more than one.
[**PostIamMfaSetupEnable**](IamAPI.md#PostIamMfaSetupEnable) | **Post** /v1/iam/mfa/setup/enable | Finishes the enrolment: from here the account&#39;s sign-ins ask for this factor.
[**PostIamMfaSetupInitiate**](IamAPI.md#PostIamMfaSetupInitiate) | **Post** /v1/iam/mfa/setup/initiate | Starts enrolling a factor and hands over whatever the person needs to prove they hold it: app a fresh secret and the otpauth:// URL to render as a QR code sms a code texted to the number on the account email a code mailed to the address on the account Nothing is switched on yet, so abandoning this step leaves the account exactly as it was.
[**PostIamMintUserKeys**](IamAPI.md#PostIamMintUserKeys) | **Post** /v1/iam/mint-user-keys | (re)generates the target user&#39;s key of the requested TYPE and returns it once, over the shared authorizeMinter + mintTarget seam.
[**PostIamOauthAuthorize**](IamAPI.md#PostIamOauthAuthorize) | **Post** /v1/iam/oauth/authorize | Starts a sign-in — the address you send a browser to, and the beginning of every OAuth and OpenID Connect flow.
[**PostIamOauthDevice**](IamAPI.md#PostIamOauthDevice) | **Post** /v1/iam/oauth/device | Starts a sign-in on a device with no browser and no keyboard — a TV, a CLI, a headless box.
[**PostIamOauthDeviceInfo**](IamAPI.md#PostIamOauthDeviceInfo) | **Post** /v1/iam/oauth/device/info | Answers \&quot;what am I approving?\&quot; for a pending device code.
[**PostIamOauthFederationMfa**](IamAPI.md#PostIamOauthFederationMfa) | **Post** /v1/iam/oauth/federation/mfa | Completes a sign-in that came in through another identity provider and still owes a second factor.
[**PostIamOauthIntrospect**](IamAPI.md#PostIamOauthIntrospect) | **Post** /v1/iam/oauth/introspect | Answers whether an access token is still good, and what it is good for — the check a resource server of yours makes before honouring a token it did not mint.
[**PostIamOauthLogout**](IamAPI.md#PostIamOauthLogout) | **Post** /v1/iam/oauth/logout | Ends a sign-in and sends the browser somewhere sensible.
[**PostIamOauthRevoke**](IamAPI.md#PostIamOauthRevoke) | **Post** /v1/iam/oauth/revoke | Retires a token before it expires — what you call when someone signs out or a credential may have leaked.
[**PostIamOauthToken**](IamAPI.md#PostIamOauthToken) | **Post** /v1/iam/oauth/token | Exchanges what your application is holding for the tokens it needs — the one-time code from a finished sign-in, a refresh token, or your own client credentials when the caller is a program rather than a person.
[**PostIamOauthUserinfo**](IamAPI.md#PostIamOauthUserinfo) | **Post** /v1/iam/oauth/userinfo | Returns the profile claims for whoever the access token belongs to — the standard OpenID Connect way to find out who is calling you without your application storing anything itself.
[**PostIamOnboard**](IamAPI.md#PostIamOnboard) | **Post** /v1/iam/onboard | Finishes setting up the account of whoever is calling — it creates their organization if they have none and puts them in it, so a person who has just signed up lands somewhere they can work.
[**PostIamPermissions**](IamAPI.md#PostIamPermissions) | **Post** /v1/iam/permissions | Grants a permission — the call that gives a person or a role the ability to do something.
[**PostIamPermissionsDelete**](IamAPI.md#PostIamPermissionsDelete) | **Post** /v1/iam/permissions/delete | Revokes a permission.
[**PostIamPermissionsUpdate**](IamAPI.md#PostIamPermissionsUpdate) | **Post** /v1/iam/permissions/update | Changes who a permission grants to, what it allows, or the resources it covers.
[**PostIamPreferences**](IamAPI.md#PostIamPreferences) | **Post** /v1/iam/preferences | Saves the calling person&#39;s own settings and returns the full set afterwards.
[**PostIamProjects**](IamAPI.md#PostIamProjects) | **Post** /v1/iam/projects | Makes a project inside your organization — the scope people pick between when their work is separated by product or client rather than by team.
[**PostIamProjectsDelete**](IamAPI.md#PostIamProjectsDelete) | **Post** /v1/iam/projects/delete | Removes a project.
[**PostIamProjectsGet**](IamAPI.md#PostIamProjectsGet) | **Post** /v1/iam/projects/get | Returns one project: what it is called and how it is set up.
[**PostIamProjectsUpdate**](IamAPI.md#PostIamProjectsUpdate) | **Post** /v1/iam/projects/update | Changes a project&#39;s settings.
[**PostIamRegistryToken**](IamAPI.md#PostIamRegistryToken) | **Post** /v1/iam/registry/token | Signs a container client in to your registry.
[**PostIamRelease**](IamAPI.md#PostIamRelease) | **Post** /v1/iam/release | Steps a platform operator back out: it returns their own access token with no organization assumed, which is the credential they had before they stepped in.
[**PostIamRevokeUserKeys**](IamAPI.md#PostIamRevokeUserKeys) | **Post** /v1/iam/revoke-user-keys | Clears the target user&#39;s key of the requested TYPE (immediate revoke).
[**PostIamRoles**](IamAPI.md#PostIamRoles) | **Post** /v1/iam/roles | Makes a role — a named group of people that permissions are granted to.
[**PostIamRolesDelete**](IamAPI.md#PostIamRolesDelete) | **Post** /v1/iam/roles/delete | Removes a role.
[**PostIamRolesGet**](IamAPI.md#PostIamRolesGet) | **Post** /v1/iam/roles/get | Returns one role: who is in it, and the roles it includes.
[**PostIamRolesUpdate**](IamAPI.md#PostIamRolesUpdate) | **Post** /v1/iam/roles/update | Changes who is in a role, or which roles it includes.
[**PostIamScimV2Users**](IamAPI.md#PostIamScimV2Users) | **Post** /v1/iam/scim/v2/Users | Provisions a person from your identity provider — how a new hire gets an account here automatically when they are added over there.
[**PostIamSendVerificationCode**](IamAPI.md#PostIamSendVerificationCode) | **Post** /v1/iam/send-verification-code | Validates the request and asks otp to get a code to the person.
[**PostIamServiceAccounts**](IamAPI.md#PostIamServiceAccounts) | **Post** /v1/iam/service-accounts | Makes a service account — an identity for a program rather than a person, for a script, a bot or a deployment that has to authenticate on its own.
[**PostIamServiceAccountsByNameKeys**](IamAPI.md#PostIamServiceAccountsByNameKeys) | **Post** /v1/iam/service-accounts/{name}/keys | Serves POST /v1/iam/service-accounts/:name/keys: mint a fresh key, invalidating the prior one, and return the new raw secret exactly once.
[**PostIamSetPreferredMfa**](IamAPI.md#PostIamSetPreferredMfa) | **Post** /v1/iam/set-preferred-mfa | Picks which second factor an account is asked for first when it has more than one.
[**PostIamSignin**](IamAPI.md#PostIamSignin) | **Post** /v1/iam/signin | Completes a sign-in: it exchanges the one-time code your application was handed at the end of the login flow for a live session, and returns the signed-in account.
[**PostIamSignup**](IamAPI.md#PostIamSignup) | **Post** /v1/iam/signup | Creates an account from the sign-up form and applies the application&#39;s own sign-up rules — whether self-service registration is open at all, and which fields it requires.
[**PostIamTokensIssue**](IamAPI.md#PostIamTokensIssue) | **Post** /v1/iam/tokens/issue | Mints an access token for the &#x60;?id&#x3D;&lt;owner&gt;/&lt;name&gt;&#x60; target user (optional &#x60;?aud&#x3D;&#x60; resource, RFC 8707), issued by the authenticated + allow-listed confidential client.
[**PostIamUnlink**](IamAPI.md#PostIamUnlink) | **Post** /v1/iam/unlink | Disconnects one sign-in identity from an account, so that provider can no longer be used to sign in as that person.
[**PostIamUpdateApplication**](IamAPI.md#PostIamUpdateApplication) | **Post** /v1/iam/update-application | Updates one of your applications — its display, its sign-in methods and the redirect URIs it is allowed to return to.
[**PostIamUpdateOrganization**](IamAPI.md#PostIamUpdateOrganization) | **Post** /v1/iam/update-organization | Updates your organization — its display, its default settings and the sign-in rules everyone in it inherits.
[**PostIamUpdatePreferences**](IamAPI.md#PostIamUpdatePreferences) | **Post** /v1/iam/update-preferences | Saves the calling person&#39;s own settings and returns the full set afterwards.
[**PostIamUpdateProvider**](IamAPI.md#PostIamUpdateProvider) | **Post** /v1/iam/update-provider | Updates a provider&#39;s settings or rotates the credentials it holds.
[**PostIamUpdateRole**](IamAPI.md#PostIamUpdateRole) | **Post** /v1/iam/update-role | Updates a role&#39;s members or the roles it includes.
[**PostIamUpdateUser**](IamAPI.md#PostIamUpdateUser) | **Post** /v1/iam/update-user | Updates one of your users&#39; profile, roles or credentials.
[**PostIamUsers**](IamAPI.md#PostIamUsers) | **Post** /v1/iam/users | Adds a person to your organization.
[**PostIamUsersDelete**](IamAPI.md#PostIamUsersDelete) | **Post** /v1/iam/users/delete | Removes a person from your organization.
[**PostIamUsersUpdate**](IamAPI.md#PostIamUsersUpdate) | **Post** /v1/iam/users/update | Changes a person&#39;s profile, their roles, or the credentials they sign in with.
[**PostIamVerificationCodes**](IamAPI.md#PostIamVerificationCodes) | **Post** /v1/iam/verification-codes | Validates the request and asks otp to get a code to the person.
[**PostIamWeb3Verify**](IamAPI.md#PostIamWeb3Verify) | **Post** /v1/iam/web3/verify | Completes a wallet sign-in: it verifies the signed challenge and, if it holds, signs the wallet&#39;s owner in.
[**PostIamWebauthnSigninFinish**](IamAPI.md#PostIamWebauthnSigninFinish) | **Post** /v1/iam/webauthn/signin/finish | Verifies the signed challenge and signs the person in.
[**PostIamWebauthnSignupFinish**](IamAPI.md#PostIamWebauthnSignupFinish) | **Post** /v1/iam/webauthn/signup/finish | Verifies the newly created passkey and stores it, so the person can sign in with their device from then on.
[**PostIamWorkspaces**](IamAPI.md#PostIamWorkspaces) | **Post** /v1/iam/workspaces | Makes a workspace inside your organization — the scope a team works in, alongside projects rather than instead of them.
[**PostIamWorkspacesDelete**](IamAPI.md#PostIamWorkspacesDelete) | **Post** /v1/iam/workspaces/delete | Removes a workspace.
[**PostIamWorkspacesGet**](IamAPI.md#PostIamWorkspacesGet) | **Post** /v1/iam/workspaces/get | Returns one workspace: what it is called and how it is set up.
[**PostIamWorkspacesUpdate**](IamAPI.md#PostIamWorkspacesUpdate) | **Post** /v1/iam/workspaces/update | Changes a workspace&#39;s settings.
[**PutIamAccount**](IamAPI.md#PutIamAccount) | **Put** /v1/iam/account | Saves the calling person&#39;s own profile — the name they are shown by, their picture, a line about themselves and a link.
[**PutIamApplication**](IamAPI.md#PutIamApplication) | **Put** /v1/iam/application | Changes an application&#39;s display, its sign-in methods and the redirect URIs it may return to — the call that makes login work from a new host.
[**PutIamConsent**](IamAPI.md#PutIamConsent) | **Put** /v1/iam/consent | Records the calling person&#39;s privacy and communication choices.
[**PutIamPassword**](IamAPI.md#PutIamPassword) | **Put** /v1/iam/password | Replaces the calling person&#39;s password.
[**PutIamScimV2UsersByOwnerByName**](IamAPI.md#PutIamScimV2UsersByOwnerByName) | **Put** /v1/iam/scim/v2/Users/{owner}/{name} | Overwrites a person&#39;s SCIM attributes with what your identity provider sends — how a change made there lands here.
[**SearchOrganizations**](IamAPI.md#SearchOrganizations) | **Get** /v1/iam/organizations/search | Returns the organizations you can act in, the ones you belong to first and the rest after, newest first, narrowed by an optional query against the name or the display name.
[**SetOrganizationAvatar**](IamAPI.md#SetOrganizationAvatar) | **Post** /v1/iam/organizations/avatar | Changes how an organization appears across Hanzo: the square mark beside its name, as an uploaded image or as a single emoji.
[**UpdateOrganization**](IamAPI.md#UpdateOrganization) | **Post** /v1/iam/organizations/update | Changes an organization&#39;s display, its defaults and the sign-in rules everyone in it inherits.
[**UpdateProvider**](IamAPI.md#UpdateProvider) | **Post** /v1/iam/providers/update | Changes a provider&#39;s settings or rotates the credentials it holds.
[**UpdateSession**](IamAPI.md#UpdateSession) | **Post** /v1/iam/sessions/update | Replaces the set of browsers a session covers — signing out the ones you leave off while the session itself stays live.
[**UpdateToken**](IamAPI.md#UpdateToken) | **Post** /v1/iam/tokens/update | Changes an access token&#39;s scope or expiry.
[**UpdateWebauthnCredential**](IamAPI.md#UpdateWebauthnCredential) | **Post** /v1/iam/webauthn-credentials/update | Renames a registered passkey or security key, so a person can tell their devices apart.
[**UpsertApplication**](IamAPI.md#UpsertApplication) | **Post** /v1/iam/admin/applications/upsert | Creates an application or updates it in place, so a deployment can declare the applications it needs and run the same declaration on every environment and on every redeploy.
[**UpsertUser**](IamAPI.md#UpsertUser) | **Post** /v1/iam/admin/users/upsert | Creates a person or updates them in place, so a deployment can declare the accounts it needs and re-run that declaration safely.



## AddProvider

> IamProviderResult AddProvider(ctx).IamProvider(iamProvider).Execute()

Adds an identity provider your people can sign in with, or a service your applications send through — a social or enterprise login, an email or SMS sender, a storage or payment connector.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamProvider := *openapiclient.NewIamProvider() // IamProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.AddProvider(context.Background()).IamProvider(iamProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.AddProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProvider`: IamProviderResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.AddProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProvider** | [**IamProvider**](IamProvider.md) |  | 

### Return type

[**IamProviderResult**](IamProviderResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddToken

> IamTokenResult AddToken(ctx).IamToken(iamToken).Execute()

Records an access token — the credential an application or integration presents on a caller's behalf.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamToken := *openapiclient.NewIamToken() // IamToken | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.AddToken(context.Background()).IamToken(iamToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.AddToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddToken`: IamTokenResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.AddToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamToken** | [**IamToken**](IamToken.md) |  | 

### Return type

[**IamTokenResult**](IamTokenResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWebauthnCredential

> IamWebauthnCredentialResult AddWebauthnCredential(ctx).IamWebauthnCredential(iamWebauthnCredential).Execute()

Registers a passkey or security key for a person, so they can sign in with their device instead of a password.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamWebauthnCredential := *openapiclient.NewIamWebauthnCredential() // IamWebauthnCredential | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.AddWebauthnCredential(context.Background()).IamWebauthnCredential(iamWebauthnCredential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.AddWebauthnCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWebauthnCredential`: IamWebauthnCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.AddWebauthnCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddWebauthnCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWebauthnCredential** | [**IamWebauthnCredential**](IamWebauthnCredential.md) |  | 

### Return type

[**IamWebauthnCredentialResult**](IamWebauthnCredentialResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> IamOrganization CreateOrganization(ctx).IamCreateOrganizationInput(iamCreateOrganizationInput).Execute()

Makes a new organization — the account your users, applications, roles, projects and workspaces are all named inside.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamCreateOrganizationInput := *openapiclient.NewIamCreateOrganizationInput() // IamCreateOrganizationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.CreateOrganization(context.Background()).IamCreateOrganizationInput(iamCreateOrganizationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.CreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganization`: IamOrganization
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamCreateOrganizationInput** | [**IamCreateOrganizationInput**](IamCreateOrganizationInput.md) |  | 

### Return type

[**IamOrganization**](IamOrganization.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> IamSession CreateSession(ctx).IamCreateSessionIn(iamCreateSessionIn).Execute()

Records a sign-in.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamCreateSessionIn := *openapiclient.NewIamCreateSessionIn("Application_example", "Name_example", "Owner_example") // IamCreateSessionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.CreateSession(context.Background()).IamCreateSessionIn(iamCreateSessionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.CreateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSession`: IamSession
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.CreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamCreateSessionIn** | [**IamCreateSessionIn**](IamCreateSessionIn.md) |  | 

### Return type

[**IamSession**](IamSession.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamApplication

> IamDeleteResult DeleteIamApplication(ctx).Owner(owner).Name(name).Execute()

Removes an application.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamApplication(context.Background()).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamApplication`: IamDeleteResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**IamDeleteResult**](IamDeleteResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamScimV2UsersByOwnerByName

> DeleteIamScimV2UsersByOwnerByName(ctx, owner, name).Execute()

Deprovisions a person — how removing someone in your identity provider removes their access here.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.DeleteIamScimV2UsersByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamScimV2UsersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamScimV2UsersByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamServiceAccountsByName

> DeleteIamServiceAccountsByName(ctx, name).Execute()

Serves DELETE /v1/iam/service-accounts/:name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.DeleteIamServiceAccountsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamServiceAccountsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamServiceAccountsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> IamDeleteOrganizationOutput DeleteOrganization(ctx).IamDeleteOrganizationInput(iamDeleteOrganizationInput).Execute()

Removes an organization and everything named inside it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamDeleteOrganizationInput := *openapiclient.NewIamDeleteOrganizationInput() // IamDeleteOrganizationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteOrganization(context.Background()).IamDeleteOrganizationInput(iamDeleteOrganizationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganization`: IamDeleteOrganizationOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamDeleteOrganizationInput** | [**IamDeleteOrganizationInput**](IamDeleteOrganizationInput.md) |  | 

### Return type

[**IamDeleteOrganizationOutput**](IamDeleteOrganizationOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProvider

> IamMutationResult DeleteProvider(ctx).IamProviderKey(iamProviderKey).Execute()

Removes a provider.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamProviderKey := *openapiclient.NewIamProviderKey("Name_example", "Owner_example") // IamProviderKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteProvider(context.Background()).IamProviderKey(iamProviderKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProvider`: IamMutationResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProviderKey** | [**IamProviderKey**](IamProviderKey.md) |  | 

### Return type

[**IamMutationResult**](IamMutationResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSession

> IamDeleteSessionOut DeleteSession(ctx).IamSessionRef(iamSessionRef).Execute()

Signs a person out of one application — the session ends and every browser carrying it stops being authenticated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamSessionRef := *openapiclient.NewIamSessionRef("Application_example", "Name_example", "Owner_example") // IamSessionRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteSession(context.Background()).IamSessionRef(iamSessionRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSession`: IamDeleteSessionOut
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamSessionRef** | [**IamSessionRef**](IamSessionRef.md) |  | 

### Return type

[**IamDeleteSessionOut**](IamDeleteSessionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> IamTokenMutation DeleteToken(ctx).IamTokenKey(iamTokenKey).Execute()

Revokes an access token.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamTokenKey := *openapiclient.NewIamTokenKey("Name_example", "Owner_example") // IamTokenKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteToken(context.Background()).IamTokenKey(iamTokenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteToken`: IamTokenMutation
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamTokenKey** | [**IamTokenKey**](IamTokenKey.md) |  | 

### Return type

[**IamTokenMutation**](IamTokenMutation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebauthnCredential

> IamWebauthnCredentialMutationResult DeleteWebauthnCredential(ctx).IamWebauthnCredentialKey(iamWebauthnCredentialKey).Execute()

Removes a passkey or security key — what you call when a device is lost.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamWebauthnCredentialKey := *openapiclient.NewIamWebauthnCredentialKey("Name_example", "Owner_example") // IamWebauthnCredentialKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteWebauthnCredential(context.Background()).IamWebauthnCredentialKey(iamWebauthnCredentialKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteWebauthnCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebauthnCredential`: IamWebauthnCredentialMutationResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteWebauthnCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebauthnCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWebauthnCredentialKey** | [**IamWebauthnCredentialKey**](IamWebauthnCredentialKey.md) |  | 

### Return type

[**IamWebauthnCredentialMutationResult**](IamWebauthnCredentialMutationResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamAccount

> GetIamAccount(ctx).Execute()

Returns the signed-in person's own account and the organization they belong to — what a console reads to draw the account menu.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamAccountRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamApplication

> IamApplication GetIamApplication(ctx).Owner(owner).Name(name).Execute()

Returns one application: its sign-in methods, its allowed redirect URIs and the client credentials your integration authenticates with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamApplication(context.Background()).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamApplication`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**IamApplication**](IamApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamApplications

> IamApplicationListResult GetIamApplications(ctx).Owner(owner).Execute()

Returns the applications in one organization, newest first — each product or site your people sign in to, with the sign-in methods and redirect URIs it allows.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamApplications(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamApplications`: IamApplicationListResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

### Return type

[**IamApplicationListResult**](IamApplicationListResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamApplicationsGet

> IamApplication GetIamApplicationsGet(ctx).Owner(owner).Name(name).Execute()

Returns one application: its sign-in methods, its allowed redirect URIs and the client credentials your integration authenticates with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamApplicationsGet(context.Background()).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamApplicationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamApplicationsGet`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamApplicationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamApplicationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**IamApplication**](IamApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamAuditLogs

> IamListOutput GetIamAuditLogs(ctx).Owner(owner).Execute()

Returns your organization's audit trail, newest first — who did what, when, and from where.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamAuditLogs(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamAuditLogs`: IamListOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

### Return type

[**IamListOutput**](IamListOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamAuthApplication

> IamAnswer GetIamAuthApplication(ctx).ClientId(clientId).ResponseType(responseType).Execute()

Returns everything a login screen needs to draw itself for one application: its branding, and each sign-in method it offers with the provider details that method needs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	clientId := "clientId_example" // string | ClientId is the application's OAuth client id — the one field that selects which login screen this is. (optional)
	responseType := "responseType_example" // string | ResponseType is the OAuth response type the screen will ask for. Only \"code\" is served; anything else is refused here rather than at the authorize leg, where the person has already typed a password. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamAuthApplication(context.Background()).ClientId(clientId).ResponseType(responseType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamAuthApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamAuthApplication`: IamAnswer
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamAuthApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamAuthApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | ClientId is the application&#39;s OAuth client id — the one field that selects which login screen this is. | 
 **responseType** | **string** | ResponseType is the OAuth response type the screen will ask for. Only \&quot;code\&quot; is served; anything else is refused here rather than at the authorize leg, where the person has already typed a password. | 

### Return type

[**IamAnswer**](IamAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamAuthMethods

> IamAnswer GetIamAuthMethods(ctx).ClientId(clientId).Execute()

Returns the sign-in methods one application actually has switched on, so a login screen can render the right buttons for it without you hard-coding a list that drifts the moment you add a provider.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	clientId := "clientId_example" // string | ClientId is the application's OAuth client id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamAuthMethods(context.Background()).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamAuthMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamAuthMethods`: IamAnswer
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamAuthMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamAuthMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | ClientId is the application&#39;s OAuth client id. | 

### Return type

[**IamAnswer**](IamAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamCerts

> IamCertsListOutput GetIamCerts(ctx).Owner(owner).Execute()

Returns your organization's signing certificates, newest first — the keys the tokens your applications verify are signed with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamCerts(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamCerts`: IamCertsListOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

### Return type

[**IamCertsListOutput**](IamCertsListOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamConsent

> GetIamConsent(ctx).Execute()

Returns the calling person's own privacy and communication choices.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamConsent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamConsent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamConsentRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetAccount

> GetIamGetAccount(ctx).Execute()

Returns the signed-in person's own account and the organization they belong to — what a console reads to draw the account menu.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetAccountRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetAppLogin

> IamAnswer GetIamGetAppLogin(ctx).ClientId(clientId).ResponseType(responseType).Execute()

Returns everything a login screen needs to draw itself for one application: its branding, and each sign-in method it offers with the provider details that method needs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	clientId := "clientId_example" // string | ClientId is the application's OAuth client id — the one field that selects which login screen this is. (optional)
	responseType := "responseType_example" // string | ResponseType is the OAuth response type the screen will ask for. Only \"code\" is served; anything else is refused here rather than at the authorize leg, where the person has already typed a password. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamGetAppLogin(context.Background()).ClientId(clientId).ResponseType(responseType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetAppLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamGetAppLogin`: IamAnswer
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamGetAppLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetAppLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | ClientId is the application&#39;s OAuth client id — the one field that selects which login screen this is. | 
 **responseType** | **string** | ResponseType is the OAuth response type the screen will ask for. Only \&quot;code\&quot; is served; anything else is refused here rather than at the authorize leg, where the person has already typed a password. | 

### Return type

[**IamAnswer**](IamAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetApplication

> GetIamGetApplication(ctx).Execute()

Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetApplication(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetApplicationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetApplications

> GetIamGetApplications(ctx).Execute()

Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetApplicationsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetCert

> GetIamGetCert(ctx).Execute()

Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetCert(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetCertRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetCerts

> GetIamGetCerts(ctx).Execute()

Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetCerts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetCertsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetGlobalUsers

> GetIamGetGlobalUsers(ctx).Execute()

Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetGlobalUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetGlobalUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetGlobalUsersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetInvitations

> GetIamGetInvitations(ctx).Execute()

Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetInvitationsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetMemberships

> IamAnswer GetIamGetMemberships(ctx).User(user).Org(org).Execute()

Answers either question about who belongs where: which organizations one person can act in, or who can act in one organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	user := "user_example" // string | User is \"<homeOrg>/<username>\" — which organizations that identity may act in. (optional)
	org := "org_example" // string | Org is an organization — who may act in it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamGetMemberships(context.Background()).User(user).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamGetMemberships`: IamAnswer
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamGetMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | User is \&quot;&lt;homeOrg&gt;/&lt;username&gt;\&quot; — which organizations that identity may act in. | 
 **org** | **string** | Org is an organization — who may act in it. | 

### Return type

[**IamAnswer**](IamAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetOrganization

> GetIamGetOrganization(ctx).Execute()

Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetOrganization(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetOrganizationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetOrganizationProjects

> GetIamGetOrganizationProjects(ctx).Execute()

Returns one organization's projects — what a scope switcher lists so somebody can move between them.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetOrganizationProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetOrganizationProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetOrganizationProjectsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetOrganizationWorkspaces

> GetIamGetOrganizationWorkspaces(ctx).Execute()

Returns one organization's workspaces — what a scope switcher lists so somebody can move between them.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetOrganizationWorkspaces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetOrganizationWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetOrganizationWorkspacesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetOrganizations

> GetIamGetOrganizations(ctx).Execute()

Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetOrganizations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetOrganizationsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetPermission

> GetIamGetPermission(ctx).Execute()

Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetPermission(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetPermissionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetPermissions

> GetIamGetPermissions(ctx).Execute()

Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetPermissions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetPermissionsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetProvider

> GetIamGetProvider(ctx).Execute()

Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetProvider(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetProviderRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetProviders

> GetIamGetProviders(ctx).Execute()

Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetProvidersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetRecords

> GetIamGetRecords(ctx).Execute()

Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetRecords(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetRecordsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetRole

> GetIamGetRole(ctx).Execute()

Reads one record — the older spelling of the single reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetRole(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetRoleRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetRoles

> GetIamGetRoles(ctx).Execute()

Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetRolesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetUser

> GetIamGetUser(ctx).Execute()

Reads one person, two ways.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetUserRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamGetUsers

> GetIamGetUsers(ctx).Execute()

Lists one kind of record in your organization — the older spelling of the collection reads on the REST surface, over the same data and the same permissions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamGetUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamGetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamGetUsersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamInvitations

> IamInvitationsListOutput GetIamInvitations(ctx).Owner(owner).Execute()

Returns your organization's invitations, newest first — who has been asked to join, on what terms, and how many seats each invitation still has left.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamInvitations(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamInvitations`: IamInvitationsListOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

### Return type

[**IamInvitationsListOutput**](IamInvitationsListOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamKeys

> IamListResponse GetIamKeys(ctx).Owner(owner).Execute()

Returns your organization's API keys, newest first — what each is called, what it may reach, and its publishable half.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamKeys(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamKeys`: IamListResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

### Return type

[**IamListResponse**](IamListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamKeysGet

> IamKey GetIamKeysGet(ctx).Owner(owner).Name(name).Execute()

Returns one API key: what it is called, what it may reach, and when it was issued.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)
	name := "name_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamKeysGet(context.Background()).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamKeysGet`: IamKey
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamKeysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**IamKey**](IamKey.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamLinkedAccounts

> GetIamLinkedAccounts(ctx).Execute()

Returns the sign-in identities linked to the calling person's account — every provider they can currently sign in with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamLinkedAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamLinkedAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamLinkedAccountsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamMemberships

> IamAnswer GetIamMemberships(ctx).User(user).Org(org).Execute()

Answers either question about who belongs where: which organizations one person can act in, or who can act in one organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	user := "user_example" // string | User is \"<homeOrg>/<username>\" — which organizations that identity may act in. (optional)
	org := "org_example" // string | Org is an organization — who may act in it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamMemberships(context.Background()).User(user).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamMemberships`: IamAnswer
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | User is \&quot;&lt;homeOrg&gt;/&lt;username&gt;\&quot; — which organizations that identity may act in. | 
 **org** | **string** | Org is an organization — who may act in it. | 

### Return type

[**IamAnswer**](IamAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamOauthAuthorize

> GetIamOauthAuthorize(ctx).Execute()

Starts a sign-in — the address you send a browser to, and the beginning of every OAuth and OpenID Connect flow.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamOauthAuthorize(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamOauthAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamOauthAuthorizeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamOauthCallback

> GetIamOauthCallback(ctx).Execute()

Completes the round-trip: it resolves and burns the single-use transaction (checking expiry + browser binding), exchanges and verifies the IdP response, links or provisions the local user, and mints the iam authorization code the relying party expects — then redirects to the original redirect_uri with code + state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamOauthCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamOauthCallbackRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamOauthLogout

> GetIamOauthLogout(ctx).Execute()

Ends a sign-in and sends the browser somewhere sensible.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamOauthLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamOauthLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamOauthLogoutRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamOauthUserinfo

> GetIamOauthUserinfo(ctx).Execute()

Returns the profile claims for whoever the access token belongs to — the standard OpenID Connect way to find out who is calling you without your application storing anything itself.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamOauthUserinfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamOauthUserinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamOauthUserinfoRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamPermissions

> IamPermissionListResponse GetIamPermissions(ctx).Owner(owner).Execute()

Returns the permissions in one organization, newest first — each one a grant saying which people or roles may do what, and to which resources.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamPermissions(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamPermissions`: IamPermissionListResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

### Return type

[**IamPermissionListResponse**](IamPermissionListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamPermissionsGet

> IamPermission GetIamPermissionsGet(ctx).Owner(owner).Name(name).Execute()

Returns one permission: who it grants to, what it allows, and the resources it covers.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)
	name := "name_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamPermissionsGet(context.Background()).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamPermissionsGet`: IamPermission
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamPermissionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**IamPermission**](IamPermission.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamProjects

> IamProjectsListOutput GetIamProjects(ctx).Owner(owner).Execute()

Returns your organization's projects, newest first — the scope people pick between when their work is separated by product or client rather than by team.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamProjects(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamProjects`: IamProjectsListOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

### Return type

[**IamProjectsListOutput**](IamProjectsListOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamRegistryJwks

> GetIamRegistryJwks(ctx).Execute()

Publishes the public key your registry uses to verify the tokens issued above — the one URL to configure so the registry trusts logins without holding any secret of its own.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamRegistryJwks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamRegistryJwks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamRegistryJwksRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamRegistryToken

> GetIamRegistryToken(ctx).Execute()

Signs a container client in to your registry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamRegistryToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamRegistryToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamRegistryTokenRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamResolveKey

> GetIamResolveKey(ctx).Execute()

Answers which organization a PUBLISHABLE key belongs to — what a service of yours calls to attribute a request that arrived carrying a key shipped in a browser.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamResolveKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamResolveKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamResolveKeyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamRoles

> IamRolesListOutput GetIamRoles(ctx).Owner(owner).Execute()

Returns your organization's roles, newest first — each a named group of people that permissions are granted to.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamRoles(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamRoles`: IamRolesListOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

### Return type

[**IamRolesListOutput**](IamRolesListOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamScimV2Resourcetypes

> IamListResponse GetIamScimV2Resourcetypes(ctx).Execute()

Returns the kinds of record this directory provisions and the address of each, so your identity provider discovers them rather than having them configured by hand.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamScimV2Resourcetypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamScimV2Resourcetypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamScimV2Resourcetypes`: IamListResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamScimV2Resourcetypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamScimV2ResourcetypesRequest struct via the builder pattern


### Return type

[**IamListResponse**](IamListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamScimV2ResourcetypesByName

> interface{} GetIamScimV2ResourcetypesByName(ctx, name).Execute()

Returns one provisionable record kind in full.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamScimV2ResourcetypesByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamScimV2ResourcetypesByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamScimV2ResourcetypesByName`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamScimV2ResourcetypesByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamScimV2ResourcetypesByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamScimV2Schemas

> IamListResponse GetIamScimV2Schemas(ctx).Execute()

Returns the attribute definitions this directory understands, so your identity provider knows which fields it may send and what they mean before it sends any.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamScimV2Schemas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamScimV2Schemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamScimV2Schemas`: IamListResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamScimV2Schemas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamScimV2SchemasRequest struct via the builder pattern


### Return type

[**IamListResponse**](IamListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamScimV2SchemasById

> interface{} GetIamScimV2SchemasById(ctx, id).Execute()

Returns one attribute definition in full.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamScimV2SchemasById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamScimV2SchemasById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamScimV2SchemasById`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamScimV2SchemasById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamScimV2SchemasByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamScimV2Serviceproviderconfig

> IamConfig GetIamScimV2Serviceproviderconfig(ctx).Execute()

Tells your identity provider which parts of SCIM this directory supports, so it configures itself instead of you filling in a form.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamScimV2Serviceproviderconfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamScimV2Serviceproviderconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamScimV2Serviceproviderconfig`: IamConfig
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamScimV2Serviceproviderconfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamScimV2ServiceproviderconfigRequest struct via the builder pattern


### Return type

[**IamConfig**](IamConfig.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamScimV2Users

> GetIamScimV2Users(ctx).Execute()

Returns the people in your organization to your identity provider, in the standard SCIM shape, so an IdP can reconcile its directory against ours.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamScimV2Users(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamScimV2Users``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamScimV2UsersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamScimV2UsersByOwnerByName

> GetIamScimV2UsersByOwnerByName(ctx, owner, name).Execute()

Returns one person in the standard SCIM shape.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamScimV2UsersByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamScimV2UsersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamScimV2UsersByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamServiceAccounts

> IamAnswer GetIamServiceAccounts(ctx).Organization(organization).P(p).PageSize(pageSize).Execute()

Returns your organization's service accounts — what each is called and when it was created.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	organization := "organization_example" // string | Organization is the organization whose service accounts to list. Required. (optional)
	p := int32(56) // int32 | P is the 1-indexed page to return. Paging takes both p and pageSize — leave either out, or send something that is not a number, and the whole list comes back. (optional)
	pageSize := int32(56) // int32 | Size is how many accounts a page holds. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamServiceAccounts(context.Background()).Organization(organization).P(p).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamServiceAccounts`: IamAnswer
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamServiceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | **string** | Organization is the organization whose service accounts to list. Required. | 
 **p** | **int32** | P is the 1-indexed page to return. Paging takes both p and pageSize — leave either out, or send something that is not a number, and the whole list comes back. | 
 **pageSize** | **int32** | Size is how many accounts a page holds. | 

### Return type

[**IamAnswer**](IamAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamUsers

> IamUsersListOutput GetIamUsers(ctx).Owner(owner).Limit(limit).Offset(offset).Execute()

Returns a page of the people in your organization, with the total so you can page through the rest.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | 
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamUsers(context.Background()).Owner(owner).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamUsers`: IamUsersListOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**IamUsersListOutput**](IamUsersListOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamUsersGet

> IamUser GetIamUsersGet(ctx).Owner(owner).Name(name).Email(email).Execute()

Returns one person in your organization, addressed by their username or by their email address.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string |  (optional)
	email := "email_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamUsersGet(context.Background()).Owner(owner).Name(name).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamUsersGet`: IamUser
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **name** | **string** |  | 
 **email** | **string** |  | 

### Return type

[**IamUser**](IamUser.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamWeb3Nonce

> GetIamWeb3Nonce(ctx).Execute()

Starts a wallet sign-in: it returns a one-time challenge for the wallet to sign.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamWeb3Nonce(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamWeb3Nonce``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamWeb3NonceRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamWebauthnSigninBegin

> GetIamWebauthnSigninBegin(ctx).Execute()

Starts a passkey sign-in: it returns the challenge the person's authenticator signs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamWebauthnSigninBegin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamWebauthnSigninBegin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamWebauthnSigninBeginRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamWebauthnSignupBegin

> GetIamWebauthnSignupBegin(ctx).Execute()

Starts enrolling a passkey for the signed-in person: it returns the options their browser hands to the authenticator.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamWebauthnSignupBegin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamWebauthnSignupBegin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamWebauthnSignupBeginRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamWellKnownJwks

> GetIamWellKnownJwks(ctx).Execute()

Publishes the public keys that verify the tokens issued here — the one URL you point a service at so it can check a token itself, offline, without calling back and without holding any secret of ours.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamWellKnownJwks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamWellKnownJwks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamWellKnownJwksRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamWellKnownOauthAuthorizationServer

> GetIamWellKnownOauthAuthorizationServer(ctx).Execute()

Returns the OpenID Connect discovery document — the one URL you point a standards-compliant client at so it can find every other endpoint on its own, instead of you configuring them by hand.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamWellKnownOauthAuthorizationServer(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamWellKnownOauthAuthorizationServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamWellKnownOauthAuthorizationServerRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamWellKnownOpenidConfiguration

> GetIamWellKnownOpenidConfiguration(ctx).Execute()

Returns the OpenID Connect discovery document — the one URL you point a standards-compliant client at so it can find every other endpoint on its own, instead of you configuring them by hand.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamWellKnownOpenidConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamWellKnownOpenidConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamWellKnownOpenidConfigurationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamWhoami

> GetIamWhoami(ctx).Execute()

Tells you who the current caller is — the lightweight check a page makes on load to decide whether to render signed-in or signed-out.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamWhoami(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamWhoami``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamWhoamiRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamWorkspaces

> IamWorkspacesListOutput GetIamWorkspaces(ctx).Owner(owner).Execute()

Returns your organization's workspaces, newest first — the scope a team works in, alongside projects rather than instead of them.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamWorkspaces(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamWorkspaces`: IamWorkspacesListOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

### Return type

[**IamWorkspacesListOutput**](IamWorkspacesListOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> IamOrganization GetOrganization(ctx).Owner(owner).Name(name).Execute()

Returns one organization: its display, its defaults and the sign-in rules everyone in it inherits.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)
	name := "name_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetOrganization(context.Background()).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganization`: IamOrganization
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**IamOrganization**](IamOrganization.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvider

> IamProviderResult GetProvider(ctx).IamProviderKey(iamProviderKey).Execute()

Returns one provider: what it connects to and how it is configured.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamProviderKey := *openapiclient.NewIamProviderKey("Name_example", "Owner_example") // IamProviderKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetProvider(context.Background()).IamProviderKey(iamProviderKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvider`: IamProviderResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProviderKey** | [**IamProviderKey**](IamProviderKey.md) |  | 

### Return type

[**IamProviderResult**](IamProviderResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSession

> IamSession GetSession(ctx).IamSessionRef(iamSessionRef).Execute()

Returns one person's session in one application — when it began and which browsers or devices are still carrying it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamSessionRef := *openapiclient.NewIamSessionRef("Application_example", "Name_example", "Owner_example") // IamSessionRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetSession(context.Background()).IamSessionRef(iamSessionRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSession`: IamSession
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamSessionRef** | [**IamSessionRef**](IamSessionRef.md) |  | 

### Return type

[**IamSession**](IamSession.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> IamTokenResult GetToken(ctx).IamTokenKey(iamTokenKey).Execute()

Returns one access token: who and what it was issued to, and when it expires.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamTokenKey := *openapiclient.NewIamTokenKey("Name_example", "Owner_example") // IamTokenKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetToken(context.Background()).IamTokenKey(iamTokenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToken`: IamTokenResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamTokenKey** | [**IamTokenKey**](IamTokenKey.md) |  | 

### Return type

[**IamTokenResult**](IamTokenResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebauthnCredential

> IamWebauthnCredentialResult GetWebauthnCredential(ctx).IamWebauthnCredentialKey(iamWebauthnCredentialKey).Execute()

Returns one passkey or security key: whose it is, what device it lives on, and when it was registered.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamWebauthnCredentialKey := *openapiclient.NewIamWebauthnCredentialKey("Name_example", "Owner_example") // IamWebauthnCredentialKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetWebauthnCredential(context.Background()).IamWebauthnCredentialKey(iamWebauthnCredentialKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetWebauthnCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebauthnCredential`: IamWebauthnCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetWebauthnCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebauthnCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWebauthnCredentialKey** | [**IamWebauthnCredentialKey**](IamWebauthnCredentialKey.md) |  | 

### Return type

[**IamWebauthnCredentialResult**](IamWebauthnCredentialResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizations

> IamListOrganizationsOutput ListOrganizations(ctx).Owner(owner).Limit(limit).Offset(offset).Execute()

Returns the organizations you can see, newest first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListOrganizations(context.Background()).Owner(owner).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.ListOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizations`: IamListOrganizationsOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.ListOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**IamListOrganizationsOutput**](IamListOrganizationsOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProviders

> IamListProvidersOut ListProviders(ctx).Owner(owner).Execute()

Returns your organization's providers, newest first — the identity providers your people sign in with, and the senders and connectors your applications go through.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListProviders(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.ListProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProviders`: IamListProvidersOut
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.ListProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

### Return type

[**IamListProvidersOut**](IamListProvidersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessions

> IamListSessionsOut ListSessions(ctx).IamListSessionsIn(iamListSessionsIn).Execute()

Returns who is currently signed in to your organization, newest first, and can be narrowed to one person or one application.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamListSessionsIn := *openapiclient.NewIamListSessionsIn("Owner_example") // IamListSessionsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListSessions(context.Background()).IamListSessionsIn(iamListSessionsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.ListSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSessions`: IamListSessionsOut
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.ListSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamListSessionsIn** | [**IamListSessionsIn**](IamListSessionsIn.md) |  | 

### Return type

[**IamListSessionsOut**](IamListSessionsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokens

> IamListTokensOut ListTokens(ctx).Owner(owner).Organization(organization).Execute()

Returns the access tokens issued in your organization, newest first, and can be narrowed to one organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string |  (optional)
	organization := "organization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListTokens(context.Background()).Owner(owner).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.ListTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTokens`: IamListTokensOut
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.ListTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **organization** | **string** |  | 

### Return type

[**IamListTokensOut**](IamListTokensOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebauthnCredentials

> IamListWebauthnCredentialsOut ListWebauthnCredentials(ctx).User(user).Execute()

Returns the passkeys and security keys registered to one person, newest first — which device each lives on and when it was registered.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	user := "user_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListWebauthnCredentials(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.ListWebauthnCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebauthnCredentials`: IamListWebauthnCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.ListWebauthnCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebauthnCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** |  | 

### Return type

[**IamListWebauthnCredentialsOut**](IamListWebauthnCredentialsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIamScimV2UsersByOwnerByName

> PatchIamScimV2UsersByOwnerByName(ctx, owner, name).Execute()

Applies a partial change from your identity provider — one attribute moved, not the whole record resent.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PatchIamScimV2UsersByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PatchIamScimV2UsersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIamScimV2UsersByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddApplication

> IamResponse PostIamAddApplication(ctx).IamApplication(iamApplication).Execute()

Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAddApplication(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAddApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddApplication`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAddApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddMembership

> PostIamAddMembership(ctx).Execute()

Lets a person or an application act in an organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamAddMembership(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAddMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddMembershipRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddOrganization

> IamResponse PostIamAddOrganization(ctx).IamCreateOrganizationInput(iamCreateOrganizationInput).Execute()

Creates an organization — the account everything else in your directory hangs from.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamCreateOrganizationInput := *openapiclient.NewIamCreateOrganizationInput() // IamCreateOrganizationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAddOrganization(context.Background()).IamCreateOrganizationInput(iamCreateOrganizationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAddOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddOrganization`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAddOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamCreateOrganizationInput** | [**IamCreateOrganizationInput**](IamCreateOrganizationInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddProject

> IamResponse PostIamAddProject(ctx).IamInput(iamInput).Execute()

Creates a project inside your organization — the scope people pick between when their work is separated by product or client rather than by team.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamInput := *openapiclient.NewIamInput() // IamInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAddProject(context.Background()).IamInput(iamInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAddProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddProject`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAddProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamInput** | [**IamInput**](IamInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddProvider

> IamResponse PostIamAddProvider(ctx).IamProvider(iamProvider).Execute()

Adds an identity provider your people can sign in with, or a service your applications send through — a social or enterprise login, an email or SMS sender, a storage or payment connector.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamProvider := *openapiclient.NewIamProvider() // IamProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAddProvider(context.Background()).IamProvider(iamProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAddProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddProvider`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAddProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProvider** | [**IamProvider**](IamProvider.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddRole

> IamResponse PostIamAddRole(ctx).IamRolesInput(iamRolesInput).Execute()

Creates a role — a named group of people that permissions are granted to.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamRolesInput := *openapiclient.NewIamRolesInput() // IamRolesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAddRole(context.Background()).IamRolesInput(iamRolesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddRole`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAddRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRolesInput** | [**IamRolesInput**](IamRolesInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddUser

> IamResponse PostIamAddUser(ctx).IamUserBody(iamUserBody).Execute()

Adds a person to your organization and, if you send a password, sets the one they will sign in with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamUserBody := *openapiclient.NewIamUserBody() // IamUserBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAddUser(context.Background()).IamUserBody(iamUserBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAddUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddUser`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUserBody** | [**IamUserBody**](IamUserBody.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddWorkspace

> IamResponse PostIamAddWorkspace(ctx).IamWorkspacesInput(iamWorkspacesInput).Execute()

Creates a workspace inside your organization — the scope a team works in, alongside projects rather than instead of them.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamWorkspacesInput := *openapiclient.NewIamWorkspacesInput() // IamWorkspacesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAddWorkspace(context.Background()).IamWorkspacesInput(iamWorkspacesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAddWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddWorkspace`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAddWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWorkspacesInput** | [**IamWorkspacesInput**](IamWorkspacesInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAdminProvision

> PostIamAdminProvision(ctx).Execute()

Sets up an account on someone's behalf — the same onboarding a person gets themselves, driven by one of your own services instead of by them.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamAdminProvision(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAdminProvision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAdminProvisionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamApplication

> IamApplication PostIamApplication(ctx).IamApplication(iamApplication).Execute()

Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamApplication(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamApplication`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamApplication**](IamApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamApplications

> IamApplication PostIamApplications(ctx).IamApplication(iamApplication).Execute()

Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamApplications(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamApplications`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamApplication**](IamApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamApplicationsDelete

> IamDeleteResult PostIamApplicationsDelete(ctx).IamApplicationRef(iamApplicationRef).Execute()

Removes an application.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamApplicationRef := *openapiclient.NewIamApplicationRef("Name_example", "Owner_example") // IamApplicationRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamApplicationsDelete(context.Background()).IamApplicationRef(iamApplicationRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamApplicationsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamApplicationsDelete`: IamDeleteResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamApplicationsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamApplicationsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplicationRef** | [**IamApplicationRef**](IamApplicationRef.md) |  | 

### Return type

[**IamDeleteResult**](IamDeleteResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamApplicationsUpdate

> IamApplication PostIamApplicationsUpdate(ctx).IamApplication(iamApplication).Execute()

Changes an application's display, its sign-in methods and the redirect URIs it may return to — the call that makes login work from a new host.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamApplicationsUpdate(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamApplicationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamApplicationsUpdate`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamApplicationsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamApplicationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamApplication**](IamApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAssume

> IamAnswer PostIamAssume(ctx).IamAssumeBody(iamAssumeBody).Authorization(authorization).XForwardedFor(xForwardedFor).Execute()

Steps a platform operator into an organization: it returns their own access token re-scoped to that tenant, so they see what the tenant sees.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamAssumeBody := *openapiclient.NewIamAssumeBody() // IamAssumeBody | 
	authorization := "authorization_example" // string |  (optional)
	xForwardedFor := "xForwardedFor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAssume(context.Background()).IamAssumeBody(iamAssumeBody).Authorization(authorization).XForwardedFor(xForwardedFor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAssume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAssume`: IamAnswer
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAssume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAssumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamAssumeBody** | [**IamAssumeBody**](IamAssumeBody.md) |  | 
 **authorization** | **string** |  | 
 **xForwardedFor** | **string** |  | 

### Return type

[**IamAnswer**](IamAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAuditLogs

> IamAuditLog PostIamAuditLogs(ctx).IamAuditlogsInput(iamAuditlogsInput).Execute()

Records an audit entry, so activity from your own systems lands in the same trail as everything the Hanzo Cloud records for you.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamAuditlogsInput := *openapiclient.NewIamAuditlogsInput() // IamAuditlogsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAuditLogs(context.Background()).IamAuditlogsInput(iamAuditlogsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAuditLogs`: IamAuditLog
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamAuditlogsInput** | [**IamAuditlogsInput**](IamAuditlogsInput.md) |  | 

### Return type

[**IamAuditLog**](IamAuditLog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAuditLogsDelete

> IamDeleteOutput PostIamAuditLogsDelete(ctx).IamRef(iamRef).Execute()

Removes an audit entry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamRef := *openapiclient.NewIamRef() // IamRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAuditLogsDelete(context.Background()).IamRef(iamRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAuditLogsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAuditLogsDelete`: IamDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAuditLogsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAuditLogsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRef** | [**IamRef**](IamRef.md) |  | 

### Return type

[**IamDeleteOutput**](IamDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAuditLogsGet

> IamAuditLog PostIamAuditLogsGet(ctx).IamRef(iamRef).Execute()

Returns one audit entry in full: the action, the person or key behind it, and the request it came in on.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamRef := *openapiclient.NewIamRef() // IamRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAuditLogsGet(context.Background()).IamRef(iamRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAuditLogsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAuditLogsGet`: IamAuditLog
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAuditLogsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAuditLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRef** | [**IamRef**](IamRef.md) |  | 

### Return type

[**IamAuditLog**](IamAuditLog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAuditLogsUpdate

> IamAuditLog PostIamAuditLogsUpdate(ctx).IamAuditlogsInput(iamAuditlogsInput).Execute()

Corrects an audit entry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamAuditlogsInput := *openapiclient.NewIamAuditlogsInput() // IamAuditlogsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAuditLogsUpdate(context.Background()).IamAuditlogsInput(iamAuditlogsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamAuditLogsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAuditLogsUpdate`: IamAuditLog
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamAuditLogsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAuditLogsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamAuditlogsInput** | [**IamAuditlogsInput**](IamAuditlogsInput.md) |  | 

### Return type

[**IamAuditLog**](IamAuditLog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamCerts

> IamCert PostIamCerts(ctx).IamCert(iamCert).Execute()

Adds a signing certificate your applications can verify tokens against — the call you make to bring your own key, or to stage the next one before a rotation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamCert := *openapiclient.NewIamCert() // IamCert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamCerts(context.Background()).IamCert(iamCert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamCerts`: IamCert
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamCert** | [**IamCert**](IamCert.md) |  | 

### Return type

[**IamCert**](IamCert.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamCertsDelete

> IamCertsDeleteOutput PostIamCertsDelete(ctx).IamCertsRef(iamCertsRef).Execute()

Removes a signing certificate.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamCertsRef := *openapiclient.NewIamCertsRef() // IamCertsRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamCertsDelete(context.Background()).IamCertsRef(iamCertsRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamCertsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamCertsDelete`: IamCertsDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamCertsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamCertsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamCertsRef** | [**IamCertsRef**](IamCertsRef.md) |  | 

### Return type

[**IamCertsDeleteOutput**](IamCertsDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamCertsGet

> IamCert PostIamCertsGet(ctx).IamCertsRef(iamCertsRef).Execute()

Returns one signing certificate — its algorithm, its validity window and its public half.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamCertsRef := *openapiclient.NewIamCertsRef() // IamCertsRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamCertsGet(context.Background()).IamCertsRef(iamCertsRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamCertsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamCertsGet`: IamCert
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamCertsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamCertsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamCertsRef** | [**IamCertsRef**](IamCertsRef.md) |  | 

### Return type

[**IamCert**](IamCert.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamCertsUpdate

> IamCert PostIamCertsUpdate(ctx).IamCert(iamCert).Execute()

Changes a signing certificate's settings.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamCert := *openapiclient.NewIamCert() // IamCert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamCertsUpdate(context.Background()).IamCert(iamCert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamCertsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamCertsUpdate`: IamCert
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamCertsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamCertsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamCert** | [**IamCert**](IamCert.md) |  | 

### Return type

[**IamCert**](IamCert.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteApplication

> IamResponse PostIamDeleteApplication(ctx).IamApplication(iamApplication).Execute()

Deletes an application.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamDeleteApplication(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamDeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteApplication`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamDeleteApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteMembership

> PostIamDeleteMembership(ctx).Execute()

Takes away a person's or an application's right to act in an organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamDeleteMembership(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamDeleteMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteMembershipRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteMfa

> PostIamDeleteMfa(ctx).Execute()

Turns a factor off, so sign-in stops asking for it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamDeleteMfa(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamDeleteMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteMfaRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteOrganization

> IamResponse PostIamDeleteOrganization(ctx).IamDeleteOrganizationInput(iamDeleteOrganizationInput).Execute()

Deletes an organization and everything named inside it — its users, applications, roles, projects and workspaces.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamDeleteOrganizationInput := *openapiclient.NewIamDeleteOrganizationInput() // IamDeleteOrganizationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamDeleteOrganization(context.Background()).IamDeleteOrganizationInput(iamDeleteOrganizationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamDeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteOrganization`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamDeleteOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamDeleteOrganizationInput** | [**IamDeleteOrganizationInput**](IamDeleteOrganizationInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteProject

> IamResponse PostIamDeleteProject(ctx).IamProjectsRef(iamProjectsRef).Execute()

Deletes a project.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamProjectsRef := *openapiclient.NewIamProjectsRef() // IamProjectsRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamDeleteProject(context.Background()).IamProjectsRef(iamProjectsRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamDeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteProject`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamDeleteProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProjectsRef** | [**IamProjectsRef**](IamProjectsRef.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteProvider

> IamResponse PostIamDeleteProvider(ctx).IamProvider(iamProvider).Execute()

Removes a provider.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamProvider := *openapiclient.NewIamProvider() // IamProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamDeleteProvider(context.Background()).IamProvider(iamProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamDeleteProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteProvider`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamDeleteProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProvider** | [**IamProvider**](IamProvider.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteRole

> IamResponse PostIamDeleteRole(ctx).IamRolesRef(iamRolesRef).Execute()

Deletes a role.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamRolesRef := *openapiclient.NewIamRolesRef() // IamRolesRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamDeleteRole(context.Background()).IamRolesRef(iamRolesRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamDeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteRole`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamDeleteRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRolesRef** | [**IamRolesRef**](IamRolesRef.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteUser

> IamResponse PostIamDeleteUser(ctx).IamUserBody(iamUserBody).Execute()

Removes a person from your organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamUserBody := *openapiclient.NewIamUserBody() // IamUserBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamDeleteUser(context.Background()).IamUserBody(iamUserBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteUser`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamDeleteUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUserBody** | [**IamUserBody**](IamUserBody.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteWorkspace

> IamResponse PostIamDeleteWorkspace(ctx).IamWorkspacesRef(iamWorkspacesRef).Execute()

Deletes a workspace.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamWorkspacesRef := *openapiclient.NewIamWorkspacesRef() // IamWorkspacesRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamDeleteWorkspace(context.Background()).IamWorkspacesRef(iamWorkspacesRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamDeleteWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteWorkspace`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamDeleteWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWorkspacesRef** | [**IamWorkspacesRef**](IamWorkspacesRef.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamInvitations

> IamInvitation PostIamInvitations(ctx).IamInvitationsInput(iamInvitationsInput).Execute()

Issues an invitation to join your organization — the code or link a new member redeems, with the role they arrive holding and the date it stops working.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamInvitationsInput := *openapiclient.NewIamInvitationsInput() // IamInvitationsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamInvitations(context.Background()).IamInvitationsInput(iamInvitationsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamInvitations`: IamInvitation
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamInvitationsInput** | [**IamInvitationsInput**](IamInvitationsInput.md) |  | 

### Return type

[**IamInvitation**](IamInvitation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamInvitationsDelete

> IamInvitationsDeleteOutput PostIamInvitationsDelete(ctx).IamInvitationsRef(iamInvitationsRef).Execute()

Withdraws an invitation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamInvitationsRef := *openapiclient.NewIamInvitationsRef() // IamInvitationsRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamInvitationsDelete(context.Background()).IamInvitationsRef(iamInvitationsRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamInvitationsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamInvitationsDelete`: IamInvitationsDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamInvitationsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamInvitationsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamInvitationsRef** | [**IamInvitationsRef**](IamInvitationsRef.md) |  | 

### Return type

[**IamInvitationsDeleteOutput**](IamInvitationsDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamInvitationsGet

> IamInvitation PostIamInvitationsGet(ctx).IamInvitationsRef(iamInvitationsRef).Execute()

Returns one invitation: who it is for, what it grants on acceptance, and when it expires.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamInvitationsRef := *openapiclient.NewIamInvitationsRef() // IamInvitationsRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamInvitationsGet(context.Background()).IamInvitationsRef(iamInvitationsRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamInvitationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamInvitationsGet`: IamInvitation
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamInvitationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamInvitationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamInvitationsRef** | [**IamInvitationsRef**](IamInvitationsRef.md) |  | 

### Return type

[**IamInvitation**](IamInvitation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamInvitationsUpdate

> IamInvitation PostIamInvitationsUpdate(ctx).IamInvitationsInput(iamInvitationsInput).Execute()

Changes an invitation's terms — the role it grants, how many may redeem it, or when it expires.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamInvitationsInput := *openapiclient.NewIamInvitationsInput() // IamInvitationsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamInvitationsUpdate(context.Background()).IamInvitationsInput(iamInvitationsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamInvitationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamInvitationsUpdate`: IamInvitation
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamInvitationsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamInvitationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamInvitationsInput** | [**IamInvitationsInput**](IamInvitationsInput.md) |  | 

### Return type

[**IamInvitation**](IamInvitation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamIssueUserToken

> PostIamIssueUserToken(ctx).Execute()

Mints an access token for the `?id=<owner>/<name>` target user (optional `?aud=` resource, RFC 8707), issued by the authenticated + allow-listed confidential client.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamIssueUserToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamIssueUserToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamIssueUserTokenRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamKeys

> IamKey PostIamKeys(ctx).IamKey(iamKey).Execute()

Issues an API key.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamKey := *openapiclient.NewIamKey() // IamKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamKeys(context.Background()).IamKey(iamKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamKeys`: IamKey
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamKey** | [**IamKey**](IamKey.md) |  | 

### Return type

[**IamKey**](IamKey.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamKeysDelete

> IamDeleteResponse PostIamKeysDelete(ctx).IamKeysRef(iamKeysRef).Execute()

Revokes an API key.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamKeysRef := *openapiclient.NewIamKeysRef() // IamKeysRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamKeysDelete(context.Background()).IamKeysRef(iamKeysRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamKeysDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamKeysDelete`: IamDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamKeysDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamKeysDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamKeysRef** | [**IamKeysRef**](IamKeysRef.md) |  | 

### Return type

[**IamDeleteResponse**](IamDeleteResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamKeysMint

> PostIamKeysMint(ctx).Execute()

(re)generates the target user's key of the requested TYPE and returns it once, over the shared authorizeMinter + mintTarget seam.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamKeysMint(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamKeysMint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamKeysMintRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamKeysRevoke

> PostIamKeysRevoke(ctx).Execute()

Clears the target user's key of the requested TYPE (immediate revoke).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamKeysRevoke(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamKeysRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamKeysRevokeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamKeysUpdate

> IamKey PostIamKeysUpdate(ctx).IamKey(iamKey).Execute()

Changes what a key is called or what it may reach.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamKey := *openapiclient.NewIamKey() // IamKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamKeysUpdate(context.Background()).IamKey(iamKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamKeysUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamKeysUpdate`: IamKey
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamKeysUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamKeysUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamKey** | [**IamKey**](IamKey.md) |  | 

### Return type

[**IamKey**](IamKey.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamLink

> PostIamLink(ctx).Execute()

Starts connecting another sign-in identity to the account you are already signed in as.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamLinkRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamLogin

> PostIamLogin(ctx).Execute()

Signs a person in with the credential they typed, and — when the request is part of an OAuth flow — hands back the one-time code that finishes it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamLogin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamLoginRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamMemberships

> PostIamMemberships(ctx).Execute()

Lets a person or an application act in an organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamMemberships(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamMembershipsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamMfaDisable

> PostIamMfaDisable(ctx).Execute()

Turns a factor off, so sign-in stops asking for it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamMfaDisable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamMfaDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamMfaDisableRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamMfaPreferred

> PostIamMfaPreferred(ctx).Execute()

Picks which second factor an account is asked for first when it has more than one.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamMfaPreferred(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamMfaPreferred``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamMfaPreferredRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamMfaSetupEnable

> PostIamMfaSetupEnable(ctx).Execute()

Finishes the enrolment: from here the account's sign-ins ask for this factor.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamMfaSetupEnable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamMfaSetupEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamMfaSetupEnableRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamMfaSetupInitiate

> PostIamMfaSetupInitiate(ctx).Execute()

Starts enrolling a factor and hands over whatever the person needs to prove they hold it: app a fresh secret and the otpauth:// URL to render as a QR code sms a code texted to the number on the account email a code mailed to the address on the account Nothing is switched on yet, so abandoning this step leaves the account exactly as it was.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamMfaSetupInitiate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamMfaSetupInitiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamMfaSetupInitiateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamMintUserKeys

> PostIamMintUserKeys(ctx).Execute()

(re)generates the target user's key of the requested TYPE and returns it once, over the shared authorizeMinter + mintTarget seam.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamMintUserKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamMintUserKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamMintUserKeysRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamOauthAuthorize

> PostIamOauthAuthorize(ctx).Execute()

Starts a sign-in — the address you send a browser to, and the beginning of every OAuth and OpenID Connect flow.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamOauthAuthorize(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamOauthAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamOauthAuthorizeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamOauthDevice

> PostIamOauthDevice(ctx).Execute()

Starts a sign-in on a device with no browser and no keyboard — a TV, a CLI, a headless box.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamOauthDevice(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamOauthDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamOauthDeviceRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamOauthDeviceInfo

> PostIamOauthDeviceInfo(ctx).Execute()

Answers \"what am I approving?\" for a pending device code.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamOauthDeviceInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamOauthDeviceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamOauthDeviceInfoRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamOauthFederationMfa

> PostIamOauthFederationMfa(ctx).Execute()

Completes a sign-in that came in through another identity provider and still owes a second factor.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamOauthFederationMfa(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamOauthFederationMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamOauthFederationMfaRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamOauthIntrospect

> PostIamOauthIntrospect(ctx).Execute()

Answers whether an access token is still good, and what it is good for — the check a resource server of yours makes before honouring a token it did not mint.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamOauthIntrospect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamOauthIntrospect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamOauthIntrospectRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamOauthLogout

> PostIamOauthLogout(ctx).Execute()

Ends a sign-in and sends the browser somewhere sensible.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamOauthLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamOauthLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamOauthLogoutRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamOauthRevoke

> PostIamOauthRevoke(ctx).Execute()

Retires a token before it expires — what you call when someone signs out or a credential may have leaked.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamOauthRevoke(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamOauthRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamOauthRevokeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamOauthToken

> PostIamOauthToken(ctx).Execute()

Exchanges what your application is holding for the tokens it needs — the one-time code from a finished sign-in, a refresh token, or your own client credentials when the caller is a program rather than a person.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamOauthToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamOauthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamOauthTokenRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamOauthUserinfo

> PostIamOauthUserinfo(ctx).Execute()

Returns the profile claims for whoever the access token belongs to — the standard OpenID Connect way to find out who is calling you without your application storing anything itself.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamOauthUserinfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamOauthUserinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamOauthUserinfoRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamOnboard

> PostIamOnboard(ctx).Execute()

Finishes setting up the account of whoever is calling — it creates their organization if they have none and puts them in it, so a person who has just signed up lands somewhere they can work.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamOnboard(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamOnboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamOnboardRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamPermissions

> IamPermission PostIamPermissions(ctx).IamPermission(iamPermission).Execute()

Grants a permission — the call that gives a person or a role the ability to do something.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamPermission := *openapiclient.NewIamPermission() // IamPermission | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamPermissions(context.Background()).IamPermission(iamPermission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamPermissions`: IamPermission
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamPermission** | [**IamPermission**](IamPermission.md) |  | 

### Return type

[**IamPermission**](IamPermission.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamPermissionsDelete

> IamPermissionDeleteResponse PostIamPermissionsDelete(ctx).IamPermissionRef(iamPermissionRef).Execute()

Revokes a permission.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamPermissionRef := *openapiclient.NewIamPermissionRef() // IamPermissionRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamPermissionsDelete(context.Background()).IamPermissionRef(iamPermissionRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamPermissionsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamPermissionsDelete`: IamPermissionDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamPermissionsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamPermissionsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamPermissionRef** | [**IamPermissionRef**](IamPermissionRef.md) |  | 

### Return type

[**IamPermissionDeleteResponse**](IamPermissionDeleteResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamPermissionsUpdate

> IamPermission PostIamPermissionsUpdate(ctx).IamPermission(iamPermission).Execute()

Changes who a permission grants to, what it allows, or the resources it covers.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamPermission := *openapiclient.NewIamPermission() // IamPermission | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamPermissionsUpdate(context.Background()).IamPermission(iamPermission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamPermissionsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamPermissionsUpdate`: IamPermission
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamPermissionsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamPermissionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamPermission** | [**IamPermission**](IamPermission.md) |  | 

### Return type

[**IamPermission**](IamPermission.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamPreferences

> PostIamPreferences(ctx).Execute()

Saves the calling person's own settings and returns the full set afterwards.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamPreferences(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamPreferencesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamProjects

> IamProject PostIamProjects(ctx).IamInput(iamInput).Execute()

Makes a project inside your organization — the scope people pick between when their work is separated by product or client rather than by team.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamInput := *openapiclient.NewIamInput() // IamInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamProjects(context.Background()).IamInput(iamInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamProjects`: IamProject
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamInput** | [**IamInput**](IamInput.md) |  | 

### Return type

[**IamProject**](IamProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamProjectsDelete

> IamProjectsDeleteOutput PostIamProjectsDelete(ctx).IamProjectsRef(iamProjectsRef).Execute()

Removes a project.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamProjectsRef := *openapiclient.NewIamProjectsRef() // IamProjectsRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamProjectsDelete(context.Background()).IamProjectsRef(iamProjectsRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamProjectsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamProjectsDelete`: IamProjectsDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamProjectsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamProjectsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProjectsRef** | [**IamProjectsRef**](IamProjectsRef.md) |  | 

### Return type

[**IamProjectsDeleteOutput**](IamProjectsDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamProjectsGet

> IamProject PostIamProjectsGet(ctx).IamProjectsRef(iamProjectsRef).Execute()

Returns one project: what it is called and how it is set up.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamProjectsRef := *openapiclient.NewIamProjectsRef() // IamProjectsRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamProjectsGet(context.Background()).IamProjectsRef(iamProjectsRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamProjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamProjectsGet`: IamProject
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamProjectsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProjectsRef** | [**IamProjectsRef**](IamProjectsRef.md) |  | 

### Return type

[**IamProject**](IamProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamProjectsUpdate

> IamProject PostIamProjectsUpdate(ctx).IamInput(iamInput).Execute()

Changes a project's settings.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamInput := *openapiclient.NewIamInput() // IamInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamProjectsUpdate(context.Background()).IamInput(iamInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamProjectsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamProjectsUpdate`: IamProject
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamProjectsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamProjectsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamInput** | [**IamInput**](IamInput.md) |  | 

### Return type

[**IamProject**](IamProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamRegistryToken

> PostIamRegistryToken(ctx).Execute()

Signs a container client in to your registry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamRegistryToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamRegistryToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamRegistryTokenRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamRelease

> IamAnswer PostIamRelease(ctx).IamAssumeBody(iamAssumeBody).Authorization(authorization).XForwardedFor(xForwardedFor).Execute()

Steps a platform operator back out: it returns their own access token with no organization assumed, which is the credential they had before they stepped in.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamAssumeBody := *openapiclient.NewIamAssumeBody() // IamAssumeBody | 
	authorization := "authorization_example" // string |  (optional)
	xForwardedFor := "xForwardedFor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamRelease(context.Background()).IamAssumeBody(iamAssumeBody).Authorization(authorization).XForwardedFor(xForwardedFor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamRelease`: IamAnswer
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamAssumeBody** | [**IamAssumeBody**](IamAssumeBody.md) |  | 
 **authorization** | **string** |  | 
 **xForwardedFor** | **string** |  | 

### Return type

[**IamAnswer**](IamAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamRevokeUserKeys

> PostIamRevokeUserKeys(ctx).Execute()

Clears the target user's key of the requested TYPE (immediate revoke).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamRevokeUserKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamRevokeUserKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamRevokeUserKeysRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamRoles

> IamRole PostIamRoles(ctx).IamRolesInput(iamRolesInput).Execute()

Makes a role — a named group of people that permissions are granted to.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamRolesInput := *openapiclient.NewIamRolesInput() // IamRolesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamRoles(context.Background()).IamRolesInput(iamRolesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamRoles`: IamRole
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRolesInput** | [**IamRolesInput**](IamRolesInput.md) |  | 

### Return type

[**IamRole**](IamRole.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamRolesDelete

> IamRolesDeleteOutput PostIamRolesDelete(ctx).IamRolesRef(iamRolesRef).Execute()

Removes a role.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamRolesRef := *openapiclient.NewIamRolesRef() // IamRolesRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamRolesDelete(context.Background()).IamRolesRef(iamRolesRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamRolesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamRolesDelete`: IamRolesDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamRolesDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamRolesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRolesRef** | [**IamRolesRef**](IamRolesRef.md) |  | 

### Return type

[**IamRolesDeleteOutput**](IamRolesDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamRolesGet

> IamRole PostIamRolesGet(ctx).IamRolesRef(iamRolesRef).Execute()

Returns one role: who is in it, and the roles it includes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamRolesRef := *openapiclient.NewIamRolesRef() // IamRolesRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamRolesGet(context.Background()).IamRolesRef(iamRolesRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamRolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamRolesGet`: IamRole
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamRolesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRolesRef** | [**IamRolesRef**](IamRolesRef.md) |  | 

### Return type

[**IamRole**](IamRole.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamRolesUpdate

> IamRole PostIamRolesUpdate(ctx).IamRolesInput(iamRolesInput).Execute()

Changes who is in a role, or which roles it includes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamRolesInput := *openapiclient.NewIamRolesInput() // IamRolesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamRolesUpdate(context.Background()).IamRolesInput(iamRolesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamRolesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamRolesUpdate`: IamRole
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamRolesUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamRolesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRolesInput** | [**IamRolesInput**](IamRolesInput.md) |  | 

### Return type

[**IamRole**](IamRole.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamScimV2Users

> PostIamScimV2Users(ctx).Execute()

Provisions a person from your identity provider — how a new hire gets an account here automatically when they are added over there.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamScimV2Users(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamScimV2Users``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamScimV2UsersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamSendVerificationCode

> PostIamSendVerificationCode(ctx).Execute()

Validates the request and asks otp to get a code to the person.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamSendVerificationCode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamSendVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamSendVerificationCodeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamServiceAccounts

> PostIamServiceAccounts(ctx).Execute()

Makes a service account — an identity for a program rather than a person, for a script, a bot or a deployment that has to authenticate on its own.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamServiceAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamServiceAccountsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamServiceAccountsByNameKeys

> PostIamServiceAccountsByNameKeys(ctx, name).Execute()

Serves POST /v1/iam/service-accounts/:name/keys: mint a fresh key, invalidating the prior one, and return the new raw secret exactly once.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamServiceAccountsByNameKeys(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamServiceAccountsByNameKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamServiceAccountsByNameKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamSetPreferredMfa

> PostIamSetPreferredMfa(ctx).Execute()

Picks which second factor an account is asked for first when it has more than one.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamSetPreferredMfa(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamSetPreferredMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamSetPreferredMfaRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamSignin

> PostIamSignin(ctx).Execute()

Completes a sign-in: it exchanges the one-time code your application was handed at the end of the login flow for a live session, and returns the signed-in account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamSignin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamSignin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamSigninRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamSignup

> PostIamSignup(ctx).Execute()

Creates an account from the sign-up form and applies the application's own sign-up rules — whether self-service registration is open at all, and which fields it requires.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamSignup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamSignup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamSignupRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamTokensIssue

> PostIamTokensIssue(ctx).Execute()

Mints an access token for the `?id=<owner>/<name>` target user (optional `?aud=` resource, RFC 8707), issued by the authenticated + allow-listed confidential client.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamTokensIssue(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamTokensIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamTokensIssueRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUnlink

> PostIamUnlink(ctx).Execute()

Disconnects one sign-in identity from an account, so that provider can no longer be used to sign in as that person.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamUnlink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUnlink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUnlinkRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdateApplication

> IamResponse PostIamUpdateApplication(ctx).IamApplication(iamApplication).Execute()

Updates one of your applications — its display, its sign-in methods and the redirect URIs it is allowed to return to.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamUpdateApplication(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUpdateApplication`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamUpdateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdateOrganization

> IamResponse PostIamUpdateOrganization(ctx).IamUpdateOrganizationInput(iamUpdateOrganizationInput).Execute()

Updates your organization — its display, its default settings and the sign-in rules everyone in it inherits.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamUpdateOrganizationInput := *openapiclient.NewIamUpdateOrganizationInput() // IamUpdateOrganizationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamUpdateOrganization(context.Background()).IamUpdateOrganizationInput(iamUpdateOrganizationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUpdateOrganization`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamUpdateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUpdateOrganizationInput** | [**IamUpdateOrganizationInput**](IamUpdateOrganizationInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdatePreferences

> PostIamUpdatePreferences(ctx).Execute()

Saves the calling person's own settings and returns the full set afterwards.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamUpdatePreferences(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUpdatePreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdatePreferencesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdateProvider

> IamResponse PostIamUpdateProvider(ctx).IamProvider(iamProvider).Execute()

Updates a provider's settings or rotates the credentials it holds.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamProvider := *openapiclient.NewIamProvider() // IamProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamUpdateProvider(context.Background()).IamProvider(iamProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUpdateProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUpdateProvider`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamUpdateProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProvider** | [**IamProvider**](IamProvider.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdateRole

> IamResponse PostIamUpdateRole(ctx).IamRolesInput(iamRolesInput).Execute()

Updates a role's members or the roles it includes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamRolesInput := *openapiclient.NewIamRolesInput() // IamRolesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamUpdateRole(context.Background()).IamRolesInput(iamRolesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUpdateRole`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamUpdateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRolesInput** | [**IamRolesInput**](IamRolesInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdateUser

> IamResponse PostIamUpdateUser(ctx).IamUserBody(iamUserBody).Execute()

Updates one of your users' profile, roles or credentials.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamUserBody := *openapiclient.NewIamUserBody() // IamUserBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamUpdateUser(context.Background()).IamUserBody(iamUserBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUpdateUser`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamUpdateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUserBody** | [**IamUserBody**](IamUserBody.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUsers

> IamUser PostIamUsers(ctx).IamCreateInput(iamCreateInput).Execute()

Adds a person to your organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamCreateInput := *openapiclient.NewIamCreateInput() // IamCreateInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamUsers(context.Background()).IamCreateInput(iamCreateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUsers`: IamUser
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamCreateInput** | [**IamCreateInput**](IamCreateInput.md) |  | 

### Return type

[**IamUser**](IamUser.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUsersDelete

> IamUsersDeleteOutput PostIamUsersDelete(ctx).IamUsersRef(iamUsersRef).Execute()

Removes a person from your organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamUsersRef := *openapiclient.NewIamUsersRef("Name_example", "Owner_example") // IamUsersRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamUsersDelete(context.Background()).IamUsersRef(iamUsersRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUsersDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUsersDelete`: IamUsersDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamUsersDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUsersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUsersRef** | [**IamUsersRef**](IamUsersRef.md) |  | 

### Return type

[**IamUsersDeleteOutput**](IamUsersDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUsersUpdate

> IamUser PostIamUsersUpdate(ctx).IamUpdateInput(iamUpdateInput).Execute()

Changes a person's profile, their roles, or the credentials they sign in with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamUpdateInput := *openapiclient.NewIamUpdateInput() // IamUpdateInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamUsersUpdate(context.Background()).IamUpdateInput(iamUpdateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUsersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUsersUpdate`: IamUser
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamUsersUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUpdateInput** | [**IamUpdateInput**](IamUpdateInput.md) |  | 

### Return type

[**IamUser**](IamUser.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamVerificationCodes

> PostIamVerificationCodes(ctx).Execute()

Validates the request and asks otp to get a code to the person.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamVerificationCodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamVerificationCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamVerificationCodesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamWeb3Verify

> PostIamWeb3Verify(ctx).Execute()

Completes a wallet sign-in: it verifies the signed challenge and, if it holds, signs the wallet's owner in.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamWeb3Verify(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamWeb3Verify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamWeb3VerifyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamWebauthnSigninFinish

> PostIamWebauthnSigninFinish(ctx).Execute()

Verifies the signed challenge and signs the person in.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamWebauthnSigninFinish(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamWebauthnSigninFinish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamWebauthnSigninFinishRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamWebauthnSignupFinish

> PostIamWebauthnSignupFinish(ctx).Execute()

Verifies the newly created passkey and stores it, so the person can sign in with their device from then on.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamWebauthnSignupFinish(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamWebauthnSignupFinish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIamWebauthnSignupFinishRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamWorkspaces

> IamWorkspace PostIamWorkspaces(ctx).IamWorkspacesInput(iamWorkspacesInput).Execute()

Makes a workspace inside your organization — the scope a team works in, alongside projects rather than instead of them.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamWorkspacesInput := *openapiclient.NewIamWorkspacesInput() // IamWorkspacesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamWorkspaces(context.Background()).IamWorkspacesInput(iamWorkspacesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamWorkspaces`: IamWorkspace
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWorkspacesInput** | [**IamWorkspacesInput**](IamWorkspacesInput.md) |  | 

### Return type

[**IamWorkspace**](IamWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamWorkspacesDelete

> IamWorkspacesDeleteOutput PostIamWorkspacesDelete(ctx).IamWorkspacesRef(iamWorkspacesRef).Execute()

Removes a workspace.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamWorkspacesRef := *openapiclient.NewIamWorkspacesRef() // IamWorkspacesRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamWorkspacesDelete(context.Background()).IamWorkspacesRef(iamWorkspacesRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamWorkspacesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamWorkspacesDelete`: IamWorkspacesDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamWorkspacesDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamWorkspacesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWorkspacesRef** | [**IamWorkspacesRef**](IamWorkspacesRef.md) |  | 

### Return type

[**IamWorkspacesDeleteOutput**](IamWorkspacesDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamWorkspacesGet

> IamWorkspace PostIamWorkspacesGet(ctx).IamWorkspacesRef(iamWorkspacesRef).Execute()

Returns one workspace: what it is called and how it is set up.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamWorkspacesRef := *openapiclient.NewIamWorkspacesRef() // IamWorkspacesRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamWorkspacesGet(context.Background()).IamWorkspacesRef(iamWorkspacesRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamWorkspacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamWorkspacesGet`: IamWorkspace
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamWorkspacesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamWorkspacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWorkspacesRef** | [**IamWorkspacesRef**](IamWorkspacesRef.md) |  | 

### Return type

[**IamWorkspace**](IamWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamWorkspacesUpdate

> IamWorkspace PostIamWorkspacesUpdate(ctx).IamWorkspacesInput(iamWorkspacesInput).Execute()

Changes a workspace's settings.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamWorkspacesInput := *openapiclient.NewIamWorkspacesInput() // IamWorkspacesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamWorkspacesUpdate(context.Background()).IamWorkspacesInput(iamWorkspacesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamWorkspacesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamWorkspacesUpdate`: IamWorkspace
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PostIamWorkspacesUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamWorkspacesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWorkspacesInput** | [**IamWorkspacesInput**](IamWorkspacesInput.md) |  | 

### Return type

[**IamWorkspace**](IamWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIamAccount

> IamAnswer PutIamAccount(ctx).IamAccountBody(iamAccountBody).Cookie(cookie).Authorization(authorization).Execute()

Saves the calling person's own profile — the name they are shown by, their picture, a line about themselves and a link.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamAccountBody := *openapiclient.NewIamAccountBody() // IamAccountBody | 
	cookie := "cookie_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamAccount(context.Background()).IamAccountBody(iamAccountBody).Cookie(cookie).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamAccount`: IamAnswer
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIamAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamAccountBody** | [**IamAccountBody**](IamAccountBody.md) |  | 
 **cookie** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**IamAnswer**](IamAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIamApplication

> IamApplication PutIamApplication(ctx).IamApplication(iamApplication).Execute()

Changes an application's display, its sign-in methods and the redirect URIs it may return to — the call that makes login work from a new host.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamApplication(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamApplication`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIamApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamApplication**](IamApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIamConsent

> PutIamConsent(ctx).Execute()

Records the calling person's privacy and communication choices.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PutIamConsent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamConsent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamConsentRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIamPassword

> IamAnswer PutIamPassword(ctx).IamPasswordBody(iamPasswordBody).Cookie(cookie).Authorization(authorization).Execute()

Replaces the calling person's password.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamPasswordBody := *openapiclient.NewIamPasswordBody() // IamPasswordBody | 
	cookie := "cookie_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamPassword(context.Background()).IamPasswordBody(iamPasswordBody).Cookie(cookie).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamPassword`: IamAnswer
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIamPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamPasswordBody** | [**IamPasswordBody**](IamPasswordBody.md) |  | 
 **cookie** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**IamAnswer**](IamAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIamScimV2UsersByOwnerByName

> PutIamScimV2UsersByOwnerByName(ctx, owner, name).Execute()

Overwrites a person's SCIM attributes with what your identity provider sends — how a change made there lands here.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PutIamScimV2UsersByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamScimV2UsersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamScimV2UsersByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrganizations

> IamSearchOrganizationsOutput SearchOrganizations(ctx).XForwardedFor(xForwardedFor).Q(q).Limit(limit).Cursor(cursor).Execute()

Returns the organizations you can act in, the ones you belong to first and the rest after, newest first, narrowed by an optional query against the name or the display name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	xForwardedFor := "xForwardedFor_example" // string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.SearchOrganizations(context.Background()).XForwardedFor(xForwardedFor).Q(q).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.SearchOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrganizations`: IamSearchOrganizationsOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.SearchOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xForwardedFor** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 

### Return type

[**IamSearchOrganizationsOutput**](IamSearchOrganizationsOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrganizationAvatar

> IamOrganization SetOrganizationAvatar(ctx).IamSetAvatarInput(iamSetAvatarInput).Execute()

Changes how an organization appears across Hanzo: the square mark beside its name, as an uploaded image or as a single emoji.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamSetAvatarInput := *openapiclient.NewIamSetAvatarInput() // IamSetAvatarInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.SetOrganizationAvatar(context.Background()).IamSetAvatarInput(iamSetAvatarInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.SetOrganizationAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrganizationAvatar`: IamOrganization
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.SetOrganizationAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetOrganizationAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamSetAvatarInput** | [**IamSetAvatarInput**](IamSetAvatarInput.md) |  | 

### Return type

[**IamOrganization**](IamOrganization.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization

> IamOrganization UpdateOrganization(ctx).IamUpdateOrganizationInput(iamUpdateOrganizationInput).Execute()

Changes an organization's display, its defaults and the sign-in rules everyone in it inherits.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamUpdateOrganizationInput := *openapiclient.NewIamUpdateOrganizationInput() // IamUpdateOrganizationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateOrganization(context.Background()).IamUpdateOrganizationInput(iamUpdateOrganizationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganization`: IamOrganization
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUpdateOrganizationInput** | [**IamUpdateOrganizationInput**](IamUpdateOrganizationInput.md) |  | 

### Return type

[**IamOrganization**](IamOrganization.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProvider

> IamMutationResult UpdateProvider(ctx).IamProvider(iamProvider).Execute()

Changes a provider's settings or rotates the credentials it holds.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamProvider := *openapiclient.NewIamProvider() // IamProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateProvider(context.Background()).IamProvider(iamProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProvider`: IamMutationResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProvider** | [**IamProvider**](IamProvider.md) |  | 

### Return type

[**IamMutationResult**](IamMutationResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSession

> IamSession UpdateSession(ctx).IamUpdateSessionIn(iamUpdateSessionIn).Execute()

Replaces the set of browsers a session covers — signing out the ones you leave off while the session itself stays live.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamUpdateSessionIn := *openapiclient.NewIamUpdateSessionIn("Application_example", "Name_example", "Owner_example") // IamUpdateSessionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateSession(context.Background()).IamUpdateSessionIn(iamUpdateSessionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSession`: IamSession
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUpdateSessionIn** | [**IamUpdateSessionIn**](IamUpdateSessionIn.md) |  | 

### Return type

[**IamSession**](IamSession.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToken

> IamTokenMutation UpdateToken(ctx).IamToken(iamToken).Execute()

Changes an access token's scope or expiry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamToken := *openapiclient.NewIamToken() // IamToken | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateToken(context.Background()).IamToken(iamToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateToken`: IamTokenMutation
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamToken** | [**IamToken**](IamToken.md) |  | 

### Return type

[**IamTokenMutation**](IamTokenMutation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebauthnCredential

> IamWebauthnCredentialMutationResult UpdateWebauthnCredential(ctx).IamWebauthnCredential(iamWebauthnCredential).Execute()

Renames a registered passkey or security key, so a person can tell their devices apart.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamWebauthnCredential := *openapiclient.NewIamWebauthnCredential() // IamWebauthnCredential | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateWebauthnCredential(context.Background()).IamWebauthnCredential(iamWebauthnCredential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateWebauthnCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebauthnCredential`: IamWebauthnCredentialMutationResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateWebauthnCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebauthnCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWebauthnCredential** | [**IamWebauthnCredential**](IamWebauthnCredential.md) |  | 

### Return type

[**IamWebauthnCredentialMutationResult**](IamWebauthnCredentialMutationResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertApplication

> IamReply UpsertApplication(ctx).IamRegistration(iamRegistration).Authorization(authorization).Execute()

Creates an application or updates it in place, so a deployment can declare the applications it needs and run the same declaration on every environment and on every redeploy.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamRegistration := *openapiclient.NewIamRegistration() // IamRegistration | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpsertApplication(context.Background()).IamRegistration(iamRegistration).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpsertApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertApplication`: IamReply
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpsertApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRegistration** | [**IamRegistration**](IamRegistration.md) |  | 
 **authorization** | **string** |  | 

### Return type

[**IamReply**](IamReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertUser

> IamReply UpsertUser(ctx).IamPerson(iamPerson).Authorization(authorization).Execute()

Creates a person or updates them in place, so a deployment can declare the accounts it needs and re-run that declaration safely.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	iamPerson := *openapiclient.NewIamPerson() // IamPerson | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpsertUser(context.Background()).IamPerson(iamPerson).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpsertUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertUser`: IamReply
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpsertUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamPerson** | [**IamPerson**](IamPerson.md) |  | 
 **authorization** | **string** |  | 

### Return type

[**IamReply**](IamReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

