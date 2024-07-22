package com.jwt.auth.jwttutorial.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class FeaturesDTO {
	

	private String id;

	@JsonProperty("outlet_id")
	private String outlet;
	
	@JsonProperty("feature_name")
	private String featureName;
	
	private double cost;


}
