/*
 * Licensed to Elasticsearch under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Package generator allows to generate the Elasticsearch APIs by fitting their JSON spec into the templates stored in
// the templates directory.
package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	_, err := New(filepath.Join("..", DefaultSpecDir), filepath.Join("..", DefaultTemplatesDir))
	// Ignore IsNotExist since we don't always have the spec available.
	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
}