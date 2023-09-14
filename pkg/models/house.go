package models

// longitude,latitude,housing_median_age,total_rooms,total_bedrooms,population,households,median_income,median_house_value,ocean_proximity
type House struct {
	Longitude          float64
	Latitude           float64
	Housing_median_age float64
	Total_rooms        float64
	Total_bedrooms     float64
	Population         float64
	Households         float64
	Median_income      float64
	Median_house_value float64
	Ocean_proximity    string
	ParseError         error
}
