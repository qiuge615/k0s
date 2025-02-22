/*
Copyright 2021 k0s authors

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
package component

import (
	"context"

	"github.com/k0sproject/k0s/pkg/apis/k0s.k0sproject.io/v1beta1"
)

// Component defines the interface each managed component implements
type Component interface {
	Init() error
	Run(context.Context) error
	Stop() error
	Healthy() error
}

// ReconcilerComponent defines the component interface that is reconciled based
// on changes on the global config CR object changes
type ReconcilerComponent interface {
	Reconcile(context.Context, *v1beta1.ClusterConfig) error
}
