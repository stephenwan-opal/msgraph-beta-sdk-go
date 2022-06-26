package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10VpnConfigurationable 
type Windows10VpnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsVpnConfigurationable
    GetAssociatedApps()([]Windows10AssociatedAppsable)
    GetAuthenticationMethod()(*Windows10VpnAuthenticationMethod)
    GetConnectionType()(*Windows10VpnConnectionType)
    GetCryptographySuite()(CryptographySuiteable)
    GetDnsRules()([]VpnDnsRuleable)
    GetDnsSuffixes()([]string)
    GetEapXml()([]byte)
    GetEnableAlwaysOn()(*bool)
    GetEnableConditionalAccess()(*bool)
    GetEnableDeviceTunnel()(*bool)
    GetEnableDnsRegistration()(*bool)
    GetEnableSingleSignOnWithAlternateCertificate()(*bool)
    GetEnableSplitTunneling()(*bool)
    GetIdentityCertificate()(WindowsCertificateProfileBaseable)
    GetMicrosoftTunnelSiteId()(*string)
    GetOnlyAssociatedAppsCanUseConnection()(*bool)
    GetProfileTarget()(*Windows10VpnProfileTarget)
    GetProxyServer()(Windows10VpnProxyServerable)
    GetRememberUserCredentials()(*bool)
    GetRoutes()([]VpnRouteable)
    GetSingleSignOnEku()(ExtendedKeyUsageable)
    GetSingleSignOnIssuerHash()(*string)
    GetTrafficRules()([]VpnTrafficRuleable)
    GetTrustedNetworkDomains()([]string)
    GetWindowsInformationProtectionDomain()(*string)
    SetAssociatedApps(value []Windows10AssociatedAppsable)()
    SetAuthenticationMethod(value *Windows10VpnAuthenticationMethod)()
    SetConnectionType(value *Windows10VpnConnectionType)()
    SetCryptographySuite(value CryptographySuiteable)()
    SetDnsRules(value []VpnDnsRuleable)()
    SetDnsSuffixes(value []string)()
    SetEapXml(value []byte)()
    SetEnableAlwaysOn(value *bool)()
    SetEnableConditionalAccess(value *bool)()
    SetEnableDeviceTunnel(value *bool)()
    SetEnableDnsRegistration(value *bool)()
    SetEnableSingleSignOnWithAlternateCertificate(value *bool)()
    SetEnableSplitTunneling(value *bool)()
    SetIdentityCertificate(value WindowsCertificateProfileBaseable)()
    SetMicrosoftTunnelSiteId(value *string)()
    SetOnlyAssociatedAppsCanUseConnection(value *bool)()
    SetProfileTarget(value *Windows10VpnProfileTarget)()
    SetProxyServer(value Windows10VpnProxyServerable)()
    SetRememberUserCredentials(value *bool)()
    SetRoutes(value []VpnRouteable)()
    SetSingleSignOnEku(value ExtendedKeyUsageable)()
    SetSingleSignOnIssuerHash(value *string)()
    SetTrafficRules(value []VpnTrafficRuleable)()
    SetTrustedNetworkDomains(value []string)()
    SetWindowsInformationProtectionDomain(value *string)()
}