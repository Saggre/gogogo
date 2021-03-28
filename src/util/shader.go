package util

// Built file, do not edit manually

var BasicFrag string = `
#version 460

uniform sampler2D tex;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = texture(tex, fragTexCoord);
}
` + "\x00"

var BasicGeom string = `
#version 460

layout(points) in;
layout(triangle_strip, max_vertices = 4) out;
out vec2 fragTexCoord;

void main()
{
    gl_Position = vec4(1.0, 1.0, 0.0, 1.0);
    fragTexCoord = vec2(1.0, 1.0);
    EmitVertex();

    gl_Position = vec4(-1.0, 1.0, 0.0, 1.0);
    fragTexCoord = vec2(0.0, 1.0);
    EmitVertex();

    gl_Position = vec4(1.0, -1.0, 0.0, 1.0);
    fragTexCoord = vec2(1.0, 0.0);
    EmitVertex();

    gl_Position = vec4(-1.0, -1.0, 0.0, 1.0);
    fragTexCoord = vec2(0.0, 0.0);
    EmitVertex();

    EndPrimitive();
}
` + "\x00"

var BasicVert string = `
#version 460

void main()
{
}
` + "\x00"

