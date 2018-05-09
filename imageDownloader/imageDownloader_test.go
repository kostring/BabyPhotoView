package imageDownloader

import (
	"testing"
)

func TestImageDownload(t *testing.T) {
	picUrl := "http://mmbiz.qpic.cn/mmbiz_jpg/J91fgIwiatneBXNRPjkwos6BkicYKpEwia7Aib2d7Rmt7D9pJGSTErMpVliae2wGR8nESKVwMPdrxQ8YLa0iaaian58kw/0"

	ImageDownloaderInit("E:\\test")
	ImageDownloaderInsertWork("Test", picUrl)
	for ; ImageDownloaderGetOutstandingWorks() != 0 ; {
		t.Error("sdfsdfsdfsd")
	}
}
