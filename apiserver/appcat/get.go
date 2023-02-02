package appcat

import (
	v1 "appcat-apiserver/apis/appcat/v1"
	"context"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
)

var _ rest.Getter = &appcatStorage{}

// Get returns an AppCat service based on its composition
func (s *appcatStorage) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	composition, err := s.compositions.GetComposition(ctx, name, options)
	if err != nil {
		return nil, convertCompositionError(err)
	}

	appcat := v1.NewAppCatFromComposition(composition)
	if appcat == nil {
		// This composition is not an AppCat service
		return nil, apierrors.NewNotFound(appcat.GetGroupVersionResource().GroupResource(), name)
	}

	return appcat, nil
}
