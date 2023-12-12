package locations

import "time"

type MappedLocationType struct {
	PublicId    string    `json:"public_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MapLocationType(locationType *LocationType) MappedLocationType {
	return MappedLocationType{
		PublicId:    locationType.PublicId,
		Name:        locationType.Name,
		Description: locationType.Description,
		CreatedAt:   locationType.CreatedAt.AsTime(),
		UpdatedAt:   locationType.UpdatedAt.AsTime(),
	}
}

type MappedWorldRegionCluster struct {
	PublicId    string    `json:"public_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MapWorldRegionCluster(worldRegionCluster *WorldRegionCluster) MappedWorldRegionCluster {
	return MappedWorldRegionCluster{
		PublicId:    worldRegionCluster.PublicId,
		Name:        worldRegionCluster.Name,
		Description: worldRegionCluster.Description,
		CreatedAt:   worldRegionCluster.CreatedAt.AsTime(),
		UpdatedAt:   worldRegionCluster.UpdatedAt.AsTime(),
	}
}

type MappedWorldRegion struct {
	PublicId           string                   `json:"public_id"`
	Name               string                   `json:"name"`
	Description        string                   `json:"description"`
	CreatedAt          time.Time                `json:"created_at"`
	UpdatedAt          time.Time                `json:"updated_at"`
	WorldRegionCluster MappedWorldRegionCluster `json:"world_region_cluster"`
}

func MapWorldRegion(worldRegion *WorldRegion) MappedWorldRegion {
	return MappedWorldRegion{
		PublicId:           worldRegion.PublicId,
		Name:               worldRegion.Name,
		Description:        worldRegion.Description,
		CreatedAt:          worldRegion.CreatedAt.AsTime(),
		UpdatedAt:          worldRegion.UpdatedAt.AsTime(),
		WorldRegionCluster: MapWorldRegionCluster(worldRegion.WorldRegionCluster),
	}
}

type MappedCountry struct {
	PublicId    string    `json:"public_id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	CallingCode string    `json:"calling_code"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MapCountry(country *Country) MappedCountry {
	return MappedCountry{
		PublicId:    country.PublicId,
		Name:        country.Name,
		Code:        country.Code,
		CallingCode: country.CallingCode,
		Longitude:   country.Longitude,
		Latitude:    country.Latitude,
		CreatedAt:   country.CreatedAt.AsTime(),
		UpdatedAt:   country.UpdatedAt.AsTime(),
	}
}

type MappedCountryRegion struct {
	PublicId  string        `json:"public_id"`
	Name      string        `json:"name"`
	Code      string        `json:"code"`
	Longitude float64       `json:"longitude"`
	Latitude  float64       `json:"latitude"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Country   MappedCountry `json:"country"`
}

func MapCountryRegion(countryRegion *CountryRegion) MappedCountryRegion {
	return MappedCountryRegion{
		PublicId:  countryRegion.PublicId,
		Name:      countryRegion.Name,
		Code:      countryRegion.Code,
		Longitude: countryRegion.Longitude,
		Latitude:  countryRegion.Latitude,
		CreatedAt: countryRegion.CreatedAt.AsTime(),
		UpdatedAt: countryRegion.UpdatedAt.AsTime(),
		Country:   MapCountry(countryRegion.Country),
	}
}

type MappedCity struct {
	PublicId      string              `json:"public_id"`
	Name          string              `json:"name"`
	Code          string              `json:"code"`
	Longitude     float64             `json:"longitude"`
	Latitude      float64             `json:"latitude"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	CountryRegion MappedCountryRegion `json:"country_region"`
	WorldRegion   MappedWorldRegion   `json:"world_region"`
}

func MapCity(city *City) MappedCity {
	return MappedCity{
		PublicId:      city.PublicId,
		Name:          city.Name,
		Code:          city.Code,
		Longitude:     city.Longitude,
		Latitude:      city.Latitude,
		CreatedAt:     city.CreatedAt.AsTime(),
		UpdatedAt:     city.UpdatedAt.AsTime(),
		CountryRegion: MapCountryRegion(city.CountryRegion),
		WorldRegion:   MapWorldRegion(city.WorldRegion),
	}
}

type MappedAddress struct {
	PublicId     string             `json:"public_id"`
	AddressLine1 string             `json:"address_line_1"`
	AddressLine2 string             `json:"address_line_2"`
	ZipCode      string             `json:"zip_code"`
	Longitude    float64            `json:"longitude"`
	Latitude     float64            `json:"latitude"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	LocationType MappedLocationType `json:"location_type"`
	City         MappedCity         `json:"city"`
}

func MapAddress(address *Address) MappedAddress {
	return MappedAddress{
		PublicId:     address.PublicId,
		AddressLine1: address.AddressLine_1,
		AddressLine2: address.AddressLine_2,
		ZipCode:      address.ZipCode,
		Longitude:    address.Longitude,
		Latitude:     address.Latitude,
		CreatedAt:    address.CreatedAt.AsTime(),
		UpdatedAt:    address.UpdatedAt.AsTime(),
		LocationType: MapLocationType(address.LocationType),
		City:         MapCity(address.City),
	}
}
