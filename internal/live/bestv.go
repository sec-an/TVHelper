package live

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func BesTvHandler(c *gin.Context) {
	ip := c.DefaultQuery("ip", "117.184.239.60")
	domain := c.DefaultQuery("domain", "liveplay-kk.rtxapp.com")
	channel := c.Param("channel")
	bitrate := c.Param("bitrate")
	stream := strings.Join([]string{"http:/", ip, domain, "live/program/live", channel, bitrate}, "/")
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))
	}
	now := time.Now().In(loc).Add(-13 * 10 * time.Second)
	content := strings.Join([]string{
		"#EXTM3U",
		"#EXT-X-VERSION:3",
		"#EXT-X-TARGETDURATION:10",
		"#EXT-X-MEDIA-SEQUENCE:" + strconv.Itoa(int(now.Unix()))[:9]}, "\r\n")
	for i := 0; i < 10; i++ {
		content = strings.Join([]string{
			content,
			"#EXTINF:10,",
			strings.Join([]string{
				stream,
				now.Format("2006010215"),
				strconv.Itoa(int(now.Unix()))[:9]}, "/") + ".ts"}, "\r\n")
		now = now.Add(10 * time.Second)
	}
	c.Header("Cache-Control", "no-cache")
	c.Header("Accept-Ranges", "none")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	//id := strings.Join([]string{channel, bitrate}, "/")
	//c.Header("X-TXlive-StreamId", id)
	//c.Header("X-TXlive-ChannelId", id)
	c.Data(200, "application/vnd.apple.mpegurl", []byte(content+"\r\n"))
}
