package gtmetrix

type ReportType string

const (
	ReportTypeLighthouse          = "lighthouse"
	ReportTypeLegacy              = "legacy"
	ReportTypeLighthouseAndLegacy = "lighthouse,legacy"
	ReportTypeNone                = "none"
)
