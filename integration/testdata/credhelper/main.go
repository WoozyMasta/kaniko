/*
Copyright 2026 OSS Container Tools

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// A docker credential helper that hands back whatever HELPER_USERNAME and
// HELPER_SECRET hold, used by TestPathScopedRegistryAuth.
package main

import (
	"encoding/json"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "get" {
		os.Exit(1)
	}
	err := json.NewEncoder(os.Stdout).Encode(map[string]string{
		"ServerURL": "",
		"Username":  os.Getenv("HELPER_USERNAME"),
		"Secret":    os.Getenv("HELPER_SECRET"),
	})
	if err != nil {
		os.Exit(1)
	}
}
