package com.jwt.auth.jwttutorial.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class LocationDTO {
	
	@JsonProperty("location_id")
	private String locationId;
	@JsonProperty("address_line")
	private String addressLine;
	@JsonProperty("postal_code")
	private String postalCode;
	private String state;
	@JsonProperty("country_code")
	private String countryName;

}
