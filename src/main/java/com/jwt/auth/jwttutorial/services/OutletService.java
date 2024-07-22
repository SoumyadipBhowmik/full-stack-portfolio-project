package com.jwt.auth.jwttutorial.services;

import java.util.List;
import java.util.stream.Collectors;

import org.modelmapper.ModelMapper;
import org.springframework.stereotype.Service;

import com.jwt.auth.jwttutorial.dto.OutletDTO;
import com.jwt.auth.jwttutorial.models.Outlets;
import com.jwt.auth.jwttutorial.repositories.OutletsRepository;

@Service
public class OutletService {

	final OutletsRepository outletsRepository;
	final ModelMapper mapper;

	OutletService(OutletsRepository outletsRepository, ModelMapper mapper){
		this.outletsRepository = outletsRepository;
		this.mapper = mapper;
	}

	public OutletDTO createNewOutlet(OutletDTO outletDTO) {
		Outlets outlet = mapper.map(outletDTO, Outlets.class);
		return mapper.map(outletsRepository.save(outlet), OutletDTO.class);
	}

	public List<OutletDTO> exploreOutlets() {
		return outletsRepository.findAll().stream().map(outlet -> mapper.map(outlet, OutletDTO.class)).collect(Collectors.toList());
		
	}


}
