#!/bin/bash
set -e

TAG=$1
REPO=$2

versions=()
while IFS= read -r line; do
  versions+=("$line")
done < <(grep '^## \[' CHANGELOG.md | sed -E 's/^## \[([^]]+)\].*/\1/')

index=-1
for i in "${!versions[@]}"; do
  if [[ "${versions[$i]}" == "$TAG" ]]; then
    index=$i
    break
  fi
done

if [[ $index -eq -1 ]]; then
  echo "❌ ERROR: Tag $TAG not found in CHANGELOG versions!"
  exit 1
fi

if [[ $index -lt $((${#versions[@]} - 1)) ]]; then
  prev_tag=${versions[$((index + 1))]}
else
  prev_tag=""
fi

sed -n "/^## \[$TAG\]/,/^---/{
  /^## \[$TAG\]/d
  /^---/q
  p
}" CHANGELOG.md >changelog_raw.txt

sed '/./,$!d' changelog_raw.txt >tmp.txt && mv tmp.txt changelog_raw.txt
sed ':a
/^[[:space:]]*$/{
  $d
  N
  ba
}' changelog_raw.txt >tmp.txt && mv tmp.txt changelog.txt

rm changelog_raw.txt

if [ ! -s changelog.txt ]; then
  echo "❌ ERROR: No CHANGELOG content found for $TAG!"
  exit 1
fi

if [[ -n "$prev_tag" ]]; then
  echo -e "\n---\n\n**Full Changelog**: https://github.com/$REPO/compare/$prev_tag...$TAG" >>changelog.txt
else
  echo -e "\n---\n\n**Full Changelog**: https://github.com/$REPO/commits/$TAG" >>changelog.txt
fi

cat changelog.txt
