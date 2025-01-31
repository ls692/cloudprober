// Copyright 2023 The Cloudprober Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alerting

import (
	"bytes"
	"html/template"
	"sort"
	"sync"
	"time"

	"github.com/cloudprober/cloudprober/probes/alerting/notifier"
)

var statusTmpl = template.Must(template.New("status").Parse(`
<h3>Current Alerts:</h3>
{{ if not .CurrentAlerts }}
  <p>No alerts.</p>
{{ else }}
<table class="status-list">
<tr>
  <th>Failing Since</th>
  <th>Probe</th>
  <th>Target</th>
  <th>Condition ID</th>
  <th>Failures / Total</th>
</tr>
{{ range .CurrentAlerts }}
<tr>
  <td>{{ .FailingSince }}</td>
  <td>{{ .ProbeName }}</td>
  <td>{{ .Target.Name }}</td>
  <td>{{ .ConditionID }}</td>
  <td>{{ .Failures }} / {{ .Total }}</td>
</tr>
{{- end }}
</table>
{{- end }}

<h3>Alerts History:</h3>
{{ if not .PreviousAlerts }}
  <p>No alerts.</p>
{{ else }}
<table class="status-list">
<tr>
  <th>Started At</th>
  <th>Probe</th>
  <th>Target</th>
  <th>Condition ID</th>
  <th>Resolved At</th>
</tr>
{{ range .PreviousAlerts }}
<tr>
  <td>{{ .AlertInfo.FailingSince }}</td>
  <td>{{ .AlertInfo.ProbeName }}</td>
  <td>{{ .AlertInfo.Target.Name }}</td>
  <td>{{ .AlertInfo.ConditionID }}</td>
  <td>{{ .ResolvedAt }}</td>
</tr>
{{- end }}
</table>
{{- end }}
`))

// resolvedAlert is used to keep track of resolved alerts, to be able to show
// alerts hitory on the alerts dashboard.
type resolvedAlert struct {
	AlertInfo  *notifier.AlertInfo
	ResolvedAt time.Time
}

var maxAlertsHistory = 20

var global = struct {
	mu             sync.RWMutex
	currentAlerts  map[string]*notifier.AlertInfo
	previousAlerts []resolvedAlert
}{
	currentAlerts: make(map[string]*notifier.AlertInfo),
}

func updateGlobalState(key string, ai *notifier.AlertInfo) {
	global.mu.Lock()
	defer global.mu.Unlock()

	if ai == nil {
		ra := resolvedAlert{global.currentAlerts[key], time.Now().Truncate(time.Second)}
		global.previousAlerts = append([]resolvedAlert{ra}, global.previousAlerts...)
		if len(global.previousAlerts) > maxAlertsHistory {
			global.previousAlerts = global.previousAlerts[:maxAlertsHistory]
		}
		delete(global.currentAlerts, key)
		return
	}

	global.currentAlerts[key] = ai
}

func resetGlobalState() {
	global.mu.Lock()
	defer global.mu.Unlock()
	global.currentAlerts = make(map[string]*notifier.AlertInfo)
	global.previousAlerts = nil
}

func currentState() ([]*notifier.AlertInfo, []resolvedAlert) {
	global.mu.RLock()
	defer global.mu.RUnlock()

	var currentAlerts []*notifier.AlertInfo
	for _, ai := range global.currentAlerts {
		currentAlerts = append(currentAlerts, ai)
	}

	sort.Slice(currentAlerts, func(i, j int) bool {
		return currentAlerts[i].FailingSince.Before(currentAlerts[j].FailingSince)
	})

	return currentAlerts, append([]resolvedAlert{}, global.previousAlerts...)
}

func StatusHTML() (string, error) {
	var statusBuf bytes.Buffer

	currentAlerts, previousAlerts := currentState()
	for _, ai := range currentAlerts {
		ai.FailingSince = ai.FailingSince.Truncate(time.Second)
	}
	for _, a := range previousAlerts {
		a.AlertInfo.FailingSince = a.AlertInfo.FailingSince.Truncate(time.Second)
	}

	err := statusTmpl.Execute(&statusBuf, struct {
		CurrentAlerts  []*notifier.AlertInfo
		PreviousAlerts []resolvedAlert
	}{
		CurrentAlerts: currentAlerts, PreviousAlerts: previousAlerts,
	})

	return statusBuf.String(), err
}
