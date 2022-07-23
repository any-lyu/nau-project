 ## build
 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

 docker build -t lv1988zhao/nau-assignment  ./

 ## push
 docker push lv1988zhao/nau-assignment