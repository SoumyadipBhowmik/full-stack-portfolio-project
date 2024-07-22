package com.jwt.auth.jwttutorial.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

import jakarta.persistence.Column;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class OutletDTO {
	
	
	private String id;
	@JsonProperty("outlet_name")
	private String outletName;
	@JsonProperty("brand_name")
	@Column(nullable = true)
	private String brandName;
	private double cost;
} 
