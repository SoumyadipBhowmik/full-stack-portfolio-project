package com.jwt.auth.jwttutorial.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;


@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class ClientDTO {

	private String id;
	@JsonProperty("client_name")
	private String clientName;
	@JsonProperty("phone_number")
	private String phoneNumber;
	@JsonProperty("provider")
	private String provider;
}
