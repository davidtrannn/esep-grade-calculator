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

// Test for the empty grade lists (numerical average should be 0, final grade "F")
func TestEmptyGradeLists(t *testing.T) {
	gc := NewGradeCalculator()
	expected_value := "F"

	actual_value := gc.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for empty grades; got '%s' instead", expected_value, actual_value)
	}
}

// Test numerical boundaries (all numerical grades are the same to simplify calculation)

func TestBoundaryA(t *testing.T) {
	gc := NewGradeCalculator()
	// All 90s results in a final grade of 90.0
	gc.AddGrade("Assignment", 90, Assignment)
	gc.AddGrade("Exam", 90, Exam)
	gc.AddGrade("Essay", 90, Essay)

	expected_value := "A"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for numerical grade 90; got '%s'", expected_value, actual_value)
	}
}

func TestBoundaryB(t *testing.T) {
	gc := NewGradeCalculator()
	// All 80s results in a final grade of 80.0
	gc.AddGrade("Assignment", 80, Assignment)
	gc.AddGrade("Exam", 80, Exam)
	gc.AddGrade("Essay", 80, Essay)

	expected_value := "B"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for numerical grade 80; got '%s'", expected_value, actual_value)
	}
}

func TestBoundaryC(t *testing.T) {
	gc := NewGradeCalculator()
	// All 70s results in a final grade of 70.0
	gc.AddGrade("Assignment", 70, Assignment)
	gc.AddGrade("Exam", 70, Exam)
	gc.AddGrade("Essay", 70, Essay)

	expected_value := "C"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for numerical grade 70; got '%s'", expected_value, actual_value)
	}
}

func TestBoundaryD(t *testing.T) {
	gc := NewGradeCalculator()
	// All 60s results in a final grade of 60.0
	gc.AddGrade("Assignment", 60, Assignment)
	gc.AddGrade("Exam", 60, Exam)
	gc.AddGrade("Essay", 60, Essay)

	expected_value := "D"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for numerical grade 60; got '%s'", expected_value, actual_value)
	}
}

func TestBoundaryF(t *testing.T) {
	gc := NewGradeCalculator()
	// All 59s results in a final grade of 59.0 (Failing grade)
	gc.AddGrade("Assignment", 59, Assignment)
	gc.AddGrade("Exam", 59, Exam)
	gc.AddGrade("Essay", 59, Essay)

	expected_value := "F"
	actual_value := gc.GetFinalGrade()

	if actual_value != expected_value {
		t.Errorf("Expected GetFinalGrade to return '%s' for numerical grade 59; got '%s'", expected_value, actual_value)
	}
}
