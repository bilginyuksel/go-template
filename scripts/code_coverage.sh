files_to_include_coverage=`go list ./... | grep -v mock | grep -v cmd`
coverage_threshold=80

go test -coverprofile=coverage.out -covermode=atomic ${files_to_include_coverage}

coverage=$(go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+')
if [ 1 -eq "$(echo "${coverage} < ${coverage_threshold}" | bc)" ] 
then
    echo "Insufficient test coverage, threshold=${coverage_threshold}, coverage= ${coverage}"
    exit 1
fi

echo "Test coverage is sufficient, threshold=${coverage_threshold}, coverage= ${coverage}"