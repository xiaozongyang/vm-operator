/*


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
// Code generated by lister-gen-v0.32. DO NOT EDIT.

package v1beta1

import (
	operatorv1beta1 "github.com/VictoriaMetrics/operator/api/operator/v1beta1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// VMAlertLister helps list VMAlerts.
// All objects returned here must be treated as read-only.
type VMAlertLister interface {
	// List lists all VMAlerts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*operatorv1beta1.VMAlert, err error)
	// VMAlerts returns an object that can list and get VMAlerts.
	VMAlerts(namespace string) VMAlertNamespaceLister
	VMAlertListerExpansion
}

// vMAlertLister implements the VMAlertLister interface.
type vMAlertLister struct {
	listers.ResourceIndexer[*operatorv1beta1.VMAlert]
}

// NewVMAlertLister returns a new VMAlertLister.
func NewVMAlertLister(indexer cache.Indexer) VMAlertLister {
	return &vMAlertLister{listers.New[*operatorv1beta1.VMAlert](indexer, operatorv1beta1.Resource("vmalert"))}
}

// VMAlerts returns an object that can list and get VMAlerts.
func (s *vMAlertLister) VMAlerts(namespace string) VMAlertNamespaceLister {
	return vMAlertNamespaceLister{listers.NewNamespaced[*operatorv1beta1.VMAlert](s.ResourceIndexer, namespace)}
}

// VMAlertNamespaceLister helps list and get VMAlerts.
// All objects returned here must be treated as read-only.
type VMAlertNamespaceLister interface {
	// List lists all VMAlerts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*operatorv1beta1.VMAlert, err error)
	// Get retrieves the VMAlert from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*operatorv1beta1.VMAlert, error)
	VMAlertNamespaceListerExpansion
}

// vMAlertNamespaceLister implements the VMAlertNamespaceLister
// interface.
type vMAlertNamespaceLister struct {
	listers.ResourceIndexer[*operatorv1beta1.VMAlert]
}
