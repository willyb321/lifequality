package command

type lifeScores []struct {
	Color        string  `json:"color"`
	Name         string  `json:"name"`
	ScoreOutOf10 float64 `json:"score_out_of_10"`
}

type searchReturn struct {
	Embedded struct {
		CitySearchResults []struct {
			Embedded struct {
				CityItem struct {
					Embedded struct {
						CityUrbanArea struct {
							Embedded struct {
								UaScores struct {
									Links struct {
										Self struct {
											Href string `json:"href"`
										} `json:"self"`
									} `json:"_links"`
									Categories []struct {
										Color        string  `json:"color"`
										Name         string  `json:"name"`
										ScoreOutOf10 float64 `json:"score_out_of_10"`
									} `json:"categories"`
									Summary           string  `json:"summary"`
									TeleportCityScore float64 `json:"teleport_city_score"`
								} `json:"ua:scores"`
							} `json:"_embedded"`
							Links struct {
								Self struct {
									Href string `json:"href"`
								} `json:"self"`
								UaAdmin1Divisions []struct {
									Href string `json:"href"`
									Name string `json:"name"`
								} `json:"ua:admin1-divisions"`
								UaCities struct {
									Href string `json:"href"`
								} `json:"ua:cities"`
								UaContinent struct {
									Href string `json:"href"`
									Name string `json:"name"`
								} `json:"ua:continent"`
								UaCountries []struct {
									Href string `json:"href"`
									Name string `json:"name"`
								} `json:"ua:countries"`
								UaDetails struct {
									Href string `json:"href"`
								} `json:"ua:details"`
								UaIdentifyingCity struct {
									Href string `json:"href"`
									Name string `json:"name"`
								} `json:"ua:identifying-city"`
								UaImages struct {
									Href string `json:"href"`
								} `json:"ua:images"`
								UaPrimaryCities []struct {
									Href string `json:"href"`
									Name string `json:"name"`
								} `json:"ua:primary-cities"`
								UaSalaries struct {
									Href string `json:"href"`
								} `json:"ua:salaries"`
								UaScores struct {
									Href string `json:"href"`
								} `json:"ua:scores"`
							} `json:"_links"`
							BoundingBox struct {
								Latlon struct {
									East  float64 `json:"east"`
									North float64 `json:"north"`
									South float64 `json:"south"`
									West  float64 `json:"west"`
								} `json:"latlon"`
							} `json:"bounding_box"`
							Continent           string `json:"continent"`
							FullName            string `json:"full_name"`
							IsGovernmentPartner bool   `json:"is_government_partner"`
							Name                string `json:"name"`
							Slug                string `json:"slug"`
							TeleportCityURL     string `json:"teleport_city_url"`
							UaID                string `json:"ua_id"`
						} `json:"city:urban_area"`
					} `json:"_embedded"`
					Links struct {
						CityAdmin1Division struct {
							Href string `json:"href"`
							Name string `json:"name"`
						} `json:"city:admin1_division"`
						CityAlternateNames struct {
							Href string `json:"href"`
						} `json:"city:alternate-names"`
						CityCountry struct {
							Href string `json:"href"`
							Name string `json:"name"`
						} `json:"city:country"`
						CityTimezone struct {
							Href string `json:"href"`
							Name string `json:"name"`
						} `json:"city:timezone"`
						CityUrbanArea struct {
							Href string `json:"href"`
							Name string `json:"name"`
						} `json:"city:urban_area"`
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"_links"`
					FullName  string `json:"full_name"`
					GeonameID int    `json:"geoname_id"`
					Location  struct {
						Geohash string `json:"geohash"`
						Latlon  struct {
							Latitude  float64 `json:"latitude"`
							Longitude float64 `json:"longitude"`
						} `json:"latlon"`
					} `json:"location"`
					Name       string `json:"name"`
					Population int    `json:"population"`
				} `json:"city:item"`
			} `json:"_embedded"`
			Links struct {
				CityItem struct {
					Href string `json:"href"`
				} `json:"city:item"`
			} `json:"_links"`
			MatchingAlternateNames []struct {
				Name string `json:"name"`
			} `json:"matching_alternate_names"`
			MatchingFullName string `json:"matching_full_name"`
		} `json:"city:search-results"`
	} `json:"_embedded"`
	Links struct {
		Curies []struct {
			Href      string `json:"href"`
			Name      string `json:"name"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Count int `json:"count"`
}

type lookupReturn struct {
	Links struct {
		CityAdmin1Division struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"city:admin1_division"`
		CityAlternateNames struct {
			Href string `json:"href"`
		} `json:"city:alternate-names"`
		CityCountry struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"city:country"`
		CityTimezone struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"city:timezone"`
		CityUrbanArea struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"city:urban_area"`
		Curies []struct {
			Href      string `json:"href"`
			Name      string `json:"name"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	FullName  string `json:"full_name"`
	GeonameID int    `json:"geoname_id"`
	Location  struct {
		Geohash string `json:"geohash"`
		Latlon  struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"latlon"`
	} `json:"location"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

type scores struct {
	Links struct {
		Curies []struct {
			Href      string `json:"href"`
			Name      string `json:"name"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Categories []struct {
		Color        string  `json:"color"`
		Name         string  `json:"name"`
		ScoreOutOf10 float64 `json:"score_out_of_10"`
	} `json:"categories"`
	Summary           string  `json:"summary"`
	TeleportCityScore float64 `json:"teleport_city_score"`
}
