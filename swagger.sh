#!/bin/bash


swagger -apiPackage="github.com/planet-work/robotsapin/sapigin/" -mainApiFile="github.com/planet-work/robotsapin/sapigin/sapigin.go" -format="swagger" -output="swagger-ui/dist/v0"
#sed -e 's;"basePath": "/doc/v0";"basePath": "/~master/api/doc/v0";g' -i swagger-ui/dist/v0/index.json

