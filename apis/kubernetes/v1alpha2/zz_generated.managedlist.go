// Code generated by angryjet. DO NOT EDIT.

package v1alpha2

import resource "github.com/crossplane/crossplane-runtime/pkg/resource"

// GetItems of this ObjectList.
func (l *ObjectList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}
