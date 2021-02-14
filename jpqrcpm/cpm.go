package jpqrcpm

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
//"errors"
	//	"log"
	"strconv"
	"unicode/utf8"
)

const (
	IDPayloadFormatIndicator = "85"
	IDApplicationTemplate    = "61"

	TagApplicationDefinitionFileName = "4F"
	TagTrack2EquivalentData          = "57"
	TagOtherTemplate                 = "99"
)

type JPQR1 struct {
	PayloadFormatIndicator string              // 85: ペイメント・フォーマット識別子
	ApplicationTemplate    template // 61: アプリケーションテンプレート
}

type template struct {
	ADFName              string // 4F: ブランド識別子 (事業者識別コード)
	Track2EquivalentData string // 57: 決済 ID
	OtherTemplate        string // 99: 自由領域
}

func NewJPQR1(adfName, equivalent, other string) JPQR1 {
	t := template{ADFName: adfName, Track2EquivalentData: equivalent, OtherTemplate: other}
	return JPQR1{ApplicationTemplate: t}
}

func (q *JPQR1) Encode() (string, error) {
	s := format(IDPayloadFormatIndicator, toHex("JPQR1"))

	tlv := formatTemplate(q.ApplicationTemplate)
	s += format(IDApplicationTemplate, tlv)

	decoded, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	s = base64.StdEncoding.EncodeToString([]byte(string(decoded)))
	return s, nil
}

func format(id, value string) string {
	length := utf8.RuneCountInString(value) / 2
	lengthStr := strconv.Itoa(length)
	lengthStr = "00" + fmt.Sprintf("%X", length)
	return id + lengthStr[len(lengthStr)-2:] + value
}

func toHex(s string) string {
	src := []byte(s)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	//hexDebugPrint(string(dst))

	return string(dst)
}

//func hexDebugPrint(src string) {
//	// 確認用に生成した hex 文字列をスペース区切りで出力
//	s := ""
//	slice := strings.Split(src, "")
//	for i := 0; i < len(slice); i += 2 {
//		s += slice[i] + slice[i+1] + " "
//	}
//
//	log.Println(s)
//}

func formatTemplate(t template) string {
	template := ""

	template += format(TagApplicationDefinitionFileName, toHex(t.ADFName))
	template += format(TagTrack2EquivalentData, toHex(t.Track2EquivalentData))

	if t.OtherTemplate != "" {
		template += format(TagOtherTemplate, toHex(t.OtherTemplate))
	}

	return template
}
