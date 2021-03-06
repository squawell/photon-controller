/*
 * Copyright 2016 VMware, Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License.  You may obtain a copy of
 * the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, without warranties or
 * conditions of any kind, EITHER EXPRESS OR IMPLIED.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package com.vmware.photon.controller.nsxclient.models;

import com.vmware.photon.controller.nsxclient.utils.ToStringHelper;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.Objects;

/**
 * This class represents a PhysicalNic JSON structure.
 */
@JsonIgnoreProperties(ignoreUnknown =  true)
public class PhysicalNic {

  @JsonProperty(value = "uplink_name", required = true)
  private String uplinkName;

  @JsonProperty(value = "device_name", required = true)
  private String deviceName;

  public String getUplinkName() {
    return this.uplinkName;
  }

  public void setUplinkName(String uplinkName) {
    this.uplinkName = uplinkName;
  }

  public String getDeviceName() {
    return this.deviceName;
  }

  public void setDeviceName(String deviceName) {
    this.deviceName = deviceName;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }

    if (o == null || getClass() != o.getClass()) {
      return false;
    }

    PhysicalNic other = (PhysicalNic) o;
    return Objects.equals(getUplinkName(), other.getUplinkName())
        && Objects.equals(getDeviceName(), other.getDeviceName());
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(),
        getUplinkName(),
        getDeviceName());
  }

  @Override
  public String toString() {
    return ToStringHelper.jsonObjectToString(this);
  }
}
