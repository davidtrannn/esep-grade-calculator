package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// --- NEW TESTS FOR 100% COVERAGE ---

// Test for the empty grade lists, hitting the computeAverage (len == 0) branch.
func TestEmptyGradeLists(t *testing.T) {
	gc := NewGradeCalculator()
	expected_value := "F"

	actual_value := gc.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for empty grades; got '%s' instead", expected_value, actual_value)
	}
}

// Test numerical boundary for A (90.0).
func TestBoundaryA(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("Assignment", 90, Assignment)
	gc.AddGrade("Exam", 90, Exam)
	gc.AddGrade("Essay", 90, Essay)

	expected_value := "A"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for numerical grade 90; got '%s'", expected_value, actual_value)
	}
}

// Test numerical boundary for B (80.0).
func TestBoundaryBEdge(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("Assignment", 80, Assignment)
	gc.AddGrade("Exam", 80, Exam)
	gc.AddGrade("Essay", 80, Essay)

	expected_value := "B"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for numerical grade 80; got '%s'", expected_value, actual_value)
	}
}

// Test numerical boundary for C (70.0).
func TestBoundaryC(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("Assignment", 70, Assignment)
	gc.AddGrade("Exam", 70, Exam)
	gc.AddGrade("Essay", 70, Essay)

	expected_value := "C"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for numerical grade 70; got '%s'", expected_value, actual_value)
	}
}

// Test numerical boundary for D (60.0).
func TestBoundaryD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("Assignment", 60, Assignment)
	gc.AddGrade("Exam", 60, Exam)
	gc.AddGrade("Essay", 60, Essay)

	expected_value := "D"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for numerical grade 60; got '%s'", expected_value, actual_value)
	}
}

// Test F grade (numerical grade 59).
func TestBoundaryF(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("Assignment", 59, Assignment)
	gc.AddGrade("Exam", 59, Exam)
	gc.AddGrade("Essay", 59, Essay)

	expected_value := "F"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for numerical grade 59; got '%s'", expected_value, actual_value)
	}
}

// Test the default case of the switch statement in AddGrade (to cover the last branch).
func TestAddGradeDefaultCase(t *testing.T) {
	gc := NewGradeCalculator()

	// Use an undefined GradeType (e.g., 99) to hit the default branch
	undefinedType := GradeType(99)
	gc.AddGrade("Unknown Item", 50, undefinedType)

	// Final grade should still be "F" since no valid grades were added.
	expected_value := "F"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' after adding unknown type; got '%s'", expected_value, actual_value)
	}
}

// Test the String() method on GradeType to cover all cases and the mapping.
func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		gType    GradeType
		expected string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, test := range tests {
		actual := test.gType.String()
		if actual != test.expected {
			t.Errorf("GradeType %d.String() failed. Expected '%s', got '%s'", test.gType, test.expected, actual)
		}
	}
}
