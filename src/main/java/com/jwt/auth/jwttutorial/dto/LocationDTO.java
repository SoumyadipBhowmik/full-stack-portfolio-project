package com.jwt.auth.jwttutorial.dto;

import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class LocationDTO {


	private UUID id;
	
	@JsonProperty("address_line")
	private String addressLine;
	
	@JsonProperty("postal_code")
	private String postalCode;
	
	private String state;
	
	@JsonProperty("country_name")
	private String countryName;

}
