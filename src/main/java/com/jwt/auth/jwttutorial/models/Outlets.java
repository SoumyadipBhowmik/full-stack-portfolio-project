package com.jwt.auth.jwttutorial.models;

import java.util.Date;
import java.util.List;
import java.util.UUID;

import org.hibernate.annotations.CreationTimestamp;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.OneToMany;
import jakarta.persistence.Table;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "outlets")
public class Outlets {
	
	@Id
	@GeneratedValue(strategy = GenerationType.UUID)
	private UUID id;
	
	@Column(name = "name")
	private String name;
	
	@Column(name = "brand_name", nullable = true)
	private String brandName;
	
	@OneToMany(cascade = CascadeType.ALL)
	private List<Features> features;
	
	@CreationTimestamp
	@Column(name = "created_at", updatable = false)
	private Date createdAt;
	
	@CreationTimestamp
	@Column(name = "updated_at")
	private Date updatedAt;
}
