// Copyright 2018 the Service Broker Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"github.com/pivotal-cf/brokerapi"
	"github.com/spf13/viper"
)

type ServiceBrokerHelper interface {
	Provision(instanceId string, details brokerapi.ProvisionDetails, plan ServicePlan) (ServiceInstanceDetails, error)
	Bind(instanceID, bindingID string, details brokerapi.BindDetails) (ServiceBindingCredentials, error)
	BuildInstanceCredentials(bindRecord ServiceBindingCredentials, instanceRecord ServiceInstanceDetails) (map[string]string, error)
	Unbind(details ServiceBindingCredentials) error
	Deprovision(instanceID string, details brokerapi.DeprovisionDetails) error
	PollInstance(instanceID string) (bool, error)
	LastOperationWasDelete(instanceID string) (bool, error)
	ProvisionsAsync() bool
	DeprovisionsAsync() bool
}

type AccountManager interface {
	CreateCredentials(instanceID string, bindingID string, details brokerapi.BindDetails, instance ServiceInstanceDetails) (ServiceBindingCredentials, error)
	DeleteCredentials(creds ServiceBindingCredentials) error
	BuildInstanceCredentials(bindRecord ServiceBindingCredentials, instanceRecord ServiceInstanceDetails) (map[string]string, error)
}

type GCPCredentials struct {
	Type                string `json:"type"`
	ProjectId           string `json:"project_id"`
	PrivateKeyId        string `json:"private_key_id"`
	PrivateKey          string `json:"private_key"`
	ClientEmail         string `json:"client_email"`
	ClientId            string `json:"client_id"`
	AuthUri             string `json:"auth_uri"`
	TokenUri            string `json:"token_uri"`
	AuthProviderCertUrl string `json:"auth_provider_x509_cert_url"`
	ClientCertUrl       string `json:"client_x509_cert_url"`
}

// This custom user agent string is added to provision calls so that Google can track the aggregated use of this tool
// We can better advocate for devoting resources to supporting cloud foundry and this service broker if we can show
// good usage statistics for it, so if you feel the need to fork this repo, please leave this string in place!
var CustomUserAgent = "cf-gcp-service-broker-test 3.6.0"

func ProductionizeUserAgent() {
	CustomUserAgent = "cf-gcp-service-broker 3.6.0"
}

const CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
const StorageName = "google-storage"
const BigqueryName = "google-bigquery"
const BigtableName = "google-bigtable"
const CloudsqlMySQLName = "google-cloudsql-mysql"
const CloudsqlPostgresName = "google-cloudsql-postgres"
const PubsubName = "google-pubsub"
const MlName = "google-ml-apis"
const SpannerName = "google-spanner"
const StackdriverTraceName = "google-stackdriver-trace"
const StackdriverDebuggerName = "google-stackdriver-debugger"
const StackdriverProfilerName = "google-stackdriver-profiler"
const DatastoreName = "google-datastore"
const rootSaEnvVar = "ROOT_SERVICE_ACCOUNT_JSON"

func init() {
	viper.BindEnv("google.account", rootSaEnvVar)
}

func GetServiceAccountJson() string {
	return viper.GetString("google.account")
}
