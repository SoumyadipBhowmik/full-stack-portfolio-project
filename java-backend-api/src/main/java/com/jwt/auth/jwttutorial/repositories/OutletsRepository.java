package com.jwt.auth.jwttutorial.repositories;

import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;

import com.jwt.auth.jwttutorial.models.Outlet;

public interface OutletsRepository extends JpaRepository<Outlet, UUID>{

}
