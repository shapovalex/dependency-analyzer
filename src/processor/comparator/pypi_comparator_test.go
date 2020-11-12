package comparator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func TestShouldReturnNoResultOnSameData(t *testing.T) {
	p := new(PyPiComparator)
	list1 := []string{"django==2.3.3", "gunicorn[gevent]==20.0.3"}
	list2 := []string{"django==2.3.3", "gunicorn[gevent]==20.0.3"}
	result := p.compare(list1, list2)
	assert.Len(t, result, 0)
}

func TestShouldReturnResultOnDifferentVersions(t *testing.T) {
	p := new(PyPiComparator)
	list1 := []string{"django==2.3.3", "gunicorn[gevent]==20.0.3"}
	list2 := []string{"django==2.3.3", "gunicorn[gevent]==20.0.4"}
	result := p.compare(list1, list2)
	assert.Contains(t, result, "gunicorn[gevent]==20.0.4")
	assert.Len(t, result, 1)
}

func TestShouldReturnResultOnNewLibrary(t *testing.T) {
	p := new(PyPiComparator)
	list1 := []string{"django==2.3.3"}
	list2 := []string{"django==2.3.3", "gunicorn[gevent]==20.0.4"}
	result := p.compare(list1, list2)
	assert.Contains(t, result, "gunicorn[gevent]==20.0.4")
	assert.Len(t, result, 1)
}

func TestShouldNotReturnResultOnDeletedLibrary(t *testing.T) {
	p := new(PyPiComparator)
	list1 := []string{"django==2.3.3", "gunicorn[gevent]==20.0.4"}
	list2 := []string{"django==2.3.3"}
	result := p.compare(list1, list2)
	assert.Len(t, result, 0)
}

func TestShouldReturnCorrectDataOnComplexCase(t *testing.T) {
	p := new(PyPiComparator)
	list1 := []string{"wheel==3.4.5", "django==2.3.3", "gunicorn[gevent]==20.0.4", "pytest==2.3.5"}
	list2 := []string{"django==2.3.4", "pytest==2.3.5", "tensorflow==1.6.7", "wheel==3.4.6"}
	result := p.compare(list1, list2)
	assert.Contains(t, result, "django==2.3.4")
	assert.Contains(t, result, "tensorflow==1.6.7")
	assert.Contains(t, result, "wheel==3.4.6")
	assert.Len(t, result, 3)
}

//complex case
