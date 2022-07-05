package entity

type ProfileOrganizer struct {
	Name            string `json:"name"`
	Site            string `json:"site"`
	ImgOrganizer    string `json:"img_organizer"`
	GenerateUrl     string `json:"generate_url"`
	OrganizerBio    string `json:"organizer_bio"`
	Description     string `json:"description"`
	FacebookUrl     string `json:"facebook_url"`
	TwitterUrl      string `json:"twitter_url"`
	StatusOrganizer bool   `json:"status_organizer"`
}
