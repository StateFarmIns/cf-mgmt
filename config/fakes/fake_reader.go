// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotalservices/cf-mgmt/config"
)

type FakeReader struct {
	OrgsStub        func() (*config.Orgs, error)
	orgsMutex       sync.RWMutex
	orgsArgsForCall []struct{}
	orgsReturns     struct {
		result1 *config.Orgs
		result2 error
	}
	OrgSpacesStub        func(orgName string) (*config.Spaces, error)
	orgSpacesMutex       sync.RWMutex
	orgSpacesArgsForCall []struct {
		orgName string
	}
	orgSpacesReturns struct {
		result1 *config.Spaces
		result2 error
	}
	SpacesStub        func() ([]config.Spaces, error)
	spacesMutex       sync.RWMutex
	spacesArgsForCall []struct{}
	spacesReturns     struct {
		result1 []config.Spaces
		result2 error
	}
	GetOrgConfigsStub        func() ([]config.OrgConfig, error)
	getOrgConfigsMutex       sync.RWMutex
	getOrgConfigsArgsForCall []struct{}
	getOrgConfigsReturns     struct {
		result1 []config.OrgConfig
		result2 error
	}
	GetSpaceConfigsStub        func() ([]config.SpaceConfig, error)
	getSpaceConfigsMutex       sync.RWMutex
	getSpaceConfigsArgsForCall []struct{}
	getSpaceConfigsReturns     struct {
		result1 []config.SpaceConfig
		result2 error
	}
	GetASGConfigsStub        func() ([]config.ASGConfig, error)
	getASGConfigsMutex       sync.RWMutex
	getASGConfigsArgsForCall []struct{}
	getASGConfigsReturns     struct {
		result1 []config.ASGConfig
		result2 error
	}
	GetDefaultASGConfigsStub        func() ([]config.ASGConfig, error)
	getDefaultASGConfigsMutex       sync.RWMutex
	getDefaultASGConfigsArgsForCall []struct{}
	getDefaultASGConfigsReturns     struct {
		result1 []config.ASGConfig
		result2 error
	}
	GetGlobalConfigStub        func() (*config.GlobalConfig, error)
	getGlobalConfigMutex       sync.RWMutex
	getGlobalConfigArgsForCall []struct{}
	getGlobalConfigReturns     struct {
		result1 *config.GlobalConfig
		result2 error
	}
	GetSpaceDefaultsStub        func() (*config.SpaceConfig, error)
	getSpaceDefaultsMutex       sync.RWMutex
	getSpaceDefaultsArgsForCall []struct{}
	getSpaceDefaultsReturns     struct {
		result1 *config.SpaceConfig
		result2 error
	}
	GetOrgConfigStub        func(orgName string) (*config.OrgConfig, error)
	getOrgConfigMutex       sync.RWMutex
	getOrgConfigArgsForCall []struct {
		orgName string
	}
	getOrgConfigReturns struct {
		result1 *config.OrgConfig
		result2 error
	}
	GetSpaceConfigStub        func(orgName, spaceName string) (*config.SpaceConfig, error)
	getSpaceConfigMutex       sync.RWMutex
	getSpaceConfigArgsForCall []struct {
		orgName   string
		spaceName string
	}
	getSpaceConfigReturns struct {
		result1 *config.SpaceConfig
		result2 error
	}
	LdapConfigStub        func(bindPassword string) (*config.LdapConfig, error)
	ldapConfigMutex       sync.RWMutex
	ldapConfigArgsForCall []struct {
		bindPassword string
	}
	ldapConfigReturns struct {
		result1 *config.LdapConfig
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReader) Orgs() (*config.Orgs, error) {
	fake.orgsMutex.Lock()
	fake.orgsArgsForCall = append(fake.orgsArgsForCall, struct{}{})
	fake.recordInvocation("Orgs", []interface{}{})
	fake.orgsMutex.Unlock()
	if fake.OrgsStub != nil {
		return fake.OrgsStub()
	} else {
		return fake.orgsReturns.result1, fake.orgsReturns.result2
	}
}

func (fake *FakeReader) OrgsCallCount() int {
	fake.orgsMutex.RLock()
	defer fake.orgsMutex.RUnlock()
	return len(fake.orgsArgsForCall)
}

func (fake *FakeReader) OrgsReturns(result1 *config.Orgs, result2 error) {
	fake.OrgsStub = nil
	fake.orgsReturns = struct {
		result1 *config.Orgs
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) OrgSpaces(orgName string) (*config.Spaces, error) {
	fake.orgSpacesMutex.Lock()
	fake.orgSpacesArgsForCall = append(fake.orgSpacesArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("OrgSpaces", []interface{}{orgName})
	fake.orgSpacesMutex.Unlock()
	if fake.OrgSpacesStub != nil {
		return fake.OrgSpacesStub(orgName)
	} else {
		return fake.orgSpacesReturns.result1, fake.orgSpacesReturns.result2
	}
}

func (fake *FakeReader) OrgSpacesCallCount() int {
	fake.orgSpacesMutex.RLock()
	defer fake.orgSpacesMutex.RUnlock()
	return len(fake.orgSpacesArgsForCall)
}

func (fake *FakeReader) OrgSpacesArgsForCall(i int) string {
	fake.orgSpacesMutex.RLock()
	defer fake.orgSpacesMutex.RUnlock()
	return fake.orgSpacesArgsForCall[i].orgName
}

func (fake *FakeReader) OrgSpacesReturns(result1 *config.Spaces, result2 error) {
	fake.OrgSpacesStub = nil
	fake.orgSpacesReturns = struct {
		result1 *config.Spaces
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) Spaces() ([]config.Spaces, error) {
	fake.spacesMutex.Lock()
	fake.spacesArgsForCall = append(fake.spacesArgsForCall, struct{}{})
	fake.recordInvocation("Spaces", []interface{}{})
	fake.spacesMutex.Unlock()
	if fake.SpacesStub != nil {
		return fake.SpacesStub()
	} else {
		return fake.spacesReturns.result1, fake.spacesReturns.result2
	}
}

func (fake *FakeReader) SpacesCallCount() int {
	fake.spacesMutex.RLock()
	defer fake.spacesMutex.RUnlock()
	return len(fake.spacesArgsForCall)
}

func (fake *FakeReader) SpacesReturns(result1 []config.Spaces, result2 error) {
	fake.SpacesStub = nil
	fake.spacesReturns = struct {
		result1 []config.Spaces
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetOrgConfigs() ([]config.OrgConfig, error) {
	fake.getOrgConfigsMutex.Lock()
	fake.getOrgConfigsArgsForCall = append(fake.getOrgConfigsArgsForCall, struct{}{})
	fake.recordInvocation("GetOrgConfigs", []interface{}{})
	fake.getOrgConfigsMutex.Unlock()
	if fake.GetOrgConfigsStub != nil {
		return fake.GetOrgConfigsStub()
	} else {
		return fake.getOrgConfigsReturns.result1, fake.getOrgConfigsReturns.result2
	}
}

func (fake *FakeReader) GetOrgConfigsCallCount() int {
	fake.getOrgConfigsMutex.RLock()
	defer fake.getOrgConfigsMutex.RUnlock()
	return len(fake.getOrgConfigsArgsForCall)
}

func (fake *FakeReader) GetOrgConfigsReturns(result1 []config.OrgConfig, result2 error) {
	fake.GetOrgConfigsStub = nil
	fake.getOrgConfigsReturns = struct {
		result1 []config.OrgConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetSpaceConfigs() ([]config.SpaceConfig, error) {
	fake.getSpaceConfigsMutex.Lock()
	fake.getSpaceConfigsArgsForCall = append(fake.getSpaceConfigsArgsForCall, struct{}{})
	fake.recordInvocation("GetSpaceConfigs", []interface{}{})
	fake.getSpaceConfigsMutex.Unlock()
	if fake.GetSpaceConfigsStub != nil {
		return fake.GetSpaceConfigsStub()
	} else {
		return fake.getSpaceConfigsReturns.result1, fake.getSpaceConfigsReturns.result2
	}
}

func (fake *FakeReader) GetSpaceConfigsCallCount() int {
	fake.getSpaceConfigsMutex.RLock()
	defer fake.getSpaceConfigsMutex.RUnlock()
	return len(fake.getSpaceConfigsArgsForCall)
}

func (fake *FakeReader) GetSpaceConfigsReturns(result1 []config.SpaceConfig, result2 error) {
	fake.GetSpaceConfigsStub = nil
	fake.getSpaceConfigsReturns = struct {
		result1 []config.SpaceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetASGConfigs() ([]config.ASGConfig, error) {
	fake.getASGConfigsMutex.Lock()
	fake.getASGConfigsArgsForCall = append(fake.getASGConfigsArgsForCall, struct{}{})
	fake.recordInvocation("GetASGConfigs", []interface{}{})
	fake.getASGConfigsMutex.Unlock()
	if fake.GetASGConfigsStub != nil {
		return fake.GetASGConfigsStub()
	} else {
		return fake.getASGConfigsReturns.result1, fake.getASGConfigsReturns.result2
	}
}

func (fake *FakeReader) GetASGConfigsCallCount() int {
	fake.getASGConfigsMutex.RLock()
	defer fake.getASGConfigsMutex.RUnlock()
	return len(fake.getASGConfigsArgsForCall)
}

func (fake *FakeReader) GetASGConfigsReturns(result1 []config.ASGConfig, result2 error) {
	fake.GetASGConfigsStub = nil
	fake.getASGConfigsReturns = struct {
		result1 []config.ASGConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetDefaultASGConfigs() ([]config.ASGConfig, error) {
	fake.getDefaultASGConfigsMutex.Lock()
	fake.getDefaultASGConfigsArgsForCall = append(fake.getDefaultASGConfigsArgsForCall, struct{}{})
	fake.recordInvocation("GetDefaultASGConfigs", []interface{}{})
	fake.getDefaultASGConfigsMutex.Unlock()
	if fake.GetDefaultASGConfigsStub != nil {
		return fake.GetDefaultASGConfigsStub()
	} else {
		return fake.getDefaultASGConfigsReturns.result1, fake.getDefaultASGConfigsReturns.result2
	}
}

func (fake *FakeReader) GetDefaultASGConfigsCallCount() int {
	fake.getDefaultASGConfigsMutex.RLock()
	defer fake.getDefaultASGConfigsMutex.RUnlock()
	return len(fake.getDefaultASGConfigsArgsForCall)
}

func (fake *FakeReader) GetDefaultASGConfigsReturns(result1 []config.ASGConfig, result2 error) {
	fake.GetDefaultASGConfigsStub = nil
	fake.getDefaultASGConfigsReturns = struct {
		result1 []config.ASGConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetGlobalConfig() (*config.GlobalConfig, error) {
	fake.getGlobalConfigMutex.Lock()
	fake.getGlobalConfigArgsForCall = append(fake.getGlobalConfigArgsForCall, struct{}{})
	fake.recordInvocation("GetGlobalConfig", []interface{}{})
	fake.getGlobalConfigMutex.Unlock()
	if fake.GetGlobalConfigStub != nil {
		return fake.GetGlobalConfigStub()
	} else {
		return fake.getGlobalConfigReturns.result1, fake.getGlobalConfigReturns.result2
	}
}

func (fake *FakeReader) GetGlobalConfigCallCount() int {
	fake.getGlobalConfigMutex.RLock()
	defer fake.getGlobalConfigMutex.RUnlock()
	return len(fake.getGlobalConfigArgsForCall)
}

func (fake *FakeReader) GetGlobalConfigReturns(result1 *config.GlobalConfig, result2 error) {
	fake.GetGlobalConfigStub = nil
	fake.getGlobalConfigReturns = struct {
		result1 *config.GlobalConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetSpaceDefaults() (*config.SpaceConfig, error) {
	fake.getSpaceDefaultsMutex.Lock()
	fake.getSpaceDefaultsArgsForCall = append(fake.getSpaceDefaultsArgsForCall, struct{}{})
	fake.recordInvocation("GetSpaceDefaults", []interface{}{})
	fake.getSpaceDefaultsMutex.Unlock()
	if fake.GetSpaceDefaultsStub != nil {
		return fake.GetSpaceDefaultsStub()
	} else {
		return fake.getSpaceDefaultsReturns.result1, fake.getSpaceDefaultsReturns.result2
	}
}

func (fake *FakeReader) GetSpaceDefaultsCallCount() int {
	fake.getSpaceDefaultsMutex.RLock()
	defer fake.getSpaceDefaultsMutex.RUnlock()
	return len(fake.getSpaceDefaultsArgsForCall)
}

func (fake *FakeReader) GetSpaceDefaultsReturns(result1 *config.SpaceConfig, result2 error) {
	fake.GetSpaceDefaultsStub = nil
	fake.getSpaceDefaultsReturns = struct {
		result1 *config.SpaceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetOrgConfig(orgName string) (*config.OrgConfig, error) {
	fake.getOrgConfigMutex.Lock()
	fake.getOrgConfigArgsForCall = append(fake.getOrgConfigArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("GetOrgConfig", []interface{}{orgName})
	fake.getOrgConfigMutex.Unlock()
	if fake.GetOrgConfigStub != nil {
		return fake.GetOrgConfigStub(orgName)
	} else {
		return fake.getOrgConfigReturns.result1, fake.getOrgConfigReturns.result2
	}
}

func (fake *FakeReader) GetOrgConfigCallCount() int {
	fake.getOrgConfigMutex.RLock()
	defer fake.getOrgConfigMutex.RUnlock()
	return len(fake.getOrgConfigArgsForCall)
}

func (fake *FakeReader) GetOrgConfigArgsForCall(i int) string {
	fake.getOrgConfigMutex.RLock()
	defer fake.getOrgConfigMutex.RUnlock()
	return fake.getOrgConfigArgsForCall[i].orgName
}

func (fake *FakeReader) GetOrgConfigReturns(result1 *config.OrgConfig, result2 error) {
	fake.GetOrgConfigStub = nil
	fake.getOrgConfigReturns = struct {
		result1 *config.OrgConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetSpaceConfig(orgName string, spaceName string) (*config.SpaceConfig, error) {
	fake.getSpaceConfigMutex.Lock()
	fake.getSpaceConfigArgsForCall = append(fake.getSpaceConfigArgsForCall, struct {
		orgName   string
		spaceName string
	}{orgName, spaceName})
	fake.recordInvocation("GetSpaceConfig", []interface{}{orgName, spaceName})
	fake.getSpaceConfigMutex.Unlock()
	if fake.GetSpaceConfigStub != nil {
		return fake.GetSpaceConfigStub(orgName, spaceName)
	} else {
		return fake.getSpaceConfigReturns.result1, fake.getSpaceConfigReturns.result2
	}
}

func (fake *FakeReader) GetSpaceConfigCallCount() int {
	fake.getSpaceConfigMutex.RLock()
	defer fake.getSpaceConfigMutex.RUnlock()
	return len(fake.getSpaceConfigArgsForCall)
}

func (fake *FakeReader) GetSpaceConfigArgsForCall(i int) (string, string) {
	fake.getSpaceConfigMutex.RLock()
	defer fake.getSpaceConfigMutex.RUnlock()
	return fake.getSpaceConfigArgsForCall[i].orgName, fake.getSpaceConfigArgsForCall[i].spaceName
}

func (fake *FakeReader) GetSpaceConfigReturns(result1 *config.SpaceConfig, result2 error) {
	fake.GetSpaceConfigStub = nil
	fake.getSpaceConfigReturns = struct {
		result1 *config.SpaceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) LdapConfig(bindPassword string) (*config.LdapConfig, error) {
	fake.ldapConfigMutex.Lock()
	fake.ldapConfigArgsForCall = append(fake.ldapConfigArgsForCall, struct {
		bindPassword string
	}{bindPassword})
	fake.recordInvocation("LdapConfig", []interface{}{bindPassword})
	fake.ldapConfigMutex.Unlock()
	if fake.LdapConfigStub != nil {
		return fake.LdapConfigStub(bindPassword)
	} else {
		return fake.ldapConfigReturns.result1, fake.ldapConfigReturns.result2
	}
}

func (fake *FakeReader) LdapConfigCallCount() int {
	fake.ldapConfigMutex.RLock()
	defer fake.ldapConfigMutex.RUnlock()
	return len(fake.ldapConfigArgsForCall)
}

func (fake *FakeReader) LdapConfigArgsForCall(i int) string {
	fake.ldapConfigMutex.RLock()
	defer fake.ldapConfigMutex.RUnlock()
	return fake.ldapConfigArgsForCall[i].bindPassword
}

func (fake *FakeReader) LdapConfigReturns(result1 *config.LdapConfig, result2 error) {
	fake.LdapConfigStub = nil
	fake.ldapConfigReturns = struct {
		result1 *config.LdapConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.orgsMutex.RLock()
	defer fake.orgsMutex.RUnlock()
	fake.orgSpacesMutex.RLock()
	defer fake.orgSpacesMutex.RUnlock()
	fake.spacesMutex.RLock()
	defer fake.spacesMutex.RUnlock()
	fake.getOrgConfigsMutex.RLock()
	defer fake.getOrgConfigsMutex.RUnlock()
	fake.getSpaceConfigsMutex.RLock()
	defer fake.getSpaceConfigsMutex.RUnlock()
	fake.getASGConfigsMutex.RLock()
	defer fake.getASGConfigsMutex.RUnlock()
	fake.getDefaultASGConfigsMutex.RLock()
	defer fake.getDefaultASGConfigsMutex.RUnlock()
	fake.getGlobalConfigMutex.RLock()
	defer fake.getGlobalConfigMutex.RUnlock()
	fake.getSpaceDefaultsMutex.RLock()
	defer fake.getSpaceDefaultsMutex.RUnlock()
	fake.getOrgConfigMutex.RLock()
	defer fake.getOrgConfigMutex.RUnlock()
	fake.getSpaceConfigMutex.RLock()
	defer fake.getSpaceConfigMutex.RUnlock()
	fake.ldapConfigMutex.RLock()
	defer fake.ldapConfigMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeReader) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ config.Reader = new(FakeReader)
