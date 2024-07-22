package com.jwt.auth.jwttutorial.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.jwt.auth.jwttutorial.dto.ClientDTO;

@RestController
@RequestMapping("clients")
public class ClientController {

	 @GetMapping
	 public ClientDTO fetchClients() {
	     return new ClientDTO("abcduuid", "Soumyadip Bhowmik", "7001392045", "Google");
	 }

	 @PostMapping
	 public ClientDTO newClients() {
		 return new ClientDTO("abcduuid1234", "Priyanka Gupta", "7001392045", "Google");
	 }

}
