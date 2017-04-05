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
package catalog

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
	"net/url"
        "path/filepath"
	"errors"
	"os"
	"strings"
	"archive/zip"
	"io/ioutil"
	"bytes"
)

func Icon(network *net.Network, itemId string) ([]byte, error) {
	url := fmt.Sprintf("/v1/catalog/icon/%s", itemId)
	body, err := network.SendGetRequest(url)
	return body, err
}

func IconWithVersion(network *net.Network, itemId, version string) ([]byte, error) {
	url := fmt.Sprintf("/v1/catalog/icon/%s/%s", itemId, version)
	body, err := network.SendGetRequest(url)
	return body, err
}

func GetEntityWithVersion(network *net.Network, entityId, version string) (models.CatalogEntitySummary, error) {
	url := fmt.Sprintf("/v1/catalog/entities/%s/%s", entityId, version)
	var catalogEntity models.CatalogEntitySummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return catalogEntity, err
	}
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity, err
}

func DeleteEntityWithVersion(network *net.Network, entityId, version string) (string, error) {
	url := fmt.Sprintf("/v1/catalog/entities/%s/%s", entityId, version)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetEntity(network *net.Network, entityId string) (models.CatalogEntitySummary, error) {
	url := fmt.Sprintf("/v1/catalog/entities/%s", entityId)
	var catalogEntity models.CatalogEntitySummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return catalogEntity, err
	}
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity, err
}

func DeleteEntity(network *net.Network, entityId string) (string, error) {
	url := fmt.Sprintf("/v1/catalog/entities/%s", entityId)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetPolicy(network *net.Network, policyId string) (models.CatalogItemSummary, error) {
	url := fmt.Sprintf("/v1/catalog/policies/%s", policyId)
	var catalogItem models.CatalogItemSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return catalogItem, err
	}
	err = json.Unmarshal(body, &catalogItem)
	return catalogItem, err
}

func GetPolicyWithVersion(network *net.Network, policyId, version string) (models.CatalogItemSummary, error) {
	url := fmt.Sprintf("/v1/catalog/policies/%s/%s", policyId, version)
	var catalogItem models.CatalogItemSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return catalogItem, err
	}
	err = json.Unmarshal(body, &catalogItem)
	return catalogItem, err
}

func DeletePolicyWithVersion(network *net.Network, policyId, version string) (string, error) {
	url := fmt.Sprintf("/v1/catalog/policies/%s/%s", policyId, version)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetApplication(network *net.Network, applicationId string) (models.CatalogEntitySummary, error) {
	url := fmt.Sprintf("/v1/catalog/applications/%s", applicationId)
	var catalogEntity models.CatalogEntitySummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return catalogEntity, err
	}
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity, err
}

func GetApplicationWithVersion(network *net.Network, applicationId, version string) (models.CatalogEntitySummary, error) {
	url := fmt.Sprintf("/v1/catalog/applications/%s/%s", applicationId, version)
	var catalogEntity models.CatalogEntitySummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return catalogEntity, err
	}
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity, err
}

func DeleteApplicationWithVersion(network *net.Network, applicationId, version string) (string, error) {
	url := fmt.Sprintf("/v1/catalog/applications/%s/%s", applicationId, version)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func DeleteLocationWithVersion(network *net.Network, locationId, version string) (string, error) {
	url := fmt.Sprintf("/v1/catalog/locations/%s/%s", locationId, version)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Policies(network *net.Network) ([]models.CatalogItemSummary, error) {
	url := "/v1/catalog/policies?allVersions"
	var policies []models.CatalogItemSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return policies, err
	}
	err = json.Unmarshal(body, &policies)
	return policies, err
}

func Locations(network *net.Network) ([]models.CatalogItemSummary, error) {
	url := "/v1/catalog/locations?allVersions=true"
	var catalogLocations []models.CatalogItemSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return catalogLocations, err
	}
	err = json.Unmarshal(body, &catalogLocations)
	return catalogLocations, err
}


func ZipResource(resource string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	defer writer.Close()

	walkFn := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(resource, path)
		if err != nil {
			return err
		}
		f, err := writer.Create(relativePath)
		if err != nil {
			return err
		}

		fileBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		_, err = f.Write(fileBytes)
		if err != nil {
			return err
		}
		return nil
	}

	err := filepath.Walk(resource, walkFn)

	return buf, err;
}

func AddCatalog(network *net.Network, resource string) (map[string]models.CatalogEntitySummary, error) {
	urlString := "/v1/catalog"
	var entities map[string]models.CatalogEntitySummary

	//Assume application/json. This is correct for http/file resources.
	//Zips will need application/x-zip
	contentType := "application/json"
	u, err := url.Parse(resource)
	if err != nil {
		return nil, err
	}

	//Only deal with the below file types
	if "" != u.Scheme && "file" != u.Scheme  && "http" != u.Scheme && "https" != u.Scheme{
		return nil, errors.New("Unrecognised protocol scheme: " + u.Scheme)
	}

	if "" == u.Scheme || "file" == u.Scheme {
		file, err := os.Open(filepath.Clean(resource))
		if err != nil {
			return nil, err
		}

		fileStat, err := file.Stat()
		if err != nil {
			return nil, err
		}

		if fileStat.IsDir() {
			//A dir is a special case, we need to zip it up, and call a different network method
			buf, err := ZipResource(resource)
			if err != nil {
				return nil, err
			}
			body, err := network.SendPostRequestWithContentType(urlString, buf.Bytes(), "application/x-zip")
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(body, &entities)
			return entities, err
		} else {
			extension := filepath.Ext(resource)
			lowercaseExtension := strings.ToLower(extension)
			if lowercaseExtension == ".zip" {
				contentType = "application/x-zip"
			} else if lowercaseExtension == ".jar" {
				contentType = "application/x-jar"
			}
		}

	}

	body, err := network.SendPostResourceRequest(urlString, resource, contentType)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &entities)

	return entities, err
}

func Reset(network *net.Network) (string, error) {
	url := "/v1/catalog/reset"
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetLocation(network *net.Network, locationId string) (models.CatalogItemSummary, error) {
	url := fmt.Sprintf("/v1/catalog/locations/%s", locationId)
	var catalogLocation models.CatalogItemSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return catalogLocation, err
	}
	err = json.Unmarshal(body, &catalogLocation)
	return catalogLocation, err
}

func GetLocationWithVersion(network *net.Network, locationId, version string) (models.CatalogItemSummary, error) {
	url := fmt.Sprintf("/v1/catalog/locations/%s/%s", locationId, version)
	var catalogLocation models.CatalogItemSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return catalogLocation, err
	}
	err = json.Unmarshal(body, &catalogLocation)
	return catalogLocation, err
}

func PostLocationWithVersion(network *net.Network, locationId, version string) (string, error) {
	url := fmt.Sprintf("/v1/catalog/locations/%s/%s", locationId, version)
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Entities(network *net.Network) ([]models.CatalogItemSummary, error) {
	url := "/v1/catalog/entities?allVersions=true"
	var entities []models.CatalogItemSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return entities, err
	}
	err = json.Unmarshal(body, &entities)
	return entities, err
}

func Catalog(network *net.Network) ([]models.CatalogItemSummary, error) {
	url := "/v1/catalog/applications?allVersions=true"
	var applications []models.CatalogItemSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return applications, err
	}
	err = json.Unmarshal(body, &applications)
	return applications, err
}

