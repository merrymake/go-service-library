package lib

import (
	"fmt"
	"log"
)

type MimeType struct {
	head string
	tail string
}

func (m MimeType) ToString() string {
	return fmt.Sprintf("%s/%s", m.head, m.tail)
}

var mimetypes = map[string]MimeType{
	"aac":    MimeType{"audio", "aac"},
	"abw":    MimeType{"application", "x-abiword"},
	"arc":    MimeType{"application", "x-freearc"},
	"avif":   MimeType{"image", "avif"},
	"avi":    MimeType{"video", "x-msvideo"},
	"azw":    MimeType{"application", "vnd.amazon.ebook"},
	"bin":    MimeType{"application", "octet-stream"},
	"bmp":    MimeType{"image", "bmp"},
	"bz":     MimeType{"application", "x-bzip"},
	"bz2":    MimeType{"application", "x-bzip2"},
	"cda":    MimeType{"application", "x-cdf"},
	"csh":    MimeType{"application", "x-csh"},
	"css":    MimeType{"text", "css"},
	"csv":    MimeType{"text", "csv"},
	"doc":    MimeType{"application", "msword"},
	"docx":   MimeType{"application", "vnd.openxmlformats-officedocument.wordprocessingml.document"},
	"eot":    MimeType{"application", "vnd.ms-fontobject"},
	"epub":   MimeType{"application", "epub+zip"},
	"gz":     MimeType{"application", "gzip"},
	"gif":    MimeType{"image", "gif"},
	"htm":    MimeType{"text", "html"},
	"html":   MimeType{"text", "html"},
	"ico":    MimeType{"image", "vnd.microsoft.icon"},
	"ics":    MimeType{"text", "calendar"},
	"jar":    MimeType{"application", "java-archive"},
	"jpeg":   MimeType{"image", "jpeg"},
	"jpg":    MimeType{"image", "jpeg"},
	"js":     MimeType{"text", "javascript"},
	"json":   MimeType{"application", "json"},
	"jsonld": MimeType{"application", "ld+json"},
	"mid":    MimeType{"audio", "midi"},
	"midi":   MimeType{"audio", "midi"},
	"mjs":    MimeType{"text", "javascript"},
	"mp3":    MimeType{"audio", "mpeg"},
	"mp4":    MimeType{"video", "mp4"},
	"mpeg":   MimeType{"video", "mpeg"},
	"mpkg":   MimeType{"application", "vnd.apple.installer+xml"},
	"odp":    MimeType{"application", "vnd.oasis.opendocument.presentation"},
	"ods":    MimeType{"application", "vnd.oasis.opendocument.spreadsheet"},
	"odt":    MimeType{"application", "vnd.oasis.opendocument.text"},
	"oga":    MimeType{"audio", "ogg"},
	"ogv":    MimeType{"video", "ogg"},
	"ogx":    MimeType{"application", "ogg"},
	"opus":   MimeType{"audio", "opus"},
	"otf":    MimeType{"font", "otf"},
	"png":    MimeType{"image", "png"},
	"pdf":    MimeType{"application", "pdf"},
	"php":    MimeType{"application", "x-httpd-php"},
	"ppt":    MimeType{"application", "vnd.ms-powerpoint"},
	"pptx":   MimeType{"application", "vnd.openxmlformats-officedocument.presentationml.presentation"},
	"rar":    MimeType{"application", "vnd.rar"},
	"rtf":    MimeType{"application", "rtf"},
	"sh":     MimeType{"application", "x-sh"},
	"svg":    MimeType{"image", "svg+xml"},
	"tar":    MimeType{"application", "x-tar"},
	"tif":    MimeType{"image", "tiff"},
	"tiff":   MimeType{"image", "tiff"},
	"ts":     MimeType{"video", "mp2t"},
	"ttf":    MimeType{"font", "ttf"},
	"txt":    MimeType{"text", "plain"},
	"vsd":    MimeType{"application", "vnd.visio"},
	"wav":    MimeType{"audio", "wav"},
	"weba":   MimeType{"audio", "webm"},
	"webm":   MimeType{"video", "webm"},
	"webp":   MimeType{"image", "webp"},
	"woff":   MimeType{"font", "woff"},
	"woff2":  MimeType{"font", "woff2"},
	"xhtml":  MimeType{"application", "xhtml+xml"},
	"xls":    MimeType{"application", "vnd.ms-excel"},
	"xlsx":   MimeType{"application", "vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
	"xml":    MimeType{"application", "xml"},
	"xul":    MimeType{"application", "vnd.mozilla.xul+xml"},
	"zip":    MimeType{"application", "zip"},
	"3gp":    MimeType{"video", "3gpp"},
	"3g2":    MimeType{"video", "3gpp2"},
	"7z":     MimeType{"application", "x-7z-compressed"},
}

func GetMimeType(mimetype string) MimeType {
	m, present := mimetypes[mimetype]

	if present == false {
		log.Fatal("unknown mimetype")
	}

	return m
}
