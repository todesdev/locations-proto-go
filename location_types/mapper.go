package location_types

import "time"

type MappedLocationType struct {
	PublicId       string    `json:"public_id"`
	NameEng        string    `json:"name_eng"`
	NameRus        string    `json:"name_rus"`
	DescriptionEng string    `json:"description_eng"`
	DescriptionRus string    `json:"description_rus"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func MapLocationType(locationType *LocationType) MappedLocationType {
	return MappedLocationType{
		PublicId:       locationType.PublicId,
		NameEng:        locationType.NameEng,
		NameRus:        locationType.NameRus,
		DescriptionEng: locationType.DescriptionEng,
		DescriptionRus: locationType.DescriptionRus,
		CreatedAt:      locationType.CreatedAt.AsTime(),
		UpdatedAt:      locationType.UpdatedAt.AsTime(),
	}
}
