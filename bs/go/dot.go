components {
  id: "dot"
  component: "/bs/scripts/dot.script"
}
embedded_components {
  id: "sprite-visual"
  type: "sprite"
  data: "default_animation: \"dot\"\n"
  "material: \"/bs/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/bs/atlases/main.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.5
  }
}
