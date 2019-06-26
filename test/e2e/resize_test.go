// +build e2e

// Copyright 2017 The nats-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2e

import (
	"context"
	"testing"

	natsv1alpha2 "github.com/nats-io/nats-operator/pkg/apis/nats/v1alpha2"
)

// TestResizeClusterFrom3To5 creates a NatsCluster resource with three members and waits for the full mesh to be formed.
// Then, it sets a size of 5 in the NatsCluster resource and waits for the scale-up operation to complete.
func TestResizeClusterFrom3To5(t *testing.T) {
	var (
		initialSize = 3
		finalSize   = 5
		version     = "2.0.0"
	)

	var (
		natsCluster *natsv1alpha2.NatsCluster
		err         error
	)

	// Create a NatsCluster resource with three members.
	if natsCluster, err = f.CreateCluster(f.Namespace, "test-nats-", initialSize, version); err != nil {
		t.Fatal(err)
	}
	// Make sure we cleanup the NatsCluster resource after we're done testing.
	defer func() {
		if err = f.DeleteCluster(natsCluster); err != nil {
			t.Error(err)
		}
	}()

	// Wait until the full mesh is formed with the initial size.
	ctx1, fn := context.WithTimeout(context.Background(), waitTimeout)
	defer fn()
	if err = f.WaitUntilFullMeshWithVersion(ctx1, natsCluster, initialSize, version); err != nil {
		t.Fatal(err)
	}

	// Scale the cluster up to five members.
	natsCluster.Spec.Size = finalSize
	if natsCluster, err = f.PatchCluster(natsCluster); err != nil {
		t.Fatal(err)
	}

	// Wait until the full mesh is formed with the final size.
	ctx2, fn := context.WithTimeout(context.Background(), waitTimeout)
	defer fn()
	if err = f.WaitUntilFullMeshWithVersion(ctx2, natsCluster, finalSize, version); err != nil {
		t.Fatal(err)
	}
}

// TestResizeClusterFrom5To3 creates a NatsCluster resource with five members and waits for the full mesh to be formed.
// Then, it sets a size of 3 in the NatsCluster resource and waits for the scale-down operation to complete.
func TestResizeClusterFrom5To3(t *testing.T) {
	var (
		initialSize = 5
		finalSize   = 3
		version     = "2.0.0"
	)

	var (
		natsCluster *natsv1alpha2.NatsCluster
		err         error
	)

	// Create a NatsCluster resource with three members.
	if natsCluster, err = f.CreateCluster(f.Namespace, "test-nats-", initialSize, version); err != nil {
		t.Fatal(err)
	}
	// Make sure we cleanup the NatsCluster resource after we're done testing.
	defer func() {
		if err = f.DeleteCluster(natsCluster); err != nil {
			t.Error(err)
		}
	}()

	// Wait until the full mesh is formed with the initial size.
	ctx1, fn := context.WithTimeout(context.Background(), waitTimeout)
	defer fn()
	if err = f.WaitUntilFullMeshWithVersion(ctx1, natsCluster, initialSize, version); err != nil {
		t.Fatal(err)
	}

	// Scale the cluster down to three members.
	natsCluster.Spec.Size = finalSize
	if natsCluster, err = f.PatchCluster(natsCluster); err != nil {
		t.Fatal(err)
	}

	// Wait until the full mesh is formed with the final size.
	ctx2, fn := context.WithTimeout(context.Background(), waitTimeout)
	defer fn()
	if err = f.WaitUntilFullMeshWithVersion(ctx2, natsCluster, finalSize, version); err != nil {
		t.Fatal(err)
	}
}
