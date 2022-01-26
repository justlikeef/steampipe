package reportevents

import "github.com/turbot/steampipe/pkg/report/reportinterfaces"

type ContainerError struct {
	Container reportinterfaces.ReportNodeRun
}

// IsReportEvent implements ReportEvent interface
func (*ContainerError) IsReportEvent() {}
