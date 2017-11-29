package geometry

//TODO: Figure out how I can use different Objects in an Array for the Scene
type Geometry interface {
	GetNormal()
	Intersect()
}