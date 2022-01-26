package reportevents

import "github.com/turbot/steampipe/pkg/report/reportinterfaces"

type ExecutionStarted struct {
	ReportNode reportinterfaces.ReportNodeRun `json:"report"`
}

// IsReportEvent implements ReportEvent interface
func (*ExecutionStarted) IsReportEvent() {}
