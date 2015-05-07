./scripts/dist.sh

VERSION=$(grep "const Version " version.go | sed -E 's/.*"(.+)"$/\1/')
BINTRAY_USER="timakin"
BINTRAY_API_KEY="924f2d307fc1f8bc0f5f8fc4e15b242c2a48d0b7"

for ARCHIVE in ./pkg/dist/*; do
    ARCHIVE_NAME=$(basename ${ARCHIVE})

    echo Uploading: ${ARCHIVE_NAME}
    curl \
        -T ${ARCHIVE} \
        -u ${BINTRAY_USER}:${BINTRAY_API_KEY} \
        "https://api.bintray.com/content/timakin/ts/ts/${VERSION}/${ARCHIVE_NAME}"
done

rm -rf pkg/
