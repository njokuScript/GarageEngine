package engine

import (
	"github.com/vova616/gl"
)

var (
	lastVAO       VAO
	lastVBO       VBO
	lastVBOTarget gl.GLenum
)

type VBO gl.Buffer
type VAO gl.VertexArray

func GenBuffer() VBO {
	return VBO(gl.GenBuffer())
}

func GenVertexArray() VAO {
	return VAO(gl.GenVertexArray())
}

func (this VBO) Bind(target gl.GLenum) {
	if lastVBO != this || lastVBOTarget != target {
		lastVBO = this
		lastVBOTarget = target
		gl.Buffer.Bind(gl.Buffer(this), target)
	}
}

func (this VAO) Bind() {
	if lastVAO != this {
		lastVAO = this
		gl.VertexArray.Bind(gl.VertexArray(this))
	}
}

type Batch interface {
	Add(position, scale Vector, rotation float32, uv UV) (index int)
	Update(index int, position, scale Vector, rotation float32, uv UV)
	Remove(index int)
	Render()
}

type StaticBatch struct {
	Tex *Texture
	VAO VAO
}

func NewStaticBatch(tex *Texture) *StaticBatch {
	return &StaticBatch{
		Tex: tex,
		VAO: GenVertexArray(),
	}
}

func (this *StaticBatch) Add(position, scale Vector, rotation float32, uv UV) (index int) {
	panic("Not implemented yet")
}

func (this *StaticBatch) Remove(index int) {
	panic("Not implemented yet")
}

func (this *StaticBatch) Render() {
	panic("Not implemented yet")
}