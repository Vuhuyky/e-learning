package connection

type Connection struct {
	Redis         string         `mapstructure:"redis"`
	Rabbitmq      string         `mapstructure:"rabbitmq"`
	QueueQuantity string         `mapstructure:"queue_quantity"`
	PublicIp      string         `mapstructure:"public_ip"`
	Psql          PsqlConnection `mapstructure:"psql"`
	Smpt          SmptConnection `mapstructure:"smpt"`

	CoreService         ServiceConnection `mapstructure:"core_service"`
	EncodingService     ServiceConnection `mapstructure:"encoding_service"`
	UploadMp4Service    ServiceConnection `mapstructure:"upload_mp4_service"`
	VideoHlsService     ServiceConnection `mapstructure:"video_hls_service"`
	QuizzService        ServiceConnection `mapstructure:"quizz_service"`
	BlobService         ServiceConnection `mapstructure:"blob_service"`
	StreamService       ServiceConnection `mapstructure:"stream_service"`
	MergeBlobSevice     ServiceConnection `mapstructure:"merge_blob_service"`
	QuantityBlobService ServiceConnection `mapstructure:"quantity_blob_service"`
	ProxyService        ServiceConnection `mapstructure:"proxy_service"`
}

type ServiceConnection struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Socket string `mapstructure:"socket"`
	Grpc   string `mapstructure:"grpc"`
}

type SmptConnection struct {
	Email    string `mapstructure:"email"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type PsqlConnection struct {
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
}
