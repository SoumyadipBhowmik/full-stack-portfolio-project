package com.jwt.auth.jwttutorial.controllers;

import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.jwt.auth.jwttutorial.dto.OutletDTO;
import com.jwt.auth.jwttutorial.services.OutletService;

@RestController
@RequestMapping("/outlets")
public class OutletsController {

	final OutletService outletService;


    OutletsController(OutletService outletService) {
        this.outletService = outletService;
    }

    @PostMapping
    public OutletDTO createNewOutlet(@RequestBody OutletDTO outlet) {
    	return outletService.createNewOutlet(outlet);
    }

    @GetMapping
    public List<OutletDTO> exploreOutlets() {
    	return outletService.exploreOutlets();
    }

}
