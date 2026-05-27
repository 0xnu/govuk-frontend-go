package there_is_a_problem_with_the_service_pages

import "testing"

func TestCompile(t *testing.T) {
	_ = Render
	var _ Model
}
