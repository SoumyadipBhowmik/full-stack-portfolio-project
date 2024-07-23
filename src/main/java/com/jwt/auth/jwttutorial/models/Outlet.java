package com.jwt.auth.jwttutorial.models;

import java.util.Date;
import java.util.UUID;

import org.hibernate.annotations.CreationTimestamp;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.OneToOne;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
public class Outlet {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    private String name;

    @Column(name = "brand_name", nullable = true)
    private String brandName;

    @OneToOne(mappedBy = "outlet", cascade = CascadeType.ALL, orphanRemoval = true)
    private Location location;

    @CreationTimestamp
    @Column(name = "created_at", updatable = false)
    private Date createdAt;

    @CreationTimestamp
    @Column(name = "updated_at")
    private Date updatedAt;
}
