package models

type Droplet struct {
	BackupIds []int64  `json:"backup_ids"`
	CreatedAt string   `json:"created_at"`
	Disk      int64    `json:"disk"`
	Features  []string `json:"features"`
	ID        int64    `json:"id"`
	Image     struct {
		CreatedAt     string        `json:"created_at"`
		Description   string        `json:"description"`
		Distribution  string        `json:"distribution"`
		ErrorMessage  string        `json:"error_message"`
		ID            int64         `json:"id"`
		MinDiskSize   int64         `json:"min_disk_size"`
		Name          string        `json:"name"`
		Public        bool          `json:"public"`
		Regions       []string      `json:"regions"`
		SizeGigabytes float64       `json:"size_gigabytes"`
		Slug          string        `json:"slug"`
		Status        string        `json:"status"`
		Tags          []interface{} `json:"tags"`
		Type          string        `json:"type"`
	} `json:"image"`
	Kernel   interface{} `json:"kernel"`
	Locked   bool        `json:"locked"`
	Memory   int64       `json:"memory"`
	Name     string      `json:"name"`
	Networks struct {
		V4 []struct {
			Gateway   string `json:"gateway"`
			IPAddress string `json:"ip_address"`
			Netmask   string `json:"netmask"`
			Type      string `json:"type"`
		} `json:"v4"`
		V6 []struct {
			Gateway   string `json:"gateway"`
			IPAddress string `json:"ip_address"`
			Netmask   int64  `json:"netmask"`
			Type      string `json:"type"`
		} `json:"v6"`
	} `json:"networks"`
	NextBackupWindow struct {
		End   string `json:"end"`
		Start string `json:"start"`
	} `json:"next_backup_window"`
	Region struct {
		Available bool     `json:"available"`
		Features  []string `json:"features"`
		Name      string   `json:"name"`
		Sizes     []string `json:"sizes"`
		Slug      string   `json:"slug"`
	} `json:"region"`
	Size struct {
		Available    bool     `json:"available"`
		Description  string   `json:"description"`
		Disk         int64    `json:"disk"`
		Memory       int64    `json:"memory"`
		PriceHourly  float64  `json:"price_hourly"`
		PriceMonthly int64    `json:"price_monthly"`
		Regions      []string `json:"regions"`
		Slug         string   `json:"slug"`
		Transfer     int64    `json:"transfer"`
		Vcpus        int64    `json:"vcpus"`
	} `json:"size"`
	SizeSlug    string        `json:"size_slug"`
	SnapshotIds []int64       `json:"snapshot_ids"`
	Status      string        `json:"status"`
	Tags        []string      `json:"tags"`
	Vcpus       int64         `json:"vcpus"`
	VolumeIds   []interface{} `json:"volume_ids"`
	VpcUUID     string        `json:"vpc_uuid"`
}



/*type Droplet struct{

	ID int `json:"id"`
	Name string `json:"name"`
	Memory int `json:"memory"`
	VCPUs int `json:"vcpus"`
	Disk int `json:"disk"`
	Locked bool `json:"locked"`
	Status string `json:"status"`
	Kernel interface{} `json:"kernel"`
	CreatedAt string `json:"created_at"`
	Features []string `json:"features"`
}

type Size struct{

}

type Network struct{

}

type Region struct{

}
*/