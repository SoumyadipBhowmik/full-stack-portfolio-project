package com.jwt.auth.jwttutorial.models;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "clients")
public class Clients {
	
	@Id
	private String id;
	@Column(name = "client_name")
	private String clientName;
	@Column(name = "phone_number")
	private String phoneNumber;
	@Column(name = "provider")
	private String provider;

}
