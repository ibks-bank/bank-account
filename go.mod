module github.com/ibks-bank/bank-account

go 1.18

require (
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/golang/glog v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/ibks-bank/bank-account/pkg/bank-account v0.0.0-00010101000000-000000000000
	github.com/ibks-bank/libs/auth v1.0.3
	github.com/ibks-bank/libs/cerr v1.0.0
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.4
	github.com/rakyll/statik v0.1.7
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220317150908-0efb43f6373e // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ibks-bank/bank-account/pkg/bank-account => ./pkg/bank_account
