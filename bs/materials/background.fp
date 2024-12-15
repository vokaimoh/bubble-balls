varying mediump vec2 var_texcoord0;

void main()
{
	gl_FragColor = mix(vec4(0.86, 0, 0.35, 1), vec4(0.13, 0.46, 1, 1), var_texcoord0.y);
}
