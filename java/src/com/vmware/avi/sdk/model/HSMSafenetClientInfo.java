package com.vmware.avi.sdk.model;

import java.util.*;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * The HSMSafenetClientInfo is a POJO class extends AviRestResource that used for creating
 * HSMSafenetClientInfo.
 *
 * @version 1.0
 * @since 
 *
 */
@JsonIgnoreProperties(ignoreUnknown = true)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class HSMSafenetClientInfo  {
    @JsonProperty("chrystoki_conf")
    private String chrystokiConf = null;

    @JsonProperty("client_cert")
    private String clientCert = null;

    @JsonProperty("client_ip")
    private String clientIp = null;

    @JsonProperty("client_priv_key")
    private String clientPrivKey = null;

    @JsonProperty("session_major_number")
    private Integer sessionMajorNumber = null;

    @JsonProperty("session_minor_number")
    private Integer sessionMinorNumber = null;



    /**
     * This is the getter method this will return the attribute value.
     * Generated file - chrystoki.conf .
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @return chrystokiConf
     */
    public String getChrystokiConf() {
        return chrystokiConf;
    }

    /**
     * This is the setter method to the attribute.
     * Generated file - chrystoki.conf .
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @param chrystokiConf set the chrystokiConf.
     */
    public void setChrystokiConf(String  chrystokiConf) {
        this.chrystokiConf = chrystokiConf;
    }

    /**
     * This is the getter method this will return the attribute value.
     * Client certificate generated by createcert.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @return clientCert
     */
    public String getClientCert() {
        return clientCert;
    }

    /**
     * This is the setter method to the attribute.
     * Client certificate generated by createcert.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @param clientCert set the clientCert.
     */
    public void setClientCert(String  clientCert) {
        this.clientCert = clientCert;
    }

    /**
     * This is the getter method this will return the attribute value.
     * Name prepended to client key and certificate filename.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @return clientIp
     */
    public String getClientIp() {
        return clientIp;
    }

    /**
     * This is the setter method to the attribute.
     * Name prepended to client key and certificate filename.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @param clientIp set the clientIp.
     */
    public void setClientIp(String  clientIp) {
        this.clientIp = clientIp;
    }

    /**
     * This is the getter method this will return the attribute value.
     * Client private key generated by createcert.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @return clientPrivKey
     */
    public String getClientPrivKey() {
        return clientPrivKey;
    }

    /**
     * This is the setter method to the attribute.
     * Client private key generated by createcert.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @param clientPrivKey set the clientPrivKey.
     */
    public void setClientPrivKey(String  clientPrivKey) {
        this.clientPrivKey = clientPrivKey;
    }

    /**
     * This is the getter method this will return the attribute value.
     * Major number of the sesseion.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @return sessionMajorNumber
     */
    public Integer getSessionMajorNumber() {
        return sessionMajorNumber;
    }

    /**
     * This is the setter method to the attribute.
     * Major number of the sesseion.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @param sessionMajorNumber set the sessionMajorNumber.
     */
    public void setSessionMajorNumber(Integer  sessionMajorNumber) {
        this.sessionMajorNumber = sessionMajorNumber;
    }

    /**
     * This is the getter method this will return the attribute value.
     * Minor number of the sesseion.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @return sessionMinorNumber
     */
    public Integer getSessionMinorNumber() {
        return sessionMinorNumber;
    }

    /**
     * This is the setter method to the attribute.
     * Minor number of the sesseion.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @param sessionMinorNumber set the sessionMinorNumber.
     */
    public void setSessionMinorNumber(Integer  sessionMinorNumber) {
        this.sessionMinorNumber = sessionMinorNumber;
    }


    @Override
    public boolean equals(java.lang.Object o) {
      if (this == o) {
          return true;
      }
      if (o == null || getClass() != o.getClass()) {
          return false;
      }
      HSMSafenetClientInfo objHSMSafenetClientInfo = (HSMSafenetClientInfo) o;
      return   Objects.equals(this.clientPrivKey, objHSMSafenetClientInfo.clientPrivKey)&&
  Objects.equals(this.clientCert, objHSMSafenetClientInfo.clientCert)&&
  Objects.equals(this.clientIp, objHSMSafenetClientInfo.clientIp)&&
  Objects.equals(this.chrystokiConf, objHSMSafenetClientInfo.chrystokiConf)&&
  Objects.equals(this.sessionMajorNumber, objHSMSafenetClientInfo.sessionMajorNumber)&&
  Objects.equals(this.sessionMinorNumber, objHSMSafenetClientInfo.sessionMinorNumber);
    }

    @Override
    public String toString() {
      StringBuilder sb = new StringBuilder();
      sb.append("class HSMSafenetClientInfo {\n");
                  sb.append("    chrystokiConf: ").append(toIndentedString(chrystokiConf)).append("\n");
                        sb.append("    clientCert: ").append(toIndentedString(clientCert)).append("\n");
                        sb.append("    clientIp: ").append(toIndentedString(clientIp)).append("\n");
                        sb.append("    clientPrivKey: ").append(toIndentedString(clientPrivKey)).append("\n");
                        sb.append("    sessionMajorNumber: ").append(toIndentedString(sessionMajorNumber)).append("\n");
                        sb.append("    sessionMinorNumber: ").append(toIndentedString(sessionMinorNumber)).append("\n");
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