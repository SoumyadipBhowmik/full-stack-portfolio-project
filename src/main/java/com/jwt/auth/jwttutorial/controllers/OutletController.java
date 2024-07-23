package com.jwt.auth.jwttutorial.controllers;

import java.util.List;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.jwt.auth.jwttutorial.dto.OutletDTO;
import com.jwt.auth.jwttutorial.services.OutletService;

@RestController
@RequestMapping("/outlets")
public class OutletController {

	final OutletService outletService;


    OutletController(OutletService outletService) {
        this.outletService = outletService;
    }

    @PostMapping
    public ResponseEntity<OutletDTO> createNewOutlet(@RequestBody OutletDTO outlet) {
    	return ResponseEntity.ok(outletService.createNewOutlet(outlet)) ;
    }

    @GetMapping
    public ResponseEntity<List<OutletDTO>> exploreOutlets() {
    	return ResponseEntity.ok(outletService.exploreOutlets());
    }
    

}
