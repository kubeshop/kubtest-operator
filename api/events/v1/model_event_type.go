/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package events

type EventType string

// List of EventType
const (
	START_TEST_EventType            EventType = "start-test"
	END_TEST_SUCCESS_EventType      EventType = "end-test-success"
	END_TEST_FAILED_EventType       EventType = "end-test-failed"
	END_TEST_ABORTED_EventType      EventType = "end-test-aborted"
	END_TEST_TIMEOUT_EventType      EventType = "end-test-timeout"
	START_TESTSUITE_EventType       EventType = "start-testsuite"
	END_TESTSUITE_SUCCESS_EventType EventType = "end-testsuite-success"
	END_TESTSUITE_FAILED_EventType  EventType = "end-testsuite-failed"
	END_TESTSUITE_ABORTED_EventType EventType = "end-testsuite-aborted"
	END_TESTSUITE_TIMEOUT_EventType EventType = "end-testsuite-timeout"
	CREATED_EventType               EventType = "created"
	UPDATED_EventType               EventType = "updated"
	DELETED_EventType               EventType = "deleted"

	TEST_EXECUTION_CREATED_EventType EventType = "test.execution.created"
	TEST_EXECUTION_UPDATED_EventType EventType = "test.execution.updated"
	TEST_EXECUTION_DELETED_EventType EventType = "test.execution.deleted"

	TEST_CREATED_EventType          EventType = "test.created"
	TEST_UPDATED_EventType          EventType = "test.updated"
	TEST_DELETED_EventType          EventType = "test.deleted"
	TEST_DELETED_ALL_EventType      EventType = "test.deleted.all"
	TEST_DELETED_FILTERED_EventType EventType = "test.deleted.filtered"

	TESTSUITE_EXECUTION_CREATED_EventType EventType = "testsuite.execution.created"
	TESTSUITE_EXECUTION_UPDATED_EventType EventType = "testsuite.execution.updated"
	TESTSUITE_EXECUTION_DELETED_EventType EventType = "testsuite.execution.deleted"

	TESTSUITE_CREATED_EventType          EventType = "testsuite.created"
	TESTSUITE_UPDATED_EventType          EventType = "testsuite.updated"
	TESTSUITE_DELETED_EventType          EventType = "testsuite.deleted"
	TESTSUITE_DELETED_ALL_EventType      EventType = "testsuite.deleted.all"
	TESTSUITE_DELETED_FILTERED_EventType EventType = "testsuite.deleted.filtered"
)
