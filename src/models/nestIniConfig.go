package models

type NestIniConfig struct {
	NestEnvSettings     NestEnvSettings
	AdminName           AdminName
	ConnectionType      ConnectionType
	IntDDName           IntDDName
	BcastDDName         BcastDDName
	IntReqDDName        IntReqDDName
	RmsGetPrsntDDName   RmsGetPrsntDDName
	TouchlineDDName     TouchlineDDName
	RmsDDName           RmsDDName
	RegistrationDetails RegistrationDetails
}

type NestEnvSettings struct {
	HostName                   string
	BackupHost                 string
	User                       string
	MmlLogType                 string
	MmlLoggerAddr              string
	MmlLoggerFile              string
	MmlDomainName              string
	MmlLocBrokAddr             string
	MmlDmnSrvrAddr             string
	MmlDsFoAddr                string
	MmlLicSrvrAddr             string
	MmlEventLoopImplementation string
}

type AdminName struct {
	AdminName string
}

type ConnectionType struct {
	IntddType    string
	BrdddType    string
	IntreqddType string
	IntrstddType string
	TlineddType  string
	RmsddType    string
}

type IntDDName struct {
	DDName string
}

type BcastDDName struct {
	DDName string
}

type IntReqDDName struct {
	DDName string
}

type RmsGetPrsntDDName struct {
	DDName string
}

type TouchlineDDName struct {
	DDName string
}

type RmsDDName struct {
	DDName string
}

type RegistrationDetails struct {
	Details string
}
