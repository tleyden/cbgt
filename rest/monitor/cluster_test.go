//  Copyright (c) 2015 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the
//  License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an "AS
//  IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
//  express or implied. See the License for the specific language
//  governing permissions and limitations under the License.

package monitor

import (
	"net/http"
	"testing"
)

func TestEmptyStartMonitorCluster(t *testing.T) {
	HttpGet = func(url string) (resp *http.Response, err error) {
		t.Errorf("expected no get")
		return nil, nil
	}
	defer func() {
		HttpGet = http.Get
	}()

	opt := MonitorClusterOptions{}

	m, err := StartMonitorCluster(nil, nil, opt)
	if err != nil || m == nil {
		t.Errorf("expected no err")
	}

	m.Stop()
}
