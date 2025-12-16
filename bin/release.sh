#!/usr/bin/env bash

# Получаем последний тег вида v*
LAST_TAG=$(git describe --tags --match "v*" --abbrev=0 2>/dev/null || echo "v0")
# Извлекаем число
NUM=${LAST_TAG#v}
# Считаем новый тег
NEW_TAG=$((NUM+1))
# Создаём новый Git-тег
git tag "v$NEW_TAG"
git push origin "v$NEW_TAG"

# Запуск GoReleaser
goreleaser release --clean
