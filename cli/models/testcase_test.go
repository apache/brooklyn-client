/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package models

const testCaseJson = `
{
  "id": "org.apache.brooklyn.test.framework.TestCase:1.0.0-SNAPSHOT",
  "name": "Test Case",
  "symbolicName": "org.apache.brooklyn.test.framework.TestCase",
  "version": "1.0.0-SNAPSHOT",
  "description": "",
  "javaType": "org.apache.brooklyn.test.framework.TestCase",
  "planYaml": "services:\n- type: org.apache.brooklyn.test.framework.TestCase\n  name: Test Case",
  "deprecated": false,
  "config": [
    {
      "reconfigurable": false,
      "possibleValues": null,
      "defaultValue": false,
      "name": "continueOnFailure",
      "description": "Whether to continue executing subsequent children if an earlier child fails",
      "links": {},
      "label": "continueOnFailure",
      "priority": 0,
      "pinned": false,
      "type": "java.lang.Boolean"
    },
    {
      "reconfigurable": false,
      "possibleValues": null,
      "defaultValue": null,
      "name": "defaultDisplayName",
      "description": "Optional default display name to use (rather than auto-generating, if no name is explicitly supplied)",
      "links": {},
      "label": "defaultDisplayName",
      "priority": 0,
      "pinned": false,
      "type": "java.lang.String"
    },
    {
      "reconfigurable": false,
      "possibleValues": null,
      "defaultValue": null,
      "name": "on.error.spec",
      "description": "Spec of entity to instantiate (and start, if startable) if the test-case fails",
      "links": {},
      "label": "on.error.spec",
      "priority": 0,
      "pinned": false,
      "type": "org.apache.brooklyn.api.entity.EntitySpec\u003c?\u003e"
    },
    {
      "reconfigurable": false,
      "possibleValues": null,
      "defaultValue": null,
      "name": "on.finally.spec",
      "description": "Spec of entity to instantiate (and start, if startable) after a test-case either passes or fails",
      "links": {},
      "label": "on.finally.spec",
      "priority": 0,
      "pinned": false,
      "type": "org.apache.brooklyn.api.entity.EntitySpec\u003c?\u003e"
    },
    {
      "reconfigurable": false,
      "possibleValues": null,
      "defaultValue": null,
      "name": "target",
      "description": "Entity under test",
      "links": {},
      "label": "target",
      "priority": 0,
      "pinned": false,
      "type": "org.apache.brooklyn.api.entity.Entity"
    },
    {
      "reconfigurable": false,
      "possibleValues": null,
      "defaultValue": null,
      "name": "targetId",
      "description": "Id of the entity under test",
      "links": {},
      "label": "targetId",
      "priority": 0,
      "pinned": false,
      "type": "java.lang.String"
    },
    {
      "reconfigurable": false,
      "possibleValues": null,
      "defaultValue": "0ms",
      "name": "targetResolutionTimeout",
      "description": "Time to wait for targetId to exist (defaults to zero, i.e. must exist immediately)",
      "links": {},
      "label": "targetResolutionTimeout",
      "priority": 0,
      "pinned": false,
      "type": "org.apache.brooklyn.util.time.Duration"
    }
  ],
  "tags": [
    "equivalent-plan(E267F5A92EED99D286E1FC053E14FC1A)",
    "equivalent-plan(4164F3DE35EA6FFA83BD5C6D899A3E70)",
    "equivalent-plan(C5C2A00B6171F9F86BA66ED440BA812D)",
    "equivalent-plan(BB6E4D5AEF7EF65F4DE98FCDE030DD16)",
    {
      "traits": [
        "org.apache.brooklyn.test.framework.TargetableTestComponent",
        "org.apache.brooklyn.api.entity.Entity",
        "org.apache.brooklyn.api.objs.BrooklynObject",
        "org.apache.brooklyn.api.objs.Identifiable",
        "org.apache.brooklyn.api.objs.Configurable",
        "org.apache.brooklyn.core.entity.trait.Startable"
      ]
    }
  ],
  "links": {
    "self": "/v1/catalog/entities/org.apache.brooklyn.test.framework.TestCase:1.0.0-SNAPSHOT/1.0.0-SNAPSHOT"
  },
  "type": "org.apache.brooklyn.test.framework.TestCase",
  "iconUrl": "",
  "effectors": [
    {
      "name": "restart",
      "description": "Restart the process/service represented by an entity",
      "links": null,
      "parameters": [],
      "returnType": "void"
    },
    {
      "name": "start",
      "description": "Start the process/service represented by an entity",
      "links": null,
      "parameters": [
        {
          "name": "locations",
          "type": "java.lang.Object",
          "description": "The location or locations to start in, as a string, a location object, a list of strings, or a list of location objects",
          "defaultValue": null
        }
      ],
      "returnType": "void"
    },
    {
      "name": "stop",
      "description": "Stop the process/service represented by an entity",
      "links": null,
      "parameters": [],
      "returnType": "void"
    }
  ],
  "sensors": [
    {
      "name": "service.isUp",
      "description": "Whether the service is active and available (confirmed and monitored)",
      "links": {},
      "type": "java.lang.Boolean"
    },
    {
      "name": "target",
      "description": "Entity under test",
      "links": {},
      "type": "org.apache.brooklyn.api.entity.Entity"
    },
    {
      "name": "test.target.entity.id",
      "description": "Id of the target entity",
      "links": {},
      "type": "java.lang.String"
    },
    {
      "name": "test.target.entity.name",
      "description": "Display name of the target entity",
      "links": {},
      "type": "java.lang.String"
    },
    {
      "name": "test.target.entity.type",
      "description": "Type of the target entity",
      "links": {},
      "type": "java.lang.String"
    }
  ]
}
`
