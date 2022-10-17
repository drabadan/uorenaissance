package harvester_test

import (
	"testing"

	m "github.com/drabadan/gostealthclient/pkg/model"
	"github.com/drabadan/uorenaissance/pkg/harvester"
)

func TestFindUniqueElement(t *testing.T) {
	t.Run("Change profile should return -4", func(t *testing.T) {
		//Arrange
		pics := make([]m.TilePic, 4)
		pics = append(pics, m.TilePic{X: 160, Y: 160, Id: 3641, Page: 0, ElemNum: 28})
		pics = append(pics, m.TilePic{X: 62, Y: 160, Id: 3908, Page: 0, ElemNum: 26})
		pics = append(pics, m.TilePic{X: 261, Y: 160, Id: 3637, Page: 0, ElemNum: 30})
		pics = append(pics, m.TilePic{X: 363, Y: 160, Id: 3637, Page: 0, ElemNum: 32})

		//Act
		r := harvester.FindUniqueCaptchaElem(pics)

		//Assert
		if r.Id != 3908 {
			t.Fatalf("Failed to find unique element, return value: %v", r.Id)
		}
	})
}
