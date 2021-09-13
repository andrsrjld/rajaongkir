package rajaongkir

type CheckCostRequest struct {
	Origin          string   `json:"origin" validate:"required"`
	OriginType      string   `json:"origin_type" validate:"required"`
	Destination     string   `json:"destination" validate:"required"`
	DestinationType string   `json:"destination_type" validate:"required"`
	Weight          int64    `json:"weight" validate:"required"`
	Courier         []string `json:"courier"`
}

type CheckCostResponse struct {
	Rajaongkir CheckCost `json:"rajaongkir"`
}

type CheckCost struct {
	Query              Query              `json:"query"`
	Status             Status             `json:"status"`
	OriginDetails      OriginDetails      `json:"origin_details"`
	DestinationDetails DestinationDetails `json:"destination_details"`
	Results            []Results          `json:"results"`
}

type Query struct {
	Origin          string `json:"origin"`
	OriginType      string `json:"originType"`
	Destination     string `json:"destination"`
	DestinationType string `json:"destinationType"`
	Weight          int64  `json:"weight"`
	Courier         string `json:"courier"`
}

type Status struct {
	Code        int64  `json:"code"`
	Description string `json:"description"`
}

type OriginDetails struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type DestinationDetails struct {
	SubdistrictID   string `json:"subdistrict_id"`
	ProvinceID      string `json:"province_id"`
	Province        string `json:"province"`
	CityID          string `json:"city_id"`
	City            string `json:"city"`
	Type            string `json:"type"`
	SubdistrictName string `json:"subdistrict_name"`
}

type Results struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Costs []Costs `json:"costs"`
}

type Costs struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []Cost `json:"cost"`
}

type Cost struct {
	Value int64  `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type WaybillRequest struct {
	NomorResi string `json:"no_resi" validate:"required"`
	Courier   string `json:"courier" validate:"required"`
}

type WaybillResponse struct {
	Rajaongkir Waybill `json:"rajaongkir"`
}

type Waybill struct {
	Query  WaybillQuery  `json:"query"`
	Status WaybillStatus `json:"status"`
	Result WaybillResult `json:"result"`
}

type WaybillQuery struct {
	Waybill string `json:"waybill"`
	Courier string `json:"courier"`
}

type WaybillStatus struct {
	Code        int64  `json:"code"`
	Description string `json:"description"`
}

type WaybillResult struct {
	Delivered      bool                  `json:"delivered"`
	Summary        WaybillSummary        `json:"summary"`
	Details        map[string]*string    `json:"details"`
	DeliveryStatus WaybillDeliveryStatus `json:"delivery_status"`
	Manifest       []WaybillManifest     `json:"manifest"`
}

type WaybillSummary struct {
	CourierCode   string `json:"courier_code"`
	CourierName   string `json:"courier_name"`
	WaybillNumber string `json:"waybill_number"`
	ServiceCode   string `json:"service_code"`
	WaybillDate   string `json:"waybill_date"`
	ShipperName   string `json:"shipper_name"`
	ReceiverName  string `json:"receiver_name"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	Status        string `json:"status"`
}

type WaybillDeliveryStatus struct {
	Status      string `json:"status"`
	PodReceiver string `json:"pod_receiver"`
	PodDate     string `json:"pod_date"`
	PodTime     string `json:"pod_time"`
}

type WaybillManifest struct {
	ManifestCode        string `json:"manifest_code"`
	ManifestDescription string `json:"manifest_description"`
	ManifestDate        string `json:"manifest_date"`
	ManifestTime        string `json:"manifest_time"`
	CityName            string `json:"city_name"`
}
