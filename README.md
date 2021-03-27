# Rendering with GO

![Build Status](https://github.com/Saggre/gogogo/workflows/Build%20release/badge.svg?style=flat-square)

Uses GLFW bindings to render with OpenGL.

## Building on Debian/Ubuntu

- `sudo apt-get install libgl1-mesa-dev xorg-dev` Install system deps
- `./embed` Embed shaders into source code
- `go get` Install go deps
- `go build && ./gogogo` Build and run