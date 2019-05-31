package material

import (
	"GoPath/color"
)

type Material interface {
	BaseColor() color.Color32
	GetSpecular() float32
}

type Diffuse struct {
	Color color.Color32
}

func (D Diffuse) BaseColor() color.Color32 {
	//TODO
	return color.Color32{1, 0, 0}
}

func (D Diffuse) GetSpecular() float32 {
	return 0
}

type StandardMaterial struct {
	Color     color.Color32
	Specular  float32
	Roughness float32
	Rim       float32
}

func (m StandardMaterial) BaseColor() color.Color32 {
	return m.Color
}

func (m StandardMaterial) GetSpecular() float32 {
	return 0.5
}

func Default() Material {
	return StandardMaterial{Color: color.Gray32(0.8)}
}
