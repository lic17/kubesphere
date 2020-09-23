/*
Copyright 2019 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VolumeSnapshotContentLister helps list VolumeSnapshotContents.
type VolumeSnapshotContentLister interface {
	// List lists all VolumeSnapshotContents in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.VolumeSnapshotContent, err error)
	// Get retrieves the VolumeSnapshotContent from the index for a given name.
	Get(name string) (*v1beta1.VolumeSnapshotContent, error)
	VolumeSnapshotContentListerExpansion
}

// volumeSnapshotContentLister implements the VolumeSnapshotContentLister interface.
type volumeSnapshotContentLister struct {
	indexer cache.Indexer
}

// NewVolumeSnapshotContentLister returns a new VolumeSnapshotContentLister.
func NewVolumeSnapshotContentLister(indexer cache.Indexer) VolumeSnapshotContentLister {
	return &volumeSnapshotContentLister{indexer: indexer}
}

// List lists all VolumeSnapshotContents in the indexer.
func (s *volumeSnapshotContentLister) List(selector labels.Selector) (ret []*v1beta1.VolumeSnapshotContent, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeSnapshotContent))
	})
	return ret, err
}

// Get retrieves the VolumeSnapshotContent from the index for a given name.
func (s *volumeSnapshotContentLister) Get(name string) (*v1beta1.VolumeSnapshotContent, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("volumesnapshotcontent"), name)
	}
	return obj.(*v1beta1.VolumeSnapshotContent), nil
}
