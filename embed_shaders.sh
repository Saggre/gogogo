#!/bin/bash

go_shader="./src/util/shader.go" # Embedded shader file path

function embed_shaders() {
  rm $go_shader || :
  touch $go_shader
  printf "package util\n\n" >>$go_shader

  for file in ./src/shaders/*; do
    suffix=""
    filename=$(basename "$file" | cut -d. -f1)
    # shellcheck disable=SC2001
    filename="$(echo "$filename" | sed 's/.*/\u&/')"
    contents=$(cat "$file")

    case "$file" in
    *.frag)
      suffix="Frag"
      ;;
    *.vert)
      suffix="Vert"
      ;;
    esac

    printf "var %s string = \`\n%s\n\` + %s\n\n" "$filename$suffix" "$contents" '"\x00"' >>$go_shader
  done
}

embed_shaders
