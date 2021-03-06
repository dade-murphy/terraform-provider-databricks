/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type ClustersTerminationCode string

// List of ClustersTerminationCode
const (
	USER_REQUEST                  ClustersTerminationCode = "USER_REQUEST"
	JOB_FINISHED                  ClustersTerminationCode = "JOB_FINISHED"
	INACTIVITY                    ClustersTerminationCode = "INACTIVITY"
	CLOUD_PROVIDER_SHUTDOWN       ClustersTerminationCode = "CLOUD_PROVIDER_SHUTDOWN"
	COMMUNICATION_LOST            ClustersTerminationCode = "COMMUNICATION_LOST"
	CLOUD_PROVIDER_LAUNCH_FAILURE ClustersTerminationCode = "CLOUD_PROVIDER_LAUNCH_FAILURE"
	SPARK_STARTUP_FAILURE         ClustersTerminationCode = "SPARK_STARTUP_FAILURE"
	INVALID_ARGUMENT              ClustersTerminationCode = "INVALID_ARGUMENT"
	UNEXPECTED_LAUNCH_FAILURE     ClustersTerminationCode = "UNEXPECTED_LAUNCH_FAILURE"
	INTERNAL_ERROR                ClustersTerminationCode = "INTERNAL_ERROR"
	INSTANCE_UNREACHABLE          ClustersTerminationCode = "INSTANCE_UNREACHABLE"
	REQUEST_REJECTED              ClustersTerminationCode = "REQUEST_REJECTED"
)
