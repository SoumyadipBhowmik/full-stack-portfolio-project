package com.jwt.auth.jwttutorial.repositories;

import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;

import com.jwt.auth.jwttutorial.models.Features;

public interface FeatureRepository extends JpaRepository<Features, UUID>{

}
