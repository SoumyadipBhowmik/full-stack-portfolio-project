package com.jwt.auth.jwttutorial.services;

import java.util.List;

import org.modelmapper.ModelMapper;
import org.springframework.stereotype.Service;

import com.jwt.auth.jwttutorial.dto.OutletDTO;
import com.jwt.auth.jwttutorial.models.Outlet;
import com.jwt.auth.jwttutorial.repositories.OutletsRepository;

@Service
public class OutletService {

	final OutletsRepository outletsRepository;
	final ModelMapper modelMapper;

	OutletService(OutletsRepository outletsRepository, ModelMapper mapper){
		this.outletsRepository = outletsRepository;
		this.modelMapper = mapper;
	}
	
	 public Outlet dTOToEntity(OutletDTO outletDTO) {
	        Outlet outlet = modelMapper.map(outletDTO, Outlet.class);
	        if (outlet.getLocation() != null) {
	            outlet.getLocation().setOutlet(outlet);
	        }
	        return outlet;
	    }
	 
	public OutletDTO entityToDTO(Outlet outlet) {
		return modelMapper.map(outlet, OutletDTO.class);
	}

	public OutletDTO createNewOutlet(OutletDTO outletDTO) {
	    
		Outlet outlet = outletsRepository.save(dTOToEntity(outletDTO));
		
	    return entityToDTO(outlet);
	}

	public List<OutletDTO> exploreOutlets() {
		return outletsRepository.findAll().stream().map(this::entityToDTO).toList();

	}


}
