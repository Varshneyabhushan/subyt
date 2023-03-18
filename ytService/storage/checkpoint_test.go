package storage

import "testing"

var sampleTestFile = "sample_test.json"
var newId = "newVideoId"

func TestLoadCheckPoint(t *testing.T) {
	checkPoint, err := LoadCheckPoint(sampleTestFile)
	if err != nil {
		t.Error("error occurred when loading invalid file", err)
		return
	}

	if checkPoint.LimitVideo.IsEmpty() {
		t.Error("empty limit loaded with invalid file")
	}
}

//func TestSaveCheckPoint(t *testing.T) {
//	checkPoint := LoadCheckPoint(sampleTestFile)
//	checkPoint.LimitVideo.Id = newId
//	err := checkPoint.Save()
//	if err != nil {
//		t.Error("saving to invalid file resulting in err", err)
//		return
//	}
//
//	checkPoint = LoadCheckPoint(sampleTestFile)
//	if checkPoint.LimitVideo.Id == newId {
//		return
//	}
//
//	t.Error("different id when saved and loaded", newId, checkPoint.LimitVideo.Id)
//}
