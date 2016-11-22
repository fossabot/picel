package image

var NameCases = []NameProvider{
	{Image{
		ID:        "help/staff",
		Extension: "jpg",
	}, "staff.jpg", "help/staff.jpg"},
	{Image{
		ID:        "section/help/staff",
		Extension: "jpg",
	}, "staff.jpg", "section/help/staff.jpg"},
	{Image{
		ID:        "dog",
		Extension: "png",
	}, "dog.png", "dog.png"},
	{Image{
		ID:        "dog",
		Extension: "",
	}, "dog", "dog"},
}

var EscapePathCases = []EscapePathProvider{
	{"", ""},
	{"_", "__"},
	{"__", "____"},
	{"x_", "x__"},
	{"_y", "__y"},
	{"x_y", "x__y"},
}

var EncodeCropCases = []EncodeCropProvider{
	{Crop{
		X:      0,
		Y:      0,
		Width:  10,
		Height: 10,
	}, "0x0:10x10"},
}

var EncodeDimensionCases = []EncodeDimensionProvider{
	{Transform{
		Width:  0,
		Height: 0,
	}, ""},
	{Transform{
		Width: 10,
	}, "10x"},
	{Transform{
		Height: 10,
	}, "x10"},
	{Transform{
		Width:  10,
		Height: 10,
	}, "10x10"},
}

var EncodeParamCases = []EncodeParamProvider{
	{"", ""},
	{"x", "_x"},
}

var ExtractCropCases = []ExtractCropProvider{
	{"10x20:400x300",
		Crop{
			X:      10,
			Y:      20,
			Width:  400,
			Height: 300,
		}},
}

var ExtractCropFailureCases = []ExtractCropFailureProvider{
	{""},
	{"10x20:x300"},
	{"10x20:400x"},
	{"10x:x300"},
	{"20:400"},
}

var GetParamsSubstringStartCases = []GetParamsSubstringProvider{
	{"", -1},
	{"little__kittens", -1},
	{"dogs_4x4.png", 5},
	{"animals/turtles__newborn_4x4.jpg", 25},
}

var GetOffsetsCases = []GetOffsetsProvider{
	{"500x100", 500, 100},
	{"300x", 300, 0},
	{"x300", 0, 300},
	{"0x0", 0, 0},
}

var GetOffsetsFailureCases = []GetOffsetsFailureProvider{
	{""},
	{"-1x10"},
	{"10x-1"},
	{"-1x-1"},
	{"x"},
	{"yx10"},
	{"10xy"},
	{"yxy"},
}

var GetDimensionsCases = []GetDimensionsProvider{
	{"500x100", 500, 100},
	{"300x", 300, 0},
	{"x300", 0, 300},
}

var GetDimensionsFailureCases = []GetDimensionsFailureProvider{
	{""},
	{"-1x10"},
	{"10x-1"},
	{"-1x-1"},
	{"x"},
	{"yx10"},
	{"10xy"},
	{"yxy"},
	{"0x0"},
}

var GetOutputCases = []GetOutputProvider{
	{"", "", ""},
	{"file", "file", ""},
	{"file.out", "file", "out"},
}

var DecodingFailureUnknownParameterCases = []DecodingFailureUnknownParameterProvider{
	{"la__office/newborn__bunnies_raw_stars.jpg"},
}

var DecodingFailureCases = []DecodingFailureProvider{
	{"_"},
	{"la__office/newborn__bunnies_.jpg"},
	{"la__office/newborn__bunnies_400x200:300_gif.jpg"},
	{"la__office/newborn__bunnies_400x200:nox300_gif.jpg"},
	{"la__office/newborn__bunnies_400x200:300xno_gif.jpg"},
}

var CompleteEncodingAndDecodingCases = []CompleteEncodingAndDecodingProvider{
	{Transform{
		Image: Image{
			ID:        "help/staff",
			Extension: "jpg",
			Source:    "help/staff.jpg",
		},
		Path:   "help/staff.jpg",
		Output: "jpg",
	}, "help/staff.jpg"},
	{Transform{
		Image: Image{
			ID:        "help/staff",
			Extension: "webp",
			Source:    "help/staff.webp",
		},
		Path:   "help/staff.webp",
		Output: "webp",
	}, "help/staff.webp"},
	{Transform{
		Image: Image{
			ID:        "help/staff",
			Extension: "webp",
			Source:    "help/staff.webp",
		},
		Path:   "help/staff_800x.webp",
		Width:  800,
		Output: "webp",
	}, "help/staff_800x.webp"},
	{Transform{
		Image: Image{
			ID:        "help/staff",
			Extension: "jpg",
			Source:    "help/staff.jpg",
		},
		Path:   "help/staff_jpg.webp",
		Output: "webp",
	}, "help/staff_jpg.webp"},
	{Transform{
		Image: Image{
			ID:        "airplane_flying_low",
			Extension: "jpg",
			Source:    "airplane_flying_low.jpg",
		},
		Path:   "airplane__flying__low_jpg.webp",
		Output: "webp",
	}, "airplane__flying__low_jpg.webp"},
	{Transform{
		Image: Image{
			ID:        "dog",
			Extension: "jpg",
			Source:    "dog.jpg",
		},
		Path:   "dog",
		Output: "",
	}, "dog"},
	{Transform{
		Image: Image{
			ID:        "help/foo",
			Extension: "jpg",
			Source:    "help/foo.jpg",
		},
		Path:   "help/foo_400x800",
		Output: "",
		Width:  400,
		Height: 800,
	}, "help/foo_400x800"},
	{Transform{
		Image: Image{
			ID:        "help/foo.my.dotted.file",
			Extension: "jpg",
			Source:    "help/foo.my.dotted.file.jpg",
		},
		Path:   "help/foo.my.dotted.file_400x800",
		Output: "",
		Width:  400,
		Height: 800,
	}, "help/foo.my.dotted.file_400x800"},
	{Transform{
		Image: Image{
			ID:        "help/foo",
			Extension: "jpg",
			Source:    "help/foo.jpg",
		},
		Path:   "help/foo_400x",
		Output: "",
		Width:  400,
	}, "help/foo_400x"},
	{Transform{
		Image: Image{
			ID:        "help/foo",
			Extension: "jpg",
			Source:    "help/foo.jpg",
		},
		Path:   "help/foo_x800",
		Output: "",
		Height: 800,
	}, "help/foo_x800"},
	{Transform{
		Image: Image{
			ID:        "adoption_shelters_in_nyc/pretty_dogs",
			Extension: "jpg",
			Source:    "adoption_shelters_in_nyc/pretty_dogs.jpg",
		},
		Path:   "adoption__shelters__in__nyc/pretty__dogs_400x800_jpg.webp",
		Output: "webp",
		Width:  400,
		Height: 800,
	}, "adoption__shelters__in__nyc/pretty__dogs_400x800_jpg.webp"},
	{Transform{
		Image: Image{
			ID:        "airplane_360",
			Extension: "gif",
			Source:    "airplane_360.gif",
		},
		Path:   "airplane__360.gif",
		Output: "gif",
	}, "airplane__360.gif"},
	{Transform{
		Image: Image{
			ID:        "airplane_360",
			Extension: "gif",
			Source:    "airplane_360.gif",
		},
		Path: "airplane__360_gif",
	}, "airplane__360_gif"},
	{Transform{
		Image: Image{
			ID:        "airplane_360",
			Extension: "gif",
			Source:    "airplane_360.gif",
		},
		Path:   "airplane__360_gif.webp",
		Output: "webp",
	}, "airplane__360_gif.webp"},
	{Transform{
		Image: Image{
			ID:        "foo",
			Extension: "jpg",
			Source:    "foo.jpg",
		},
		Path:   "foo",
		Output: "",
	}, "foo"},
	{Transform{
		Image: Image{
			ID:        "foo_bah_h",
			Extension: "jpg",
			Source:    "foo_bah_h.jpg",
		},
		Path: "foo__bah__h_0x0:800x400.jpg",
		Crop: Crop{
			X:      0,
			Y:      0,
			Width:  800,
			Height: 400,
		},
		Output: "jpg",
	}, "foo__bah__h_0x0:800x400.jpg"},
	{Transform{
		Image: Image{
			ID:        "foo_bah_h",
			Extension: "jpg",
			Source:    "foo_bah_h.jpg",
		},
		Path: "foo__bah__h_300x300:800x400.jpg",
		Crop: Crop{
			X:      300,
			Y:      300,
			Width:  800,
			Height: 400,
		},
		Output: "jpg",
	}, "foo__bah__h_300x300:800x400.jpg"},
	{Transform{
		Image: Image{
			ID:        "foo",
			Extension: "jpg",
			Source:    "foo.jpg",
		},
		Path:   "foo_137x0:737x450_800x600_jpg.webp",
		Width:  800,
		Height: 600,
		Crop: Crop{
			X:      137,
			Y:      0,
			Width:  737,
			Height: 450,
		},
		Output: "webp",
	}, "foo_137x0:737x450_800x600_jpg.webp",
	},
	{Transform{
		Image: Image{
			ID:        "adoption_shelters_in_nyc/pretty_dogs",
			Extension: "jpg",
			Source:    "adoption_shelters_in_nyc/pretty_dogs.jpg",
		},
		Path:   "adoption__shelters__in__nyc/pretty__dogs_137x1:737x451_800x600_jpg.webp",
		Width:  800,
		Height: 600,
		Crop: Crop{
			X:      137,
			Y:      1,
			Width:  737,
			Height: 451,
		},
		Output: "webp",
	}, "adoption__shelters__in__nyc/pretty__dogs_137x1:737x451_800x600_jpg.webp",
	},
	{Transform{
		Image: Image{
			ID:        "adoption_shelters_in_nyc/pretty_dogs",
			Extension: "jpg",
			Source:    "adoption_shelters_in_nyc/pretty_dogs.jpg",
		},
		Path: "adoption__shelters__in__nyc/pretty__dogs_137x1:737x451_jpg.webp",
		Crop: Crop{
			X:      137,
			Y:      1,
			Width:  737,
			Height: 451,
		},
		Output: "webp",
	}, "adoption__shelters__in__nyc/pretty__dogs_137x1:737x451_jpg.webp",
	},
	{Transform{
		Image: Image{
			ID:        "la_office/newborn_bunnies",
			Extension: "jpg",
			Source:    "la_office/newborn_bunnies.jpg",
		},
		Path:   "la__office/newborn__bunnies_raw.jpg",
		Raw:    true,
		Output: "jpg",
	}, "la__office/newborn__bunnies_raw.jpg",
	},
}

var DecodingToDefaultOutputFormatCases = []DecodingToDefaultOutputFormatProvider{
	{Transform{
		Image: Image{
			ID:        "la_office/newborn_bunnies",
			Extension: "jpg",
			Source:    "la_office/newborn_bunnies.jpg",
		},
		Path:   "la__office/newborn__bunnies",
		Raw:    false,
		Output: "other",
	}, "la__office/newborn__bunnies",
	},
	{Transform{
		Image: Image{
			ID:        "la_office/newborn_bunnies",
			Extension: "jpg",
			Source:    "la_office/newborn_bunnies.jpg",
		},
		Path:   "la__office/newborn__bunnies_jpg",
		Raw:    false,
		Output: "other",
	}, "la__office/newborn__bunnies_jpg",
	},
	{Transform{
		Image: Image{
			ID:        "la_office/newborn_bunnies",
			Extension: "other",
			Source:    "la_office/newborn_bunnies.other",
		},
		Path:   "la__office/newborn__bunnies_other",
		Raw:    false,
		Output: "other",
	}, "la__office/newborn__bunnies_other",
	},
	{Transform{
		Image: Image{
			ID:        "la_office/newborn_bunnies",
			Extension: "gif",
			Source:    "la_office/newborn_bunnies.gif",
		},
		Path:   "la__office/newborn__bunnies.gif",
		Raw:    false,
		Output: "gif",
	}, "la__office/newborn__bunnies.gif",
	},
	{Transform{
		Image: Image{
			ID:        "big_sur",
			Extension: "jpg",
			Source:    "big_sur.jpg",
		},
		Path: "big__sur_0x0:600x600",
		Crop: Crop{
			X:      0,
			Y:      0,
			Width:  600,
			Height: 600,
		},
		Output: "other",
	}, "big__sur_0x0:600x600",
	},
}

var IncompleteEncodingCases = []IncompleteEncodingProvider{
	{Transform{
		Image: Image{
			ID:        "la_office/newborn_bunnies",
			Extension: "jpg",
		},
		Raw: true,
	}, "la__office/newborn__bunnies_raw.jpg",
	},
	{Transform{
		Image: Image{
			ID: "foo",
		},
		Output: "",
	}, "foo"},
	{Transform{
		Image: Image{
			ID: "help/staff",
		},
		Output: "webp",
	}, "help/staff_jpg.webp"},
}
