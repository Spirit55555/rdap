package rdap

type NoridDomainCount struct {
	DecodeData *DecodeData

	Common
	Conformance     []string `rdap:"rdapConformance"`
	ObjectClassName string

	Notices []Notice

	DomainCount []DomainCount `rdap:"domain_count"`
	Identity    string
}

type DomainCount struct {
	Count            *uint8
	ParentDomainName string
}
