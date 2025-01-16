#!/bin/bash

OUTPUT_DIR="bin"
APP_NAME="yh"

mkdir -p $OUTPUT_DIR

PLATFORMS=("linux/amd64" "darwin/amd64" "windows/amd64")

for PLATFORM in "${PLATFORMS[@]}"; do
  OS=$(echo $PLATFORM | cut -d'/' -f1)
  ARCH=$(echo $PLATFORM | cut -d'/' -f2)
  OUTPUT_FILE="$OUTPUT_DIR/$APP_NAME-$OS-$ARCH"

  if [ "$OS" == "windows" ]; then
    OUTPUT_FILE+=".exe"
  fi

  echo "Building for $PLATFORM..."
  GOOS=$OS GOARCH=$ARCH go build -o $OUTPUT_FILE ./cmd/litGit
done

echo "Builds completed and stored in $OUTPUT_DIR/"
