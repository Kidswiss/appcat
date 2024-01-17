// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v0.0.0-00010101000000-000000000000 DO NOT EDIT.
package v1

import "k8s.io/apimachinery/pkg/runtime"

// SGPoolingConfigSpec defines model for SGPoolingConfigSpec.
type SGPoolingConfigSpec struct {
	// Connection pooling configuration based on PgBouncer.
	PgBouncer *SGPoolingConfigSpecPgBouncer `json:"pgBouncer,omitempty"`
}

// SGPoolingConfigSpecPgBouncer defines model for SGPoolingConfigSpecPgBouncer.
type SGPoolingConfigSpecPgBouncer struct {
	// The `pgbouncer.ini` parameters the configuration contains, represented as an object where the keys are valid names for the `pgbouncer.ini` configuration file parameters.
	//
	// Check [pgbouncer configuration](https://www.pgbouncer.org/config.html#generic-settings) for more information about supported parameters.
	PgbouncerIni *SGPoolingConfigSpecPgBouncerPgbouncerIni `json:"pgbouncer.ini,omitempty"`
}

// SGPoolingConfigSpecPgBouncerPgbouncerIni defines model for SGPoolingConfigSpecPgBouncerPgbouncerIni.
type SGPoolingConfigSpecPgBouncerPgbouncerIni struct {
	// The `pgbouncer.ini` (Section [databases]) parameters the configuration contains, represented as an object where the keys are valid names for the `pgbouncer.ini` configuration file parameters.
	//
	// Check [pgbouncer configuration](https://www.pgbouncer.org/config.html#section-databases) for more information about supported parameters.
	Databases runtime.RawExtension `json:"databases,omitempty"`

	// The `pgbouncer.ini` (Section [pgbouncer]) parameters the configuration contains, represented as an object where the keys are valid names for the `pgbouncer.ini` configuration file parameters.
	//
	// Check [pgbouncer configuration](https://www.pgbouncer.org/config.html#generic-settings) for more information about supported parameters
	Pgbouncer runtime.RawExtension `json:"pgbouncer,omitempty"`

	// The `pgbouncer.ini` (Section [users]) parameters the configuration contains, represented as an object where the keys are valid names for the `pgbouncer.ini` configuration file parameters.
	//
	// Check [pgbouncer configuration](https://www.pgbouncer.org/config.html#section-users) for more information about supported parameters.
	Users runtime.RawExtension `json:"users,omitempty"`
}

// SGPoolingConfigStatus defines model for SGPoolingConfigStatus.
type SGPoolingConfigStatus struct {
	// Connection pooling configuration status based on PgBouncer.
	PgBouncer *SGPoolingConfigStatusPgBouncer `json:"pgBouncer,omitempty"`
}

// SGPoolingConfigStatusPgBouncer defines model for SGPoolingConfigStatusPgBouncer.
type SGPoolingConfigStatusPgBouncer struct {
	// The `pgbouncer.ini` default parameters parameters which are used if not set.
	DefaultParameters SGPoolingConfigSpecPgBouncerPgbouncerIni `json:"defaultParameters"`
}
