package data

// INTERMEDIATE MODEL TYPES
var INTERMEDIATE = map[string][2]float32{
	"ORGANIC":    {3.2, 1.05},
	"SEMIDETACH": {3.0, 1.12},
	"EMBEDDED":   {2.8, 1.2},
}

// RATINGS OF CHARACTERISTICS
var RATING = [...]string{"VERY LOW", "LOW", "INTERMEDIATE", "HIGH", "VERY HIGH", "CRITICAL"}

// ATTRIBUTES DESCRIPTIONS
var ATTRIBS = [...]string{
	"Required Software Reliability", "Size of application database", "Complexity of the product",
	"Run-time preformance constraints", "Memory constraints", "Volatility of the virual machine environment", "Required turnabout time",
	"Analyst capability", "Applications exprerience", "Software engineer capability", "Virtual machine experience", "Programming language experience",
	"Application of software experience methods", "Use of software tools", "Required development schedule",
}

// TAGS
var ATTRIBUTES_TAGS = [...]string{
	"RELY", "DATA", "CPLX",
	"TIME", "STOR", "VIRT", "TURN",
	"ACAP", "AEXP", "PCAP", "VEXT", "LEXT",
	"MODP", "TOOL", "SCED",
}

// ALL ATTRIBUTES
var ATTRIBUTES = map[string]map[string]float64{
	//PRODUCT ATTRIBUTES
	"RELY": { // Required Software Reliability
		RATING[0]: 0.75,
		RATING[1]: 0.88,
		RATING[2]: 1,
		RATING[3]: 1.15,
		RATING[4]: 1.40,
		RATING[5]: 1,
	},
	"DATA": { // Size of application database
		RATING[0]: 1,
		RATING[1]: 0.94,
		RATING[2]: 1,
		RATING[3]: 1.08,
		RATING[4]: 1.16,
		RATING[5]: 1,
	},
	"CPLX": { // Complexity of the product
		RATING[0]: 0.7,
		RATING[1]: 0.85,
		RATING[2]: 1,
		RATING[3]: 1.15,
		RATING[4]: 1.3,
		RATING[5]: 1.65,
	},

	//HARDWARE ATTRIBUTES
	"TIME": { // Run-time preformance constraints
		RATING[0]: 1,
		RATING[1]: 1,
		RATING[2]: 1,
		RATING[3]: 1.11,
		RATING[4]: 1.3,
		RATING[5]: 1.66,
	},
	"STOR": { // Memory constraints
		RATING[0]: 1,
		RATING[1]: 1,
		RATING[2]: 1,
		RATING[3]: 1.06,
		RATING[4]: 1.21,
		RATING[5]: 1.56,
	},
	"VIRT": { // Volatility of the virual machine environment
		RATING[0]: 1,
		RATING[1]: 0.87,
		RATING[2]: 1,
		RATING[3]: 1.15,
		RATING[4]: 1.3,
		RATING[5]: 1,
	},
	"TURN": { // Required turnabout time
		RATING[0]: 1,
		RATING[1]: 0.87,
		RATING[2]: 1,
		RATING[3]: 1.07,
		RATING[4]: 1.15,
		RATING[5]: 1,
	},

	//PERSONNEL ATTRIBUTES
	"ACAP": { // Analyst capability
		RATING[0]: 1.46,
		RATING[1]: 1.19,
		RATING[2]: 1,
		RATING[3]: 0.86,
		RATING[4]: 0.71,
		RATING[5]: 1,
	},
	"AEXP": { // Applications exprerience
		RATING[0]: 1.29,
		RATING[1]: 1.13,
		RATING[2]: 1,
		RATING[3]: 0.91,
		RATING[4]: 0.82,
		RATING[5]: 1,
	},
	"PCAP": { // Software engineer capability
		RATING[0]: 1.42,
		RATING[1]: 1.17,
		RATING[2]: 1,
		RATING[3]: 0.86,
		RATING[4]: 0.7,
		RATING[5]: 1,
	},
	"VEXT": { // Virtual machine experience
		RATING[0]: 1.21,
		RATING[1]: 1.1,
		RATING[2]: 1,
		RATING[3]: 0.9,
		RATING[4]: 1,
		RATING[5]: 1,
	},
	"LEXT": { // Programming language experience
		RATING[0]: 1.14,
		RATING[1]: 1.07,
		RATING[2]: 1,
		RATING[3]: 0.95,
		RATING[4]: 1,
		RATING[5]: 1,
	},

	//PROJECT ATTRIBUTES
	"MODP": { // Application of software experience methods
		RATING[0]: 1.24,
		RATING[1]: 1.1,
		RATING[2]: 1,
		RATING[3]: 0.91,
		RATING[4]: 0.82,
		RATING[5]: 1,
	},
	"TOOL": { // Use of software tools
		RATING[0]: 1.24,
		RATING[1]: 1.10,
		RATING[2]: 1,
		RATING[3]: 0.91,
		RATING[4]: 0.83,
		RATING[5]: 1,
	},
	"SCED": { // Required development schedule
		RATING[0]: 1.23,
		RATING[1]: 1.08,
		RATING[2]: 1,
		RATING[3]: 0.91,
		RATING[4]: 1.1,
		RATING[5]: 1,
	},
}
