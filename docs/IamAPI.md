# \IamAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProvider**](IamAPI.md#AddProvider) | **Post** /v1/iam/providers | Adds an identity provider your people can sign in with, or a service your applications send through — a social or enterprise login, an email or SMS sender, a storage or payment connector.
[**AddToken**](IamAPI.md#AddToken) | **Post** /v1/iam/tokens | Records an access token — the credential an application or integration presents on a caller&#39;s behalf.
[**AddWebauthnCredential**](IamAPI.md#AddWebauthnCredential) | **Post** /v1/iam/webauthn-credentials | Registers a passkey or security key for a person, so they can sign in with their device instead of a password.
[**CreateOrganization**](IamAPI.md#CreateOrganization) | **Post** /v1/iam/organizations | Makes a new organization — the account your users, applications, roles, projects and workspaces are all named inside.
[**CreateSession**](IamAPI.md#CreateSession) | **Post** /v1/iam/sessions | Records a sign-in.
[**DeleteIamApplicationsByOwnerByName**](IamAPI.md#DeleteIamApplicationsByOwnerByName) | **Delete** /v1/iam/applications/{owner}/{name} | Removes an application.
[**DeleteIamAuditLogsByOwnerByName**](IamAPI.md#DeleteIamAuditLogsByOwnerByName) | **Delete** /v1/iam/audit-logs/{owner}/{name} | Removes an audit entry.
[**DeleteIamCertsByOwnerByName**](IamAPI.md#DeleteIamCertsByOwnerByName) | **Delete** /v1/iam/certs/{owner}/{name} | Removes a signing certificate.
[**DeleteIamInvitationsByOwnerByName**](IamAPI.md#DeleteIamInvitationsByOwnerByName) | **Delete** /v1/iam/invitations/{owner}/{name} | Withdraws an invitation.
[**DeleteIamKeysByOwnerByName**](IamAPI.md#DeleteIamKeysByOwnerByName) | **Delete** /v1/iam/keys/{owner}/{name} | Revokes an API key.
[**DeleteIamMfa**](IamAPI.md#DeleteIamMfa) | **Delete** /v1/iam/mfa | Turns a factor off, so sign-in stops asking for it.
[**DeleteIamPermissionsByOwnerByName**](IamAPI.md#DeleteIamPermissionsByOwnerByName) | **Delete** /v1/iam/permissions/{owner}/{name} | Revokes a permission.
[**DeleteIamProjectsByOwnerByName**](IamAPI.md#DeleteIamProjectsByOwnerByName) | **Delete** /v1/iam/projects/{owner}/{name} | Removes a project.
[**DeleteIamRolesByOwnerByName**](IamAPI.md#DeleteIamRolesByOwnerByName) | **Delete** /v1/iam/roles/{owner}/{name} | Removes a role.
[**DeleteIamScimV2UsersByOwnerByName**](IamAPI.md#DeleteIamScimV2UsersByOwnerByName) | **Delete** /v1/iam/scim/v2/Users/{owner}/{name} | Deprovisions a person — how removing someone in your identity provider removes their access here.
[**DeleteIamServiceAccountsByName**](IamAPI.md#DeleteIamServiceAccountsByName) | **Delete** /v1/iam/service-accounts/{name} | Serves DELETE /v1/iam/service-accounts/:name.
[**DeleteIamUsersByOwnerByName**](IamAPI.md#DeleteIamUsersByOwnerByName) | **Delete** /v1/iam/users/{owner}/{name} | Removes a person from your organization.
[**DeleteIamUsersByOwnerByNameKeys**](IamAPI.md#DeleteIamUsersByOwnerByNameKeys) | **Delete** /v1/iam/users/{owner}/{name}/keys | Clears the target user&#39;s key of the requested TYPE (immediate revoke).
[**DeleteIamWorkspacesByOwnerByName**](IamAPI.md#DeleteIamWorkspacesByOwnerByName) | **Delete** /v1/iam/workspaces/{owner}/{name} | Removes a workspace.
[**DeleteOrganization**](IamAPI.md#DeleteOrganization) | **Delete** /v1/iam/organizations/{owner}/{name} | Removes an organization and everything named inside it.
[**DeleteProvider**](IamAPI.md#DeleteProvider) | **Delete** /v1/iam/providers/{owner}/{name} | Removes a provider.
[**DeleteSession**](IamAPI.md#DeleteSession) | **Delete** /v1/iam/sessions/{owner}/{name}/{application} | Signs a person out of one application — the session ends and every browser carrying it stops being authenticated.
[**DeleteToken**](IamAPI.md#DeleteToken) | **Delete** /v1/iam/tokens/{owner}/{name} | Revokes an access token.
[**DeleteWebauthnCredential**](IamAPI.md#DeleteWebauthnCredential) | **Delete** /v1/iam/webauthn-credentials/{owner}/{name} | Removes a passkey or security key — what you call when a device is lost.
[**GetIamAccount**](IamAPI.md#GetIamAccount) | **Get** /v1/iam/account | Returns the signed-in person&#39;s own account and the organization they belong to — what a console reads to draw the account menu.
[**GetIamApplications**](IamAPI.md#GetIamApplications) | **Get** /v1/iam/applications | Returns the applications in one organization, newest first — each product or site your people sign in to, with the sign-in methods and redirect URIs it allows.
[**GetIamApplicationsByOwnerByName**](IamAPI.md#GetIamApplicationsByOwnerByName) | **Get** /v1/iam/applications/{owner}/{name} | Returns one application: its sign-in methods, its allowed redirect URIs and the client credentials your integration authenticates with.
[**GetIamAuditLogs**](IamAPI.md#GetIamAuditLogs) | **Get** /v1/iam/audit-logs | Returns your organization&#39;s audit trail, newest first — who did what, when, and from where.
[**GetIamAuditLogsByOwnerByName**](IamAPI.md#GetIamAuditLogsByOwnerByName) | **Get** /v1/iam/audit-logs/{owner}/{name} | Returns one audit entry in full: the action, the person or key behind it, and the request it came in on.
[**GetIamAuthApplication**](IamAPI.md#GetIamAuthApplication) | **Get** /v1/iam/auth/application | Returns everything a login screen needs to draw itself for one application: its branding, and each sign-in method it offers with the provider details that method needs.
[**GetIamAuthMethods**](IamAPI.md#GetIamAuthMethods) | **Get** /v1/iam/auth/methods | Returns the sign-in methods one application actually has switched on, so a login screen can render the right buttons for it without you hard-coding a list that drifts the moment you add a provider.
[**GetIamCerts**](IamAPI.md#GetIamCerts) | **Get** /v1/iam/certs | Returns your organization&#39;s signing certificates, newest first — the keys the tokens your applications verify are signed with.
[**GetIamCertsByOwnerByName**](IamAPI.md#GetIamCertsByOwnerByName) | **Get** /v1/iam/certs/{owner}/{name} | Returns one signing certificate — its algorithm, its validity window and its public half.
[**GetIamConsent**](IamAPI.md#GetIamConsent) | **Get** /v1/iam/consent | Returns the calling person&#39;s own privacy and communication choices.
[**GetIamInvitations**](IamAPI.md#GetIamInvitations) | **Get** /v1/iam/invitations | Returns your organization&#39;s invitations, newest first — who has been asked to join, on what terms, and how many seats each invitation still has left.
[**GetIamInvitationsByOwnerByName**](IamAPI.md#GetIamInvitationsByOwnerByName) | **Get** /v1/iam/invitations/{owner}/{name} | Returns one invitation: who it is for, what it grants on acceptance, and when it expires.
[**GetIamKeys**](IamAPI.md#GetIamKeys) | **Get** /v1/iam/keys | Returns your organization&#39;s API keys, newest first — what each is called, what it may reach, and its publishable half.
[**GetIamKeysByOwnerByName**](IamAPI.md#GetIamKeysByOwnerByName) | **Get** /v1/iam/keys/{owner}/{name} | Returns one API key: what it is called, what it may reach, and when it was issued.
[**GetIamKeysOrg**](IamAPI.md#GetIamKeysOrg) | **Get** /v1/iam/keys/org | Resolve a PUBLISHABLE key to the organization that owns it
[**GetIamKeysPrincipal**](IamAPI.md#GetIamKeysPrincipal) | **Get** /v1/iam/keys/principal | Resolve a SECRET key to the principal it authenticates
[**GetIamLinkedAccounts**](IamAPI.md#GetIamLinkedAccounts) | **Get** /v1/iam/linked-accounts | Returns the sign-in identities linked to the calling person&#39;s account — every provider they can currently sign in with.
[**GetIamMemberships**](IamAPI.md#GetIamMemberships) | **Get** /v1/iam/memberships | Answers either question about who belongs where: which organizations one person can act in, or who can act in one organization.
[**GetIamOauthAuthorize**](IamAPI.md#GetIamOauthAuthorize) | **Get** /v1/iam/oauth/authorize | Starts a sign-in — the address you send a browser to, and the beginning of every OAuth and OpenID Connect flow.
[**GetIamOauthCallback**](IamAPI.md#GetIamOauthCallback) | **Get** /v1/iam/oauth/callback | Completes the round-trip: it resolves and burns the single-use transaction (checking expiry + browser binding), exchanges and verifies the IdP response, links or provisions the local user, and mints the iam authorization code the relying party expects — then redirects to the original redirect_uri with code + state.
[**GetIamOauthLogout**](IamAPI.md#GetIamOauthLogout) | **Get** /v1/iam/oauth/logout | Ends a sign-in and sends the browser somewhere sensible.
[**GetIamOauthUserinfo**](IamAPI.md#GetIamOauthUserinfo) | **Get** /v1/iam/oauth/userinfo | Returns the profile claims for whoever the access token belongs to — the standard OpenID Connect way to find out who is calling you without your application storing anything itself.
[**GetIamPermissions**](IamAPI.md#GetIamPermissions) | **Get** /v1/iam/permissions | Returns the permissions in one organization, newest first — each one a grant saying which people or roles may do what, and to which resources.
[**GetIamPermissionsByOwnerByName**](IamAPI.md#GetIamPermissionsByOwnerByName) | **Get** /v1/iam/permissions/{owner}/{name} | Returns one permission: who it grants to, what it allows, and the resources it covers.
[**GetIamProjects**](IamAPI.md#GetIamProjects) | **Get** /v1/iam/projects | Returns your organization&#39;s projects, newest first — the scope people pick between when their work is separated by product or client rather than by team.
[**GetIamProjectsByOwnerByName**](IamAPI.md#GetIamProjectsByOwnerByName) | **Get** /v1/iam/projects/{owner}/{name} | Returns one project: what it is called and how it is set up.
[**GetIamRegistryJwks**](IamAPI.md#GetIamRegistryJwks) | **Get** /v1/iam/registry/jwks | Publishes the public key your registry uses to verify the tokens issued above — the one URL to configure so the registry trusts logins without holding any secret of its own.
[**GetIamRegistryToken**](IamAPI.md#GetIamRegistryToken) | **Get** /v1/iam/registry/token | Signs a container client in to your registry.
[**GetIamRoles**](IamAPI.md#GetIamRoles) | **Get** /v1/iam/roles | Returns your organization&#39;s roles, newest first — each a named group of people that permissions are granted to.
[**GetIamRolesByOwnerByName**](IamAPI.md#GetIamRolesByOwnerByName) | **Get** /v1/iam/roles/{owner}/{name} | Returns one role: who is in it, and the roles it includes.
[**GetIamScimV2Resourcetypes**](IamAPI.md#GetIamScimV2Resourcetypes) | **Get** /v1/iam/scim/v2/ResourceTypes | Returns the kinds of record this directory provisions and the address of each, so your identity provider discovers them rather than having them configured by hand.
[**GetIamScimV2ResourcetypesByName**](IamAPI.md#GetIamScimV2ResourcetypesByName) | **Get** /v1/iam/scim/v2/ResourceTypes/{name} | Returns one provisionable record kind in full.
[**GetIamScimV2Schemas**](IamAPI.md#GetIamScimV2Schemas) | **Get** /v1/iam/scim/v2/Schemas | Returns the attribute definitions this directory understands, so your identity provider knows which fields it may send and what they mean before it sends any.
[**GetIamScimV2SchemasById**](IamAPI.md#GetIamScimV2SchemasById) | **Get** /v1/iam/scim/v2/Schemas/{id} | Returns one attribute definition in full.
[**GetIamScimV2Serviceproviderconfig**](IamAPI.md#GetIamScimV2Serviceproviderconfig) | **Get** /v1/iam/scim/v2/ServiceProviderConfig | Tells your identity provider which parts of SCIM this directory supports, so it configures itself instead of you filling in a form.
[**GetIamScimV2Users**](IamAPI.md#GetIamScimV2Users) | **Get** /v1/iam/scim/v2/Users | Returns the people in your organization to your identity provider, in the standard SCIM shape, so an IdP can reconcile its directory against ours.
[**GetIamScimV2UsersByOwnerByName**](IamAPI.md#GetIamScimV2UsersByOwnerByName) | **Get** /v1/iam/scim/v2/Users/{owner}/{name} | Returns one person in the standard SCIM shape.
[**GetIamServiceAccounts**](IamAPI.md#GetIamServiceAccounts) | **Get** /v1/iam/service-accounts | Returns your organization&#39;s service accounts — what each is called and when it was created.
[**GetIamUsers**](IamAPI.md#GetIamUsers) | **Get** /v1/iam/users | Returns a page of the people in your organization, with the total so you can page through the rest.
[**GetIamUsersByOwnerByName**](IamAPI.md#GetIamUsersByOwnerByName) | **Get** /v1/iam/users/{owner}/{name} | Returns one person in your organization, addressed by their username or by their email address.
[**GetIamWeb3Nonce**](IamAPI.md#GetIamWeb3Nonce) | **Get** /v1/iam/web3/nonce | Starts a wallet sign-in: it returns a one-time challenge for the wallet to sign.
[**GetIamWebauthnSigninBegin**](IamAPI.md#GetIamWebauthnSigninBegin) | **Get** /v1/iam/webauthn/signin/begin | Starts a passkey sign-in: it returns the challenge the person&#39;s authenticator signs.
[**GetIamWebauthnSignupBegin**](IamAPI.md#GetIamWebauthnSignupBegin) | **Get** /v1/iam/webauthn/signup/begin | Starts enrolling a passkey for the signed-in person: it returns the options their browser hands to the authenticator.
[**GetIamWellKnownJwks**](IamAPI.md#GetIamWellKnownJwks) | **Get** /v1/iam/.well-known/jwks | Publishes the public keys that verify the tokens issued here — the one URL you point a service at so it can check a token itself, offline, without calling back and without holding any secret of ours.
[**GetIamWellKnownOauthAuthorizationServer**](IamAPI.md#GetIamWellKnownOauthAuthorizationServer) | **Get** /v1/iam/.well-known/oauth-authorization-server | Returns the OpenID Connect discovery document — the one URL you point a standards-compliant client at so it can find every other endpoint on its own, instead of you configuring them by hand.
[**GetIamWellKnownOpenidConfiguration**](IamAPI.md#GetIamWellKnownOpenidConfiguration) | **Get** /v1/iam/.well-known/openid-configuration | Returns the OpenID Connect discovery document — the one URL you point a standards-compliant client at so it can find every other endpoint on its own, instead of you configuring them by hand.
[**GetIamWhoami**](IamAPI.md#GetIamWhoami) | **Get** /v1/iam/whoami | Tells you who the current caller is — the lightweight check a page makes on load to decide whether to render signed-in or signed-out.
[**GetIamWorkspaces**](IamAPI.md#GetIamWorkspaces) | **Get** /v1/iam/workspaces | Returns your organization&#39;s workspaces, newest first — the scope a team works in, alongside projects rather than instead of them.
[**GetIamWorkspacesByOwnerByName**](IamAPI.md#GetIamWorkspacesByOwnerByName) | **Get** /v1/iam/workspaces/{owner}/{name} | Returns one workspace: what it is called and how it is set up.
[**GetOrganization**](IamAPI.md#GetOrganization) | **Get** /v1/iam/organizations/{owner}/{name} | Returns one organization: its display, its defaults and the sign-in rules everyone in it inherits.
[**GetProvider**](IamAPI.md#GetProvider) | **Get** /v1/iam/providers/{owner}/{name} | Returns one provider: what it connects to and how it is configured.
[**GetSession**](IamAPI.md#GetSession) | **Get** /v1/iam/sessions/{owner}/{name}/{application} | Returns one person&#39;s session in one application — when it began and which browsers or devices are still carrying it.
[**GetToken**](IamAPI.md#GetToken) | **Get** /v1/iam/tokens/{owner}/{name} | Returns one access token: who and what it was issued to, and when it expires.
[**GetWebauthnCredential**](IamAPI.md#GetWebauthnCredential) | **Get** /v1/iam/webauthn-credentials/{owner}/{name} | Returns one passkey or security key: whose it is, what device it lives on, and when it was registered.
[**ListOrganizations**](IamAPI.md#ListOrganizations) | **Get** /v1/iam/organizations | Returns the organizations you can act in, the ones you belong to first and the rest after, newest first, narrowed by an optional query against the name or the display name.
[**ListProviders**](IamAPI.md#ListProviders) | **Get** /v1/iam/providers | Returns your organization&#39;s providers, newest first — the identity providers your people sign in with, and the senders and connectors your applications go through.
[**ListSessions**](IamAPI.md#ListSessions) | **Get** /v1/iam/sessions | Returns who is currently signed in to your organization, newest first, and can be narrowed to one person or one application.
[**ListTokens**](IamAPI.md#ListTokens) | **Get** /v1/iam/tokens | Returns the access tokens issued in your organization, newest first, and can be narrowed to one organization.
[**ListWebauthnCredentials**](IamAPI.md#ListWebauthnCredentials) | **Get** /v1/iam/webauthn-credentials | Returns the passkeys and security keys registered to one person, newest first — which device each lives on and when it was registered.
[**PatchIamScimV2UsersByOwnerByName**](IamAPI.md#PatchIamScimV2UsersByOwnerByName) | **Patch** /v1/iam/scim/v2/Users/{owner}/{name} | Applies a partial change from your identity provider — one attribute moved, not the whole record resent.
[**PostIamAdminProvision**](IamAPI.md#PostIamAdminProvision) | **Post** /v1/iam/admin/provision | Sets up an account on someone&#39;s behalf — the same onboarding a person gets themselves, driven by one of your own services instead of by them.
[**PostIamApplications**](IamAPI.md#PostIamApplications) | **Post** /v1/iam/applications | Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.
[**PostIamAssume**](IamAPI.md#PostIamAssume) | **Post** /v1/iam/assume | Steps a platform operator into an organization: it returns their own access token re-scoped to that tenant, so they see what the tenant sees.
[**PostIamAuditLogs**](IamAPI.md#PostIamAuditLogs) | **Post** /v1/iam/audit-logs | Records an audit entry, so activity from your own systems lands in the same trail as everything the Hanzo Cloud records for you.
[**PostIamCerts**](IamAPI.md#PostIamCerts) | **Post** /v1/iam/certs | Adds a signing certificate your applications can verify tokens against — the call you make to stage the next one before a rotation.
[**PostIamDeleteMembership**](IamAPI.md#PostIamDeleteMembership) | **Post** /v1/iam/delete-membership | Takes away a person&#39;s or an application&#39;s right to act in an organization.
[**PostIamInvitations**](IamAPI.md#PostIamInvitations) | **Post** /v1/iam/invitations | Issues an invitation to join your organization — the code or link a new member redeems, with the role they arrive holding and the date it stops working.
[**PostIamKeys**](IamAPI.md#PostIamKeys) | **Post** /v1/iam/keys | Issues an API key.
[**PostIamLink**](IamAPI.md#PostIamLink) | **Post** /v1/iam/link | Starts connecting another sign-in identity to the account you are already signed in as.
[**PostIamLogin**](IamAPI.md#PostIamLogin) | **Post** /v1/iam/login | Signs a person in with the credential they typed, and — when the request is part of an OAuth flow — hands back the one-time code that finishes it.
[**PostIamMemberships**](IamAPI.md#PostIamMemberships) | **Post** /v1/iam/memberships | Lets a person or an application act in an organization.
[**PostIamMfaPreferred**](IamAPI.md#PostIamMfaPreferred) | **Post** /v1/iam/mfa/preferred | Picks which second factor an account is asked for first when it has more than one.
[**PostIamMfaSetupEnable**](IamAPI.md#PostIamMfaSetupEnable) | **Post** /v1/iam/mfa/setup/enable | Finishes the enrolment: from here the account&#39;s sign-ins ask for this factor.
[**PostIamMfaSetupInitiate**](IamAPI.md#PostIamMfaSetupInitiate) | **Post** /v1/iam/mfa/setup/initiate | Starts enrolling a factor and hands over whatever the person needs to prove they hold it: app a fresh secret and the otpauth:// URL to render as a QR code sms a code texted to the number on the account email a code mailed to the address on the account Nothing is switched on yet, so abandoning this step leaves the account exactly as it was.
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
[**PostIamPreferences**](IamAPI.md#PostIamPreferences) | **Post** /v1/iam/preferences | Saves the calling person&#39;s own settings and returns the full set afterwards.
[**PostIamProjects**](IamAPI.md#PostIamProjects) | **Post** /v1/iam/projects | Makes a project inside your organization — the scope people pick between when their work is separated by product or client rather than by team.
[**PostIamRegistryToken**](IamAPI.md#PostIamRegistryToken) | **Post** /v1/iam/registry/token | Signs a container client in to your registry.
[**PostIamRelease**](IamAPI.md#PostIamRelease) | **Post** /v1/iam/release | Steps a platform operator back out: it returns their own access token with no organization assumed, which is the credential they had before they stepped in.
[**PostIamRoles**](IamAPI.md#PostIamRoles) | **Post** /v1/iam/roles | Makes a role — a named group of people that permissions are granted to.
[**PostIamScimV2Users**](IamAPI.md#PostIamScimV2Users) | **Post** /v1/iam/scim/v2/Users | Provisions a person from your identity provider — how a new hire gets an account here automatically when they are added over there.
[**PostIamServiceAccounts**](IamAPI.md#PostIamServiceAccounts) | **Post** /v1/iam/service-accounts | Makes a service account — an identity for a program rather than a person, for a script, a bot or a deployment that has to authenticate on its own.
[**PostIamServiceAccountsByNameKeys**](IamAPI.md#PostIamServiceAccountsByNameKeys) | **Post** /v1/iam/service-accounts/{name}/keys | Serves POST /v1/iam/service-accounts/:name/keys: mint a fresh key, invalidating the prior one, and return the new raw secret exactly once.
[**PostIamSignin**](IamAPI.md#PostIamSignin) | **Post** /v1/iam/signin | Completes a sign-in: it exchanges the one-time code your application was handed at the end of the login flow for a live session, and returns the signed-in account.
[**PostIamSignup**](IamAPI.md#PostIamSignup) | **Post** /v1/iam/signup | Creates an account from the sign-up form and applies the application&#39;s own sign-up rules — whether self-service registration is open at all, and which fields it requires.
[**PostIamTokensIssue**](IamAPI.md#PostIamTokensIssue) | **Post** /v1/iam/tokens/issue | Mints an access token for the &#x60;?id&#x3D;&lt;owner&gt;/&lt;name&gt;&#x60; target user (optional &#x60;?aud&#x3D;&#x60; resource, RFC 8707), issued by the authenticated + allow-listed confidential client.
[**PostIamUnlink**](IamAPI.md#PostIamUnlink) | **Post** /v1/iam/unlink | Disconnects one sign-in identity from an account, so that provider can no longer be used to sign in as that person.
[**PostIamUsers**](IamAPI.md#PostIamUsers) | **Post** /v1/iam/users | Adds a person to your organization.
[**PostIamUsersByOwnerByNameKeys**](IamAPI.md#PostIamUsersByOwnerByNameKeys) | **Post** /v1/iam/users/{owner}/{name}/keys | (re)generates the target user&#39;s key of the requested TYPE and returns it once, over the shared authorizeMinter + mintTarget seam.
[**PostIamVerificationCodes**](IamAPI.md#PostIamVerificationCodes) | **Post** /v1/iam/verification-codes | Validates the request and asks otp to get a code to the person.
[**PostIamWeb3Verify**](IamAPI.md#PostIamWeb3Verify) | **Post** /v1/iam/web3/verify | Completes a wallet sign-in: it verifies the signed challenge and, if it holds, signs the wallet&#39;s owner in.
[**PostIamWebauthnSigninFinish**](IamAPI.md#PostIamWebauthnSigninFinish) | **Post** /v1/iam/webauthn/signin/finish | Verifies the signed challenge and signs the person in.
[**PostIamWebauthnSignupFinish**](IamAPI.md#PostIamWebauthnSignupFinish) | **Post** /v1/iam/webauthn/signup/finish | Verifies the newly created passkey and stores it, so the person can sign in with their device from then on.
[**PostIamWorkspaces**](IamAPI.md#PostIamWorkspaces) | **Post** /v1/iam/workspaces | Makes a workspace inside your organization — the scope a team works in, alongside projects rather than instead of them.
[**PutIamAccount**](IamAPI.md#PutIamAccount) | **Put** /v1/iam/account | Saves the calling person&#39;s own profile — the name they are shown by, their picture, a line about themselves and a link.
[**PutIamApplicationsByOwnerByName**](IamAPI.md#PutIamApplicationsByOwnerByName) | **Put** /v1/iam/applications/{owner}/{name} | Changes an application&#39;s display, its sign-in methods and the redirect URIs it may return to — the call that makes login work from a new host.
[**PutIamAuditLogsByOwnerByName**](IamAPI.md#PutIamAuditLogsByOwnerByName) | **Put** /v1/iam/audit-logs/{owner}/{name} | Corrects an audit entry.
[**PutIamCertsByOwnerByName**](IamAPI.md#PutIamCertsByOwnerByName) | **Put** /v1/iam/certs/{owner}/{name} | Changes a signing certificate&#39;s settings.
[**PutIamConsent**](IamAPI.md#PutIamConsent) | **Put** /v1/iam/consent | Records the calling person&#39;s privacy and communication choices.
[**PutIamInvitationsByOwnerByName**](IamAPI.md#PutIamInvitationsByOwnerByName) | **Put** /v1/iam/invitations/{owner}/{name} | Changes an invitation&#39;s terms — the role it grants, how many may redeem it, or when it expires.
[**PutIamKeysByOwnerByName**](IamAPI.md#PutIamKeysByOwnerByName) | **Put** /v1/iam/keys/{owner}/{name} | Changes what a key is called or what it may reach.
[**PutIamPassword**](IamAPI.md#PutIamPassword) | **Put** /v1/iam/password | Replaces the calling person&#39;s password.
[**PutIamPermissionsByOwnerByName**](IamAPI.md#PutIamPermissionsByOwnerByName) | **Put** /v1/iam/permissions/{owner}/{name} | Changes who a permission grants to, what it allows, or the resources it covers.
[**PutIamProjectsByOwnerByName**](IamAPI.md#PutIamProjectsByOwnerByName) | **Put** /v1/iam/projects/{owner}/{name} | Changes a project&#39;s settings.
[**PutIamRolesByOwnerByName**](IamAPI.md#PutIamRolesByOwnerByName) | **Put** /v1/iam/roles/{owner}/{name} | Changes who is in a role, or which roles it includes.
[**PutIamScimV2UsersByOwnerByName**](IamAPI.md#PutIamScimV2UsersByOwnerByName) | **Put** /v1/iam/scim/v2/Users/{owner}/{name} | Overwrites a person&#39;s SCIM attributes with what your identity provider sends — how a change made there lands here.
[**PutIamUsersByOwnerByName**](IamAPI.md#PutIamUsersByOwnerByName) | **Put** /v1/iam/users/{owner}/{name} | Changes a person&#39;s profile, their roles, or the credentials they sign in with.
[**PutIamWorkspacesByOwnerByName**](IamAPI.md#PutIamWorkspacesByOwnerByName) | **Put** /v1/iam/workspaces/{owner}/{name} | Changes a workspace&#39;s settings.
[**SetOrganizationAvatar**](IamAPI.md#SetOrganizationAvatar) | **Post** /v1/iam/organizations/avatar | Changes how an organization appears across Hanzo: the square mark beside its name, as an uploaded image or as a single emoji.
[**UpdateOrganization**](IamAPI.md#UpdateOrganization) | **Put** /v1/iam/organizations/{owner}/{name} | Changes an organization&#39;s display, its defaults and the sign-in rules everyone in it inherits.
[**UpdateProvider**](IamAPI.md#UpdateProvider) | **Put** /v1/iam/providers/{owner}/{name} | Changes a provider&#39;s settings or rotates the credentials it holds.
[**UpdateSession**](IamAPI.md#UpdateSession) | **Put** /v1/iam/sessions/{owner}/{name}/{application} | Replaces the set of browsers a session covers — signing out the ones you leave off while the session itself stays live.
[**UpdateToken**](IamAPI.md#UpdateToken) | **Put** /v1/iam/tokens/{owner}/{name} | Changes an access token&#39;s scope or expiry.
[**UpdateWebauthnCredential**](IamAPI.md#UpdateWebauthnCredential) | **Put** /v1/iam/webauthn-credentials/{owner}/{name} | Renames a registered passkey or security key, so a person can tell their devices apart.
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## DeleteIamApplicationsByOwnerByName

> IamDeleteResult DeleteIamApplicationsByOwnerByName(ctx, owner, name).Execute()

Removes an application.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamApplicationsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamApplicationsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamApplicationsByOwnerByName`: IamDeleteResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamApplicationsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamApplicationsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DeleteIamAuditLogsByOwnerByName

> IamDeleteOutput DeleteIamAuditLogsByOwnerByName(ctx, owner, name).Execute()

Removes an audit entry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamAuditLogsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamAuditLogsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamAuditLogsByOwnerByName`: IamDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamAuditLogsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamAuditLogsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamDeleteOutput**](IamDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamCertsByOwnerByName

> IamCertsDeleteOutput DeleteIamCertsByOwnerByName(ctx, owner, name).Execute()

Removes a signing certificate.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamCertsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamCertsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamCertsByOwnerByName`: IamCertsDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamCertsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamCertsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamCertsDeleteOutput**](IamCertsDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamInvitationsByOwnerByName

> IamInvitationsDeleteOutput DeleteIamInvitationsByOwnerByName(ctx, owner, name).Execute()

Withdraws an invitation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamInvitationsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamInvitationsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamInvitationsByOwnerByName`: IamInvitationsDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamInvitationsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamInvitationsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamInvitationsDeleteOutput**](IamInvitationsDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamKeysByOwnerByName

> IamDeleteResponse DeleteIamKeysByOwnerByName(ctx, owner, name).Execute()

Revokes an API key.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamKeysByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamKeysByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamKeysByOwnerByName`: IamDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamKeysByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamKeysByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamDeleteResponse**](IamDeleteResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamMfa

> DeleteIamMfa(ctx).Execute()

Turns a factor off, so sign-in stops asking for it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.DeleteIamMfa(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamMfaRequest struct via the builder pattern


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


## DeleteIamPermissionsByOwnerByName

> IamPermissionDeleteResponse DeleteIamPermissionsByOwnerByName(ctx, owner, name).Execute()

Revokes a permission.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamPermissionsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamPermissionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamPermissionsByOwnerByName`: IamPermissionDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamPermissionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamPermissionsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamPermissionDeleteResponse**](IamPermissionDeleteResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamProjectsByOwnerByName

> IamProjectsDeleteOutput DeleteIamProjectsByOwnerByName(ctx, owner, name).Execute()

Removes a project.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamProjectsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamProjectsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamProjectsByOwnerByName`: IamProjectsDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamProjectsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamProjectsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamProjectsDeleteOutput**](IamProjectsDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamRolesByOwnerByName

> IamRolesDeleteOutput DeleteIamRolesByOwnerByName(ctx, owner, name).Execute()

Removes a role.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamRolesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamRolesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamRolesByOwnerByName`: IamRolesDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamRolesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamRolesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamRolesDeleteOutput**](IamRolesDeleteOutput.md)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## DeleteIamUsersByOwnerByName

> IamUsersDeleteOutput DeleteIamUsersByOwnerByName(ctx, owner, name).Execute()

Removes a person from your organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamUsersByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamUsersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamUsersByOwnerByName`: IamUsersDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamUsersByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamUsersByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamUsersDeleteOutput**](IamUsersDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamUsersByOwnerByNameKeys

> DeleteIamUsersByOwnerByNameKeys(ctx, owner, name).Execute()

Clears the target user's key of the requested TYPE (immediate revoke).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.DeleteIamUsersByOwnerByNameKeys(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamUsersByOwnerByNameKeys``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIamUsersByOwnerByNameKeysRequest struct via the builder pattern


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


## DeleteIamWorkspacesByOwnerByName

> IamWorkspacesDeleteOutput DeleteIamWorkspacesByOwnerByName(ctx, owner, name).Execute()

Removes a workspace.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteIamWorkspacesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIamWorkspacesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamWorkspacesByOwnerByName`: IamWorkspacesDeleteOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteIamWorkspacesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamWorkspacesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamWorkspacesDeleteOutput**](IamWorkspacesDeleteOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> IamDeleteOrganizationOutput DeleteOrganization(ctx, owner, name).Execute()

Removes an organization and everything named inside it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteOrganization(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganization`: IamDeleteOrganizationOutput
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamDeleteOrganizationOutput**](IamDeleteOrganizationOutput.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProvider

> IamMutationResult DeleteProvider(ctx, owner, name).Execute()

Removes a provider.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteProvider(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProvider`: IamMutationResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamMutationResult**](IamMutationResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSession

> IamDeleteSessionOut DeleteSession(ctx, owner, name, application).Execute()

Signs a person out of one application — the session ends and every browser carrying it stops being authenticated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	application := "application_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteSession(context.Background(), owner, name, application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSession`: IamDeleteSessionOut
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 
**application** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamDeleteSessionOut**](IamDeleteSessionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> IamTokenMutation DeleteToken(ctx, owner, name).Execute()

Revokes an access token.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteToken(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteToken`: IamTokenMutation
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamTokenMutation**](IamTokenMutation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebauthnCredential

> IamWebauthnCredentialMutationResult DeleteWebauthnCredential(ctx, owner, name).Execute()

Removes a passkey or security key — what you call when a device is lost.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.DeleteWebauthnCredential(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteWebauthnCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebauthnCredential`: IamWebauthnCredentialMutationResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.DeleteWebauthnCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebauthnCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamWebauthnCredentialMutationResult**](IamWebauthnCredentialMutationResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## GetIamApplicationsByOwnerByName

> IamApplication GetIamApplicationsByOwnerByName(ctx, owner, name).Execute()

Returns one application: its sign-in methods, its allowed redirect URIs and the client credentials your integration authenticates with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamApplicationsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamApplicationsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamApplicationsByOwnerByName`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamApplicationsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamApplicationsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## GetIamAuditLogsByOwnerByName

> IamAuditLog GetIamAuditLogsByOwnerByName(ctx, owner, name).Execute()

Returns one audit entry in full: the action, the person or key behind it, and the request it came in on.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamAuditLogsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamAuditLogsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamAuditLogsByOwnerByName`: IamAuditLog
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamAuditLogsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamAuditLogsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamAuditLog**](IamAuditLog.md)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## GetIamCertsByOwnerByName

> IamCert GetIamCertsByOwnerByName(ctx, owner, name).Execute()

Returns one signing certificate — its algorithm, its validity window and its public half.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamCertsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamCertsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamCertsByOwnerByName`: IamCert
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamCertsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamCertsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamCert**](IamCert.md)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## GetIamInvitationsByOwnerByName

> IamInvitation GetIamInvitationsByOwnerByName(ctx, owner, name).Execute()

Returns one invitation: who it is for, what it grants on acceptance, and when it expires.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamInvitationsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamInvitationsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamInvitationsByOwnerByName`: IamInvitation
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamInvitationsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamInvitationsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamInvitation**](IamInvitation.md)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## GetIamKeysByOwnerByName

> IamKey GetIamKeysByOwnerByName(ctx, owner, name).Execute()

Returns one API key: what it is called, what it may reach, and when it was issued.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamKeysByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamKeysByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamKeysByOwnerByName`: IamKey
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamKeysByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamKeysByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetIamKeysOrg

> GetIamKeysOrg(ctx).Execute()

Resolve a PUBLISHABLE key to the organization that owns it



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamKeysOrg(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamKeysOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamKeysOrgRequest struct via the builder pattern


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


## GetIamKeysPrincipal

> GetIamKeysPrincipal(ctx).Execute()

Resolve a SECRET key to the principal it authenticates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.GetIamKeysPrincipal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamKeysPrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamKeysPrincipalRequest struct via the builder pattern


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## GetIamPermissionsByOwnerByName

> IamPermission GetIamPermissionsByOwnerByName(ctx, owner, name).Execute()

Returns one permission: who it grants to, what it allows, and the resources it covers.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamPermissionsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamPermissionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamPermissionsByOwnerByName`: IamPermission
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamPermissionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamPermissionsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## GetIamProjectsByOwnerByName

> IamProject GetIamProjectsByOwnerByName(ctx, owner, name).Execute()

Returns one project: what it is called and how it is set up.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamProjectsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamProjectsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamProjectsByOwnerByName`: IamProject
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamProjectsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamProjectsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamProject**](IamProject.md)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## GetIamRolesByOwnerByName

> IamRole GetIamRolesByOwnerByName(ctx, owner, name).Execute()

Returns one role: who is in it, and the roles it includes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamRolesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamRolesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamRolesByOwnerByName`: IamRole
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamRolesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamRolesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamRole**](IamRole.md)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

> IamUsersListOutput GetIamUsers(ctx).Owner(owner).Email(email).Limit(limit).Offset(offset).Execute()

Returns a page of the people in your organization, with the total so you can page through the rest.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	email := "email_example" // string | Email narrows the page to the accounts carrying one address. Looking a person up by their address is a QUERY over the collection, not an item read: an address is not the natural key, two rows in one org can carry one, and a caller that gets a page SEES both — where a single-item read would have to choose, and choosing is how somebody joins a team under a colleague's identity. (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamUsers(context.Background()).Owner(owner).Email(email).Limit(limit).Offset(offset).Execute()
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
 **email** | **string** | Email narrows the page to the accounts carrying one address. Looking a person up by their address is a QUERY over the collection, not an item read: an address is not the natural key, two rows in one org can carry one, and a caller that gets a page SEES both — where a single-item read would have to choose, and choosing is how somebody joins a team under a colleague&#39;s identity. | 
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


## GetIamUsersByOwnerByName

> IamUser GetIamUsersByOwnerByName(ctx, owner, name).Email(email).Execute()

Returns one person in your organization, addressed by their username or by their email address.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	email := "email_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamUsersByOwnerByName(context.Background(), owner, name).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamUsersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamUsersByOwnerByName`: IamUser
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamUsersByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamUsersByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## GetIamWorkspacesByOwnerByName

> IamWorkspace GetIamWorkspacesByOwnerByName(ctx, owner, name).Execute()

Returns one workspace: what it is called and how it is set up.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIamWorkspacesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIamWorkspacesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamWorkspacesByOwnerByName`: IamWorkspace
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIamWorkspacesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamWorkspacesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamWorkspace**](IamWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> IamOrganization GetOrganization(ctx, owner, name).Execute()

Returns one organization: its display, its defaults and the sign-in rules everyone in it inherits.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetOrganization(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganization`: IamOrganization
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> IamProviderResult GetProvider(ctx, owner, name).Execute()

Returns one provider: what it connects to and how it is configured.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetProvider(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvider`: IamProviderResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamProviderResult**](IamProviderResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSession

> IamSession GetSession(ctx, owner, name, application).Execute()

Returns one person's session in one application — when it began and which browsers or devices are still carrying it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	application := "application_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetSession(context.Background(), owner, name, application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSession`: IamSession
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 
**application** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamSession**](IamSession.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> IamTokenResult GetToken(ctx, owner, name).Execute()

Returns one access token: who and what it was issued to, and when it expires.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetToken(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToken`: IamTokenResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamTokenResult**](IamTokenResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebauthnCredential

> IamWebauthnCredentialResult GetWebauthnCredential(ctx, owner, name).Execute()

Returns one passkey or security key: whose it is, what device it lives on, and when it was registered.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetWebauthnCredential(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetWebauthnCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebauthnCredential`: IamWebauthnCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetWebauthnCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebauthnCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamWebauthnCredentialResult**](IamWebauthnCredentialResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizations

> IamListOrganizationsOutput ListOrganizations(ctx).XForwardedFor(xForwardedFor).Q(q).Limit(limit).Cursor(cursor).Execute()

Returns the organizations you can act in, the ones you belong to first and the rest after, newest first, narrowed by an optional query against the name or the display name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	xForwardedFor := "xForwardedFor_example" // string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListOrganizations(context.Background()).XForwardedFor(xForwardedFor).Q(q).Limit(limit).Cursor(cursor).Execute()
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
 **xForwardedFor** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

> IamListSessionsOut ListSessions(ctx).Owner(owner).Name(name).Application(application).Execute()

Returns who is currently signed in to your organization, newest first, and can be narrowed to one person or one application.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string |  (optional)
	application := "application_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListSessions(context.Background()).Owner(owner).Name(name).Application(application).Execute()
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
 **owner** | **string** |  | 
 **name** | **string** |  | 
 **application** | **string** |  | 

### Return type

[**IamListSessionsOut**](IamListSessionsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

> IamAuditLog PostIamAuditLogs(ctx).IamInput(iamInput).Execute()

Records an audit entry, so activity from your own systems lands in the same trail as everything the Hanzo Cloud records for you.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	iamInput := *openapiclient.NewIamInput() // IamInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamAuditLogs(context.Background()).IamInput(iamInput).Execute()
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
 **iamInput** | [**IamInput**](IamInput.md) |  | 

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

Adds a signing certificate your applications can verify tokens against — the call you make to stage the next one before a rotation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

> IamProject PostIamProjects(ctx).IamProjectsInput(iamProjectsInput).Execute()

Makes a project inside your organization — the scope people pick between when their work is separated by product or client rather than by team.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	iamProjectsInput := *openapiclient.NewIamProjectsInput() // IamProjectsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PostIamProjects(context.Background()).IamProjectsInput(iamProjectsInput).Execute()
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
 **iamProjectsInput** | [**IamProjectsInput**](IamProjectsInput.md) |  | 

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## PostIamUsersByOwnerByNameKeys

> PostIamUsersByOwnerByNameKeys(ctx, owner, name).Execute()

(re)generates the target user's key of the requested TYPE and returns it once, over the shared authorizeMinter + mintTarget seam.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.PostIamUsersByOwnerByNameKeys(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PostIamUsersByOwnerByNameKeys``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostIamUsersByOwnerByNameKeysRequest struct via the builder pattern


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## PutIamApplicationsByOwnerByName

> IamApplication PutIamApplicationsByOwnerByName(ctx, owner, name).IamApplication(iamApplication).Execute()

Changes an application's display, its sign-in methods and the redirect URIs it may return to — the call that makes login work from a new host.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamApplicationsByOwnerByName(context.Background(), owner, name).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamApplicationsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamApplicationsByOwnerByName`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamApplicationsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamApplicationsByOwnerByNameRequest struct via the builder pattern


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


## PutIamAuditLogsByOwnerByName

> IamAuditLog PutIamAuditLogsByOwnerByName(ctx, owner, name).IamInput(iamInput).Execute()

Corrects an audit entry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamInput := *openapiclient.NewIamInput() // IamInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamAuditLogsByOwnerByName(context.Background(), owner, name).IamInput(iamInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamAuditLogsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamAuditLogsByOwnerByName`: IamAuditLog
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamAuditLogsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamAuditLogsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamInput** | [**IamInput**](IamInput.md) |  | 

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


## PutIamCertsByOwnerByName

> IamCert PutIamCertsByOwnerByName(ctx, owner, name).IamCert(iamCert).Execute()

Changes a signing certificate's settings.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamCert := *openapiclient.NewIamCert() // IamCert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamCertsByOwnerByName(context.Background(), owner, name).IamCert(iamCert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamCertsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamCertsByOwnerByName`: IamCert
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamCertsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamCertsByOwnerByNameRequest struct via the builder pattern


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## PutIamInvitationsByOwnerByName

> IamInvitation PutIamInvitationsByOwnerByName(ctx, owner, name).IamInvitationsInput(iamInvitationsInput).Execute()

Changes an invitation's terms — the role it grants, how many may redeem it, or when it expires.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamInvitationsInput := *openapiclient.NewIamInvitationsInput() // IamInvitationsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamInvitationsByOwnerByName(context.Background(), owner, name).IamInvitationsInput(iamInvitationsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamInvitationsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamInvitationsByOwnerByName`: IamInvitation
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamInvitationsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamInvitationsByOwnerByNameRequest struct via the builder pattern


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


## PutIamKeysByOwnerByName

> IamKey PutIamKeysByOwnerByName(ctx, owner, name).IamKey(iamKey).Execute()

Changes what a key is called or what it may reach.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | Owner is the tenant that holds the key; Name is unique within Owner.
	name := "name_example" // string | 
	iamKey := *openapiclient.NewIamKey() // IamKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamKeysByOwnerByName(context.Background(), owner, name).IamKey(iamKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamKeysByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamKeysByOwnerByName`: IamKey
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamKeysByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owner is the tenant that holds the key; Name is unique within Owner. | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamKeysByOwnerByNameRequest struct via the builder pattern


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## PutIamPermissionsByOwnerByName

> IamPermission PutIamPermissionsByOwnerByName(ctx, owner, name).IamPermission(iamPermission).Execute()

Changes who a permission grants to, what it allows, or the resources it covers.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | Identity — the (owner, name) natural key.
	name := "name_example" // string | 
	iamPermission := *openapiclient.NewIamPermission() // IamPermission | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamPermissionsByOwnerByName(context.Background(), owner, name).IamPermission(iamPermission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamPermissionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamPermissionsByOwnerByName`: IamPermission
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamPermissionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Identity — the (owner, name) natural key. | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamPermissionsByOwnerByNameRequest struct via the builder pattern


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


## PutIamProjectsByOwnerByName

> IamProject PutIamProjectsByOwnerByName(ctx, owner, name).IamProjectsInput(iamProjectsInput).Execute()

Changes a project's settings.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamProjectsInput := *openapiclient.NewIamProjectsInput() // IamProjectsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamProjectsByOwnerByName(context.Background(), owner, name).IamProjectsInput(iamProjectsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamProjectsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamProjectsByOwnerByName`: IamProject
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamProjectsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamProjectsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectsInput** | [**IamProjectsInput**](IamProjectsInput.md) |  | 

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


## PutIamRolesByOwnerByName

> IamRole PutIamRolesByOwnerByName(ctx, owner, name).IamRolesInput(iamRolesInput).Execute()

Changes who is in a role, or which roles it includes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamRolesInput := *openapiclient.NewIamRolesInput() // IamRolesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamRolesByOwnerByName(context.Background(), owner, name).IamRolesInput(iamRolesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamRolesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamRolesByOwnerByName`: IamRole
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamRolesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamRolesByOwnerByNameRequest struct via the builder pattern


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## PutIamUsersByOwnerByName

> IamUser PutIamUsersByOwnerByName(ctx, owner, name).IamUpdateInput(iamUpdateInput).Execute()

Changes a person's profile, their roles, or the credentials they sign in with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamUpdateInput := *openapiclient.NewIamUpdateInput() // IamUpdateInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamUsersByOwnerByName(context.Background(), owner, name).IamUpdateInput(iamUpdateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamUsersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamUsersByOwnerByName`: IamUser
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamUsersByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamUsersByOwnerByNameRequest struct via the builder pattern


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


## PutIamWorkspacesByOwnerByName

> IamWorkspace PutIamWorkspacesByOwnerByName(ctx, owner, name).IamWorkspacesInput(iamWorkspacesInput).Execute()

Changes a workspace's settings.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamWorkspacesInput := *openapiclient.NewIamWorkspacesInput() // IamWorkspacesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.PutIamWorkspacesByOwnerByName(context.Background(), owner, name).IamWorkspacesInput(iamWorkspacesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.PutIamWorkspacesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamWorkspacesByOwnerByName`: IamWorkspace
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.PutIamWorkspacesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIamWorkspacesByOwnerByNameRequest struct via the builder pattern


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

> IamOrganization UpdateOrganization(ctx, owner, name).IamUpdateOrganizationInput(iamUpdateOrganizationInput).Execute()

Changes an organization's display, its defaults and the sign-in rules everyone in it inherits.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamUpdateOrganizationInput := *openapiclient.NewIamUpdateOrganizationInput() // IamUpdateOrganizationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateOrganization(context.Background(), owner, name).IamUpdateOrganizationInput(iamUpdateOrganizationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganization`: IamOrganization
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

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

> IamMutationResult UpdateProvider(ctx, owner, name).IamProvider(iamProvider).Execute()

Changes a provider's settings or rotates the credentials it holds.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamProvider := *openapiclient.NewIamProvider() // IamProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateProvider(context.Background(), owner, name).IamProvider(iamProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProvider`: IamMutationResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

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

> IamSession UpdateSession(ctx, owner, name, application).IamUpdateSessionIn(iamUpdateSessionIn).Execute()

Replaces the set of browsers a session covers — signing out the ones you leave off while the session itself stays live.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	application := "application_example" // string | 
	iamUpdateSessionIn := *openapiclient.NewIamUpdateSessionIn("Application_example", "Name_example", "Owner_example") // IamUpdateSessionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateSession(context.Background(), owner, name, application).IamUpdateSessionIn(iamUpdateSessionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSession`: IamSession
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 
**application** | **string** |  | 

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

> IamTokenMutation UpdateToken(ctx, owner, name).IamToken(iamToken).Execute()

Changes an access token's scope or expiry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamToken := *openapiclient.NewIamToken() // IamToken | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateToken(context.Background(), owner, name).IamToken(iamToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateToken`: IamTokenMutation
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

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

> IamWebauthnCredentialMutationResult UpdateWebauthnCredential(ctx, owner, name).IamWebauthnCredential(iamWebauthnCredential).Execute()

Renames a registered passkey or security key, so a person can tell their devices apart.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	iamWebauthnCredential := *openapiclient.NewIamWebauthnCredential() // IamWebauthnCredential | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateWebauthnCredential(context.Background(), owner, name).IamWebauthnCredential(iamWebauthnCredential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateWebauthnCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebauthnCredential`: IamWebauthnCredentialMutationResult
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateWebauthnCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**name** | **string** |  | 

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

