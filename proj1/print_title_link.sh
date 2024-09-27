#!/bin/bash

TITLE=$(xmllint --xpath '//item/title/text()' rss_feed.xml)
LINK=$(xmllint --xpath '//item/link/text()' rss_feed.xml)

IFS=$'\n' read -d '' -r -a array1 <<< "$TITLE"
IFS=$'\n' read -d '' -r -a array2 <<< "$LINK"

LENGTH=${#array1[@]}

for i in $(seq 0 "$((LENGTH-1))")  ; do 
	T="${array1[$i]}"
	echo $i: $T
	L="${array2[$i]}"
	curl -H "Authorization: api_key 763623006edc2eedaf41b33bb4dce105b0c74246c10849fcbc6b69f6f5f4a308" -X POST -d "{\"name\" : \"$T\" , \"url\" : \"$L\" }" http://127.0.0.1:8000/v1/feeds
done
