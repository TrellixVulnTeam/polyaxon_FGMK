// Copyright 2018-2021 Polyaxon, Inc.
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

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.11.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import org.openapitools.client.model.V1ArtifactsType;
import org.openapitools.client.model.V1DockerfileType;
import org.openapitools.client.model.V1FileType;
import org.openapitools.client.model.V1GitType;

/**
 * V1Init
 */

public class V1Init {
  public static final String SERIALIZED_NAME_ARTIFACTS = "artifacts";
  @SerializedName(SERIALIZED_NAME_ARTIFACTS)
  private V1ArtifactsType artifacts;

  public static final String SERIALIZED_NAME_GIT = "git";
  @SerializedName(SERIALIZED_NAME_GIT)
  private V1GitType git;

  public static final String SERIALIZED_NAME_DOCKERFILE = "dockerfile";
  @SerializedName(SERIALIZED_NAME_DOCKERFILE)
  private V1DockerfileType dockerfile;

  public static final String SERIALIZED_NAME_FILE = "file";
  @SerializedName(SERIALIZED_NAME_FILE)
  private V1FileType file;

  public static final String SERIALIZED_NAME_CONNECTION = "connection";
  @SerializedName(SERIALIZED_NAME_CONNECTION)
  private String connection;

  public static final String SERIALIZED_NAME_PATH = "path";
  @SerializedName(SERIALIZED_NAME_PATH)
  private String path;

  public static final String SERIALIZED_NAME_CONTAINER = "container";
  @SerializedName(SERIALIZED_NAME_CONTAINER)
  private Object container;


  public V1Init artifacts(V1ArtifactsType artifacts) {
    
    this.artifacts = artifacts;
    return this;
  }

   /**
   * Get artifacts
   * @return artifacts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1ArtifactsType getArtifacts() {
    return artifacts;
  }


  public void setArtifacts(V1ArtifactsType artifacts) {
    this.artifacts = artifacts;
  }


  public V1Init git(V1GitType git) {
    
    this.git = git;
    return this;
  }

   /**
   * Get git
   * @return git
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1GitType getGit() {
    return git;
  }


  public void setGit(V1GitType git) {
    this.git = git;
  }


  public V1Init dockerfile(V1DockerfileType dockerfile) {
    
    this.dockerfile = dockerfile;
    return this;
  }

   /**
   * Get dockerfile
   * @return dockerfile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1DockerfileType getDockerfile() {
    return dockerfile;
  }


  public void setDockerfile(V1DockerfileType dockerfile) {
    this.dockerfile = dockerfile;
  }


  public V1Init file(V1FileType file) {
    
    this.file = file;
    return this;
  }

   /**
   * Get file
   * @return file
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1FileType getFile() {
    return file;
  }


  public void setFile(V1FileType file) {
    this.file = file;
  }


  public V1Init connection(String connection) {
    
    this.connection = connection;
    return this;
  }

   /**
   * Get connection
   * @return connection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getConnection() {
    return connection;
  }


  public void setConnection(String connection) {
    this.connection = connection;
  }


  public V1Init path(String path) {
    
    this.path = path;
    return this;
  }

   /**
   * Get path
   * @return path
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPath() {
    return path;
  }


  public void setPath(String path) {
    this.path = path;
  }


  public V1Init container(Object container) {
    
    this.container = container;
    return this;
  }

   /**
   * Get container
   * @return container
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getContainer() {
    return container;
  }


  public void setContainer(Object container) {
    this.container = container;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Init v1Init = (V1Init) o;
    return Objects.equals(this.artifacts, v1Init.artifacts) &&
        Objects.equals(this.git, v1Init.git) &&
        Objects.equals(this.dockerfile, v1Init.dockerfile) &&
        Objects.equals(this.file, v1Init.file) &&
        Objects.equals(this.connection, v1Init.connection) &&
        Objects.equals(this.path, v1Init.path) &&
        Objects.equals(this.container, v1Init.container);
  }

  @Override
  public int hashCode() {
    return Objects.hash(artifacts, git, dockerfile, file, connection, path, container);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Init {\n");
    sb.append("    artifacts: ").append(toIndentedString(artifacts)).append("\n");
    sb.append("    git: ").append(toIndentedString(git)).append("\n");
    sb.append("    dockerfile: ").append(toIndentedString(dockerfile)).append("\n");
    sb.append("    file: ").append(toIndentedString(file)).append("\n");
    sb.append("    connection: ").append(toIndentedString(connection)).append("\n");
    sb.append("    path: ").append(toIndentedString(path)).append("\n");
    sb.append("    container: ").append(toIndentedString(container)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

