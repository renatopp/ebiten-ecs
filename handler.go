package sk

type Handler struct {
	id      string
	data    []byte
	err     error
	ready   bool
	loading bool
}

func NewHandler(id string) *Handler {
	return &Handler{id: id}
}

func (h *Handler) Id() string {
	return h.id
}

func (h *Handler) SetData(data []byte) {
	h.data = data
	h.ready = true
	h.err = nil
}

func (h *Handler) Error() error {
	return h.err
}

func (h *Handler) SetError(err error) {
	h.err = err
	h.ready = false
	h.data = nil
}

func (h *Handler) IsReady() bool {
	return h.ready
}

func (h *Handler) AsRaw() []byte {
	return h.data
}

func (h *Handler) AsString() string {
	return string(h.data)
}

func (h *Handler) AsTexture() *Texture {
	img, _ := NewTextureFromBytes(h.data)
	return img
}
