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
package org.apache.brooklyn.rest.client;

import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.Response.Status;

import org.apache.brooklyn.rest.domain.ApiError;
import org.apache.brooklyn.rest.domain.TaskSummary;
import org.apache.brooklyn.test.Asserts;
import org.jboss.resteasy.specimpl.BuiltResponse;
import org.testng.Assert;
import org.testng.annotations.Test;

public class BrooklynApiGetEntityTest {

    @Test(expectedExceptions=NullPointerException.class)
    public void testGetEntityDisallowsNull() {
        Assert.assertNull( BrooklynApi.getEntity(null, TaskSummary.class) );
    }
    
    @Test(expectedExceptions=Exception.class)
    public void testGetEntityFailsOnNonJson() {
        BrooklynApi.getEntity(
            new BuiltResponse(200, null, "I'm a string not JSON", null),
            TaskSummary.class);
    }

    @Test
    public void testGetEntityAllowsEmptyMaps() {
        BrooklynApi.getEntity(
            new BuiltResponse(200, null, "{}", null),
            TaskSummary.class);
    }

    @Test
    public void testGetEntityIgnoresExtraFields() {
        BrooklynApi.getEntity(
            new BuiltResponse(200, null, "{ foo: \"This should cause an error\" }", null),
            TaskSummary.class);
    }

    @Test
    public void testGetEntityAllowsNull() {
        BrooklynApi.getEntity(
            new BuiltResponse(200, null, null, null),
            TaskSummary.class);
    }

    @Test
    public void testGetEntityIgnoresErrorResponseCodeInBuiltResponse() {
        BrooklynApi.getEntity(
            new BuiltResponse(400, null, "{}", null),
            TaskSummary.class);
    }

    @Test
    public void testGetEntityIgnoresErrorResponseCodeInBasicResponse() {
        BrooklynApi.getEntity(
            newBadRequestResponse("{}"),
            TaskSummary.class);
    }

    @Test(expectedExceptions=Exception.class)
    public void testGetEntityFailsIfLooksLikeApiError() {
        BrooklynApi.getEntity(
            newBadRequestResponse("{ error: 400 }"),
            TaskSummary.class);
    }

    @Test
    public void testGetEntityFailsWithMessageIfLooksLikeApiErrorWithMessage() {
        try {
            BrooklynApi.getEntity(
                newBadRequestResponse("{ error: 400, message: \"Foo\" }"),
                TaskSummary.class);
            Asserts.shouldHaveFailedPreviously();
        } catch (Exception e) {
            Asserts.expectedFailureContains(e, "Foo");
        }
    }

    @Test
    public void testGetEntitySucceedsIfLooksLikeApiErrorWhenWantingApiError() {
        BrooklynApi.getEntity(
            newBadRequestResponse("{ error: 400 }"),
            ApiError.class);
    }

    @Test(expectedExceptions=Exception.class)
    public void testGetSuccessfulEntityFailsOnAnyError() {
        BrooklynApi.getEntityOnSuccess(
            newBadRequestResponse("{ error: 400 }"),
            ApiError.class);
    }

    @Test
    public void testGetEntitySucceedsIfNoErrorCodeWhenApiError() {
        BrooklynApi.getEntity(
            newJsonResponse(Status.OK, "{ error: 400 }"),
            TaskSummary.class);
    }
    
    @Test
    public void testGetSuccessfulEntitySucceedsOnNoErrorApiError() {
        BrooklynApi.getEntityOnSuccess(
            newJsonResponse(Status.OK, "{ error: 400 }"),
            ApiError.class);
    }

    public Response newBadRequestResponse(Object entity) {
        return newJsonResponse(Status.BAD_REQUEST, entity);
    }

    public static Response newJsonResponse(Status status, Object entity) {
        return newResponse(status, MediaType.APPLICATION_JSON_TYPE, entity);
    }
    
    public static Response newResponse(Status status, MediaType type, Object entity) {
        return Response.status(status)
            .type(type)
            .entity(entity)
            .build();
    }

}
