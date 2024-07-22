package com.jwt.auth.jwttutorial.controllers;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.jwt.auth.jwttutorial.dto.ClientDTO;

@RestController
@RequestMapping("clients")
public class ClientController {

	 @GetMapping
	 public ResponseEntity<ClientDTO> fetchClients() {
	     return ResponseEntity.ok(new ClientDTO("abcduuid", "Soumyadip Bhowmik", "7001392045", "Google")) ;
	 }

	 @PostMapping
	 public ResponseEntity<ClientDTO> newClients() {
		 return ResponseEntity.ok( new ClientDTO("abcduuid1234", "Priyanka Gupta", "7001392045", "Google"));
	 }

}
