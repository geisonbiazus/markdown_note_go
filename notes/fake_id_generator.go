package notes

type FakeIDGenerator struct {
	NextID string
}

func NewFakeIdGenerator(id string) *FakeIDGenerator {
	return &FakeIDGenerator{
		NextID: id,
	}
}

func (g *FakeIDGenerator) Generate() string {
	return g.NextID
}
