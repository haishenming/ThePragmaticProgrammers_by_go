package Area

func Area(length int, width int) (int, float32) {
	feetArea := length * width
	metersArea := float32(feetArea) * 0.09290304

	return feetArea, metersArea
}
