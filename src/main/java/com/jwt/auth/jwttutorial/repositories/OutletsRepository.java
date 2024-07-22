package com.jwt.auth.jwttutorial.repositories;

import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;

import com.jwt.auth.jwttutorial.models.Outlets;

public interface OutletsRepository extends JpaRepository<Outlets, UUID>{

}
